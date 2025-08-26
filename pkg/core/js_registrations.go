package core

import "github.com/ohayocorp/anemos/pkg/js"

func RegisterCore(jsRuntime *js.JsRuntime) {
	registerBuilder(jsRuntime)
	registerBuildContext(jsRuntime)
	registerBuilderOptions(jsRuntime)
	registerComponent(jsRuntime)
	registerDocument(jsRuntime)
	registerDocumentGroup(jsRuntime)
	registerFile(jsRuntime)
	registerHelm(jsRuntime)
	registerKubernetesResourceInfo(jsRuntime)
	registerProvisioner(jsRuntime)
	registerStep(jsRuntime)
	registerStringExtensions(jsRuntime)
	registerYamlParsing(jsRuntime)
	registerYamlSerialization(jsRuntime)
}
