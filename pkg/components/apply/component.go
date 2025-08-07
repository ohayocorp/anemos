package apply

import (
	"fmt"
	"log/slog"
	"sort"
	"time"

	"github.com/ohayocorp/anemos/pkg/client"
	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
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

	if options.Timeout == 0 {
		options.Timeout, _ = time.ParseDuration("5m0s")
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
		documentGroups := context.GetDocumentGroups()

		sort.Slice(documentGroups, func(i, j int) bool {
			return documentGroups[i].Path < documentGroups[j].Path
		})

		for _, documentGroup := range documentGroups {
			path := documentGroup.Path
			if path == "" {
				path = "default"
			}

			path = core.ToKubernetesIdentifier(path)

			slog.Info("Applying document group: ${path}", slog.String("path", path))

			err := kubernetesClient.Apply(
				documentGroup.Documents,
				path,
				"",
				options.SkipConfirmation,
				options.Timeout)

			if err != nil {
				if _, ok := err.(client.NoChangesError); ok {
					slog.Info("No changes to apply for document group ${path}", slog.String("path", path))
				} else {
					js.Throw(fmt.Errorf("failed to apply document group '%s': %w", path, err))
				}
			}

			err = kubernetesClient.WaitDocuments(documentGroup.Documents, status.CurrentStatus, options.Timeout)
			if err != nil {
				js.Throw(fmt.Errorf("failed to wait for Kubernetes manifests: %w", err))
			}

			numberOfAppliedChanges++
		}

		if numberOfAppliedChanges == 0 {
			return
		}
	} else {
		err := kubernetesClient.Apply(
			documents,
			options.ApplySetParentName,
			options.ApplySetParentNamespace,
			options.SkipConfirmation,
			options.Timeout)

		if err != nil {
			if _, ok := err.(client.NoChangesError); ok {
				slog.Info("No changes to apply")
			} else {
				js.Throw(fmt.Errorf("failed to apply Kubernetes manifests: %w", err))
			}
		}

		err = kubernetesClient.WaitDocuments(documents, status.CurrentStatus, options.Timeout)
		if err != nil {
			js.Throw(fmt.Errorf("failed to wait for Kubernetes manifests: %w", err))
		}
	}

	slog.Info("Successfully applied Kubernetes manifests")
}
