package js

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"

	"github.com/grafana/sobek"
)

type VariableRegistration struct {
	jsModule string
	jsName   string
	value    reflect.Value
}

type TypeRegistration struct {
	objectType           reflect.Type
	jsModule             string
	jsName               string
	typeConversions      map[reflect.Type][]*TypeConversion
	fields               []*FieldRegistration
	methods              []*MethodRegistration
	constructors         []*ConstructorRegistration
	extensionMethods     []*ExtensionMethodRegistration
	aliases              []*TypeAlias
	disableObjectMapping bool
}

type TypeAlias struct {
	jsModule string
	jsName   string
}

type FunctionRegistration struct {
	function     reflect.Value
	jsModule     string
	jsName       string
	functionType FunctionType
}

type ExtensionMethodRegistration struct {
	function reflect.Value
	jsName   string
}

type ConstructorRegistration struct {
	function reflect.Value
}

type FieldRegistration struct {
	fieldName string
	jsName    string
}

type MethodRegistration struct {
	methodName string
	jsName     string
}

type TypeConversion struct {
	converter reflect.Value
}

func (jsRuntime *JsRuntime) registerTypes() {
	for _, typeRegistration := range jsRuntime.typeRegistrations {
		// Enable type conversions for non-struct types. These types cannot have fields, methods, or constructors.
		if typeRegistration.objectType.Kind() != reflect.Struct &&
			len(typeRegistration.constructors) == 0 &&
			len(typeRegistration.fields) == 0 &&
			len(typeRegistration.methods) == 0 &&
			len(typeRegistration.extensionMethods) == 0 {

			for targetType, conversions := range typeRegistration.typeConversions {
				jsRuntime.typeConversions[targetType] = append(jsRuntime.typeConversions[targetType], conversions...)
			}

			continue
		}

		template := jsRuntime.createTemplate(typeRegistration.objectType)
		template.jsModule = typeRegistration.jsModule
		template.jsName = typeRegistration.jsName

		if template.jsName == "" {
			template.jsName = typeToJsTypeName(template.objectType)
		}

		if typeRegistration.disableObjectMapping {
			jsRuntime.disabledObjectMappings.Add(typeRegistration.objectType)
		}

		for targetType, conversions := range typeRegistration.typeConversions {
			jsRuntime.typeConversions[targetType] = append(jsRuntime.typeConversions[targetType], conversions...)
		}

		fields := typeRegistration.fields
		sort.Slice(fields, func(i, j int) bool {
			return fields[i].fieldName < fields[j].fieldName
		})

		objectType := typeRegistration.objectType

		for _, field := range fields {
			if field.jsName == "" {
				field.jsName = toCamelCase(field.fieldName)
			}

			f, ok := objectType.FieldByName(field.fieldName)

			if !ok {
				panic(fmt.Errorf("field %s not found in type %s", field.fieldName, objectType.String()))
			}

			template.jsToGoNameMappings[field.jsName] = append(template.jsToGoNameMappings[field.jsName], f.Name)
			template.goToJsNameMappings[f.Name] = field.jsName
			template.exportedFields.Add(field.fieldName)
		}

		for _, method := range typeRegistration.methods {
			if method.jsName == "" {
				method.jsName = toCamelCase(method.methodName)
			}

			m, ok := objectType.MethodByName(method.methodName)
			if !ok {
				m, ok = reflect.PointerTo(objectType).MethodByName(method.methodName)
			}

			if !ok {
				panic(fmt.Errorf("method %s not found in type %s", method.methodName, objectType.String()))
			}

			function := &DynamicFunction{
				jsModule:     "",
				jsName:       method.jsName,
				functionType: jsFunction,
				function:     m.Func,
			}

			template.functions = append(template.functions, function)

			template.jsToGoNameMappings[method.jsName] = append(template.jsToGoNameMappings[method.jsName], m.Name)
			template.goToJsNameMappings[m.Name] = method.jsName
		}

		for _, extensionMethod := range typeRegistration.extensionMethods {
			signature := runtime.FuncForPC(extensionMethod.function.Pointer()).Name()
			functionName := typeNameToJsTypeName(signature)

			if extensionMethod.jsName == "" {
				extensionMethod.jsName = toCamelCase(functionName)
			}

			function := &DynamicFunction{
				jsModule:     "",
				jsName:       extensionMethod.jsName,
				functionType: jsFunction,
				function:     extensionMethod.function,
			}

			template.functions = append(template.functions, function)

			template.jsToGoNameMappings[extensionMethod.jsName] = append(template.jsToGoNameMappings[extensionMethod.jsName], functionName)
			template.goToJsNameMappings[functionName] = extensionMethod.jsName
		}

		for _, constructor := range typeRegistration.constructors {
			function := &DynamicFunction{
				jsModule:     template.jsModule,
				jsName:       typeRegistration.jsName,
				functionType: jsConstructor,
				function:     constructor.function,
			}

			template.functions = append(template.functions, function)

			for _, alias := range typeRegistration.aliases {
				function := &DynamicFunction{
					jsModule:     alias.jsModule,
					jsName:       alias.jsName,
					functionType: jsConstructor,
					function:     constructor.function,
				}

				template.functions = append(template.functions, function)
			}
		}
	}
}

func (jsRuntime *JsRuntime) registerFunctions() {
	for _, function := range jsRuntime.functionRegistrations {
		jsModule := function.jsModule
		jsName := function.jsName

		if jsName == "" {
			functionSignature := runtime.FuncForPC(function.function.Pointer()).Name()
			jsName = toCamelCase(typeNameToJsTypeName(functionSignature))
		}

		function := &DynamicFunction{
			jsModule:     jsModule,
			jsName:       jsName,
			functionType: function.functionType,
			function:     function.function,
		}

		jsRuntime.functions = append(jsRuntime.functions, function)
	}
}

