package core

import "github.com/ohayocorp/anemos/pkg/js"

func RegisterCore(jsRuntime *js.JsRuntime) {
	registerBuilder(jsRuntime)
	registerBuildContext(jsRuntime)
	registerBuilderOptions(jsRuntime)
	registerComponent(jsRuntime)
	registerDiagnostic(jsRuntime)
	registerDocument(jsRuntime)
	registerDocumentGroup(jsRuntime)
	registerFile(jsRuntime)
	registerHelm(jsRuntime)
	registerKubernetesResourceInfo(jsRuntime)
	registerProvisioner(jsRuntime)
	registerQuantity(jsRuntime)
	registerReport(jsRuntime)
	registerStep(jsRuntime)
	registerStringExtensions(jsRuntime)
	registerYamlParsing(jsRuntime)
	registerYamlSerialization(jsRuntime)
}
