package js

import (
	"errors"
	"fmt"
	"go/ast"
	"reflect"

	"github.com/grafana/sobek"
)

func (jsRuntime *JsRuntime) MarshalToGo(jsArg sobek.Value, expectedType reflect.Type) (reflect.Value, error) {
	if jsArg == nil || jsArg == sobek.Null() || jsArg == sobek.Undefined() {
		return jsRuntime.marshalToGoNil(expectedType)
	}

	jsType := jsArg.ExportType()
	if jsType == reflect.TypeFor[DynamicObject]() || jsType == reflect.TypeFor[*DynamicObject]() {
		return jsRuntime.marshalToGoDynamicObject(jsArg, expectedType)
	}

	if jsType == reflect.TypeFor[DynamicArray]() || jsType == reflect.TypeFor[*DynamicArray]() {
		return jsRuntime.marshalToGoDynamicArray(jsArg, expectedType)
	}

	if expectedType.Kind() == reflect.Ptr {
		return jsRuntime.marshalToGoPointer(jsArg, expectedType)
	}

	if expectedType.Kind() == reflect.Interface {
		return reflect.ValueOf(jsArg), nil
	}

	ok, result, err := jsRuntime.tryConvert(expectedType, jsType, jsArg)
	if ok {
		return result, err
	}

	if jsRuntime.disabledObjectMappings.Contains(expectedType) {
		return reflect.Value{}, fmt.Errorf(
			"cannot convert %s to %s, object mapping is disabled for this type",
			jsType.String(),
			expectedType.String())
	}

	jsKind := jsType.Kind()
	switch jsKind {
	case reflect.Map:
		return jsRuntime.marshalToGoMap(jsArg, expectedType)
	case reflect.Slice:
		return jsRuntime.marshalToGoSlice(jsArg, expectedType)
	case reflect.Func:
		return jsRuntime.marshalToGoFunc(jsArg, expectedType)
	case reflect.Bool:
		return jsRuntime.marshalToGoBool(jsArg, expectedType)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return jsRuntime.marshalToGoInt(jsArg, expectedType)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return jsRuntime.marshalToGoUint(jsArg, expectedType)
	case reflect.Float32, reflect.Float64:
		return jsRuntime.marshalToGoFloat(jsArg, expectedType)
	case reflect.String:
		return jsRuntime.marshalToGoString(jsArg, expectedType)
	}

	return reflect.Value{}, fmt.Errorf("unsupported kind %s", jsKind.String())
}

func (jsRuntime *JsRuntime) marshalToGoNil(expectedType reflect.Type) (reflect.Value, error) {
	if expectedType.Kind() == reflect.Ptr {
		pointer := reflect.New(expectedType)
		return pointer.Elem(), nil
	}

	if expectedType.Kind() == reflect.Interface {
		return reflect.New(expectedType).Elem(), nil
	}

	if expectedType.Kind() == reflect.Slice {
		return reflect.New(expectedType).Elem(), nil
	}

	return reflect.Value{}, fmt.Errorf("cannot convert nil to %s", expectedType.String())
}

func (jsRuntime *JsRuntime) marshalToGoPointer(jsArg sobek.Value, expectedType reflect.Type) (reflect.Value, error) {
	pointer := reflect.New(expectedType.Elem())

	object, err := jsRuntime.MarshalToGo(jsArg, expectedType.Elem())
	if err != nil {
		return reflect.Value{}, err
	}

	pointer.Elem().Set(object)
	return pointer, nil
}

func (jsRuntime *JsRuntime) marshalToGoDynamicObject(jsArg sobek.Value, expectedType reflect.Type) (reflect.Value, error) {
	var backingObject reflect.Value

	if dynamicObject, ok := jsArg.Export().(DynamicObject); ok {
		backingObject = dynamicObject.backingObject
	} else {
		backingObject = jsArg.Export().(*DynamicObject).backingObject
	}

	if expectedType.Kind() == reflect.Interface {
		if backingObject.Type().Implements(expectedType) {
			return backingObject, nil
		}

		return reflect.Value{}, fmt.Errorf("backingObject must implement %s", expectedType.String())
	}

	if expectedType != backingObject.Type() && expectedType != reflect.TypeFor[interface{}]() {
		return reflect.Value{}, fmt.Errorf("backingObject must be of type %s", expectedType.String())
	}

	return backingObject, nil
}

