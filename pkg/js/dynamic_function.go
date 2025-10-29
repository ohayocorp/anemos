package js

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"slices"
	"sort"
	"strings"

	"github.com/grafana/sobek"
)

const (
	jsFunction    FunctionType = "function"
	jsConstructor FunctionType = "constructor"
)

type overloadError struct {
	error error
}

type FunctionType string

type DynamicFunction struct {
	jsModule     string
	jsName       string
	functionType FunctionType
	function     reflect.Value
}

func (e overloadError) Error() string {
	return e.error.Error()
}

func (jsRuntime *JsRuntime) initializeFunctions(module *sobek.Object, functions []*DynamicFunction, prototype *sobek.Object) {
	keysInOrder := []string{}
	functionOverloads := map[string][]*DynamicFunction{}
	constructorOverloads := map[string][]*DynamicFunction{}

	for _, function := range functions {
		key := function.jsName
		if function.jsModule != "" {
			key = fmt.Sprintf("%s/%s", function.jsModule, key)
		}

		keysInOrder = append(keysInOrder, key)

		switch function.functionType {
		case jsFunction:
			functionOverloads[key] = append(functionOverloads[key], function)
		case jsConstructor:
			constructorOverloads[key] = append(constructorOverloads[key], function)
		default:
			panic(fmt.Errorf("unknown function type %s", function.functionType))
		}
	}

	sortOverloads := func(m map[string][]*DynamicFunction) {
		for _, overloads := range m {
			sort.SliceStable(overloads, func(i, j int) bool {
				return overloads[i].function.Type().NumIn() > overloads[j].function.Type().NumIn()
			})
		}
	}

	sortOverloads(functionOverloads)
	sortOverloads(constructorOverloads)

	for _, key := range keysInOrder {
		if overloads, ok := functionOverloads[key]; ok {
			fn := func(call sobek.FunctionCall) sobek.Value {
				defer func() {
					if r := recover(); r != nil {
						if jsErr, ok := r.(JsError); ok {
							panic(jsRuntime.Runtime.ToValue(jsErr.Err))
						}

						// This is not an expected error, panic with the error so that the users can report it.
						// Using a JS exception here since a Golang panic will pollute the stack trace with Sobek
						// runtime internals and occasionally cause an invalid memory access error which hides the real error.
						err := fmt.Errorf("unexpected error in function %s: %v\n%s", overloads[0].jsName, r, string(debug.Stack()))
						panic(jsRuntime.Runtime.ToValue(err))
					}
				}()

				value, err := jsRuntime.handleFunction(overloads, call, prototype != nil)
				if err != nil {
					Throw(err)
				}

				return value
			}

			if prototype != nil {
				err := prototype.Set(overloads[0].jsName, fn)
				if err != nil {
					panic(fmt.Errorf("failed to add function %s to prototype: %w", overloads[0].jsName, err))
				}
			} else {
				err := module.Set(overloads[0].jsName, reflect.ValueOf(fn).Interface())

				if err != nil {
					panic(fmt.Errorf("failed to add function %s to module %s: %w", overloads[0].jsName, overloads[0].jsModule, err))
				}
			}

			continue
		}

		if overloads, ok := constructorOverloads[key]; ok {
			fn := func(call sobek.ConstructorCall) *sobek.Object {
				object, err := jsRuntime.handleConstructor(overloads, call)
				if err != nil {
					panic(jsRuntime.Runtime.ToValue(err))
				}

				return object
			}

			err := module.Set(overloads[0].jsName, reflect.ValueOf(fn).Interface())
			if err != nil {
				panic(fmt.Errorf("failed to add constructor %s to module %s: %w", overloads[0].jsName, overloads[0].jsModule, err))
			}

			continue
		}

		panic(fmt.Errorf("no overloads for %s", key))
	}
}

func (jsRuntime *JsRuntime) handleFunction(overloads []*DynamicFunction, call sobek.FunctionCall, containsThisArg bool) (sobek.Value, error) {
	overloadErrors := map[*DynamicFunction]error{}

	for _, overload := range overloads {
		value, err := jsRuntime.handleOverload(overload, call.This, call.Arguments, containsThisArg)
		if err != nil {
			if _, ok := err.(overloadError); !ok {
				return nil, err
			}

			overloadErrors[overload] = err
			continue
		}

		return value, nil
	}

	return nil, overloadsToError(overloads, overloadErrors, containsThisArg)
}

