package setdefaultprovisionerdependencies

import (
	"fmt"

	"github.com/ohayocorp/anemos/pkg/core"
)

const (
	ComponentType = "set-default-provisioner-dependencies"
)

type component struct {
	*core.Component
	options *Options
}

type resource struct {
	apiVersion      string
	kind            string
	document        *core.Document
	checkDependency func(resource *resource, document *core.Document) bool
}

func NewComponent(options *Options) *core.Component {
	component := &component{
		Component: core.NewComponent(),
		options:   options,
	}

	component.AddAction(core.StepSanitize, component.sanitizeOptions)
	component.AddAction(core.StepSpecifyProvisionerDependencies, component.specifyProvisionerDependencies)

	return component.Component
}

func (component *component) sanitizeOptions(context *core.BuildContext) {
	options := component.options

	if options == nil {
		options = &Options{}
		component.options = options
	}
}

func (component *component) specifyProvisionerDependencies(context *core.BuildContext) {
	// Find document groups that depend on other resources in other document groups such as
	// namespaces or CRDs and add provisioner dependencies accordingly.
	resources := []*resource{}

	for _, document := range context.GetAllDocuments() {
		apiVersion := document.GetApiVersion()
		kind := document.GetKind()

		if apiVersion == nil || kind == nil {
			continue
		}

		if *apiVersion == "v1" && *kind == "Namespace" {
			resources = append(resources, &resource{
				apiVersion:      *apiVersion,
				kind:            *kind,
				document:        document,
				checkDependency: checkNamespace,
			})
		} else if *apiVersion == "apiextensions.k8s.io/v1" && *kind == "CustomResourceDefinition" {
			resources = append(resources, &resource{
				apiVersion:      *apiVersion,
				kind:            *kind,
				document:        document,
				checkDependency: checkCustomResourceDefinition,
			})
		}
	}

	for _, document := range context.GetAllDocuments() {
		for _, resource := range resources {
			if resource.checkDependency(resource, document) {
				prerequisite := resource.document.Group
				dependent := document.Group

				if prerequisite != nil && dependent != nil && prerequisite != dependent {
					prerequisite.ProvisionBefore(dependent)
				}
			}
		}
	}
}

func checkNamespace(resource *resource, document *core.Document) bool {
	namespace := document.GetNamespace()
	if namespace == nil {
		return false
	}

	return *resource.document.GetName() == *namespace
}

func checkCustomResourceDefinition(resource *resource, document *core.Document) bool {
	apiVersion := document.GetApiVersion()
	kind := document.GetKind()

	if apiVersion == nil || kind == nil {
		return false
	}

	resourceKind := resource.document.GetRoot().GetValueChain("spec", "names", "kind")
	if resourceKind == nil || *resourceKind != *kind {
		return false
	}

	group := resource.document.GetRoot().GetValueChain("spec", "group")
	if group == nil {
		return false
	}

	versions := resource.document.GetRoot().GetSequenceChain("spec", "versions")
	for i := 0; i < versions.Length(); i++ {
		version := versions.GetMapping(i).GetValue("name")
		if version != nil && *apiVersion == fmt.Sprintf("%s/%s", *group, *version) {
			return true
		}
	}

	return false
}
