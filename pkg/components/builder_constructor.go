package components

import (
	"reflect"

	"github.com/Masterminds/semver/v3"
	"github.com/ohayocorp/anemos/pkg/components/collectcrds"
	"github.com/ohayocorp/anemos/pkg/components/collectnamespaces"
	"github.com/ohayocorp/anemos/pkg/components/deleteoutputdirectory"
	"github.com/ohayocorp/anemos/pkg/components/writedocuments"
	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

func NewBuilder(builderOptions *core.BuilderOptions, jsRuntime *js.JsRuntime) *core.Builder {
	builder := core.NewBuilder(builderOptions, jsRuntime)

	collectcrds.Add(builder)
	collectnamespaces.Add(builder)
	deleteoutputdirectory.Add(builder)
	writedocuments.Add(builder)

	return builder
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

	return NewBuilder(options, jsRuntime)
}

func registerBuilderConstructor(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[core.Builder]()).Constructors(
		js.Constructor(reflect.ValueOf(NewBuilder)),
		js.Constructor(reflect.ValueOf(NewBuilderVersionDistributionEnvironmentType)),
	)
}
