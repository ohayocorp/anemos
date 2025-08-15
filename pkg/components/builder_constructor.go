package components

import (
	"reflect"

	"github.com/Masterminds/semver/v3"
	"github.com/ohayocorp/anemos/pkg/components/apply"
	"github.com/ohayocorp/anemos/pkg/components/deleteoutputdirectory"
	"github.com/ohayocorp/anemos/pkg/components/setdefaultprovisionerdependencies"
	"github.com/ohayocorp/anemos/pkg/components/writedocuments"
	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

const (
	JsRuntimeMetadataBuilderApply            = "builder/apply"
	JsRuntimeMetadataBuilderSkipConfirmation = "builder/skipConfirmation"
)

func NewBuilderWithOptions(builderOptions *core.BuilderOptions, jsRuntime *js.JsRuntime) *core.Builder {
	builder := core.NewBuilderWithOptions(builderOptions, jsRuntime)

	deleteoutputdirectory.Add(builder)
	writedocuments.Add(builder)

	setdefaultprovisionerdependencies.Add(builder)

	if jsRuntime.Flags[JsRuntimeMetadataBuilderApply] == "true" {
		applyOptions := apply.NewOptions()
		if jsRuntime.Flags[JsRuntimeMetadataBuilderSkipConfirmation] == "true" {
			applyOptions.SkipConfirmation = true
		}

		apply.AddWithOptions(builder, applyOptions)
	}

	return builder
}

func NewBuilder(jsRuntime *js.JsRuntime) *core.Builder {
	return NewBuilderWithOptions(nil, jsRuntime)
}

func NewBuilderVersionDistributionEnvironmentType(
	version *semver.Version,
	distribution core.KubernetesDistribution,
	environment core.EnvironmentType,
	jsRuntime *js.JsRuntime,
) *core.Builder {
	options := core.NewBuilderOptions(
		core.NewKubernetesCluster(version, distribution),
		core.NewEnvironment(string(environment), environment),
	)

	return NewBuilderWithOptions(options, jsRuntime)
}

func registerBuilderConstructor(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[core.Builder]()).Constructors(
		js.Constructor(reflect.ValueOf(NewBuilder)),
		js.Constructor(reflect.ValueOf(NewBuilderWithOptions)),
		js.Constructor(reflect.ValueOf(NewBuilderVersionDistributionEnvironmentType)),
	)
}
