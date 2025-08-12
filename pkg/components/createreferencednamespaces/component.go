package createreferencednamespaces

import (
	"fmt"
	"log/slog"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ohayocorp/anemos/pkg/core"
)

const (
	DocumentGroupPath = "namespaces"
	ComponentType     = "create-referenced-namespaces"
)

var existingNamespaces mapset.Set[string] = mapset.NewSet(
	"kube-system",
	"kube-public",
	"kube-node-lease",
	"default",
)

type component struct {
	*core.Component
	options *Options
}

func NewComponent(options *Options) *core.Component {
	component := &component{
		Component: core.NewComponent(),
		options:   options,
	}

	component.AddAction(core.StepSanitize, component.sanitizeOptions)
	component.AddAction(core.StepGenerateResourcesBasedOnOtherResources, component.generate)

	return component.Component
}

func (component *component) sanitizeOptions(context *core.BuildContext) {
	options := component.options

	if options == nil {
		options = &Options{}
		component.options = options
	}
}

func (component *component) generate(context *core.BuildContext) {
	options := component.options
	namespaces := mapset.NewSet[string]()

	for _, documentGroup := range context.GetDocumentGroups() {
		for _, document := range documentGroup.Documents {
			namespace := document.GetNamespace()
			if namespace == nil {
				continue
			}

			if existingNamespaces.Contains(*namespace) {
				continue
			}

			if options.Predicate != nil && !options.Predicate(*namespace) {
				continue
			}

			namespaces.Add(*namespace)
		}
	}

	documentGroup := core.NewDocumentGroup(DocumentGroupPath)

	for namespace := range namespaces.Iter() {
		slog.Info("Adding namespace resource ${namespace}", slog.String("namespace", namespace))

		template := `
			apiVersion: v1
			kind: Namespace
			metadata:
			  name: {{ . }}`

		document := core.ParseTemplateAsDocument(fmt.Sprintf("%s.yaml", namespace), template, namespace)
		documentGroup.AddDocument(document)
	}

	if len(documentGroup.Documents) > 0 {
		context.AddDocumentGroup(documentGroup)
	}
}
