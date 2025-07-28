package apply

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

func RegisterJsDeclarations(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Options]()).JsNamespace(
		"apply",
	).Fields(
		js.Field("Documents"),
		js.Field("ApplySetParentName"),
		js.Field("ApplySetParentNamespace"),
		js.Field("SkipConfirmation"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewOptions)),
		js.Constructor(reflect.ValueOf(NewOptionsWithDocumentGroup)),
		js.Constructor(reflect.ValueOf(NewOptionsWithDocumentGroupAndNamespace)),
		js.Constructor(reflect.ValueOf(NewOptionsWithDocumentsAndName)),
		js.Constructor(reflect.ValueOf(NewOptionsWithDocumentsNameAndNamespace)),
	)

	jsRuntime.Type(reflect.TypeFor[core.Builder]()).ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(Add)).JsName("apply"),
		js.ExtensionMethod(reflect.ValueOf(AddWithOptions)).JsName("apply"),
	)
}
