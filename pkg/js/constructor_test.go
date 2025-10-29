package js_test

import (
	"reflect"
	"testing"

	"github.com/ohayocorp/anemos/pkg/cmd"
	"github.com/ohayocorp/anemos/pkg/js"
)

func TestConstructors(t *testing.T) {
	jsRuntime, err := cmd.InitializeNewRuntime(&cmd.AnemosProgram{
		RegisterRuntimeCallback: func(jsRuntime *js.JsRuntime) error {
			jsRuntime.Type(reflect.TypeFor[ConstructorTest]()).JsName(
				"Test",
			).Constructors(
				js.Constructor(reflect.ValueOf(EmptyConstructor)),
				js.Constructor(reflect.ValueOf(ConstructorPrimitives)),
			)

			return nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = jsRuntime.Run(ReadScript(t, "tests/constructors.js"), nil)
	if err != nil {
		t.Error(err)
	}
}

type ConstructorTest struct {
}

func EmptyConstructor() *ConstructorTest {
	return &ConstructorTest{}
}

func ConstructorPrimitives(a1 bool, a2 int, a3 float64, a4 string) *ConstructorTest {
	return &ConstructorTest{}
}
