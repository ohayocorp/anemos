package collectnamespaces

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

func RegisterJsDeclarations(jsRuntime *js.JsRuntime) {
	jsRuntime.Variable("collectNamespaces", "componentType", reflect.ValueOf(componentType))

	jsRuntime.Type(reflect.TypeFor[Options]()).JsNamespace(
		"collectNamespaces",
	).Fields(
		js.Field("DocumentGroupPath"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewOptions)),
		js.Constructor(reflect.ValueOf(NewOptionsWithDocumentGroupPath)),
	)

	jsRuntime.Type(reflect.TypeFor[core.Builder]()).ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(Add)).JsName("collectNamespaces"),
		js.ExtensionMethod(reflect.ValueOf(AddWithOptions)).JsName("collectNamespaces"),
	)
}
