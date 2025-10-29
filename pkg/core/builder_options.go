package core

import (
	"fmt"
	"reflect"

	"github.com/Masterminds/semver/v3"
	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg/js"
)

const (
	EnvironmentTypeUnknown     EnvironmentType = "unknown"
	EnvironmentTypeDevelopment EnvironmentType = "dev"
	EnvironmentTypeTesting     EnvironmentType = "test"
	EnvironmentTypeProduction  EnvironmentType = "prod"
)

const (
	KubernetesDistributionUnknown   KubernetesDistribution = "unknown"
	KubernetesDistributionAKS       KubernetesDistribution = "aks"
	KubernetesDistributionEKS       KubernetesDistribution = "eks"
	KubernetesDistributionGKE       KubernetesDistribution = "gke"
	KubernetesDistributionK3S       KubernetesDistribution = "k3s"
	KubernetesDistributionKubeadm   KubernetesDistribution = "kubeadm"
	KubernetesDistributionMicroK8s  KubernetesDistribution = "microk8s"
	KubernetesDistributionMinikube  KubernetesDistribution = "minikube"
	KubernetesDistributionOpenShift KubernetesDistribution = "openshift"
)

const (
	DocumentsDir = "manifests"
)

// EnvironmentType represents the type of the target environment such as Development or Production.
type EnvironmentType string

// Environment contains information about the target environment.
type Environment struct {
	Name string
	Type EnvironmentType
}

// KubernetesDistribution represents the distribution of the target Kubernetes cluster such as MicroK8s or OpenShift.
type KubernetesDistribution string

// KubernetesCluster contains information about the target Kubernetes cluster.
type KubernetesCluster struct {
	Version             *semver.Version
	Distribution        KubernetesDistribution
	AdditionalResources []*KubernetesResource
}

// OutputConfiguration specifies the output paths.
type OutputConfiguration struct {
	OutputPath string
}

// BuilderOptions contains common options and global services that are used by all components.
type BuilderOptions struct {
	KubernetesCluster   *KubernetesCluster
	Environment         *Environment
	OutputConfiguration *OutputConfiguration
}

func NewKubernetesCluster(version *semver.Version, distribution KubernetesDistribution) *KubernetesCluster {
	return &KubernetesCluster{
		Version:      version,
		Distribution: distribution,
	}
}

func NewKubernetesClusterWithAdditionalResources(
	version *semver.Version,
	distribution KubernetesDistribution,
	additionalResources []*KubernetesResource,
) *KubernetesCluster {
	return &KubernetesCluster{
		Version:             version,
		Distribution:        distribution,
		AdditionalResources: additionalResources,
	}
}

func NewEnvironment(name string, environmentType EnvironmentType) *Environment {
	return &Environment{
		Name: name,
		Type: environmentType,
	}
}

func NewOutputConfiguration() *OutputConfiguration {
	return &OutputConfiguration{}
}

func NewBuilderOptions(kubernetesCluster *KubernetesCluster, environment *Environment) *BuilderOptions {
	return NewBuilderOptionsWithOutputConfiguration(kubernetesCluster, environment, NewOutputConfiguration())
}

func NewBuilderOptionsWithOutputConfiguration(kubernetesCluster *KubernetesCluster, environment *Environment, outputConfiguration *OutputConfiguration) *BuilderOptions {
	return &BuilderOptions{
		KubernetesCluster:   kubernetesCluster,
		Environment:         environment,
		OutputConfiguration: outputConfiguration,
	}
}

func jsToVersion(jsRuntime *js.JsRuntime, jsValue sobek.Value) (*semver.Version, error) {
	value, err := jsRuntime.MarshalToGo(jsValue, reflect.TypeFor[string]())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JS value to string: %w", err)
	}

	return semver.NewVersion(value.Interface().(string))
}

func registerBuilderOptions(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Environment]()).JsModule(
		"builderOptions",
	).Fields(
		js.Field("Name"),
		js.Field("Type"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewEnvironment)),
	)

	jsRuntime.Type(reflect.TypeFor[KubernetesCluster]()).JsModule(
		"builderOptions",
	).Fields(
		js.Field("Version"),
		js.Field("Distribution"),
		js.Field("AdditionalResources"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewKubernetesCluster)),
		js.Constructor(reflect.ValueOf(NewKubernetesClusterWithAdditionalResources)),
	)

	jsRuntime.Type(reflect.TypeFor[OutputConfiguration]()).JsModule(
		"builderOptions",
	).Fields(
		js.Field("OutputPath"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewOutputConfiguration)),
	)

	jsRuntime.Type(reflect.TypeFor[BuilderOptions]()).JsModule(
		"builderOptions",
	).Fields(
		js.Field("KubernetesCluster"),
		js.Field("Environment"),
		js.Field("OutputConfiguration"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewBuilderOptions)),
		js.Constructor(reflect.ValueOf(NewBuilderOptionsWithOutputConfiguration)),
	)

	jsRuntime.Variable("environmentType", "unknown", reflect.ValueOf(EnvironmentTypeUnknown))
	jsRuntime.Variable("environmentType", "development", reflect.ValueOf(EnvironmentTypeDevelopment))
	jsRuntime.Variable("environmentType", "testing", reflect.ValueOf(EnvironmentTypeTesting))
	jsRuntime.Variable("environmentType", "production", reflect.ValueOf(EnvironmentTypeProduction))

	jsRuntime.Variable("kubernetesDistribution", "unknown", reflect.ValueOf(KubernetesDistributionUnknown))
	jsRuntime.Variable("kubernetesDistribution", "aks", reflect.ValueOf(KubernetesDistributionAKS))
	jsRuntime.Variable("kubernetesDistribution", "eks", reflect.ValueOf(KubernetesDistributionEKS))
	jsRuntime.Variable("kubernetesDistribution", "gke", reflect.ValueOf(KubernetesDistributionGKE))
	jsRuntime.Variable("kubernetesDistribution", "k3s", reflect.ValueOf(KubernetesDistributionK3S))
	jsRuntime.Variable("kubernetesDistribution", "kubeadm", reflect.ValueOf(KubernetesDistributionKubeadm))
	jsRuntime.Variable("kubernetesDistribution", "microk8s", reflect.ValueOf(KubernetesDistributionMicroK8s))
	jsRuntime.Variable("kubernetesDistribution", "minikube", reflect.ValueOf(KubernetesDistributionMinikube))
	jsRuntime.Variable("kubernetesDistribution", "openshift", reflect.ValueOf(KubernetesDistributionOpenShift))

	jsRuntime.Type(reflect.TypeFor[semver.Version]()).JsModule(
		"builderOptions",
	).Constructors(
		js.Constructor(reflect.ValueOf(semver.NewVersion)),
	).TypeConversion(reflect.ValueOf(jsToVersion))
}
