package components

import (
	"github.com/ohayocorp/anemos/pkg/components/apply"
	"github.com/ohayocorp/anemos/pkg/components/deleteoutputdirectory"
	"github.com/ohayocorp/anemos/pkg/components/writedocuments"
	"github.com/ohayocorp/anemos/pkg/js"
)

func RegisterComponents(jsRuntime *js.JsRuntime) {
	apply.RegisterJsDeclarations(jsRuntime)
	deleteoutputdirectory.RegisterJsDeclarations(jsRuntime)
	writedocuments.RegisterJsDeclarations(jsRuntime)
}
