package writereports

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

func RegisterJsDeclarations(jsRuntime *js.JsRuntime) {
	jsRuntime.Variable("writeReports", "componentType", reflect.ValueOf(componentType))

	jsRuntime.Variable("writeReports", "Html", reflect.ValueOf(ReportOutputTypeHtml))
	jsRuntime.Variable("writeReports", "Markdown", reflect.ValueOf(ReportOutputTypeMarkdown))

	jsRuntime.Type(reflect.TypeFor[Options]()).JsModule(
		"writeReports",
	).Fields(
		js.Field("OutputTypes"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewOptions)),
		js.Constructor(reflect.ValueOf(NewOptionsWithOutputTypes)),
	)

	jsRuntime.Type(reflect.TypeFor[core.Builder]()).ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(Add)).JsName("writeReports"),
		js.ExtensionMethod(reflect.ValueOf(AddWithOptions)).JsName("writeReports"),
	)
}
