package createreferencednamespaces

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

func RegisterJsDeclarations(jsRuntime *js.JsRuntime) {
	jsRuntime.Variable("createReferencedNamespaces", "componentType", reflect.ValueOf(ComponentType))

	jsRuntime.Type(reflect.TypeFor[Options]()).JsNamespace(
		"createReferencedNamespaces",
	).Fields(
		js.Field("Predicate"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewOptions)),
		js.Constructor(reflect.ValueOf(NewOptionsWithPredicate)),
	)

	jsRuntime.Type(reflect.TypeFor[core.Builder]()).ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(Add)).JsName("createReferencedNamespaces"),
		js.ExtensionMethod(reflect.ValueOf(AddWithOptions)).JsName("createReferencedNamespaces"),
	)
}
