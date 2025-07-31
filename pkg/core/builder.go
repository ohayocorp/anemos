package core

import (
	"bytes"
	"fmt"
	"log/slog"
	"path/filepath"
	"reflect"
	"runtime/debug"
	"slices"
	"strings"

	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg/js"
)

// Builder is a collection of components.
type Builder struct {
	Components []*Component
	Options    *BuilderOptions

	jsRuntime *js.JsRuntime
}

// Creates a new [Builder] instance with given options.
func NewBuilder(options *BuilderOptions, jsRuntime *js.JsRuntime) *Builder {
	return &Builder{
		Options:   options,
		jsRuntime: jsRuntime,
	}
}

// Appends given component to the list of components.
func (builder *Builder) AddComponent(component *Component) {
	if builder.jsRuntime != nil {
		buffer := &bytes.Buffer{}

		frames := builder.jsRuntime.GetStackTrace()
		for _, frame := range frames {
			frame.Write(buffer)
			buffer.WriteByte('\n')
		}

		component.stackTrace = buffer.String()
	}

	builder.Components = append(builder.Components, component)
}

// Removes given component from the list of components.
func (builder *Builder) RemoveComponent(component *Component) {
	for i, c := range builder.Components {
		if c == component {
			builder.Components = slices.Delete(builder.Components, i, i+1)
			break
		}
	}
}

func (builder *Builder) AddDocument(document *Document) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddDocument(document)
	})
}

func (builder *Builder) AddDocumentWithGroupName(documentGroupName string, document *Document) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddDocumentWithGroupName(documentGroupName, document)
	})
}

func (builder *Builder) AddDocumentParse(path string, yamlContent string) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddDocumentParse(path, yamlContent)
	})
}

func (builder *Builder) AddDocumentParseWithGroupName(documentGroupName string, path string, yamlContent string) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddDocumentParseWithGroupName(documentGroupName, path, yamlContent)
	})
}

func (builder *Builder) AddDocumentMapping(path string, root *Mapping) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddDocumentMapping(path, root)
	})
}

func (builder *Builder) AddDocumentMappingWithGroupName(documentGroupName string, path string, root *Mapping) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddDocumentMappingWithGroupName(documentGroupName, path, root)
	})
}

func (builder *Builder) AddAdditionalFile(additionalFile *AdditionalFile) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddAdditionalFile(additionalFile)
	})
}

func (builder *Builder) AddAdditionalFileWithGroupName(documentGroupName string, additionalFile *AdditionalFile) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddAdditionalFileWithGroupName(documentGroupName, additionalFile)
	})
}

// Creates a new component with the given action and adds it to the list of components.
func (builder *Builder) OnStep(step *Step, callback func(context *BuildContext)) *Component {
	component := NewComponent()
	component.AddAction(step, callback)

	builder.AddComponent(component)
	return component
}

// Build method is at the heart of the all process. It collects all actions from all components
// and sorts them by their steps. Then it applies each action sequentially.
func (builder *Builder) Build() {
	slog.Info("Starting to build documents")

	context := BuildContext{
		BuilderOptions:         builder.Options,
		KubernetesResourceInfo: NewKubernetesResourceInfo(builder.Options.KubernetesCluster.Version),
		CustomData:             map[string]any{},
		JsRuntime:              builder.jsRuntime,
		builder:                builder,
		documentGroups:         map[*Component][]*DocumentGroup{},
	}

	for _, resource := range builder.Options.KubernetesCluster.AdditionalResources {
		context.KubernetesResourceInfo.AddKubernetesResource(resource)
	}

	builder.sanitizeBuilderOptions(context.BuilderOptions)

	steps := builder.getSteps()
	components := builder.Components

	defer func() {
		if r := recover(); r != nil {
			cleanupStackTrace := func(stackTrace string) string {
				lines := strings.Split(stackTrace, "\n")

				var builder strings.Builder
				for _, line := range lines {
					if strings.Contains(line, "initializeFunctions.func2 (native)") {
						continue
					}

					line = strings.TrimSpace(line)
					line = strings.TrimPrefix(line, "at")
					line = strings.TrimSpace(line)

					if line == "" {
						continue
					}

					builder.WriteString(fmt.Sprintf("\tat %s\n", line))
				}

				return builder.String()
			}

			var err error

			if jsErr, ok := r.(js.JsError); ok && context.currentComponent != nil {
				err = fmt.Errorf(
					"%s\n%s\n%s",
					strings.TrimPrefix(cleanupStackTrace(jsErr.Err.Error()), "\tat "),
					"Component registration stack trace:",
					cleanupStackTrace(context.currentComponent.stackTrace))
			} else if jsObj, ok := r.(*sobek.Object); ok && context.currentComponent != nil {
				err = fmt.Errorf(
					"%s\n%s\n%s",
					strings.TrimPrefix(cleanupStackTrace(jsObj.ToString().String()), "\tat "),
					"Component registration stack trace:",
					cleanupStackTrace(context.currentComponent.stackTrace))
			} else {
				// This is not an expected error, panic with the error so that the users can report it.
				// Using a JS exception here since a Golang panic will pollute the stack trace with Sobek
				// runtime internals and occasionally cause an invalid memory access error which hides the real error.
				err = fmt.Errorf("unexpected error: %v\n%s", r, string(debug.Stack()))
			}

			js.Throw(err)
		}
	}()

	var lastAppliedStep *Step = nil

	for i := 0; i < len(steps); i++ {
		step := steps[i]
		if lastAppliedStep != nil && step.Compare(*lastAppliedStep) < 0 {
			js.Throw(fmt.Errorf(
				"cannot add an action that will be run before the step it is added in, last applied step: %s",
				lastAppliedStep.String()))
		}

		slog.Info(
			"Applying actions for step: '${step}' - ${description}",
			slog.String("description", step.Description),
			slog.String("step", step.String()))

		for _, component := range components {
			context.currentComponent = component

			for _, action := range component.Actions {
				if !action.Step.Equals(step) {
					continue
				}

				if action.Callback != nil {
					action.Callback(&context)
				}
			}
		}

		lastAppliedStep = &step
		// Components may have added new actions, so we need to recompute the steps.
		steps = builder.getSteps()
	}
}

