package core

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"

	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg/js"
	"gopkg.in/yaml.v3"
)

// Parses a template as a string by calling [MultilineString] on template text beforehand.
// Panics if template is invalid.
func ParseTemplate(templateText string, data any) string {
	templateText = MultilineString(templateText)
	template := template.Must(template.New("template").Parse(templateText))

	var buffer bytes.Buffer

	if err := template.Execute(&buffer, data); err != nil {
		panic(fmt.Errorf("can't parse template, %v\n%s", err, templateText))
	}

	return buffer.String()
}

// Deserializes given string into an object of given type. Dedents the data using [Dedent] so that the
// multiline strings with indentation are handled properly. Trims the newlines before deserialization.
func ParseYaml[T any](data string) (T, error) {
	data = Dedent(data)
	data = strings.Trim(data, "\n")

	var result T
	if err := yaml.Unmarshal([]byte(data), &result); err != nil {
		return result, fmt.Errorf("can't parse yaml, %v:\n%s", err, data)
	}

	return result, nil
}

// Parses given text as a [Document].
func ParseDocument(jsRuntime *js.JsRuntime, yaml string) (*Document, error) {
	object, err := Parse(jsRuntime, yaml)
	if err != nil {
		return nil, err
	}

	return NewDocumentWithContent(object), nil
}

func Parse(jsRuntime *js.JsRuntime, yamlText string) (*sobek.Object, error) {
	node, err := ParseYaml[yaml.Node](yamlText)
	if err != nil {
		return nil, err
	}

	value, err := parseYamlNode(jsRuntime, &node)
	if err != nil {
		return nil, err
	}

	return value.ToObject(jsRuntime.Runtime), nil
}

func parseYamlNode(jsRuntime *js.JsRuntime, node *yaml.Node) (sobek.Value, error) {
	if node.Kind == yaml.DocumentNode && len(node.Content) > 0 {
		node = node.Content[0]
	}

	if scalar := tryParseScalar(jsRuntime, node); scalar != nil {
		return scalar, nil
	}

	mapping, err := tryParseMapping(jsRuntime, node)
	if err != nil {
		return nil, err
	}

	if mapping != nil {
		return mapping, nil
	}

	if sequence, err := tryParseSequence(jsRuntime, node); err == nil {
		return sequence, nil
	}

	return nil, fmt.Errorf("can't parse yaml node of kind %s", getYamlNodeKind(node))
}

func tryParseMapping(jsRuntime *js.JsRuntime, node *yaml.Node) (sobek.Value, error) {
	if node.Kind != yaml.MappingNode {
		return nil, nil
	}

	object := jsRuntime.Runtime.NewObject()

	for i := 0; i < len(node.Content); i += 2 {
		key := node.Content[i]
		value := node.Content[i+1]

		valueObject, err := parseYamlNode(jsRuntime, value)
		if err != nil {
			return nil, err
		}

		object.Set(key.Value, valueObject)
	}

	return object, nil
}

func tryParseSequence(jsRuntime *js.JsRuntime, node *yaml.Node) (sobek.Value, error) {
	if node.Kind != yaml.SequenceNode {
		return nil, nil
	}

	array := jsRuntime.Runtime.NewArray()

	for i, content := range node.Content {
		valueObject, err := parseYamlNode(jsRuntime, content)
		if err != nil {
			return nil, err
		}

		array.Set(strconv.Itoa(i), valueObject)
	}

	return array.ToObject(jsRuntime.Runtime), nil
}

func tryParseScalar(jsRuntime *js.JsRuntime, node *yaml.Node) sobek.Value {
	if node.Kind != yaml.ScalarNode {
		return nil
	}

	tag := node.Tag

	if tag == "!!bool" {
		if value, err := strconv.ParseBool(node.Value); err == nil {
			return jsRuntime.Runtime.ToValue(value)
		}
	}

	if tag == "!!int" {
		if value, err := strconv.ParseInt(node.Value, 10, 64); err == nil {
			return jsRuntime.Runtime.ToValue(value)
		}
	}

	if tag == "!!float" {
		if value, err := strconv.ParseFloat(node.Value, 64); err == nil {
			return jsRuntime.Runtime.ToValue(value)
		}
	}

	return jsRuntime.Runtime.ToValue(node.Value)
}

func getYamlNodeKind(node *yaml.Node) string {
	switch node.Kind {
	case yaml.DocumentNode:
		return "document"
	case yaml.SequenceNode:
		return "sequence"
	case yaml.MappingNode:
		return "mapping"
	case yaml.ScalarNode:
		return "scalar"
	case yaml.AliasNode:
		return "alias"
	}

	return "unknown"
}

func registerYamlParsing(jsRuntime *js.JsRuntime) {
	jsRuntime.Function(reflect.ValueOf(Parse))
	jsRuntime.Function(reflect.ValueOf(ParseDocument))
}
