package setdefaultprovisionerdependencies

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

func RegisterJsDeclarations(jsRuntime *js.JsRuntime) {
	jsRuntime.Variable("setDefaultProvisionerDependencies", "componentType", reflect.ValueOf(ComponentType))

	jsRuntime.Type(reflect.TypeFor[Options]()).JsNamespace(
		"setDefaultProvisionerDependencies",
	).Constructors(
		js.Constructor(reflect.ValueOf(NewOptions)),
	)

	jsRuntime.Type(reflect.TypeFor[core.Builder]()).ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(Add)).JsName("setDefaultProvisionerDependencies"),
		js.ExtensionMethod(reflect.ValueOf(AddWithOptions)).JsName("setDefaultProvisionerDependencies"),
	)
}
