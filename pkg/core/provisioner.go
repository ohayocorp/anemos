package core

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/js"
)

const (
	ProvisionerTypeApply ProvisionerType = "apply"
	ProvisionerTypeWait  ProvisionerType = "wait"
)

type ProvisionerType string

type Provisioner struct {
	Type          ProvisionerType
	DocumentGroup *DocumentGroup
	Dependencies  *Dependencies[*Provisioner]
}

func ApplyDocuments(documentGroup *DocumentGroup) *Provisioner {
	return &Provisioner{
		Type:          ProvisionerTypeApply,
		DocumentGroup: documentGroup,
		Dependencies:  NewDependencies[*Provisioner](),
	}
}

func WaitDocuments(documentGroup *DocumentGroup) *Provisioner {
	return &Provisioner{
		Type:          ProvisionerTypeWait,
		DocumentGroup: documentGroup,
		Dependencies:  NewDependencies[*Provisioner](),
	}
}

// Makes the given provisioner a prerequisite of this provisioner.
func (provisioner *Provisioner) RunAfter(p *Provisioner) {
	provisioner.Dependencies.RunAfter(p)
}

// Makes this provisioner a prerequisite of the given provisioner.
func (provisioner *Provisioner) RunBefore(p *Provisioner) {
	provisioner.Dependencies.RunBefore(p)
}

func registerProvisioner(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Provisioner]()).JsModule(
		"provisioner",
	).Fields(
		js.Field("Type"),
		js.Field("DocumentGroup"),
	).Methods(
		js.Method("RunAfter"),
		js.Method("RunBefore"),
	).DisableObjectMapping()
}
