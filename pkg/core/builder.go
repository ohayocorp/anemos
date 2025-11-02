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

	"github.com/Masterminds/semver/v3"
	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg/js"
)

const (
	JsRuntimeMetadataBuilderApply            = "builder/apply"
	JsRuntimeMetadataBuilderSkipConfirmation = "builder/skipConfirmation"
)

// Builder is a collection of components.
type Builder struct {
	Components []*Component
	Options    *BuilderOptions

	jsRuntime *js.JsRuntime
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

// Adds a component that creates a document group with the given name during [StepGenerateResources].
// Document group doesn't contain any documents, it serves as a placeholder for provision dependencies.
func (builder *Builder) AddProvisionCheckpoint(name string) *Component {
	component := NewComponent()

	component.SetIdentifier(name)
	component.AddAction(StepGenerateResources, func(context *BuildContext) {
		documentGroup := NewDocumentGroup(name)
		context.AddDocumentGroup(documentGroup)
	})

	builder.AddComponent(component)

	return component
}

func (builder *Builder) AddDocument(document *Document) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddDocument(document)
	})
}

func (builder *Builder) AddDocumentString(jsRuntime *js.JsRuntime, yaml string) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddDocumentWithOptions(jsRuntime, &NewDocumentOptions{
			Yaml: &yaml,
		})
	})
}

func (builder *Builder) AddDocumentWithOptions(jsRuntime *js.JsRuntime, options *NewDocumentOptions) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddDocumentWithOptions(jsRuntime, options)
	})
}

func (builder *Builder) AddAdditionalFile(additionalFile *AdditionalFile) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddAdditionalFile(additionalFile)
	})
}

func (builder *Builder) AddAdditionalFileWithGroupPath(documentGroupPath string, additionalFile *AdditionalFile) {
	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		context.AddAdditionalFileWithGroupPath(documentGroupPath, additionalFile)
	})
}

// Creates a new component with the given action and adds it to the list of components.
func (builder *Builder) OnStep(step *Step, callback func(context *BuildContext)) *Component {
	component := NewComponent()
	component.AddAction(step, callback)

	builder.AddComponent(component)
	return component
}

// Creates a new component with the given action that will be run during [StepPopulateKubernetesResources] and adds it to the list of components.
func (builder *Builder) OnPopulateKubernetesResources(callback func(context *BuildContext)) *Component {
	return builder.OnStep(StepPopulateKubernetesResources, callback)
}

// Creates a new component with the given action that will be run during [StepSanitize] and adds it to the list of components.
func (builder *Builder) OnSanitize(callback func(context *BuildContext)) *Component {
	return builder.OnStep(StepSanitize, callback)
}

// Creates a new component with the given action that will be run during [StepGenerateResources] and adds it to the list of components.
func (builder *Builder) OnGenerateResources(callback func(context *BuildContext)) *Component {
	return builder.OnStep(StepGenerateResources, callback)
}

// Creates a new component with the given action that will be run during [StepGenerateResourcesBasedOnOtherResources] and adds it to the list of components.
func (builder *Builder) OnGenerateResourcesBasedOnOtherResources(callback func(context *BuildContext)) *Component {
	return builder.OnStep(StepGenerateResourcesBasedOnOtherResources, callback)
}

// Creates a new component with the given action that will be run during [StepModify] and adds it to the list of components.
func (builder *Builder) OnModify(callback func(context *BuildContext)) *Component {
	return builder.OnStep(StepModify, callback)
}

// Creates a new component with the given action that will be run during [StepSpecifyProvisionerDependencies] and adds it to the list of components.
func (builder *Builder) OnSpecifyProvisionerDependencies(callback func(context *BuildContext)) *Component {
	return builder.OnStep(StepSpecifyProvisionerDependencies, callback)
}

