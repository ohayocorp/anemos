package components

import (
	"github.com/ohayocorp/anemos/pkg/components/apply"
	"github.com/ohayocorp/anemos/pkg/components/collectcrds"
	"github.com/ohayocorp/anemos/pkg/components/collectnamespaces"
	"github.com/ohayocorp/anemos/pkg/components/createreferencednamespaces"
	"github.com/ohayocorp/anemos/pkg/components/deleteoutputdirectory"
	"github.com/ohayocorp/anemos/pkg/components/writedocuments"
	"github.com/ohayocorp/anemos/pkg/js"
)

func RegisterComponents(jsRuntime *js.JsRuntime) {
	registerBuilderConstructor(jsRuntime)

	apply.RegisterJsDeclarations(jsRuntime)
	collectcrds.RegisterJsDeclarations(jsRuntime)
	collectnamespaces.RegisterJsDeclarations(jsRuntime)
	createreferencednamespaces.RegisterJsDeclarations(jsRuntime)
	deleteoutputdirectory.RegisterJsDeclarations(jsRuntime)
	writedocuments.RegisterJsDeclarations(jsRuntime)
}
