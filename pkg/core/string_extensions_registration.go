package core

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/js"
	"github.com/ohayocorp/anemos/pkg/util"
)

func registerStringExtensions(jsRuntime *js.JsRuntime) {
	jsRuntime.Function(reflect.ValueOf(util.Indent))
	jsRuntime.Function(reflect.ValueOf(util.Dedent))
	jsRuntime.Function(reflect.ValueOf(util.MultilineString))
	jsRuntime.Function(reflect.ValueOf(util.ToKubernetesIdentifier))
}