func (jsRuntime *JsRuntime) Variable(jsModule, jsName string, value reflect.Value) *VariableRegistration {
	v := &VariableRegistration{
		jsModule: jsModule,
		jsName:   jsName,
		value:    value,
	}

	jsRuntime.variableRegistrations = append(jsRuntime.variableRegistrations, v)

	return v
}

func (jsRuntime *JsRuntime) Type(objectType reflect.Type) *TypeRegistration {
	if t, ok := jsRuntime.typeRegistrations[objectType]; ok {
		return t
	}

	t := &TypeRegistration{
		objectType:      objectType,
		jsName:          typeToJsTypeName(objectType),
		typeConversions: map[reflect.Type][]*TypeConversion{},
	}

	jsRuntime.typeRegistrations[objectType] = t

	return t
}

func (t *TypeRegistration) Alias(module, name string) *TypeRegistration {
	alias := &TypeAlias{
		jsModule: module,
		jsName:   name,
	}

	t.aliases = append(t.aliases, alias)

	return t
}

func (t *TypeRegistration) JsModule(module string) *TypeRegistration {
	t.jsModule = module
	return t
}

func (t *TypeRegistration) JsName(name string) *TypeRegistration {
	t.jsName = name
	return t
}

func (t *TypeRegistration) Fields(fields ...*FieldRegistration) *TypeRegistration {
	t.fields = append(t.fields, fields...)
	return t
}

func (t *TypeRegistration) Methods(methods ...*MethodRegistration) *TypeRegistration {
	t.methods = append(t.methods, methods...)
	return t
}

func (t *TypeRegistration) ExtensionMethods(extensionMethods ...*ExtensionMethodRegistration) *TypeRegistration {
	t.extensionMethods = append(t.extensionMethods, extensionMethods...)
	return t
}

func (t *TypeRegistration) Constructors(constructors ...*ConstructorRegistration) *TypeRegistration {
	t.constructors = append(t.constructors, constructors...)
	return t
}

func (t *TypeRegistration) ClearConstructors() *TypeRegistration {
	t.constructors = nil
	return t
}

func (t *TypeRegistration) TypeConversion(converter reflect.Value) *TypeRegistration {
	converterType := converter.Type()
	if converterType.NumOut() == 0 {
		panic(fmt.Errorf("converter must have at least one output %v", converterType))
	}

	if converterType.NumOut() > 2 {
		panic(fmt.Errorf("converter must return at most two values %v", converterType))
	}

	outType := converterType.Out(0)

	if outType.Kind() == reflect.Struct && outType != reflect.PointerTo(t.objectType) {
		panic(fmt.Errorf("converter must return a pointer to %s as first output %v", t.objectType, converterType))
	}

	if converterType.NumOut() == 2 && converterType.Out(1) != reflect.TypeFor[error]() {
		panic(fmt.Errorf("converter must return error as second output %v", converterType))
	}

	if converterType.NumIn() != 1 && converterType.NumIn() != 2 {
		panic(fmt.Errorf("converter must have one or two inputs %v", converterType))
	}

	convertFrom := converterType.In(0)

	if converterType.NumIn() == 2 {
		if converterType.In(0) != reflect.TypeFor[*JsRuntime]() {
			panic(fmt.Errorf("converter first input type must be *JsRuntime %v", converterType))
		}

		convertFrom = converterType.In(1)
	}

	if convertFrom != reflect.TypeFor[sobek.Value]() {
		panic(fmt.Errorf("converter first input type must be sobek.Value %v", converterType))
	}

	t.typeConversions[t.objectType] = append(t.typeConversions[t.objectType], &TypeConversion{
		converter: converter,
	})

	return t
}

func (t *TypeRegistration) DisableObjectMapping() *TypeRegistration {
	t.disableObjectMapping = true
	return t
}

func (jsRuntime *JsRuntime) Function(function reflect.Value) *FunctionRegistration {
	f := &FunctionRegistration{
		function:     function,
		functionType: jsFunction,
	}

	jsRuntime.functionRegistrations = append(jsRuntime.functionRegistrations, f)

	return f
}

func (jsRuntime *JsRuntime) Constructor(function reflect.Value) *FunctionRegistration {
	f := &FunctionRegistration{
		function:     function,
		functionType: jsConstructor,
	}

	jsRuntime.functionRegistrations = append(jsRuntime.functionRegistrations, f)

	return f
}

func (f *FunctionRegistration) JsModule(module string) *FunctionRegistration {
	f.jsModule = module
	return f
}

func (f *FunctionRegistration) JsName(name string) *FunctionRegistration {
	f.jsName = name
	return f
}

func Field(name string) *FieldRegistration {
	return &FieldRegistration{
		fieldName: name,
	}
}

func (f *FieldRegistration) JsName(name string) *FieldRegistration {
	f.jsName = name
	return f
}

func Method(name string) *MethodRegistration {
	return &MethodRegistration{
		methodName: name,
	}
}

func (m *MethodRegistration) JsName(name string) *MethodRegistration {
	m.jsName = name
	return m
}

func ExtensionMethod(function reflect.Value) *ExtensionMethodRegistration {
	return &ExtensionMethodRegistration{
		function: function,
	}
}

func (e *ExtensionMethodRegistration) JsName(name string) *ExtensionMethodRegistration {
	e.jsName = name
	return e
}

func Constructor(function reflect.Value) *ConstructorRegistration {
	return &ConstructorRegistration{
		function: function,
	}
}
