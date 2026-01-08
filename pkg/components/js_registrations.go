package components

import (
	"github.com/ohayocorp/anemos/pkg/components/apply"
	"github.com/ohayocorp/anemos/pkg/components/deleteoutputdirectory"
	"github.com/ohayocorp/anemos/pkg/components/reportdiagnostics"
	"github.com/ohayocorp/anemos/pkg/components/writedocuments"
	"github.com/ohayocorp/anemos/pkg/components/writereports"
	"github.com/ohayocorp/anemos/pkg/js"
)

func RegisterComponents(jsRuntime *js.JsRuntime) {
	apply.RegisterJsDeclarations(jsRuntime)
	deleteoutputdirectory.RegisterJsDeclarations(jsRuntime)
	reportdiagnostics.RegisterJsDeclarations(jsRuntime)
	writedocuments.RegisterJsDeclarations(jsRuntime)
	writereports.RegisterJsDeclarations(jsRuntime)
}