func (jsRuntime *JsRuntime) handleConstructor(overloads []*DynamicFunction, call sobek.ConstructorCall) (*sobek.Object, error) {
	overloadErrors := map[*DynamicFunction]error{}

	for _, overload := range overloads {
		value, err := jsRuntime.handleOverload(overload, call.This, call.Arguments, false)
		if err != nil {
			if _, ok := err.(overloadError); !ok {
				return nil, err
			}

			overloadErrors[overload] = err
			continue
		}

		return value.ToObject(jsRuntime.Runtime), nil
	}

	return nil, overloadsToError(overloads, overloadErrors, false)
}

func (jsRuntime *JsRuntime) handleOverload(overload *DynamicFunction, this sobek.Value, callArgs []sobek.Value, containsThisArg bool) (sobek.Value, error) {
	jsArgs := []sobek.Value{}
	if containsThisArg {
		jsArgs = append(jsArgs, this)
	}

	jsArgs = append(jsArgs, callArgs...)

	functionType := overload.function.Type()
	argCount := functionType.NumIn()

	for i := 0; i < functionType.NumIn(); i++ {
		if functionType.In(i) == reflect.TypeFor[*JsRuntime]() {
			argCount--
		}
	}

	if argCount != len(jsArgs) {
		expected := argCount
		got := len(jsArgs)

		if containsThisArg {
			expected--
			got--
		}

		return nil, overloadError{error: fmt.Errorf("expected %d arguments, got %d", expected, got)}
	}

	args := []reflect.Value{}
	jsArgIndex := 0

	for i := 0; i < functionType.NumIn(); i++ {
		if functionType.In(i) == reflect.TypeFor[*JsRuntime]() {
			args = append(args, reflect.ValueOf(jsRuntime))
			continue
		}

		jsArg := jsArgs[jsArgIndex]
		arg, err := jsRuntime.MarshalToGo(jsArg, functionType.In(i))
		if err != nil {
			return nil, overloadError{error: err}
		}

		args = append(args, arg)
		jsArgIndex++
	}

	if functionType.IsVariadic() && len(args) > 0 {
		// Flatten variadic arguments as they will be converted to a slice by the reflect package.
		slice := args[len(args)-1]
		if slice.Kind() != reflect.Slice {
			return nil, overloadError{error: fmt.Errorf("expected variadic argument to be a slice, got %s", slice.Kind())}
		}

		sliceArgs := []reflect.Value{}
		for i := 0; i < slice.Len(); i++ {
			sliceArgs = append(sliceArgs, slice.Index(i))
		}

		args = append(args[:len(args)-1], sliceArgs...)
	}

	result := overload.function.Call(args)

	if len(result) == 0 {
		return nil, nil
	}

	if len(result) == 1 {
		r := result[0]
		if r.Type() == reflect.TypeFor[error]() {
			if r.IsNil() {
				return nil, nil
			}

			return nil, r.Interface().(error)
		}

		value, err := jsRuntime.MarshalToJs(r)
		if err != nil {
			return nil, overloadError{error: err}
		}

		return value, nil
	}

	if len(result) == 2 {
		r := result[0]
		err := result[1]
		if err.Type() != reflect.TypeFor[error]() {
			return nil, overloadError{error: fmt.Errorf("expected error as second return value")}
		}

		if err.IsNil() {
			value, err := jsRuntime.MarshalToJs(r)
			if err != nil {
				return nil, overloadError{error: err}
			}

			return value, nil
		}

		return nil, err.Interface().(error)
	}

	return nil, overloadError{error: fmt.Errorf("unexpected number of return values. expected 0, 1, or 2, got %d", len(result))}
}

func overloadsToError(overloads []*DynamicFunction, overloadErrors map[*DynamicFunction]error, containsThisArg bool) error {
	errors := []string{}

	for _, overload := range overloads {
		err := overloadErrors[overload]
		if err == nil {
			continue
		}

		argTypes := []string{}
		for i := 0; i < overload.function.Type().NumIn(); i++ {
			if i == 0 && containsThisArg {
				continue
			}

			argType := typeNameToJsTypeName(overload.function.Type().In(i).String())
			argTypes = append(argTypes, argType)
		}

		e := fmt.Sprintf("%s -> %s", fmt.Sprintf("%s(%s)", overload.jsName, strings.Join(argTypes, ", ")), err.Error())
		errors = append(errors, e)
	}

	slices.Reverse(errors)

	return fmt.Errorf("no overload found with given arguments. overload errors:\n%s", strings.Join(errors, "\n"))
}
