package core

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"path"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg/js"
	"github.com/ohayocorp/anemos/pkg/util"
	"gopkg.in/yaml.v3"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/releaseutil"
)

// Options to create documents using Helm.
type HelmOptions struct {
	ReleaseName string
	Namespace   string
	Values      string
}

func NewHelmOptions(releaseName string, namespace string) *HelmOptions {
	return &HelmOptions{
		ReleaseName: releaseName,
		Namespace:   namespace,
	}
}

func NewHelmOptionsWithValues(releaseName string, namespace string, values string) *HelmOptions {
	return &HelmOptions{
		ReleaseName: releaseName,
		Namespace:   namespace,
		Values:      values,
	}
}

// Sanitizes Helm options and throws on invalid configuration.
func (options *HelmOptions) sanitize() {
	if options.ReleaseName == "" {
		js.Throw(fmt.Errorf("release name is not defined on helm options, %#+v", options))
	}
}

// Loads the given chart file into memory.
func LoadChart(data []byte) *chart.Chart {
	chart, err := loader.LoadArchive(bytes.NewReader(data))
	if err != nil {
		js.Throw(err)
	}

	return chart
}

// Loads the chart from the given path. The path can be a local file or directory.
func LoadChartFromPath(path string) *chart.Chart {
	chart, err := loader.Load(path)
	if err != nil {
		js.Throw(fmt.Errorf("can't load chart from path %s, %v", path, err))
	}

	return chart
}

func AddHelmChart(builder *Builder, chartIdentifier string, releaseName string, values string) {
	if chartIdentifier == "" {
		js.Throw(fmt.Errorf("chart identifier is not defined"))
	}

	slog.Info(
		"Adding Helm chart: ${chart}, release name: ${releaseName}",
		slog.String("chart", chartIdentifier),
		slog.String("releaseName", releaseName))

	builder.OnStep(StepGenerateResources, func(context *BuildContext) {
		var chart *chart.Chart

		if strings.HasPrefix(chartIdentifier, "http://") || strings.HasPrefix(chartIdentifier, "https://") {
			response, err := http.Get(chartIdentifier)
			if err != nil {
				js.Throw(fmt.Errorf("can't load chart from URL %s, %v", chartIdentifier, err))
			}
			defer response.Body.Close()

			if response.StatusCode != http.StatusOK {
				js.Throw(fmt.Errorf("can't load chart from URL %s, status code: %d", chartIdentifier, response.StatusCode))
			}

			data, err := io.ReadAll(response.Body)
			if err != nil {
				js.Throw(fmt.Errorf("can't read chart data from URL %s, %v", chartIdentifier, err))
			}

			chart = LoadChart(data)
		} else {
			chart = LoadChartFromPath(chartIdentifier)
		}

		if chart == nil {
			js.Throw(fmt.Errorf("can't load chart from path %s", chartIdentifier))
		}

		options := NewHelmOptionsWithValues(releaseName, "", values)

		documentGroup := GenerateFromChart(chart, context, options)
		context.AddDocumentGroup(documentGroup)
	})
}

func AddHelmChartObject(builder *Builder, chartIdentifier string, releaseName string, values *sobek.Object) {
	valuesString, err := SerializeSobekObjectToYaml(builder.jsRuntime, values)
	if err != nil {
		js.Throw(fmt.Errorf("can't serialize values object to yaml, %v", err))
	}

	AddHelmChart(builder, chartIdentifier, releaseName, valuesString)
}

func AddHelmChartNoValues(builder *Builder, chartIdentifier string, releaseName string) {
	AddHelmChart(builder, chartIdentifier, releaseName, "")
}

