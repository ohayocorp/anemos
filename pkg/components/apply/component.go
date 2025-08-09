package apply

import (
	"fmt"
	"log/slog"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/ohayocorp/anemos/pkg/client"
	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
	"helm.sh/helm/v3/pkg/releaseutil"
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

		provisioners := []*core.Provisioner{}
		for _, group := range documentGroups {
			if group.ApplyProvisioner != nil {
				provisioners = append(provisioners, group.ApplyProvisioner)
			}
			if group.WaitProvisioner != nil {
				provisioners = append(provisioners, group.WaitProvisioner)
			}
		}

		provisioners = getSortedProvisioners(provisioners)

		slog.Info("Provision plan:")
		for _, provisioner := range provisioners {
			slog.Info(
				"  ${type} -> ${applySetName}",
				slog.Any("type", strings.ToUpper(fmt.Sprintf("%-5s", provisioner.Type))),
				slog.String("applySetName", getApplySetName(provisioner.DocumentGroup)))
		}

		for _, provisioner := range provisioners {
			documentGroup := provisioner.DocumentGroup
			documents := getSortedDocuments(documentGroup)

			if len(documents) == 0 {
				slog.Info("No documents to apply in document group: ${path}", slog.String("path", documentGroup.Path))
				continue
			}

			if provisioner.Type == core.ProvisionerTypeApply {
				slog.Debug("Document apply order:")

				for _, document := range documents {
					slog.Info("  ${path}", slog.String("path", document.Path))
				}
			}

			switch provisioner.Type {
			case core.ProvisionerTypeApply:
				applySetName := getApplySetName(documentGroup)

				slog.Info("Applying document group: ${applySetName}", slog.String("applySetName", applySetName))

				err := kubernetesClient.Apply(
					documents,
					applySetName,
					"",
					options.SkipConfirmation,
					options.Timeout)

				if err != nil {
					if _, ok := err.(client.NoChangesError); ok {
						slog.Info("No changes to apply for document group ${applySetName}", slog.String("applySetName", applySetName))
					} else {
						js.Throw(fmt.Errorf("failed to apply document group '%s': %w", applySetName, err))
					}
				}

				slog.Info("Successfully applied document group: ${applySetName}", slog.String("applySetName", applySetName))

				numberOfAppliedChanges++
			case core.ProvisionerTypeWait:
				err = kubernetesClient.WaitDocuments(documents, status.CurrentStatus, options.Timeout)
				if err != nil {
					js.Throw(fmt.Errorf("failed to wait for Kubernetes manifests: %w", err))
				}
			}
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

func getSortedProvisioners(provisioners []*core.Provisioner) []*core.Provisioner {
	sort.SliceStable(provisioners, func(i, j int) bool {
		iId := fmt.Sprintf("%s-%s", provisioners[i].Type, provisioners[i].DocumentGroup.Path)
		jId := fmt.Sprintf("%s-%s", provisioners[j].Type, provisioners[j].DocumentGroup.Path)

		return iId < jId
	})

	d := core.DependencyGraph[*core.Provisioner]{
		Elements: provisioners,
		IdentifierGetter: func(p *core.Provisioner) string {
			return fmt.Sprintf("%s-%s", p.Type, p.DocumentGroup.Path)
		},
		DependenciesGetter: func(p *core.Provisioner) *core.Dependencies[*core.Provisioner] {
			return p.Dependencies
		},
	}

	return d.GetSortedElements()
}

func getSortedDocuments(documentGroup *core.DocumentGroup) []*core.Document {
	documents := make([]*core.Document, len(documentGroup.Documents))
	copy(documents, documentGroup.Documents)

	// First, sort the documents by their path. Subsequent sorts will preserve this order
	// for the documents that do not have a specific dependency.
	sort.SliceStable(documents, func(i, j int) bool {
		return documents[i].Path < documents[j].Path
	})

	// Next, sort the documents by their kind. Use the pre-defined order from Helm.
	sort.SliceStable(documents, func(i, j int) bool {
		iKind := documents[i].GetKind()
		jKind := documents[j].GetKind()

		if iKind == nil && jKind == nil {
			return true
		}

		if iKind == nil {
			return false
		}
		if jKind == nil {
			return true
		}

		iIndex := slices.Index(releaseutil.InstallOrder, *iKind)
		jIndex := slices.Index(releaseutil.InstallOrder, *jKind)

		return iIndex < jIndex
	})

	d := core.DependencyGraph[*core.Document]{
		Elements: documents,
		IdentifierGetter: func(d *core.Document) string {
			return d.Path
		},
		DependenciesGetter: func(d *core.Document) *core.Dependencies[*core.Document] {
			return d.Dependencies
		},
	}

	return d.GetSortedElements()
}

func getApplySetName(documentGroup *core.DocumentGroup) string {
	path := documentGroup.Path
	if path == "" {
		path = "default"
	}

	return core.ToKubernetesIdentifier(path)
}
