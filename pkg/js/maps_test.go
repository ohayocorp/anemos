package js_test

import (
	"reflect"
	"testing"

	"github.com/ohayocorp/anemos/pkg/js"
)

type MapElem struct {
	Property string
}

type MapObject struct {
	Bool   bool
	Int    int
	Float  float64
	String string
	Array  []*MapElem
}

func TestMaps(t *testing.T) {
	jsRuntime := js.NewJsRuntime()

	jsRuntime.Type(reflect.TypeFor[MapElem]()).Fields(
		js.Field("Property"),
	)

	jsRuntime.Type(reflect.TypeFor[MapObject]()).Fields(
		js.Field("Array"),
		js.Field("Bool"),
		js.Field("Int"),
		js.Field("Float"),
		js.Field("String"),
	)

	jsRuntime.Function(reflect.ValueOf(MapParam))

	object := &MapObject{
		Bool:   true,
		Int:    1,
		Float:  1.2,
		String: "a",
		Array: []*MapElem{{
			Property: "a",
		}, {
			Property: "b",
		}},
	}

	jsRuntime.Variable("", "object", reflect.ValueOf(object))

	err := jsRuntime.Run(ReadScript(t, "tests/maps.js"), nil)
	if err != nil {
		t.Error(err)
	}
}

func MapParam(object *MapObject) *MapObject {
	return object
}