// Runs helm template with values from the options and parses the generated documents.
func GenerateFromChart(chart *chart.Chart, context *BuildContext, options *HelmOptions) *DocumentGroup {
	options.sanitize()

	config := action.Configuration{}
	client := action.NewInstall(&config)

	client.DryRun = true
	client.Replace = true
	client.ClientOnly = true
	client.ReleaseName = options.ReleaseName
	client.Namespace = options.Namespace

	for resource := range context.KubernetesResourceInfo.resources.Iterator().C {
		client.APIVersions = append(client.APIVersions, resource.ApiVersion)
		client.APIVersions = append(client.APIVersions, fmt.Sprintf("%s/%s", resource.ApiVersion, resource.Kind))
	}

	slog.Info(
		"Generating documents using Helm, chart: ${chart}, version: ${version}",
		slog.String("chart", chart.Metadata.Name),
		slog.String("version", chart.Metadata.Version))

	kubeVersion := context.BuilderOptions.KubernetesCluster.Version.String()

	if kubeVersion != "" {
		parsedKubeVersion, err := chartutil.ParseKubeVersion(kubeVersion)

		if err != nil {
			js.Throw(fmt.Errorf("invalid kube version '%s': %v", kubeVersion, err))
		}

		client.KubeVersion = parsedKubeVersion
	}

	values := options.getValues()

	helmRelease, err := client.Run(chart, values)
	if err != nil {
		js.Throw(fmt.Errorf("helm returned error, %v", err))
	}

	documentGroup := NewDocumentGroup(options.ReleaseName)
	documents := HelmManifestToDocuments(context.JsRuntime, helmRelease.Manifest, options.ReleaseName, "no-name.yaml")

	for _, document := range documents {
		documentGroup.AddDocument(document)
	}

	for _, hook := range helmRelease.Hooks {
		if slices.Contains(hook.Events, release.HookTest) {
			continue
		}

		hookDocuments := HelmManifestToDocuments(context.JsRuntime, hook.Manifest, options.ReleaseName, hook.Path)

		for _, document := range hookDocuments {
			documentGroup.AddDocument(document)
		}
	}

	for _, crd := range chart.CRDObjects() {
		manifest := bytes.NewBuffer(crd.File.Data).String()

		crdDocuments := HelmManifestToDocuments(context.JsRuntime, manifest, options.ReleaseName, crd.Filename)

		for _, document := range crdDocuments {
			documentGroup.AddDocument(document)
		}
	}

	fixNameClashes(documentGroup)

	return documentGroup
}

func (options *HelmOptions) getValues() (values map[string]interface{}) {
	valuesYaml := util.MultilineString(options.Values)

	slog.Debug("Values for helm chart, release name: ${releaseName}, values:\n${values}",
		slog.String("releaseName", options.ReleaseName),
		slog.String("values", valuesYaml))

	if valuesYaml == "" {
		return map[string]interface{}{}
	}

	if err := yaml.Unmarshal([]byte(valuesYaml), &values); err != nil {
		js.Throw(fmt.Errorf("can't parse values yaml, %v", err))
	}

	return values
}

var yamlDocumentSeparator = regexp.MustCompile(`(?m)^---\s*$`)

func HelmManifestToDocuments(jsRuntime *js.JsRuntime, manifests string, releaseName string, defaultName string) []*Document {
	var documents []*Document
	var splitManifests map[string]string

	numberOfDocuments := len(yamlDocumentSeparator.FindAllString(manifests, -1))

	if numberOfDocuments > 1 {
		splitManifests = releaseutil.SplitManifests(manifests)
	} else {
		splitManifests = map[string]string{
			"manifest-0": manifests,
		}
	}

	manifestsKeys := make([]string, 0, len(splitManifests))

	for k := range splitManifests {
		manifestsKeys = append(manifestsKeys, k)
	}

	manifestNameRegex := regexp.MustCompile("# Source: [^/]+/(.+)")

	for _, manifestKey := range manifestsKeys {
		manifest := splitManifests[manifestKey]
		submatch := manifestNameRegex.FindStringSubmatch(manifest)
		var manifestName string

		if len(submatch) == 0 {
			manifestName = defaultName
		} else {
			manifestName = submatch[1]
			manifest, _ = strings.CutPrefix(manifest, submatch[0])
		}

		document := createDocumentFromHelmManifest(jsRuntime, manifest, releaseName, manifestName)
		if document != nil {
			documents = append(documents, document)
		}
	}

	return documents
}

