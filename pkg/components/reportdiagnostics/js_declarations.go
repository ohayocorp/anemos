package reportdiagnostics

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

func RegisterJsDeclarations(jsRuntime *js.JsRuntime) {
	jsRuntime.Variable("reportDiagnostics", "componentType", reflect.ValueOf(componentType))

	jsRuntime.Type(reflect.TypeFor[Options]()).JsModule(
		"reportDiagnostics",
	).Constructors(
		js.Constructor(reflect.ValueOf(NewOptions)),
	)

	jsRuntime.Type(reflect.TypeFor[core.Builder]()).JsModule(
		"builder",
	).ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(Add)).JsName("reportDiagnostics"),
		js.ExtensionMethod(reflect.ValueOf(AddWithOptions)).JsName("reportDiagnostics"),
	)
}
