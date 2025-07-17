package core

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg/js"
	"gopkg.in/yaml.v3"
)

// Sequence wraps a [yaml.Node] with kind [yaml.SequenceNode] and provides convenience methods for YAML modification.
type Sequence struct {
	YamlNode *yaml.Node
}

// Retuns a new [Sequence] with the content set to an empty yaml node with type [yaml.SequenceNode].
func NewEmptySequence() *Sequence {
	yamlNode := NewYamlSequenceNode()
	return &Sequence{YamlNode: yamlNode}
}

// Returns a new [yaml.Node] with kind [yaml.SequenceNode] and style [yaml.TaggedStyle].
func NewYamlSequenceNode() *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.SequenceNode,
		Style: yaml.TaggedStyle,
	}
}

// Retuns a new [Sequence] with the content set to the given yaml sequence. Panics if the
// yaml node is nil or its type is not [yaml.SequenceNode].
func NewSequence(yamlNode *yaml.Node) *Sequence {
	if yamlNode == nil {
		panic(fmt.Errorf("passed yaml node is nil"))
	}

	if yamlNode.Kind != yaml.SequenceNode {
		panic(fmt.Errorf("yaml node is not sequence, but %s", getYamlNodeKind(yamlNode)))
	}

	return &Sequence{
		YamlNode: yamlNode,
	}
}

// Custom YAML marshaler that actually marshals the internal yaml node.
func (sequence *Sequence) MarshalYAML() (interface{}, error) {
	return sequence.YamlNode, nil
}

// Returns a clone of the sequence.
func (sequence *Sequence) Clone() *Sequence {
	clone := cloneYamlNode(sequence.YamlNode)

	if clone.Kind == yaml.DocumentNode {
		clone = clone.Content[0]
	}

	return NewSequence(clone)
}

// Returns the number of children.
func (sequence *Sequence) Length() int {
	return len(sequence.YamlNode.Content)
}

// Removes all children.
func (sequence *Sequence) Clear() {
	sequence.YamlNode.Content = nil
}

// Returns the child at given index as [Mapping]. Throws if the child is not a [Mapping] or index
// is out of bounds.
func (sequence *Sequence) GetMapping(index int) *Mapping {
	yamlNode := sequence.YamlNode
	contents := yamlNode.Content

	if index < 0 || index >= len(contents) {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", index, len(contents)))
	}

	result := contents[index]
	return NewMapping(result)
}

// Returns the child at given index as [Sequence]. Throws if the child is not a [Sequence] or index
// is out of bounds.
func (sequence *Sequence) GetSequence(index int) *Sequence {
	yamlNode := sequence.YamlNode
	contents := yamlNode.Content

	if index < 0 || index >= len(contents) {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", index, len(contents)))
	}

	result := contents[index]
	return NewSequence(result)
}

// Returns the child at given index as [Scalar]. Throws if the child is not a [Scalar] or index
// is out of bounds.
func (sequence *Sequence) GetScalar(index int) *Scalar {
	yamlNode := sequence.YamlNode
	contents := yamlNode.Content

	if index < 0 || index >= len(contents) {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", index, len(contents)))
	}

	result := contents[index]
	return NewScalar(result)
}

// Returns the value of the node at given index. Throws if the child is not a [Scalar] or index
// is out of bounds.
func (sequence *Sequence) GetValue(index int) string {
	scalar := sequence.GetScalar(index)
	return scalar.GetValue()
}

// Returns the first child that returns true for the filter function. Returns nil if the node is not found.
func (sequence *Sequence) GetMappingFunc(filter func(*Mapping) bool) *Mapping {
	yamlNode := sequence.YamlNode
	contents := yamlNode.Content

	for _, content := range contents {
		if content.Kind != yaml.MappingNode {
			continue
		}

		mapping := NewMapping(content)
		ok := filter(mapping)
		if ok {
			return mapping
		}
	}

	return nil
}

// Returns the first child that returns true for the filter function. Returns nil if the node is not found.
func (sequence *Sequence) GetSequenceFunc(filter func(*Sequence) bool) *Sequence {
	yamlNode := sequence.YamlNode
	contents := yamlNode.Content

	for _, content := range contents {
		if content.Kind != yaml.SequenceNode {
			continue
		}

		sequence := NewSequence(content)
		ok := filter(sequence)
		if ok {
			return sequence
		}
	}

	return nil
}

// Returns the first child that returns true for the filter function. Returns nil if the node is not found.
func (sequence *Sequence) GetScalarFunc(filter func(*Scalar) bool) *Scalar {
	yamlNode := sequence.YamlNode
	contents := yamlNode.Content

	for _, content := range contents {
		if content.Kind != yaml.ScalarNode {
			continue
		}

		scalar := NewScalar(content)
		ok := filter(scalar)
		if ok {
			return scalar
		}
	}

	return nil
}

