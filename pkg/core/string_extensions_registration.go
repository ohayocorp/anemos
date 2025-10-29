package core

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/js"
	"github.com/ohayocorp/anemos/pkg/util"
)

func registerStringExtensions(jsRuntime *js.JsRuntime) {
	jsRuntime.Function(reflect.ValueOf(util.Indent)).JsModule("stringExtensions")
	jsRuntime.Function(reflect.ValueOf(util.Dedent)).JsModule("stringExtensions")
	jsRuntime.Function(reflect.ValueOf(util.MultilineString)).JsModule("stringExtensions")
	jsRuntime.Function(reflect.ValueOf(util.ToKubernetesIdentifier)).JsModule("stringExtensions")
}
