package js_test

import (
	"reflect"
	"testing"

	"github.com/ohayocorp/anemos/pkg/js"
)

func TestSlices(t *testing.T) {
	jsRuntime := js.NewJsRuntime()

	type elem struct {
		Property string
	}

	type object struct {
		Array []*elem
	}

	jsRuntime.Type(reflect.TypeFor[elem]()).Fields(
		js.Field("Property"),
	)

	jsRuntime.Type(reflect.TypeFor[object]()).Fields(
		js.Field("Array"),
	)

	boolArray := []bool{true, false}
	intArray := []int{1, 2}
	floatArray := []float64{1.2, 2.3}
	stringArray := []string{"a", "b"}
	objectArray := &object{
		Array: []*elem{{
			Property: "a",
		}, {
			Property: "b",
		}},
	}

	jsRuntime.Variable("", "boolArray", reflect.ValueOf(boolArray))
	jsRuntime.Variable("", "intArray", reflect.ValueOf(intArray))
	jsRuntime.Variable("", "floatArray", reflect.ValueOf(floatArray))
	jsRuntime.Variable("", "stringArray", reflect.ValueOf(stringArray))
	jsRuntime.Variable("", "object", reflect.ValueOf(objectArray))

	err := jsRuntime.Run(ReadScript(t, "tests/slices.js"), nil)
	if err != nil {
		t.Error(err)
	}
}
