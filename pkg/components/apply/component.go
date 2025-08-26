package apply

import (
	"fmt"
	"log/slog"
	"regexp"
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

	component.SetComponentType("apply")
	component.SetIdentifier("apply")

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
}

func (component *component) apply(context *core.BuildContext) {
	options := component.options

	kubernetesClient, err := client.NewKubernetesClient()
	if err != nil {
		js.Throw(fmt.Errorf("failed to create Kubernetes client: %w", err))
	}

	regexList, err := component.getRegexList()
	if err != nil {
		js.Throw(fmt.Errorf("failed to get regex list: %w", err))
	}

	numberOfAppliedChanges := 0
	documentGroups := context.GetDocumentGroups()

	provisioners := []*core.Provisioner{}
	for _, group := range documentGroups {
		if !component.shouldApplyDocumentGroup(regexList, group) {
			continue
		}

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

		documentYamls := make([]string, 0, len(documents))
		for _, document := range documents {
			yaml, err := core.SerializeSobekObjectToYaml(context.JsRuntime, document.Object)
			if err != nil {
				js.Throw(fmt.Errorf("failed to serialize document to YAML: %w", err))
			}

			documentYamls = append(documentYamls, yaml)
		}

		switch provisioner.Type {
		case core.ProvisionerTypeApply:
			applySetName := getApplySetName(documentGroup)

			slog.Info("")
			slog.Info("Applying document group: ${applySetName}", slog.String("applySetName", applySetName))

			slog.Debug("Document apply order:")

			for _, document := range documents {
				slog.Debug("  ${path}", slog.String("path", document.GetPath()))
			}

			err := kubernetesClient.Apply(
				documentYamls,
				applySetName,
				"",
				options.SkipConfirmation,
				options.ForceConflicts,
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
			err = kubernetesClient.WaitDocuments(documentYamls, status.CurrentStatus, options.Timeout)
			if err != nil {
				js.Throw(fmt.Errorf("failed to wait for Kubernetes manifests: %w", err))
			}
		}
	}

	if numberOfAppliedChanges == 0 {
		return
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
			d := core.NewDependencies[*core.Provisioner]()

			for _, p := range p.Dependencies.Dependents {
				if slices.Contains(provisioners, p) {
					d.Dependents = append(d.Dependents, p)
				}
			}

			for _, p := range p.Dependencies.Prerequisites {
				if slices.Contains(provisioners, p) {
					d.Prerequisites = append(d.Prerequisites, p)
				}
			}

			return d
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
		return documents[i].GetPath() < documents[j].GetPath()
	})

	// Next, sort the documents by their kind. Use the pre-defined order from Helm.
	sort.SliceStable(documents, func(i, j int) bool {
		iKind := core.SobekObjectGetString(documents[i].Object, "kind")
		jKind := core.SobekObjectGetString(documents[j].Object, "kind")

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
			return d.GetPath()
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

func (component *component) getRegexList() ([]*regexp.Regexp, error) {
	patterns := component.options.DocumentGroups
	if len(patterns) == 0 {
		return nil, nil
	}

	var regexList []*regexp.Regexp
	for _, pattern := range patterns {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("failed to compile regex pattern '%s': %w", pattern, err)
		}
		regexList = append(regexList, re)
	}
	return regexList, nil
}

func (component *component) shouldApplyDocumentGroup(regexList []*regexp.Regexp, documentGroup *core.DocumentGroup) bool {
	if regexList == nil {
		return true
	}

	// Check if the document group is applicable for this component.
	for _, re := range regexList {
		if re.MatchString(documentGroup.Path) {
			return true
		}
	}

	return false
}
