package core

import (
	"cmp"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg/js"
	"gopkg.in/yaml.v3"
)

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

// Helper method that clones a [yaml.Node] by serializing and then deserializing it.
func cloneYamlNode(node *yaml.Node) *yaml.Node {
	serialized, err := yaml.Marshal(node)
	if err != nil {
		js.Throw(fmt.Errorf("failed to clone yaml node: %w", err))
	}

	deserialized := &yaml.Node{}
	err = yaml.Unmarshal(serialized, deserialized)
	if err != nil {
		js.Throw(fmt.Errorf("failed to clone yaml node: %w", err))
	}

	return deserialized
}

func tryGetScalar(jsRuntime *js.JsRuntime, value sobek.Value) *Scalar {
	i, err := jsRuntime.MarshalToGo(jsRuntime.Runtime.ToValue(value), reflect.TypeFor[int]())
	if err == nil {
		return NewScalarFromIntValue(i.Interface().(int))
	}

	f, err := jsRuntime.MarshalToGo(jsRuntime.Runtime.ToValue(value), reflect.TypeFor[float64]())
	if err == nil {
		return NewScalarFromFloatValue(f.Interface().(float64))
	}

	b, err := jsRuntime.MarshalToGo(jsRuntime.Runtime.ToValue(value), reflect.TypeFor[bool]())
	if err == nil {
		return NewScalarFromBoolValue(b.Interface().(bool))
	}

	str, err := jsRuntime.MarshalToGo(jsRuntime.Runtime.ToValue(value), reflect.TypeFor[string]())
	if err == nil {
		stringValue := str.Interface().(string)
		scalar := NewScalarFromStringValue(stringValue)

		SetScalarNodeStyle(scalar, stringValue)

		return scalar
	}

	return nil
}

func SetScalarNodeStyle(scalar *Scalar, value string) {
	if strings.Contains(value, "\n") {
		scalar.SetStyle(yaml.LiteralStyle)
	}

	// If the value can be converted to a int, float or bool, set the style to double quoted
	// so that it is interpreted as a string in YAML.
	if _, err := strconv.Atoi(value); err == nil {
		scalar.SetStyle(yaml.DoubleQuotedStyle)
	} else if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		scalar.SetStyle(yaml.DoubleQuotedStyle)
	} else if _, err := strconv.ParseUint(value, 10, 64); err == nil {
		scalar.SetStyle(yaml.DoubleQuotedStyle)
	} else if _, err := strconv.ParseFloat(value, 64); err == nil {
		scalar.SetStyle(yaml.DoubleQuotedStyle)
	} else if _, err := strconv.ParseBool(value); err == nil {
		scalar.SetStyle(yaml.DoubleQuotedStyle)
	}
}

func Pointer[T any](input T) *T {
	return &input
}

func GetAsPointer[T any](o any) *T {
	if t, ok := o.(T); ok {
		return &t
	}

	if t, ok := o.(*T); ok {
		return t
	}

	return nil
}

// GetImageTag returns the tag of the image. If the image does not have a tag, empty string is returned.
func GetImageTag(image string) string {
	index := strings.LastIndex(image, ":")
	if index == -1 {
		return ""
	}

	return image[index+1:]
}

func SortedKeys[TKey cmp.Ordered, TValue any](m map[TKey]TValue) []TKey {
	keys := make([]TKey, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}