func (jsRuntime *JsRuntime) marshalToGoDynamicArray(jsArg sobek.Value, expectedType reflect.Type) (reflect.Value, error) {
	var backingSlice reflect.Value

	if dynamicArray, ok := jsArg.Export().(DynamicArray); ok {
		backingSlice = dynamicArray.backingSlice
	} else {
		backingSlice = jsArg.Export().(*DynamicArray).backingSlice
	}

	if expectedType.Kind() == reflect.Interface {
		if backingSlice.Type().Implements(expectedType) {
			return backingSlice, nil
		}

		return reflect.Value{}, fmt.Errorf("backingSlice must implement %s", expectedType.String())
	}

	if expectedType != backingSlice.Type() && expectedType != reflect.TypeFor[interface{}]() {
		return reflect.Value{}, fmt.Errorf("backingSlice must be of type %s", expectedType.String())
	}

	return backingSlice, nil
}

func (jsRuntime *JsRuntime) marshalToGoMap(jsArg sobek.Value, expectedType reflect.Type) (reflect.Value, error) {
	object, ok := jsArg.(*sobek.Object)
	if !ok {
		return reflect.Value{}, fmt.Errorf("expected object, got %s", jsArg.ExportType().Kind())
	}

	if expectedType.Kind() == reflect.Map {
		keyType := expectedType.Key()
		valueType := expectedType.Elem()

		if keyType.Kind() != reflect.String {
			return reflect.Value{}, fmt.Errorf("map key must be of type string, got %s", keyType.String())
		}

		result := reflect.MakeMap(expectedType)

		for _, propertyName := range object.GetOwnPropertyNames() {
			property := object.Get(propertyName)

			key := reflect.New(keyType).Elem()
			key.SetString(propertyName)

			value, err := jsRuntime.MarshalToGo(property, valueType)
			if err != nil {
				return reflect.Value{}, fmt.Errorf("could not convert %v to %v for map value %s: %w", property, valueType, propertyName, err)
			}

			result.SetMapIndex(key, value)
		}

		return result, nil
	}

	result := reflect.New(expectedType).Elem()
	template := jsRuntime.getTemplate(expectedType)

	if template == nil {
		return reflect.Value{}, fmt.Errorf("no template found for type %s", expectedType.String())
	}

	marshalErrors := map[string][]error{}
	tryCounts := map[string]int{}

	for i := 0; i < expectedType.NumField(); i++ {
		field := expectedType.Field(i)

		if !ast.IsExported(field.Name) {
			continue
		}

		name := template.goToJsNameMappings[field.Name]

		var value sobek.Value
		if field.Anonymous {
			value = object
		} else {
			value = object.Get(name)
		}

		if value != nil {
			tryCounts[field.Name]++

			fieldValue, err := jsRuntime.MarshalToGo(value, field.Type)
			if err != nil {
				marshalErrors[name] = append(marshalErrors[name], fmt.Errorf("could not convert %v to %v for field %s: %w", value, field.Type, field.Name, err))
				continue
			}

			result.Field(i).Set(fieldValue)
		}
	}

	for name, tryCount := range tryCounts {
		if tryCount == 0 {
			continue
		}

		fieldErrors := marshalErrors[name]
		if len(fieldErrors) == tryCount {
			return reflect.Value{}, errors.Join(fieldErrors...)
		}
	}

	return result, nil
}

func (jsRuntime *JsRuntime) marshalToGoSlice(jsArg sobek.Value, expectedType reflect.Type) (reflect.Value, error) {
	if expectedType.Kind() != reflect.Slice {
		return reflect.Value{}, fmt.Errorf("cannot convert slice to %s", expectedType.String())
	}

	jsArgKind := jsArg.ExportType().Kind()

	object, ok := jsArg.(*sobek.Object)
	if !ok {
		return reflect.Value{}, fmt.Errorf("expected object, got %s", jsArgKind)
	}

	sym := object.GetSymbol(sobek.SymIterator)
	if sym == nil || sym == sobek.Undefined() {
		return reflect.Value{}, fmt.Errorf("expected iterable, got %s", jsArgKind)
	}

	slice := reflect.MakeSlice(expectedType, 0, 0)

	err := jsRuntime.runtime.Try(func() {
		jsRuntime.runtime.ForOf(jsArg, func(value sobek.Value) bool {
			result, err := jsRuntime.MarshalToGo(value, expectedType.Elem())
			if err != nil {
				panic(jsRuntime.runtime.ToValue(err))
			}

			slice = reflect.Append(slice, result)

			return true
		})
	})

	if err != nil {
		return reflect.Value{}, err
	}

	return slice, nil
}

