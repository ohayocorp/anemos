package js_test

import (
	"reflect"
	"testing"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

func TestPrimitivesString(t *testing.T) {
	jsRuntime := js.NewJsRuntime()

	globalVariable := "globalVariable"
	globalVariablePointer := core.Pointer("globalVariablePointer")
	globalVariableNamespace := "globalVariableNamespace"

	type object struct {
		Property string
		Pointer  *string
	}

	jsRuntime.Type(reflect.TypeFor[object]()).Fields(
		js.Field("Property"),
		js.Field("Pointer"),
	)

	instance := &object{
		Property: "instanceProperty",
		Pointer:  core.Pointer("instancePointer"),
	}

	jsRuntime.Variable("", "globalVariable", reflect.ValueOf(globalVariable))
	jsRuntime.Variable("", "globalVariablePointer", reflect.ValueOf(globalVariablePointer))
	jsRuntime.Variable("ns", "globalVariable", reflect.ValueOf(globalVariableNamespace))
	jsRuntime.Variable("", "globalObject", reflect.ValueOf(instance))

	err := jsRuntime.Run(ReadScript(t, "tests/primitives-string.js"), nil)
	if err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt(t *testing.T) {
	jsRuntime := js.NewJsRuntime()

	globalVariable := 1
	globalVariablePointer := core.Pointer(2)
	globalVariableNamespace := 3

	type object struct {
		Property int
		Pointer  *int
	}

	jsRuntime.Type(reflect.TypeFor[object]()).Fields(
		js.Field("Property"),
		js.Field("Pointer"),
	)

	instance := &object{
		Property: 4,
		Pointer:  core.Pointer(5),
	}

	jsRuntime.Variable("", "globalVariable", reflect.ValueOf(globalVariable))
	jsRuntime.Variable("", "globalVariablePointer", reflect.ValueOf(globalVariablePointer))
	jsRuntime.Variable("ns", "globalVariable", reflect.ValueOf(globalVariableNamespace))
	jsRuntime.Variable("", "globalObject", reflect.ValueOf(instance))

	err := jsRuntime.Run(ReadScript(t, "tests/primitives-int.js"), nil)
	if err != nil {
		t.Error(err)
	}
}

func TestPrimitivesBool(t *testing.T) {
	jsRuntime := js.NewJsRuntime()

	globalVariable := true
	globalVariablePointer := core.Pointer(true)
	globalVariableNamespace := true

	type object struct {
		Property bool
		Pointer  *bool
	}

	jsRuntime.Type(reflect.TypeFor[object]()).Fields(
		js.Field("Property"),
		js.Field("Pointer"),
	)

	instance := &object{
		Property: true,
		Pointer:  core.Pointer(true),
	}

	jsRuntime.Variable("", "globalVariable", reflect.ValueOf(globalVariable))
	jsRuntime.Variable("", "globalVariablePointer", reflect.ValueOf(globalVariablePointer))
	jsRuntime.Variable("ns", "globalVariable", reflect.ValueOf(globalVariableNamespace))
	jsRuntime.Variable("", "globalObject", reflect.ValueOf(instance))

	err := jsRuntime.Run(ReadScript(t, "tests/primitives-bool.js"), nil)
	if err != nil {
		t.Error(err)
	}
}