// Build method is at the heart of the all process. It collects all actions from all components
// and sorts them by their steps. Then it applies each action sequentially.
func (builder *Builder) Build() {
	slog.Info("Starting to build documents")

	builder.sanitizeBuilderOptions(builder.Options)

	context := NewBuildContext(builder, builder.Options)

	for _, resource := range builder.Options.KubernetesCluster.AdditionalResources {
		context.KubernetesResourceInfo.AddKubernetesResource(resource)
	}

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
					action.Callback(context)
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
		options.KubernetesCluster = NewKubernetesCluster(DefaultKubernetesVersion, KubernetesDistributionUnknown)
	}

	if options.KubernetesCluster.Version == nil {
		slog.Debug(
			"Using Kubernetes ${defaultVersion} resources as base since version is not specified.",
			slog.String("defaultVersion", DefaultKubernetesVersion.String()),
		)

		options.KubernetesCluster.Version = DefaultKubernetesVersion
	}

	if options.Environment == nil {
		options.Environment = NewEnvironment("none", EnvironmentTypeUnknown)
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

// Creates a new [Builder] instance with default options.
func NewBuilder(jsRuntime *js.JsRuntime) *Builder {
	return NewBuilderWithOptions(nil, jsRuntime)
}

// Creates a new [Builder] instance with given options.
func NewBuilderWithOptions(options *BuilderOptions, jsRuntime *js.JsRuntime) *Builder {
	if options == nil {
		options = &BuilderOptions{}
	}

	builder := &Builder{
		Options:   options,
		jsRuntime: jsRuntime,
	}

	builder.sanitizeBuilderOptions(builder.Options)

	builderJs, err := jsRuntime.MarshalToJs(reflect.ValueOf(builder))
	if err != nil {
		panic(fmt.Errorf("failed to marshal builder to JS: %w", err))
	}

	runtime := jsRuntime.Runtime
	runtime.Set("__anemos__builder", builderJs)
	runtime.Set("__anemos__flags__apply", jsRuntime.Flags[JsRuntimeMetadataBuilderApply] == "true")
	runtime.Set("__anemos__flags__skipConfirmation", jsRuntime.Flags[JsRuntimeMetadataBuilderSkipConfirmation] == "true")

	_, err = runtime.RunScript("builderDefaults.js", `
		__anemos__require = require("@ohayocorp/anemos");

		// Register default native components.
		__anemos__builder.deleteOutputDirectory();
		__anemos__builder.writeDocuments();

		// Register default components from the JavaScript libraries.
		__anemos__require.sortFields.add(__anemos__builder);
		__anemos__require.setDefaultProvisionerDependencies.add(__anemos__builder);

		if (__anemos__flags__apply) {
			const applyOptions = {
				skipConfirmation: __anemos__flags__skipConfirmation
			};

			__anemos__builder.apply(applyOptions);
		}

		delete __anemos__builder;
		delete __anemos__require;
		`)

	if err != nil {
		panic(fmt.Errorf("failed to initialize builder defaults: %w", err))
	}

	return builder
}

func NewBuilderVersionDistributionEnvironmentType(version *semver.Version, distribution KubernetesDistribution, environment EnvironmentType, jsRuntime *js.JsRuntime) *Builder {
	options := NewBuilderOptions(
		NewKubernetesCluster(version, distribution),
		NewEnvironment(string(environment), environment),
	)

	return NewBuilderWithOptions(options, jsRuntime)
}

func registerBuilder(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Builder]()).JsModule(
		"builder",
	).Fields(
		js.Field("Components"),
		js.Field("Options"),
	).Methods(
		js.Method("AddComponent"),
		js.Method("RemoveComponent"),
		js.Method("AddProvisionCheckpoint"),
		js.Method("AddDocument"),
		js.Method("AddDocumentString").JsName("addDocument"),
		js.Method("AddDocumentWithOptions").JsName("addDocument"),
		js.Method("AddAdditionalFile"),
		js.Method("AddAdditionalFileWithGroupPath").JsName("addAdditionalFile"),
		js.Method("OnStep"),
		js.Method("OnPopulateKubernetesResources"),
		js.Method("OnSanitize"),
		js.Method("OnGenerateResources"),
		js.Method("OnGenerateResourcesBasedOnOtherResources"),
		js.Method("OnModify"),
		js.Method("OnSpecifyProvisionerDependencies"),
		js.Method("Build"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewBuilder)),
		js.Constructor(reflect.ValueOf(NewBuilderWithOptions)),
		js.Constructor(reflect.ValueOf(NewBuilderVersionDistributionEnvironmentType)),
	)
}
