package apply

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

func RegisterJsDeclarations(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Options]()).JsModule(
		"apply",
	).Fields(
		js.Field("DocumentGroups"),
		js.Field("SkipConfirmation"),
		js.Field("ForceConflicts"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewOptions)),
	)

	jsRuntime.Type(reflect.TypeFor[core.Builder]()).JsModule(
		"builder",
	).ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(Add)).JsName("apply"),
		js.ExtensionMethod(reflect.ValueOf(AddWithOptions)).JsName("apply"),
	)
}