// Adds the given [Mapping] at the end of the children.
func (sequence *Sequence) AddMapping(child *Mapping) {
	if child == nil {
		js.Throw(fmt.Errorf("cannot add nil Mapping to Sequence"))
	}

	yaml := sequence.YamlNode
	yaml.Content = append(yaml.Content, child.YamlNode)
}

// Adds the given [Mapping] slice at the end of the children.
func (sequence *Sequence) AddMappings(children []*Mapping) {
	for _, child := range children {
		sequence.AddMapping(child)
	}
}

// Adds the given [Sequence] at the end of the children.
func (sequence *Sequence) AddSequence(child *Sequence) {
	if child == nil {
		js.Throw(fmt.Errorf("cannot add nil Sequence to Sequence"))
	}

	yaml := sequence.YamlNode
	yaml.Content = append(yaml.Content, child.YamlNode)
}

// Adds the given [Sequence] slice at the end of the children.
func (sequence *Sequence) AddSequences(children []*Sequence) {
	for _, child := range children {
		sequence.AddSequence(child)
	}
}

// Adds the given [Scalar] at the end of the children.
func (sequence *Sequence) AddScalar(child *Scalar) {
	if child == nil {
		js.Throw(fmt.Errorf("cannot add nil Scalar to Sequence"))
	}

	yaml := sequence.YamlNode
	yaml.Content = append(yaml.Content, child.YamlNode)
}

// Adds the given [Scalar] slice at the end of the children.
func (sequence *Sequence) AddScalars(children []*Scalar) {
	for _, child := range children {
		sequence.AddScalar(child)
	}
}

// Adds the given value as a [Scalar] at the end of the children and returns the created [Scalar].
func (sequence *Sequence) AddValue(value string) *Scalar {
	yaml := sequence.YamlNode
	scalar := NewYamlScalarNode(value)

	yaml.Content = append(yaml.Content, scalar)

	return NewScalar(scalar)
}

// Adds the given values at the end of the children.
func (sequence *Sequence) AddValues(children []string) {
	for _, child := range children {
		sequence.AddValue(child)
	}
}

// Set the given [Mapping] at the given index. Throws if the index is out of bounds.
func (sequence *Sequence) SetMapping(index int, value *Mapping) {
	if value == nil {
		js.Throw(fmt.Errorf("cannot set nil Mapping at index %d", index))
	}

	if index < 0 || index >= len(sequence.YamlNode.Content) {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", index, len(sequence.YamlNode.Content)))
	}

	yaml := sequence.YamlNode
	yaml.Content[index] = value.YamlNode
}

// Sets the given [Sequence] at the given index. Throws if the index is out of bounds.
func (sequence *Sequence) SetSequence(index int, value *Sequence) {
	if value == nil {
		js.Throw(fmt.Errorf("cannot set nil Sequence at index %d", index))
	}

	if index < 0 || index >= len(sequence.YamlNode.Content) {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", index, len(sequence.YamlNode.Content)))
	}

	yaml := sequence.YamlNode
	yaml.Content[index] = value.YamlNode
}

// Sets the given [Scalar] at the given index. Throws if the index is out of bounds.
func (sequence *Sequence) SetScalar(index int, value *Scalar) {
	if value == nil {
		js.Throw(fmt.Errorf("cannot set nil Scalar at index %d", index))
	}

	if index < 0 || index >= len(sequence.YamlNode.Content) {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", index, len(sequence.YamlNode.Content)))
	}

	yaml := sequence.YamlNode
	yaml.Content[index] = value.YamlNode
}

// Sets the given value as a [Scalar] at the given index and returns the created [Scalar]. Throws if the index is out of bounds.
func (sequence *Sequence) SetValue(index int, value string) *Scalar {
	valueNode := NewYamlScalarNode(value)
	scalar := NewScalar(valueNode)

	sequence.SetScalar(index, scalar)

	return scalar
}

// Inserts the given [Mapping] at the given index. Throws if the index is out of bounds.
func (sequence *Sequence) InsertMapping(index int, value *Mapping) {
	if value == nil {
		js.Throw(fmt.Errorf("cannot insert nil Mapping at index %d", index))
	}

	if index < 0 || index > len(sequence.YamlNode.Content) {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", index, len(sequence.YamlNode.Content)))
	}

	yaml := sequence.YamlNode
	yaml.Content = slices.Insert(yaml.Content, index, value.YamlNode)
}

// Inserts the given [Sequence] at the given index. Throws if the index is out of bounds.
func (sequence *Sequence) InsertSequence(index int, value *Sequence) {
	if value == nil {
		js.Throw(fmt.Errorf("cannot insert nil Sequence at index %d", index))
	}

	if index < 0 || index > len(sequence.YamlNode.Content) {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", index, len(sequence.YamlNode.Content)))
	}

	yaml := sequence.YamlNode
	yaml.Content = slices.Insert(yaml.Content, index, value.YamlNode)
}

