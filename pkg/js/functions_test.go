package js_test

import (
	"reflect"
	"testing"

	"github.com/ohayocorp/anemos/pkg/js"
)

func TestFunctions(t *testing.T) {
	jsRuntime, err := js.NewJsRuntime()
	if err != nil {
		t.Errorf("NewJsRuntime() failed: %v", err)
	}

	jsRuntime.Function(reflect.ValueOf(NoParams))
	jsRuntime.Function(reflect.ValueOf(ReturnBool))
	jsRuntime.Function(reflect.ValueOf(ReturnBoolParam)).JsName("returnBool")
	jsRuntime.Function(reflect.ValueOf(ReturnBoolPointer))
	jsRuntime.Function(reflect.ValueOf(ReturnInt))
	jsRuntime.Function(reflect.ValueOf(ReturnIntParam)).JsName("returnInt")
	jsRuntime.Function(reflect.ValueOf(ReturnIntPointer))
	jsRuntime.Function(reflect.ValueOf(ReturnFloat))
	jsRuntime.Function(reflect.ValueOf(ReturnFloatParam)).JsName("returnFloat")
	jsRuntime.Function(reflect.ValueOf(ReturnFloatPointer))
	jsRuntime.Function(reflect.ValueOf(ReturnString))
	jsRuntime.Function(reflect.ValueOf(ReturnStringParam)).JsName("returnString")
	jsRuntime.Function(reflect.ValueOf(ReturnStringPointer))

	err = jsRuntime.Run(ReadScript(t, "tests/functions.js"), nil)
	if err != nil {
		t.Error(err)
	}
}

func NoParams() {
}

func ReturnBool() bool {
	return true
}

func ReturnBoolParam(arg bool) bool {
	return arg
}

func ReturnBoolPointer(arg *bool) *bool {
	return arg
}

func ReturnInt() int {
	return 1
}

func ReturnIntParam(arg int) int {
	return arg
}

func ReturnIntPointer(arg *int) *int {
	return arg
}

func ReturnFloat() float64 {
	return 1.2
}

func ReturnFloatParam(arg float64) float64 {
	return arg
}

func ReturnFloatPointer(arg *float64) *float64 {
	return arg
}

func ReturnString() string {
	return "test"
}

func ReturnStringParam(arg string) string {
	return arg
}

func ReturnStringPointer(arg *string) *string {
	return arg
}
