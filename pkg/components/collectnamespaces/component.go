package collectnamespaces

import (
	"fmt"
	"log/slog"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

type component struct {
	*core.Component
	options *Options

	documents *core.DocumentGroup
}

func NewComponent(options *Options) *core.Component {
	component := &component{
		Component: core.NewComponent(),
		options:   options,
	}

	component.AddAction(core.StepSanitize, component.sanitizeOptions)
	component.AddAction(core.NewStep("Collect namespaces", append(core.StepModify.Numbers, 1)...), component.modify)

	return component.Component
}

func (component *component) sanitizeOptions(context *core.BuildContext) {
	options := component.options

	if options == nil {
		options = &Options{}
		component.options = options
	}

	if options.Directory == "" {
		options.Directory = "namespaces"
	}

	component.SetIdentifier("collect-namespaces")
}

func (component *component) modify(context *core.BuildContext) {
	options := component.options

	namespaces := core.NewDocumentGroup(options.Directory)
	documentGroupsToRemove := []*core.DocumentGroup{}

	for _, documentGroup := range context.GetDocumentGroups() {
		documentsToMove := []*core.Document{}

		if documentGroup == nil {
			continue
		}

		for _, document := range documentGroup.Documents {
			if !document.IsNamespace() {
				continue
			}

			slog.Debug(
				"Moving document ${document} to ${to}",
				slog.String("document", document.FullPath()),
				slog.String("to", options.Directory))

			documentsToMove = append(documentsToMove, document)
		}

		if len(documentsToMove) == 0 {
			continue
		}

		for _, document := range documentsToMove {
			name := document.GetName()
			if name == nil {
				js.Throw(fmt.Errorf("failed to get name of Namespace: %s", document.FullPath()))
			}

			documentGroup.RemoveDocument(document)
			namespaces.AddDocument(document)

			document.Path = fmt.Sprintf("%s.yaml", *name)
		}

		if len(documentGroup.Documents) == 0 {
			documentGroupsToRemove = append(documentGroupsToRemove, documentGroup)
		}
	}

	for _, documentGroup := range documentGroupsToRemove {
		context.RemoveDocumentGroup(documentGroup)
	}

	if len(namespaces.Documents) > 0 {
		context.AddDocumentGroup(namespaces)
	}

	component.documents = namespaces
}
