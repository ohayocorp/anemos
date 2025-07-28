package js_test

import (
	"reflect"
	"testing"

	"github.com/ohayocorp/anemos/pkg/js"
)

type ExtensionObject struct {
	Property string
}

func TestExtensionMethods(t *testing.T) {
	jsRuntime, err := js.NewJsRuntime()
	if err != nil {
		t.Errorf("NewJsRuntime() failed: %v", err)
	}

	jsRuntime.Type(reflect.TypeFor[ExtensionObject]()).Fields(
		js.Field("Property"),
	).ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(ExtensionNoParams)),
		js.ExtensionMethod(reflect.ValueOf(ExtensionReturnProperty)).JsName("returnProperty"),
		js.ExtensionMethod(reflect.ValueOf(ExtensionReturnPropertyOverload)).JsName("returnProperty"),
	)

	object := &ExtensionObject{
		Property: "test",
	}

	jsRuntime.Variable("", "object", reflect.ValueOf(object))

	err = jsRuntime.Run(ReadScript(t, "tests/extension-methods.js"), nil)
	if err != nil {
		t.Error(err)
	}
}

func ExtensionNoParams(e *ExtensionObject) {
}

func ExtensionReturnProperty(e *ExtensionObject) string {
	return e.Property
}

func ExtensionReturnPropertyOverload(e *ExtensionObject, s string) string {
	return e.Property + s
}
