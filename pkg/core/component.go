package core

import (
	"fmt"
	"reflect"

	"github.com/ohayocorp/anemos/pkg/js"
)

const (
	MetadataIdentifier    = "identifier"
	MetadataComponentType = "component-type"
)

// Action is added to a [Component] to be run during the build process.
type Action struct {
	Step     *Step
	Callback func(context *BuildContext)
}

// Component is collection of actions that are executed in sequence.
type Component struct {
	// Actions are ordered by their steps and executed in sequence.
	Actions []*Action

	metadata   map[string]string
	customData map[string]any

	stackTrace string
}

func NewComponent() *Component {
	return &Component{
		metadata:   make(map[string]string),
		customData: make(map[string]any),
	}
}

// Adds given action to the list of actions.
func (component *Component) AddAction(step *Step, callback func(context *BuildContext)) {
	if step == nil {
		js.Throw(fmt.Errorf("step cannot be nil"))
	}

	if callback == nil {
		js.Throw(fmt.Errorf("callback cannot be nil"))
	}

	action := &Action{
		Step:     step,
		Callback: callback,
	}

	component.Actions = append(component.Actions, action)
}

func (component *Component) GetCustomData(key string) any {
	return component.customData[key]
}

func (component *Component) SetCustomData(key string, value any) {
	component.customData[key] = value
}

func (component *Component) GetMetadata(key string) *string {
	result, ok := component.metadata[key]
	if !ok {
		return nil
	}

	return &result
}

func (component *Component) SetMetadata(key, value string) {
	component.metadata[key] = value
}

func (component *Component) GetIdentifier() *string {
	return component.GetMetadata(MetadataIdentifier)
}

func (component *Component) SetIdentifier(identifier string) {
	component.SetMetadata(MetadataIdentifier, identifier)
}

func (component *Component) GetComponentType() *string {
	return component.GetMetadata(MetadataComponentType)
}

func (component *Component) SetComponentType(componentType string) {
	component.SetMetadata(MetadataComponentType, componentType)
}

func (component *Component) ProvisionAfter(provisioner *Provisioner) {
	component.AddAction(StepSpecifyProvisionerDependencies, func(context *BuildContext) {
		for _, documentGroup := range context.GetDocumentGroupsForComponent(component) {
			if documentGroup.ApplyProvisioner != nil {
				documentGroup.ApplyProvisioner.RunAfter(provisioner)
			}
			if documentGroup.WaitProvisioner != nil {
				documentGroup.WaitProvisioner.RunAfter(provisioner)
			}
		}
	})
}

func (component *Component) ProvisionBefore(provisioner *Provisioner) {
	component.AddAction(StepSpecifyProvisionerDependencies, func(context *BuildContext) {
		for _, documentGroup := range context.GetDocumentGroupsForComponent(component) {
			if documentGroup.ApplyProvisioner != nil {
				documentGroup.ApplyProvisioner.RunBefore(provisioner)
			}
			if documentGroup.WaitProvisioner != nil {
				documentGroup.WaitProvisioner.RunBefore(provisioner)
			}
		}
	})
}

func (component *Component) ProvisionAfterComponent(other *Component) {
	component.AddAction(StepSpecifyProvisionerDependencies, func(context *BuildContext) {
		for _, documentGroup := range context.GetDocumentGroupsForComponent(component) {
			for _, otherDocumentGroup := range context.GetDocumentGroupsForComponent(other) {
				documentGroup.ProvisionAfter(otherDocumentGroup)
			}
		}
	})
}

func (component *Component) ProvisionBeforeComponent(other *Component) {
	component.AddAction(StepSpecifyProvisionerDependencies, func(context *BuildContext) {
		for _, documentGroup := range context.GetDocumentGroupsForComponent(component) {
			for _, otherDocumentGroup := range context.GetDocumentGroupsForComponent(other) {
				documentGroup.ProvisionBefore(otherDocumentGroup)
			}
		}
	})
}

func registerComponent(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Component]()).JsModule(
		"component",
	).Fields(
		js.Field("Actions"),
	).Methods(
		js.Method("AddAction"),
		js.Method("GetCustomData"),
		js.Method("SetCustomData"),
		js.Method("GetMetadata"),
		js.Method("SetMetadata"),
		js.Method("GetIdentifier"),
		js.Method("SetIdentifier"),
		js.Method("GetComponentType"),
		js.Method("SetComponentType"),
		js.Method("ProvisionAfter"),
		js.Method("ProvisionBefore"),
		js.Method("ProvisionAfterComponent").JsName("provisionAfter"),
		js.Method("ProvisionBeforeComponent").JsName("provisionBefore"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewComponent)),
	)
}
