package core

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg/js"
	"gopkg.in/yaml.v3"
)

func SerializeSobekObjectToYaml(jsRuntime *js.JsRuntime, object *sobek.Object) (string, error) {
	node, err := serializeSobekValueToYamlNode(jsRuntime, object)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer

	encoder := yaml.NewEncoder(&buffer)
	encoder.SetIndent(2)

	if err := encoder.Encode(node); err != nil {
		return "", fmt.Errorf("can't serialize object to yaml, %v", err)
	}

	encoder.Close()

	return buffer.String(), nil
}

func serializeSobekValueToYamlNode(jsRuntime *js.JsRuntime, value sobek.Value) (*yaml.Node, error) {
	errs := []error{}

	// First try to marshal the object into primitive types.
	scalarNode, err := serializeSobekValueToScalar(jsRuntime, value)
	if err == nil {
		return scalarNode, nil
	}
	errs = append(errs, err)

	sequenceNode, err := serializeSobekValueToSequence(jsRuntime, value)
	if err == nil {
		return sequenceNode, nil
	}
	errs = append(errs, err)

	mappingNode, err := serializeSobekValueToMapping(jsRuntime, value)
	if err == nil {
		return mappingNode, nil
	}
	errs = append(errs, err)

	return nil, fmt.Errorf(
		"unsupported Sobek value type: %s, conversion errors:\n%w",
		value.ExportType().String(),
		errors.Join(errs...))
}

func serializeSobekValueToMapping(jsRuntime *js.JsRuntime, value sobek.Value) (*yaml.Node, error) {
	object, ok := value.(*sobek.Object)
	if !ok {
		return nil, fmt.Errorf("expected sobek.Object, got %s", value.ExportType().String())
	}

	mapping := &yaml.Node{
		Kind: yaml.MappingNode,
	}

	for _, key := range object.GetOwnPropertyNames() {
		value := object.Get(key)
		node, err := serializeSobekValueToYamlNode(jsRuntime, value)
		if err != nil {
			return nil, err
		}

		// Append the key first.
		mapping.Content = append(mapping.Content, &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: key,
		})

		// Append the value next.
		mapping.Content = append(mapping.Content, node)
	}

	return mapping, nil
}

func serializeSobekValueToSequence(jsRuntime *js.JsRuntime, value sobek.Value) (*yaml.Node, error) {
	slice, err := jsRuntime.MarshalToGo(value, reflect.TypeFor[[]sobek.Value]())
	if err != nil {
		return nil, fmt.Errorf("can't marshal to go slice: %w", err)
	}

	sequence := &yaml.Node{
		Kind: yaml.SequenceNode,
	}

	for _, value := range slice.Interface().([]sobek.Value) {
		if value == nil {
			continue
		}

		node, err := serializeSobekValueToYamlNode(jsRuntime, value)
		if err != nil {
			return nil, err
		}

		sequence.Content = append(sequence.Content, node)
	}

	return sequence, nil
}

func serializeSobekValueToScalar(jsRuntime *js.JsRuntime, value sobek.Value) (*yaml.Node, error) {
	boolValue, err := jsRuntime.MarshalToGo(value, reflect.TypeFor[bool]())
	if err == nil {
		return &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: fmt.Sprintf("%t", boolValue.Bool()),
		}, nil
	}

	intValue, err := jsRuntime.MarshalToGo(value, reflect.TypeFor[int]())
	if err == nil {
		return &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: fmt.Sprintf("%d", intValue.Int()),
		}, nil
	}

	floatValue, err := jsRuntime.MarshalToGo(value, reflect.TypeFor[float64]())
	if err == nil {
		return &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: fmt.Sprintf("%f", floatValue.Float()),
		}, nil
	}

	stringValue, err := jsRuntime.MarshalToGo(value, reflect.TypeFor[string]())
	if err == nil {
		node := &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: stringValue.String(),
		}

		if strings.Contains(node.Value, "\n") {
			node.Style = yaml.LiteralStyle
		} else {
			node.Style = yaml.DoubleQuotedStyle
		}

		return node, nil
	}

	return nil, fmt.Errorf("unsupported scalar type: %s", value.ExportType().String())
}

func registerYamlSerialization(jsRuntime *js.JsRuntime) {
	jsRuntime.Function(reflect.ValueOf(SerializeSobekObjectToYaml)).JsName("serializeToYaml")
}