func (jsRuntime *JsRuntime) marshalToGoFunc(jsArg sobek.Value, expectedType reflect.Type) (reflect.Value, error) {
	if expectedType.Kind() != reflect.Func {
		return reflect.Value{}, fmt.Errorf("cannot convert function to %s", expectedType.String())
	}

	jsFunction, ok := sobek.AssertFunction(jsArg)
	if !ok {
		return reflect.Value{}, fmt.Errorf("expected function, got %s", jsArg.ExportType().Kind())
	}

	wrapper := func(args []reflect.Value) []reflect.Value {
		numOut := expectedType.NumOut()
		lastReturnIsError := numOut > 0 && expectedType.Out(numOut-1) == reflect.TypeFor[error]()

		jsArgs := []sobek.Value{}
		isVariadic := expectedType.IsVariadic()

		for i, arg := range args {
			// Append elements of variadic argument to the end of the arguments array.
			if i == expectedType.NumIn()-1 && isVariadic {
				for j := 0; j < arg.Len(); j++ {
					marshalledArg, err := jsRuntime.MarshalToJs(arg.Index(j))
					if err != nil {
						panic(jsRuntime.runtime.ToValue(err))
					}

					jsArgs = append(jsArgs, marshalledArg)
				}

				continue
			}

			marshalledArg, err := jsRuntime.MarshalToJs(arg)
			if err != nil {
				panic(jsRuntime.runtime.ToValue(err))
			}

			jsArgs = append(jsArgs, marshalledArg)
		}

		jsResult, err := jsFunction(sobek.Undefined(), jsArgs...)

		if numOut == 0 {
			if err != nil {
				panic(jsRuntime.runtime.ToValue(err))
			}

			return []reflect.Value{}
		}

		if numOut == 1 {
			if err != nil && !lastReturnIsError {
				panic(jsRuntime.runtime.ToValue(err))
			}

			if err != nil {
				return []reflect.Value{reflect.ValueOf(jsRuntime.unwindSobekException(err))}
			}

			goResult, marshalErr := jsRuntime.MarshalToGo(jsResult, expectedType.Out(0))
			if marshalErr != nil {
				panic(jsRuntime.runtime.ToValue(marshalErr))
			}

			return []reflect.Value{goResult}
		}

		if numOut == 2 {
			if !lastReturnIsError {
				panic(jsRuntime.runtime.ToValue(fmt.Errorf("second return value must be of type error")))
			}

			goResult, marshalErr := jsRuntime.MarshalToGo(jsResult, expectedType.Out(0))
			if marshalErr != nil {
				panic(jsRuntime.runtime.ToValue(marshalErr))
			}

			return []reflect.Value{
				goResult,
				reflect.ValueOf(jsRuntime.unwindSobekException(err)),
			}
		}

		panic(jsRuntime.runtime.ToValue(fmt.Errorf("function with more than 2 return values is not supported")))
	}

	return reflect.MakeFunc(expectedType, wrapper), nil
}

func (jsRuntime *JsRuntime) unwindSobekException(err error) error {
	if err == nil {
		return nil
	}

	exception, ok := err.(*sobek.Exception)
	if !ok {
		return err
	}

	object, ok := exception.Value().(*sobek.Object)
	if !ok {
		return err
	}

	value := object.Get("value")
	if value == nil {
		return err
	}

	if value.ExportType().AssignableTo(reflect.TypeFor[error]()) {
		return value.Export().(error)
	}

	return err
}

func (jsRuntime *JsRuntime) marshalToGoBool(jsArg sobek.Value, expectedType reflect.Type) (reflect.Value, error) {
	if expectedType.Kind() != reflect.Bool {
		return reflect.Value{}, fmt.Errorf("cannot convert bool to %s", expectedType.String())
	}

	result := reflect.ValueOf(jsArg.ToBoolean())

	if expectedType.Kind() != reflect.Bool {
		result = result.Convert(expectedType)
	}

	return result, nil
}

