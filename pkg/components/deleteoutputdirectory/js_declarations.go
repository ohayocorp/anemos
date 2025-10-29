package deleteoutputdirectory

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

func RegisterJsDeclarations(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Options]()).JsModule(
		"deleteOutputDirectory",
	).Constructors(
		js.Constructor(reflect.ValueOf(NewOptions)),
	)

	jsRuntime.Type(reflect.TypeFor[core.Builder]()).JsModule(
		"builder",
	).ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(Add)).JsName("deleteOutputDirectory"),
		js.ExtensionMethod(reflect.ValueOf(AddWithOptions)).JsName("deleteOutputDirectory"),
	)
}
