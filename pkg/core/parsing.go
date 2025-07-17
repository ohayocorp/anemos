package core

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg/js"
	"gopkg.in/yaml.v3"
	k8sYaml "sigs.k8s.io/yaml"
)

// Serializes given Kubernetes to YAML string using [k8sYaml.Marshal].
func serializeKubernetesObjectToYaml(data interface{}) string {
	value, err := k8sYaml.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("can't serialize Kubernetes object to yaml, %v", err))
	}

	return string(value)
}

func serializeSobekObjectToYaml(object *sobek.Object, jsRuntime *js.JsRuntime) string {
	mapping, err := jsToMapping(jsRuntime, object)
	if err != nil {
		js.Throw(fmt.Errorf("can't convert Sobek object to mapping, %v", err))
	}

	return SerializeToYaml(mapping)
}

func serializeToYamlJs(jsRuntime *js.JsRuntime, data interface{}) string {
	if object, ok := data.(*sobek.Object); ok {
		return serializeSobekObjectToYaml(object, jsRuntime)
	}

	return SerializeToYaml(data)
}

// Serializes given object to YAML string.
func SerializeToYaml(data interface{}) string {
	dataType := reflect.TypeOf(data)
	if dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
	}

	pkgPath := dataType.PkgPath()

	if strings.HasPrefix(pkgPath, "k8s.io") {
		return serializeKubernetesObjectToYaml(data)
	}

	if doc := GetAsPointer[Document](data); doc != nil {
		if doc.YamlNode != nil && doc.YamlNode.Content == nil {
			return ""
		}

		data = doc.YamlNode
	}

	var buffer bytes.Buffer

	encoder := yaml.NewEncoder(&buffer)
	encoder.SetIndent(2)

	if err := encoder.Encode(data); err != nil {
		js.Throw(fmt.Errorf("can't serialize object to yaml, %v", err))
	}

	encoder.Close()

	return buffer.String()
}

// Deserializes given string into a [yaml.Node]. Dedents the data using [Dedent] so that the
// multiline strings with indentation are handled properly. Trims the newlines before deserialization.
func ParseYamlNode(data string) *yaml.Node {
	return Pointer(ParseYaml[yaml.Node](data))
}

// Deserializes given string into an object of given type. Dedents the data using [Dedent] so that the
// multiline strings with indentation are handled properly. Trims the newlines before deserialization.
func ParseYaml[T any](data string) T {
	data = Dedent(data)
	data = strings.Trim(data, "\n")

	var result T
	if err := yaml.Unmarshal([]byte(data), &result); err != nil {
		js.Throw(fmt.Errorf("can't parse yaml, %v:\n%s", err, data))
	}

	return result
}

// Parses a template as a string.
func ParseTemplate(template string, data any) string {
	bytes := ParseTemplateAsBytes(template, data)
	return string(bytes)
}

// Parses a template by calling [MultilineString] on template text beforehand.
func ParseTemplateAsBytes(templateText string, data any) []byte {
	templateText = MultilineString(templateText)
	template := template.Must(template.New("template").Parse(templateText))

	var buffer bytes.Buffer

	if err := template.Execute(&buffer, data); err != nil {
		js.Throw(fmt.Errorf("can't parse template, %v\n%s", err, templateText))
	}

	return buffer.Bytes()
}

// Parses a template as a [yaml.Node].
func ParseTemplateAsYamlNode(template string, data any) *yaml.Node {
	bytes := ParseTemplateAsBytes(template, data)

	var node yaml.Node
	if err := yaml.Unmarshal(bytes, &node); err != nil {
		js.Throw(fmt.Errorf("can't parse yaml, %v:\n%s", err, string(bytes)))
	}

	return &node
}

// Parses given text as a [Document].
func ParseDocument(path string, text string) *Document {
	yamlNode := ParseYamlNode(text)
	return NewDocument(path, yamlNode)
}

// Parses given text as a [Mapping].
func ParseMapping(text string) *Mapping {
	yamlNode := ParseYamlNode(text)

	if yamlNode.Kind == yaml.DocumentNode && len(yamlNode.Content) == 1 {
		yamlNode = yamlNode.Content[0]
	}

	return NewMapping(yamlNode)
}

// Parses given text as a [Sequence].
func ParseSequence(text string) *Sequence {
	yamlNode := ParseYamlNode(text)

	if yamlNode.Kind == yaml.DocumentNode && len(yamlNode.Content) == 1 {
		yamlNode = yamlNode.Content[0]
	}

	return NewSequence(yamlNode)
}

// Parses a template as a [Document].
func ParseTemplateAsDocument(path string, text string, data any) *Document {
	yaml := ParseTemplateAsYamlNode(text, data)
	return NewDocument(path, yaml)
}

// Parses a template as a [Mapping].
func ParseTemplateAsMapping(text string, data any) *Mapping {
	yamlNode := ParseTemplateAsYamlNode(text, data)

	if yamlNode.Kind == yaml.DocumentNode && len(yamlNode.Content) == 1 {
		yamlNode = yamlNode.Content[0]
	}

	return NewMapping(yamlNode)
}

// Parses a template as a [Sequence].
func ParseTemplateAsSequence(text string, data any) *Sequence {
	yaml := ParseTemplateAsYamlNode(text, data)
	return NewSequence(yaml)
}

func registerYamlParsing(jsRuntime *js.JsRuntime) {
	jsRuntime.Function(reflect.ValueOf(ParseDocument))
	jsRuntime.Function(reflect.ValueOf(ParseMapping))
	jsRuntime.Function(reflect.ValueOf(ParseSequence))
	jsRuntime.Function(reflect.ValueOf(serializeToYamlJs)).JsName("serializeToYaml")
}