// Inserts the given [Scalar] at the given index. Throws if the index is out of bounds.
func (sequence *Sequence) InsertScalar(index int, value *Scalar) {
	if value == nil {
		js.Throw(fmt.Errorf("cannot insert nil Scalar at index %d", index))
	}

	if index < 0 || index > len(sequence.YamlNode.Content) {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", index, len(sequence.YamlNode.Content)))
	}

	yaml := sequence.YamlNode
	yaml.Content = slices.Insert(yaml.Content, index, value.YamlNode)
}

// Inserts the given value as a [Scalar] at the given index and returns the created [Scalar]. Error is not nil
// if the index is out of bounds.
func (sequence *Sequence) InsertValue(index int, value string) *Scalar {
	valueNode := NewYamlScalarNode(value)
	scalar := NewScalar(valueNode)

	sequence.InsertScalar(index, scalar)

	return scalar
}

// Removes the node at the given index. Throws if the index is out of bounds.
func (sequence *Sequence) Remove(index int) {
	if index < 0 || index >= len(sequence.YamlNode.Content) {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", index, len(sequence.YamlNode.Content)))
	}

	yamlNode := sequence.YamlNode
	contents := yamlNode.Content

	yamlNode.Content = append(contents[:index], contents[index+1:]...)
}

func jsToSequence(jsRuntime *js.JsRuntime, jsValue sobek.Value) (*Sequence, error) {
	slice, err := jsRuntime.MarshalToGo(jsValue, reflect.TypeFor[[]sobek.Value]())
	if err != nil {
		return nil, fmt.Errorf("failed to convert JS value to []sobek.Value: %w", err)
	}

	sequence := NewEmptySequence()

	for _, value := range slice.Interface().([]sobek.Value) {
		if value == nil {
			continue
		}

		i, err := jsRuntime.MarshalToGo(value, reflect.TypeFor[int]())
		if err == nil {
			sequence.AddValue(fmt.Sprintf("%d", i.Interface().(int)))
			continue
		}

		f, err := jsRuntime.MarshalToGo(value, reflect.TypeFor[float64]())
		if err == nil {
			sequence.AddValue(fmt.Sprintf("%f", f.Interface().(float64)))
			continue
		}

		b, err := jsRuntime.MarshalToGo(value, reflect.TypeFor[bool]())
		if err == nil {
			sequence.AddValue(fmt.Sprintf("%t", b.Interface().(bool)))
			continue
		}

		str, err := jsRuntime.MarshalToGo(value, reflect.TypeFor[string]())
		if err == nil {
			stringValue := str.Interface().(string)
			scalar := sequence.AddValue(stringValue)

			SetScalarNodeStyle(scalar, stringValue)

			continue
		}

		childSequence, err := jsToSequence(jsRuntime, value)
		if err == nil {
			sequence.AddSequence(childSequence)
			continue
		}

		mapping, err := jsToMapping(jsRuntime, value)
		if err == nil {
			sequence.AddMapping(mapping)
			continue
		}

		return nil, fmt.Errorf("unsupported type %T", value)
	}

	return sequence, nil
}

func registerYamlSequence(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Sequence]()).Methods(
		js.Method("AddMapping").JsName("add"),
		js.Method("AddMappings").JsName("add"),
		js.Method("AddScalar").JsName("add"),
		js.Method("AddScalars").JsName("add"),
		js.Method("AddSequence").JsName("add"),
		js.Method("AddSequences").JsName("add"),
		js.Method("AddValue").JsName("add"),
		js.Method("AddValues").JsName("add"),
		js.Method("SetMapping").JsName("set"),
		js.Method("SetScalar").JsName("set"),
		js.Method("SetSequence").JsName("set"),
		js.Method("SetValue").JsName("set"),
		js.Method("InsertMapping").JsName("insert"),
		js.Method("InsertScalar").JsName("insert"),
		js.Method("InsertSequence").JsName("insert"),
		js.Method("InsertValue").JsName("insert"),
		js.Method("GetMapping"),
		js.Method("GetMappingFunc").JsName("getMapping"),
		js.Method("GetScalar"),
		js.Method("GetScalarFunc").JsName("getScalar"),
		js.Method("GetSequence"),
		js.Method("GetSequenceFunc").JsName("getSequence"),
		js.Method("GetValue"),
		js.Method("Length"),
		js.Method("Remove"),
		js.Method("Clear"),
		js.Method("Clone"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewEmptySequence)),
		js.Constructor(reflect.ValueOf(jsToSequence)),
	).TypeConversion(
		reflect.ValueOf(jsToSequence),
	).DisableObjectMapping()
}
