package core

import "github.com/ohayocorp/anemos/pkg/js"

func RegisterCore(jsRuntime *js.JsRuntime) {
	registerBuilder(jsRuntime)
	registerBuildContext(jsRuntime)
	registerBuilderOptions(jsRuntime)
	registerComponent(jsRuntime)
	registerDocumentGroup(jsRuntime)
	registerFile(jsRuntime)
	registerHelm(jsRuntime)
	registerKubernetesResourceInfo(jsRuntime)
	registerStep(jsRuntime)
	registerStringExtensions(jsRuntime)
	registerYamlDocument(jsRuntime)
	registerYamlDocumentExtensions(jsRuntime)
	registerYamlMapping(jsRuntime)
	registerYamlParsing(jsRuntime)
	registerYamlScalar(jsRuntime)
	registerYamlSequence(jsRuntime)
}
