package core

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/ohayocorp/anemos/pkg/js"
)

var (
	// Use this step to populate [KubernetesResource] resources so that other components can rely on this
	// information to modify existing resources or generate extra resources.
	// E.g. when ServiceMonitor is added on this step, other components can generate ServiceMonitor resources
	// to monitor the services.
	StepPopulateKubernetesResources = NewStep("Populate Kubernetes resources", 1)

	// Sanitize the options and the component properties in this step.
	StepSanitize = NewStep("Sanitize", 2)

	// Use this step to generate the documents and additional files.
	StepGenerateResources = NewStep("Generate resources", 5)

	// Use this step to generate the documents and additional files based on other documents and additional files.
	StepGenerateResourcesBasedOnOtherResources = NewStep("Generate based on other resources", 5, 1)

	// Use this step to modify the generated documents, e.g. set labels, annotations, etc. This is useful when
	// other services add documents to the document group on the [StepGenerateResources] step and you want to modify them.
	StepModify = NewStep("Modify", 6)

	// Specify provisioner dependencies in this step.
	StepSpecifyProvisionerDependencies = NewStep("Specify provisioner dependencies", 7)

	// Write the outputs, e.g. documents and additional files in this step.
	StepOutput = NewStep("Output", 99)

	// Apply the resources to the Kubernetes cluster in this step.
	StepApply = NewStep("Apply", 100)
)

type Step struct {
	Description string
	Numbers     []int
}

func NewStep(description string, numbers ...int) *Step {
	if len(numbers) == 0 {
		js.Throw(fmt.Errorf("step must have at least one number"))
	}

	return &Step{
		Description: description,
		Numbers:     numbers,
	}
}

func (s *Step) String() string {
	result := ""

	for i, n := range s.Numbers {
		if i > 0 {
			result += ","
		}

		result += strconv.Itoa(n)
	}

	return result
}

func (s *Step) Equals(other Step) bool {
	return s.Compare(other) == 0
}

func (s *Step) Compare(other Step) int {
	maxCount := len(s.Numbers)
	if len(other.Numbers) > maxCount {
		maxCount = len(other.Numbers)
	}

	for i := 0; i < maxCount; i++ {
		var a, b int
		if i < len(s.Numbers) {
			a = s.Numbers[i]
		}
		if i < len(other.Numbers) {
			b = other.Numbers[i]
		}

		if a < b {
			return -1
		} else if a > b {
			return 1
		}
	}

	return 0
}

func registerStep(jsRuntime *js.JsRuntime) {
	jsRuntime.Variable("steps", "populateKubernetesResources", reflect.ValueOf(StepPopulateKubernetesResources))
	jsRuntime.Variable("steps", "sanitize", reflect.ValueOf(StepSanitize))
	jsRuntime.Variable("steps", "generateResources", reflect.ValueOf(StepGenerateResources))
	jsRuntime.Variable("steps", "generateResourcesBasedOnOtherResources", reflect.ValueOf(StepGenerateResourcesBasedOnOtherResources))
	jsRuntime.Variable("steps", "modify", reflect.ValueOf(StepModify))
	jsRuntime.Variable("steps", "specifyProvisionerDependencies", reflect.ValueOf(StepSpecifyProvisionerDependencies))
	jsRuntime.Variable("steps", "output", reflect.ValueOf(StepOutput))

	jsRuntime.Type(reflect.TypeFor[Step]()).Constructors(
		js.Constructor(reflect.ValueOf(NewStep)),
	).Fields(
		js.Field("Description"),
		js.Field("Numbers"),
	).Methods(
		js.Method("String").JsName("toString"),
		js.Method("Equals"),
		js.Method("Compare").JsName("compareTo"),
	)
}
