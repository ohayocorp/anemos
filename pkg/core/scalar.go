package core

import (
	"fmt"
	"reflect"

	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg/js"
	"gopkg.in/yaml.v3"
)

// Scalar wraps a [yaml.Node] with kind [yaml.ScalarNode] and provides convenience methods for YAML modification.
type Scalar struct {
	YamlNode *yaml.Node
}

// Retuns a new [Scalar] with the content set to an empty yaml node with type [yaml.ScalarNode].
func NewEmptyScalar() *Scalar {
	yamlNode := NewYamlScalarNode("")
	return &Scalar{YamlNode: yamlNode}
}

// Retuns a new [yaml.Node] with kind [yaml.ScalarNode] and style [yaml.FlowStyle] containing the given value.
func NewYamlScalarNode(value string) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Style: yaml.FlowStyle,
		Value: value,
	}
}

// Retuns a new [Scalar] with the content set to the given yaml scalar. Panics if the
// node is nil or its type is not [yaml.ScalarNode].
func NewScalar(yamlNode *yaml.Node) *Scalar {
	if yamlNode == nil {
		panic(fmt.Errorf("passed yaml node is nil"))
	}

	if yamlNode.Kind != yaml.ScalarNode {
		panic(fmt.Errorf("yaml node is not scalar, but %s", getYamlNodeKind(yamlNode)))
	}

	return &Scalar{
		YamlNode: yamlNode,
	}
}

// Custom YAML marshaler that actually marshals the internal node.
func (scalar *Scalar) MarshalYAML() (interface{}, error) {
	return scalar.YamlNode, nil
}

// Returns a clone of the scalar.
func (scalar *Scalar) Clone() *Scalar {
	clone := cloneYamlNode(scalar.YamlNode)

	if clone.Kind == yaml.DocumentNode {
		clone = clone.Content[0]
	}

	return NewScalar(clone)
}

// Returns the value as string.
func (scalar *Scalar) GetValue() string {
	yamlNode := scalar.YamlNode
	return yamlNode.Value
}

// Sets the value to the given string.
func (scalar *Scalar) SetValue(value string) {
	scalar.YamlNode.Value = value
}

func (scalar *Scalar) SetStyle(style yaml.Style) {
	scalar.YamlNode.Style = style
}

func NewScalarFromStringValue(value string) *Scalar {
	scalar := NewEmptyScalar()
	scalar.SetValue(value)

	return scalar
}

func NewScalarFromStringValueAndStyle(value string, style yaml.Style) *Scalar {
	scalar := NewEmptyScalar()
	scalar.SetValue(value)
	scalar.SetStyle(style)

	return scalar
}

func NewScalarFromIntValue(value int) *Scalar {
	return NewScalarFromStringValue(fmt.Sprintf("%d", value))
}

func NewScalarFromIntValueAndStyle(value int, style yaml.Style) *Scalar {
	return NewScalarFromStringValueAndStyle(fmt.Sprintf("%d", value), style)
}

func NewScalarFromFloatValue(value float64) *Scalar {
	return NewScalarFromStringValue(fmt.Sprintf("%f", value))
}

func NewScalarFromFloatValueAndStyle(value float64, style yaml.Style) *Scalar {
	return NewScalarFromStringValueAndStyle(fmt.Sprintf("%f", value), style)
}

func NewScalarFromBoolValue(value bool) *Scalar {
	return NewScalarFromStringValue(fmt.Sprintf("%t", value))
}

func NewScalarFromBoolValueAndStyle(value bool, style yaml.Style) *Scalar {
	return NewScalarFromStringValueAndStyle(fmt.Sprintf("%t", value), style)
}

func (scalar *Scalar) SetValueInt(value int) {
	scalar.SetValue(fmt.Sprintf("%d", value))
}

func (scalar *Scalar) SetValueFloat(value float64) {
	scalar.SetValue(fmt.Sprintf("%f", value))
}

func (scalar *Scalar) SetValueBool(value bool) {
	scalar.SetValue(fmt.Sprintf("%t", value))
}

func jsToScalar(jsRuntime *js.JsRuntime, jsValue sobek.Value) (*Scalar, error) {
	value, err := jsRuntime.MarshalToGo(jsValue, reflect.TypeFor[string]())
	if err != nil {
		return nil, fmt.Errorf("failed to convert JS value to string: %w", err)
	}

	return NewScalarFromStringValue(value.Interface().(string)), nil
}

func (scalar *Scalar) ToJSON(jsRuntime *js.JsRuntime, dummy string) sobek.Value {
	return jsRuntime.ToSobekValue(scalar.GetValue())
}

func registerYamlScalar(jsRuntime *js.JsRuntime) {
	jsRuntime.Variable("YamlStyle", "Plain", reflect.ValueOf(yaml.TaggedStyle))
	jsRuntime.Variable("YamlStyle", "SingleQuoted", reflect.ValueOf(yaml.SingleQuotedStyle))
	jsRuntime.Variable("YamlStyle", "DoubleQuoted", reflect.ValueOf(yaml.DoubleQuotedStyle))
	jsRuntime.Variable("YamlStyle", "Literal", reflect.ValueOf(yaml.LiteralStyle))
	jsRuntime.Variable("YamlStyle", "Folded", reflect.ValueOf(yaml.FoldedStyle))

	jsRuntime.Type(reflect.TypeFor[Scalar]()).Methods(
		js.Method("Clone"),
		js.Method("GetValue"),
		js.Method("SetValue"),
		js.Method("SetValueInt").JsName("setValue"),
		js.Method("SetValueFloat").JsName("setValue"),
		js.Method("SetValueBool").JsName("setValue"),
		js.Method("SetStyle"),
		js.Method("ToJSON"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewEmptyScalar)),
		js.Constructor(reflect.ValueOf(NewScalarFromStringValue)),
		js.Constructor(reflect.ValueOf(NewScalarFromStringValueAndStyle)),
		js.Constructor(reflect.ValueOf(NewScalarFromIntValue)),
		js.Constructor(reflect.ValueOf(NewScalarFromIntValueAndStyle)),
		js.Constructor(reflect.ValueOf(NewScalarFromFloatValue)),
		js.Constructor(reflect.ValueOf(NewScalarFromFloatValueAndStyle)),
		js.Constructor(reflect.ValueOf(NewScalarFromBoolValue)),
		js.Constructor(reflect.ValueOf(NewScalarFromBoolValueAndStyle)),
	).TypeConversion(
		reflect.ValueOf(jsToScalar),
	).DisableObjectMapping()
}