func (builder *Builder) getSteps() []Step {
	stepsMap := map[string]Step{}

	for _, component := range builder.Components {
		for _, action := range component.Actions {
			stepsMap[action.Step.String()] = *action.Step
		}
	}

	steps := []Step{}
	for _, step := range stepsMap {
		steps = append(steps, step)
	}

	slices.SortStableFunc(steps, func(a, b Step) int {
		return a.Compare(b)
	})

	return steps
}

func (builder *Builder) sanitizeBuilderOptions(options *BuilderOptions) {
	if options.KubernetesCluster == nil {
		js.Throw(fmt.Errorf("KubernetesCluster is not set in builder options"))
	}

	if options.Environment == nil {
		js.Throw(fmt.Errorf("Environment is not set in builder options"))
	}

	outputConfiguration := options.OutputConfiguration
	if outputConfiguration == nil {
		outputConfiguration = &OutputConfiguration{}
		options.OutputConfiguration = outputConfiguration
	}

	if outputConfiguration.OutputPath == "" {
		// If output path is not set, use the main script directory as the output path.
		outputPath := filepath.Dir(builder.jsRuntime.MainScriptPath)

		outputPathEnv := builder.jsRuntime.GetEnv("ANEMOS_OUTPUT_PATH")
		if outputPathEnv != nil && *outputPathEnv != "" {
			// If ANEMOS_OUTPUT_PATH is set, use it as the output path.
			// If it is relative, join it with the main script directory.
			// If it is absolute, use it as is.

			outputPathEnvCleaned := filepath.Clean(*outputPathEnv)
			if filepath.IsAbs(outputPathEnvCleaned) {
				outputPath = outputPathEnvCleaned
			} else {
				outputPath = filepath.Join(outputPath, outputPathEnvCleaned)
			}
		} else {
			// If ANEMOS_OUTPUT_PATH is not set, use the output directory inside the main script directory.
			outputPath = filepath.Join(outputPath, "output")
		}

		outputConfiguration.OutputPath = outputPath
	}
}

func registerBuilder(jsRuntime *js.JsRuntime) {
	// Don't register constructor here as we want to add default components and referencing
	// them here will cause circular dependency issues.
	// Instead, we will register the constructor in the components/builder_constructor.go file.
	jsRuntime.Type(reflect.TypeFor[Builder]()).Fields(
		js.Field("Components"),
		js.Field("Options"),
	).Methods(
		js.Method("AddComponent"),
		js.Method("RemoveComponent"),
		js.Method("AddDocument"),
		js.Method("AddDocumentWithGroupName").JsName("addDocument"),
		js.Method("AddDocumentParse").JsName("addDocument"),
		js.Method("AddDocumentParseWithGroupName").JsName("addDocument"),
		js.Method("AddDocumentMapping").JsName("addDocument"),
		js.Method("AddDocumentMappingWithGroupName").JsName("addDocument"),
		js.Method("AddAdditionalFile"),
		js.Method("AddAdditionalFileWithGroupName").JsName("addAdditionalFile"),
		js.Method("OnStep"),
		js.Method("Build"),
	)
}
