package apply

import (
	"fmt"
	"log/slog"

	"github.com/ohayocorp/anemos/pkg/client"
	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
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
	component.AddAction(core.StepApply, component.apply)

	return component.Component
}

func (component *component) sanitizeOptions(context *core.BuildContext) {
	options := component.options

	if options == nil {
		options = &Options{}
		component.options = options
	}

	component.SetIdentifier(fmt.Sprintf("apply-%s-%s", options.ApplySetParentNamespace, options.ApplySetParentName))
}

func (component *component) apply(context *core.BuildContext) {
	options := component.options
	documents := options.Documents

	kubernetesClient, err := client.NewKubernetesClient()
	if err != nil {
		js.Throw(fmt.Errorf("failed to create Kubernetes client: %w", err))
	}

	// If no documents are provided, apply all document groups from the context.
	if len(documents) == 0 {
		numberOfAppliedChanges := 0

		for _, documentGroup := range context.GetDocumentGroups() {
			err = kubernetesClient.Apply(documentGroup.Documents, documentGroup.Name, "", options.SkipConfirmation)
			if err != nil {
				if _, ok := err.(client.NoChangesError); ok {
					slog.Info("No changes to apply for document group ${name}", slog.String("name", documentGroup.Name))
					continue
				}

				js.Throw(fmt.Errorf("failed to apply document group '%s': %w", documentGroup.Name, err))
			}

			numberOfAppliedChanges++
		}

		if numberOfAppliedChanges == 0 {
			return
		}
	} else {
		err = kubernetesClient.Apply(documents, options.ApplySetParentName, options.ApplySetParentNamespace, options.SkipConfirmation)
		if err != nil {
			if _, ok := err.(client.NoChangesError); ok {
				slog.Info("No changes to apply")
				return
			}

			js.Throw(fmt.Errorf("failed to apply Kubernetes manifests: %w", err))
		}
	}

	slog.Info("Successfully applied Kubernetes manifests")
}