func createDocumentFromHelmManifest(jsRuntime *js.JsRuntime, manifest string, releaseName string, path string) *Document {
	if manifest == "" {
		return nil
	}

	document, err := ParseDocument(jsRuntime, manifest)
	if err != nil {
		js.Throw(err)
	}

	// Skip empty manifests.
	if document == nil {
		return nil
	}

	name := fixDocumentName(path, releaseName)
	document.SetPath(&name)

	return document
}

func fixDocumentName(name string, releaseName string) string {
	name = strings.TrimPrefix(name, fmt.Sprintf("%s/", releaseName))
	name = strings.TrimPrefix(name, "templates/")
	name = strings.ReplaceAll(name, "/templates/", "/")

	return name
}

func RemoveTestHooks(helmRelease *release.Release) {
	helmRelease.Hooks = slices.DeleteFunc(helmRelease.Hooks, func(hook *release.Hook) bool {
		for _, e := range hook.Events {
			if e == release.HookTest {
				return true
			}
		}

		return false
	})
}

type compareFieldGetter = func(*Document) *string

var compareFieldGetters []compareFieldGetter = []func(*Document) *string{
	func(x *Document) *string { return SobekObjectGetString(x.Object, "apiVersion") },
	func(x *Document) *string { return SobekObjectGetString(x.Object, "kind") },
	func(x *Document) *string { return SobekObjectGetStringChain(x.Object, "metadata", "namespace") },
	func(x *Document) *string { return SobekObjectGetStringChain(x.Object, "metadata", "name") },
}

// Fixes duplicate document file paths by adding an index suffix to the end the names. Sorts the documents
// by their apiVersion, kind, namespace, and name to ensure consistent naming.
// This is useful for documents that are generated by Helm charts.
func fixNameClashes(group *DocumentGroup) {
	groups := map[string][]*Document{}

	for _, document := range group.Documents {
		groups[document.GetPath()] = append(groups[document.GetPath()], document)
	}

	for documentPath, documents := range groups {
		if len(documents) == 1 {
			continue
		}

		sort.SliceStable(documents, func(i, j int) bool {
			for _, getter := range compareFieldGetters {
				x := getter(documents[i])
				y := getter(documents[j])

				if x == nil && y == nil {
					continue
				}

				if x == nil && y != nil {
					return true
				}

				if x != nil && y == nil {
					return false
				}

				if *x == *y {
					continue
				}

				return *x < *y
			}

			return false
		})

		for index, document := range documents {
			if index == 0 {
				continue
			}

			extension := path.Ext(documentPath)
			fileName := path.Base(documentPath)
			fileName, _ = strings.CutSuffix(fileName, extension)

			path := fmt.Sprintf("%s-%d%s", fileName, index, extension)
			document.SetPath(&path)
		}
	}
}

func registerHelm(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Builder]()).ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(AddHelmChart)),
		js.ExtensionMethod(reflect.ValueOf(AddHelmChartObject)).JsName("addHelmChart"),
		js.ExtensionMethod(reflect.ValueOf(AddHelmChartNoValues)).JsName("addHelmChart"),
	)

	jsRuntime.Type(reflect.TypeFor[chart.Chart]()).JsName("HelmChart").ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(GenerateFromChart)).JsName("generate"),
	).Constructors(
		js.Constructor(reflect.ValueOf(LoadChart)),
		js.Constructor(reflect.ValueOf(LoadChartFromPath)),
	)

	jsRuntime.Type(reflect.TypeFor[HelmOptions]()).Fields(
		js.Field("ReleaseName"),
		js.Field("Namespace"),
		js.Field("Values"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewHelmOptions)),
		js.Constructor(reflect.ValueOf(NewHelmOptionsWithValues)),
	)
}