func (jsRuntime *JsRuntime) marshalToGoInt(jsArg sobek.Value, expectedType reflect.Type) (reflect.Value, error) {
	if expectedType.Kind() != reflect.Int &&
		expectedType.Kind() != reflect.Int8 &&
		expectedType.Kind() != reflect.Int16 &&
		expectedType.Kind() != reflect.Int32 &&
		expectedType.Kind() != reflect.Int64 &&
		expectedType.Kind() != reflect.Uint &&
		expectedType.Kind() != reflect.Uint8 &&
		expectedType.Kind() != reflect.Uint16 &&
		expectedType.Kind() != reflect.Uint32 &&
		expectedType.Kind() != reflect.Uint64 &&
		expectedType.Kind() != reflect.Float32 &&
		expectedType.Kind() != reflect.Float64 {

		return reflect.Value{}, fmt.Errorf("cannot convert int to %s", expectedType.String())
	}

	result := reflect.ValueOf(jsArg.ToInteger())

	if expectedType.Kind() != reflect.Int64 {
		result = result.Convert(expectedType)
	}

	return result, nil
}

func (jsRuntime *JsRuntime) marshalToGoUint(jsArg sobek.Value, expectedType reflect.Type) (reflect.Value, error) {
	if expectedType.Kind() != reflect.Uint &&
		expectedType.Kind() != reflect.Uint8 &&
		expectedType.Kind() != reflect.Uint16 &&
		expectedType.Kind() != reflect.Uint32 &&
		expectedType.Kind() != reflect.Uint64 &&
		expectedType.Kind() != reflect.Int &&
		expectedType.Kind() != reflect.Int16 &&
		expectedType.Kind() != reflect.Int32 &&
		expectedType.Kind() != reflect.Int64 {

		return reflect.Value{}, fmt.Errorf("cannot convert uint to %s", expectedType.String())
	}

	result := reflect.ValueOf(uint64(jsArg.ToInteger()))

	if expectedType.Kind() != reflect.Uint64 {
		result = result.Convert(expectedType)
	}

	return result, nil
}

func (jsRuntime *JsRuntime) marshalToGoFloat(jsArg sobek.Value, expectedType reflect.Type) (reflect.Value, error) {
	if expectedType.Kind() != reflect.Float32 && expectedType.Kind() != reflect.Float64 {
		return reflect.Value{}, fmt.Errorf("cannot convert float to %s", expectedType.String())
	}

	result := reflect.ValueOf(jsArg.ToFloat())

	if expectedType.Kind() != reflect.Float64 {
		result = result.Convert(expectedType)
	}

	return result, nil
}

func (jsRuntime *JsRuntime) marshalToGoString(jsArg sobek.Value, expectedType reflect.Type) (reflect.Value, error) {
	if expectedType.Kind() != reflect.String {
		return reflect.Value{}, fmt.Errorf("cannot convert string to %s", expectedType.String())
	}

	pointer := reflect.New(expectedType)
	result := pointer.Elem()

	result.SetString(jsArg.String())

	if expectedType.Kind() != reflect.String {
		result = result.Convert(expectedType)
	}

	return result, nil
}

func (jsRuntime *JsRuntime) tryConvert(expectedType reflect.Type, jsType reflect.Type, jsArg sobek.Value) (bool, reflect.Value, error) {
	conversions, ok := jsRuntime.typeConversions[expectedType]
	if !ok {
		return false, reflect.Value{}, nil
	}

	for _, conversion := range conversions {
		inputs := []reflect.Value{}
		if conversion.converter.Type().NumIn() == 2 {
			inputs = append(inputs, reflect.ValueOf(jsRuntime))
		}

		inputs = append(inputs, reflect.ValueOf(jsArg))

		values := conversion.converter.Call(inputs)
		numOut := conversion.converter.Type().NumOut()

		if numOut == 0 {
			panic(fmt.Errorf("conversion function must return at least one value"))
		}

		returnValue := values[0]

		if expectedType.Kind() == reflect.Pointer || expectedType.Kind() == reflect.Interface || expectedType.Kind() == reflect.Struct {
			returnValue = returnValue.Elem()
		}

		if numOut == 1 {
			return true, returnValue, nil
		}

		if numOut == 2 {
			if values[1].Type() != reflect.TypeFor[error]() {
				panic(fmt.Errorf("conversion function can only return an error as second value"))
			}

			if values[1].IsNil() {
				return true, returnValue, nil
			}

			return true, reflect.Value{}, values[1].Interface().(error)
		}

		panic(fmt.Errorf("conversion function must return at most two values"))
	}

	return false, reflect.Value{}, nil
}
