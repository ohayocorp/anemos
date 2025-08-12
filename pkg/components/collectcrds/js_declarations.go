package collectcrds

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

func RegisterJsDeclarations(jsRuntime *js.JsRuntime) {
	jsRuntime.Variable("collectCRDs", "componentType", reflect.ValueOf(componentType))

	jsRuntime.Type(reflect.TypeFor[Options]()).JsNamespace(
		"collectCRDs",
	).Fields(
		js.Field("DocumentGroupPath"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewOptions)),
		js.Constructor(reflect.ValueOf(NewOptionsWithDocumentGroupPath)),
	)

	jsRuntime.Type(reflect.TypeFor[core.Builder]()).ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(Add)).JsName("collectCRDs"),
		js.ExtensionMethod(reflect.ValueOf(AddWithOptions)).JsName("collectCRDs"),
	)
}
