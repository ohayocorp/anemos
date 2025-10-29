package js_test

import (
	"reflect"
	"testing"

	"github.com/ohayocorp/anemos/pkg/cmd"
	"github.com/ohayocorp/anemos/pkg/js"
)

func TestMethods(t *testing.T) {
	jsRuntime, err := cmd.InitializeNewRuntime(&cmd.AnemosProgram{
		RegisterRuntimeCallback: func(jsRuntime *js.JsRuntime) error {
			jsRuntime.Type(reflect.TypeFor[MethodTest]()).Methods(
				js.Method("NoParams"),
				js.Method("ReturnBool"),
				js.Method("ReturnBoolParam").JsName("returnBool"),
				js.Method("ReturnBoolPointer"),
				js.Method("ReturnInt"),
				js.Method("ReturnIntParam").JsName("returnInt"),
				js.Method("ReturnIntPointer"),
				js.Method("ReturnFloat"),
				js.Method("ReturnFloatParam").JsName("returnFloat"),
				js.Method("ReturnFloatPointer"),
				js.Method("ReturnString"),
				js.Method("ReturnStringParam").JsName("returnString"),
				js.Method("ReturnStringPointer"),
			)

			instance := &MethodTest{}
			jsRuntime.Variable("", "test", reflect.ValueOf(instance))

			return nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = jsRuntime.Run(ReadScript(t, "tests/primitives-method.js"), nil)
	if err != nil {
		t.Error(err)
	}
}

type MethodTest struct{}

func (m *MethodTest) NoParams() {
}

func (m *MethodTest) ReturnBool() bool {
	return true
}

func (m *MethodTest) ReturnBoolParam(arg bool) bool {
	return arg
}

func (m *MethodTest) ReturnBoolPointer(arg *bool) *bool {
	return arg
}

func (m *MethodTest) ReturnInt() int {
	return 1
}

func (m *MethodTest) ReturnIntParam(arg int) int {
	return arg
}

func (m *MethodTest) ReturnIntPointer(arg *int) *int {
	return arg
}

func (m *MethodTest) ReturnFloat() float64 {
	return 1.2
}

func (m *MethodTest) ReturnFloatParam(arg float64) float64 {
	return arg
}

func (m *MethodTest) ReturnFloatPointer(arg *float64) *float64 {
	return arg
}

func (m *MethodTest) ReturnString() string {
	return "test"
}

func (m *MethodTest) ReturnStringParam(arg string) string {
	return arg
}

func (m *MethodTest) ReturnStringPointer(arg *string) *string {
	return arg
}
