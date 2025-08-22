package core

import (
	"fmt"
	"log/slog"
	"reflect"
	"slices"
	"sort"

	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg/js"
	"gopkg.in/yaml.v3"
)

var _ sort.Interface = &Mapping{}
var _ js.DynamicObjectCustomGetterSetter = &Mapping{}

// Mapping wraps a [yaml.Node] with kind [yaml.MappingNode] and provides convenience methods for YAML modification.
type Mapping struct {
	YamlNode *yaml.Node
}

// Retuns a new [Mapping] with the content set to an empty yaml node with type [yaml.MappingNode].
func NewEmptyMapping() *Mapping {
	yamlNode := NewYamlMappingNode()
	return &Mapping{YamlNode: yamlNode}
}

// Returns a new [Mapping] with the content set to the given yaml mapping. Panics if the
// yaml node is nil or its type is not [yaml.MappingNode].
func NewMapping(yamlNode *yaml.Node) *Mapping {
	if yamlNode == nil {
		panic(fmt.Errorf("passed yaml node is nil"))
	}

	if yamlNode.Kind != yaml.MappingNode {
		panic(fmt.Errorf("yaml node is not mapping, but %s", getYamlNodeKind(yamlNode)))
	}

	return &Mapping{
		YamlNode: yamlNode,
	}
}

// Retuns a new [yaml.Node] with kind [yaml.MappingNode] and style [yaml.TaggedStyle].
func NewYamlMappingNode() *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.MappingNode,
		Style: yaml.TaggedStyle,
	}
}

// Custom YAML marshaler that actually marshals the internal yaml node.
func (mapping *Mapping) MarshalYAML() (interface{}, error) {
	return mapping.YamlNode, nil
}

// Returns a clone of the mapping.
func (mapping *Mapping) Clone() *Mapping {
	clone := cloneYamlNode(mapping.YamlNode)
	if clone.Kind == yaml.DocumentNode {
		clone = clone.Content[0]
	}

	return NewMapping(clone)
}

// Returns the number of children.
func (mapping *Mapping) Length() int {
	return len(mapping.YamlNode.Content) / 2
}

// Returns the keys as a slice of [Scalar].
func (mapping *Mapping) Keys() []*Scalar {
	yamlNode := mapping.YamlNode
	contents := yamlNode.Content

	keys := make([]*Scalar, 0, len(contents)/2)

	for index := 0; index < len(contents); index = index + 2 {
		key := contents[index]
		keys = append(keys, NewScalar(key))
	}

	return keys
}

// Returns the key at given index as a [Scalar].
func (mapping *Mapping) KeyAt(i int) *Scalar {
	if i < 0 || i >= mapping.Length() {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", i, mapping.Length()))
	}

	key := mapping.YamlNode.Content[i*2]
	return NewScalar(key)
}

// Returns true if the key exists in the keys list.
func (mapping *Mapping) ContainsKey(key string) bool {
	for _, iterator := range mapping.Keys() {
		if iterator.GetValue() == key {
			return true
		}
	}

	return false
}

// Sets the value of the given key to the given [Mapping].
func (mapping *Mapping) SetMapping(key string, value *Mapping) {
	if value == nil {
		js.Throw(fmt.Errorf("mapping to add can't be nil"))
	}

	index := mapping.IndexOfKey(key)
	if index >= 0 {
		mapping.setKeyValue(index, value.YamlNode)
	} else {
		mapping.YamlNode.Content = append(mapping.YamlNode.Content, NewYamlScalarNode(key), value.YamlNode)
	}
}

// Sets the value of the given key to the given [Sequence].
func (mapping *Mapping) SetSequence(key string, value *Sequence) {
	if value == nil {
		js.Throw(fmt.Errorf("sequence to add can't be nil"))
	}

	index := mapping.IndexOfKey(key)
	if index >= 0 {
		mapping.setKeyValue(index, value.YamlNode)
	} else {
		mapping.YamlNode.Content = append(mapping.YamlNode.Content, NewYamlScalarNode(key), value.YamlNode)
	}
}

// Sets the value of the given key to the given [Scalar].
func (mapping *Mapping) SetScalar(key string, value *Scalar) {
	if value == nil {
		js.Throw(fmt.Errorf("scalar to add can't be nil"))
	}

	index := mapping.IndexOfKey(key)
	if index >= 0 {
		mapping.setKeyValue(index, value.YamlNode)
	} else {
		mapping.YamlNode.Content = append(mapping.YamlNode.Content, NewYamlScalarNode(key), value.YamlNode)
	}
}

// Sets the value of the given key to the given string as a [Scalar] and returns the created [Scalar] for value.
func (mapping *Mapping) SetValue(key string, value string) *Scalar {
	scalarNode := NewYamlScalarNode(value)
	scalar := NewScalar(scalarNode)

	mapping.SetScalar(key, scalar)

	return scalar
}

// Sets the value of the given key to the given int as a [Scalar] and returns the created [Scalar] for value.
func (mapping *Mapping) SetValueInt(key string, value int) *Scalar {
	scalarNode := NewYamlScalarNode(fmt.Sprintf("%d", value))
	scalar := NewScalar(scalarNode)

	mapping.SetScalar(key, scalar)

	return scalar
}

// Sets the value of the given key to the given float as a [Scalar] and returns the created [Scalar] for value.
func (mapping *Mapping) SetValueFloat(key string, value float64) *Scalar {
	scalarNode := NewYamlScalarNode(fmt.Sprintf("%f", value))
	scalar := NewScalar(scalarNode)

	mapping.SetScalar(key, scalar)

	return scalar
}

// Sets the value of the given key to the given bool as a [Scalar] and returns the created [Scalar] for value.
func (mapping *Mapping) SetValueBool(key string, value bool) *Scalar {
	scalarNode := NewYamlScalarNode(fmt.Sprintf("%t", value))
	scalar := NewScalar(scalarNode)

	mapping.SetScalar(key, scalar)

	return scalar
}

// Inserts the given key value pair at the given index.
func (mapping *Mapping) InsertMapping(index int, key string, value *Mapping) {
	if value == nil {
		js.Throw(fmt.Errorf("mapping to add can't be nil"))
	}

	mapping.insertKeyValue(index, key, value.YamlNode)
}

// Inserts the given key value pair at the given index.
func (mapping *Mapping) InsertSequence(index int, key string, value *Sequence) {
	if value == nil {
		js.Throw(fmt.Errorf("sequence to add can't be nil"))
	}

	mapping.insertKeyValue(index, key, value.YamlNode)
}

// Inserts the given key value pair at the given index.
func (mapping *Mapping) InsertScalar(index int, key string, value *Scalar) {
	if value == nil {
		js.Throw(fmt.Errorf("scalar to add can't be nil"))
	}

	mapping.insertKeyValue(index, key, value.YamlNode)
}

// Inserts the given key value pair at the given index and returns the created [Scalar].
func (mapping *Mapping) InsertValue(index int, key string, value string) *Scalar {
	valueNode := NewYamlScalarNode(value)
	scalar := NewScalar(valueNode)

	mapping.insertKeyValue(index, key, scalar.YamlNode)

	return scalar
}

// Returns the index of the key, or -1 if it doesn't exist.
func (mapping *Mapping) IndexOfKey(key string) int {
	index := 0

	for _, iterator := range mapping.Keys() {
		if iterator.GetValue() == key {
			return index
		}

		index = index + 1
	}

	return -1
}

// Moves the key and the corresponding value to the given index.
func (mapping *Mapping) Move(key string, index int) {
	if index < 0 || index > mapping.Length() {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", index, mapping.Length()))
	}

	currentIndex := mapping.IndexOfKey(key)
	if currentIndex < 0 {
		js.Throw(fmt.Errorf("key %s not found", key))
	}

	if currentIndex == index {
		return
	}

	keyNode := mapping.YamlNode.Content[currentIndex*2]
	valueNode := mapping.YamlNode.Content[currentIndex*2+1]

	mapping.YamlNode.Content = slices.Delete(mapping.YamlNode.Content, currentIndex*2, currentIndex*2+2)
	mapping.YamlNode.Content = slices.Insert(mapping.YamlNode.Content, index*2, keyNode, valueNode)
}

// Removes the key and corresponding value.
func (mapping *Mapping) Remove(key string) {
	yamlNode := mapping.YamlNode
	contents := yamlNode.Content

	for index := 0; index < len(contents); index = index + 2 {
		contentKey := contents[index]

		if contentKey.Value == key {
			yamlNode.Content = slices.Delete(contents, index, index+2)
			break
		}
	}
}

// Clears all the contents of the mapping.
func (mapping *Mapping) Clear() {
	mapping.YamlNode.Content = nil
}

// Returns the [Mapping] corresponding to the given key. Returns nil if the key doesn't exist
// or it doesn't correspond to a [Mapping].
func (mapping *Mapping) GetMapping(key string) *Mapping {
	yamlNode := mapping.YamlNode
	contents := yamlNode.Content

	for index := 0; index < len(contents); index = index + 2 {
		contentKey := contents[index]
		contentValue := contents[index+1]

		if contentKey.Value == key {
			if contentValue.Kind != yaml.MappingNode {
				return nil
			}

			return NewMapping(contentValue)
		}
	}

	return nil
}

// Returns the [Sequence] corresponding to the given key. Returns nil if the key doesn't exist
// or it doesn't correspond to a [Sequence].
func (mapping *Mapping) GetSequence(key string) *Sequence {
	yamlNode := mapping.YamlNode
	contents := yamlNode.Content

	for index := 0; index < len(contents); index = index + 2 {
		contentKey := contents[index]
		contentValue := contents[index+1]

		if contentKey.Value == key {
			if contentValue.Kind != yaml.SequenceNode {
				return nil
			}

			return NewSequence(contentValue)
		}
	}

	return nil
}

// Returns the [Scalar] corresponding to the given key. Returns nil if the key doesn't exist
// or it doesn't correspond to a [Scalar].
func (mapping *Mapping) GetScalar(key string) *Scalar {
	yamlNode := mapping.YamlNode
	contents := yamlNode.Content

	for index := 0; index < len(contents); index = index + 2 {
		contentKey := contents[index]
		contentValue := contents[index+1]

		if contentKey.Value == key {
			if contentValue.Kind != yaml.ScalarNode {
				return nil
			}

			return NewScalar(contentValue)
		}
	}

	return nil
}

// Returns the value corresponding to the given key. Returns nil if the key doesn't exist.
// Throws if the key doesn't correspond to a [Scalar].
func (mapping *Mapping) GetValue(key string) *string {
	scalar := mapping.GetScalar(key)
	if scalar == nil {
		return nil
	}

	return Pointer(scalar.GetValue())
}

// Returns a [Mapping] by following each one of the keys. Expects that each node is a [Mapping], otherwise throws.
// Returns nil if the key doesn't exist at any point in the chain.
func (mapping *Mapping) GetMappingChain(keys ...string) *Mapping {
	if len(keys) == 0 {
		js.Throw(fmt.Errorf("empty keys passed"))
	}

	var result *Mapping = nil

	for _, key := range keys {
		if result == nil {
			result = mapping.GetMapping(key)
		} else {
			result = result.GetMapping(key)
		}

		if result == nil {
			break
		}
	}

	return result
}

// Returns a [Sequence] by following each one of the keys. Expects that each intermediate node is a [Mapping]
// and the last node is a [Sequence], otherwise throws. Returns nil if a key doesn't exist at any point in the chain.
func (mapping *Mapping) GetSequenceChain(keys ...string) *Sequence {
	if len(keys) == 0 {
		js.Throw(fmt.Errorf("empty keys passed"))
	}

	if len(keys) == 1 {
		return mapping.GetSequence(keys[0])
	}

	node := mapping.GetMappingChain(keys[:len(keys)-1]...)
	if node == nil {
		return nil
	}

	return node.GetSequence(keys[len(keys)-1])
}

// Returns a [Scalar] by following each one of the keys. Expects that each intermediate node is a [Mapping]
// and the last node is a [Scalar], otherwise throws. Returns nil if the key doesn't exist at any point in the chain.
func (mapping *Mapping) GetScalarChain(keys ...string) *Scalar {
	if len(keys) == 0 {
		js.Throw(fmt.Errorf("empty keys passed"))
	}

	if len(keys) == 1 {
		return mapping.GetScalar(keys[0])
	}

	node := mapping.GetMappingChain(keys[:len(keys)-1]...)
	if node == nil {
		return nil
	}

	return node.GetScalar(keys[len(keys)-1])
}

// Returns the resulting value by following each one of the keys. Expects that each intermediate node is a [Mapping]
// and the last node is a [Scalar], otherwise throws. Returns nil if the key doesn't exist at any point in the chain.
func (mapping *Mapping) GetValueChain(keys ...string) *string {
	if len(keys) == 0 {
		js.Throw(fmt.Errorf("empty keys passed"))
	}

	if len(keys) == 1 {
		return mapping.GetValue(keys[0])
	}

	node := mapping.GetMappingChain(keys[:len(keys)-1]...)
	if node == nil {
		return nil
	}

	return node.GetValue(keys[len(keys)-1])
}

// Ensures that the given key corresponds to a [Mapping]. Returns immediately if the key already points to a [Mapping].
// Creates an empty [Mapping] if the key doesn't exist or it exists but correspond to an empty [Scalar]. Throws otherwise.
func (mapping *Mapping) EnsureMapping(key string) *Mapping {
	result := mapping.GetMapping(key)
	if result != nil {
		return result
	}

	value := mapping.GetValue(key)
	if value != nil && *value == "" {
		slog.Debug("replacing empty scalar with mapping for ${key}", slog.String("key", key))
		mapping.Remove(key)
	}

	if mapping.ContainsKey(key) {
		js.Throw(fmt.Errorf("key %s already exists and is not a mapping", key))
	}

	result = NewMapping(NewYamlMappingNode())
	mapping.SetMapping(key, result)

	return result
}

// Ensures that the given key corresponds to a [Sequence]. Returns immediately if the key already points to a [Sequence].
// Creates an empty [Sequence] if the key doesn't exist or it exists but correspond to an empty [Scalar]. Throws otherwise.
func (mapping *Mapping) EnsureSequence(key string) *Sequence {
	result := mapping.GetSequence(key)
	if result != nil {
		return result
	}

	value := mapping.GetValue(key)
	if value != nil && *value == "" {
		slog.Debug("replacing empty scalar with sequence for ${key}", slog.String("key", key))
		mapping.Remove(key)
	}

	if mapping.ContainsKey(key) {
		js.Throw(fmt.Errorf("key %s already exists and is not a sequence", key))
	}

	result = NewSequence(NewYamlSequenceNode())
	mapping.SetSequence(key, result)

	return result
}

// Ensures that the last key corresponds to a [Mapping]. Expects that each intermediate node is a [Mapping].
// Returns immediately if the key already points to a [Mapping]. Creates an empty [Mapping] if the key doesn't
// exist or it exists but correspond to an empty [Scalar]. Throws otherwise.
func (mapping *Mapping) EnsureMappingChain(keys ...string) *Mapping {
	if len(keys) == 0 {
		js.Throw(fmt.Errorf("empty keys passed"))
	}

	iterator := mapping

	for _, key := range keys {
		iterator = iterator.EnsureMapping(key)
	}

	return iterator
}

// Ensures that the last key corresponds to a [Sequence]. Expects that each intermediate node is a [Mapping].
// Returns immediately if the key already points to a [Sequence]. Creates an empty [Sequence] if the key doesn't
// exist or it exists but correspond to an empty [Scalar]. Throws otherwise.
func (mapping *Mapping) EnsureSequenceChain(keys ...string) *Sequence {
	if len(keys) == 0 {
		js.Throw(fmt.Errorf("empty keys passed"))
	}

	iterator := mapping.EnsureMappingChain(keys[:len(keys)-1]...)
	return iterator.EnsureSequence(keys[len(keys)-1])
}

// Sorts the child nodes by their keys.
func (mapping *Mapping) SortByKey() {
	sort.Stable(mapping)
}

func (x *Mapping) Less(i, j int) bool {
	return x.YamlNode.Content[i*2].Value < x.YamlNode.Content[j*2].Value
}

func (x *Mapping) Len() int {
	return len(x.YamlNode.Content) / 2
}

func (x *Mapping) Swap(i, j int) {
	content := x.YamlNode.Content

	content[i*2], content[j*2] = content[j*2], content[i*2]
	content[i*2+1], content[j*2+1] = content[j*2+1], content[i*2+1]
}

func (mapping *Mapping) setKeyValue(indexOfKey int, node *yaml.Node) {
	// Content indexes are not the same as the indexes of the keys. Yaml nodes keep the keys and values
	// as a one contiguous array and even indexes are used for keys and odd indexes are used for values.
	mapping.YamlNode.Content[2*indexOfKey+1] = node
}

func (mapping *Mapping) insertKeyValue(indexOfKey int, key string, node *yaml.Node) {
	if indexOfKey < 0 || indexOfKey > mapping.Length() {
		js.Throw(fmt.Errorf("index %d is out of bounds, length: %d", indexOfKey, mapping.Length()))
	}

	yaml := mapping.YamlNode
	keyNode := NewYamlScalarNode(key)

	// See the comment above for indexes.
	yaml.Content = slices.Insert(yaml.Content, 2*indexOfKey, keyNode, node)
}

func jsToMapping(jsRuntime *js.JsRuntime, jsValue sobek.Value) (*Mapping, error) {
	// Marshal the object to a map[string]any, but iterate over the keys of the JS object
	// to ensure we preserve the order of the keys.
	objectMarshalled, err := jsRuntime.MarshalToGo(jsValue, reflect.TypeFor[map[string]any]())
	if err != nil {
		return nil, fmt.Errorf("failed to convert JS value to map[string]any: %w", err)
	}

	mapping := NewEmptyMapping()
	object := jsRuntime.ToSobekObject(jsValue)
	objectMap := objectMarshalled.Interface().(map[string]any)

	for _, key := range object.Keys() {
		value := objectMap[key]
		if value == nil {
			continue
		}

		if scalar := tryGetScalar(jsRuntime, jsRuntime.ToSobekValue(value)); scalar != nil {
			mapping.SetScalar(key, scalar)
			continue
		}

		seq, err := jsToSequence(jsRuntime, jsRuntime.ToSobekValue(value))
		if err == nil {
			mapping.SetSequence(key, seq)
			continue
		}

		child, err := jsToMapping(jsRuntime, jsRuntime.ToSobekValue(value))
		if err == nil {
			mapping.SetMapping(key, child)
			continue
		}

		return nil, fmt.Errorf("unsupported type %T", value)
	}

	return mapping, nil
}

func (m *Mapping) GetKeys(jsRuntime *js.JsRuntime) []string {
	mappingKeys := m.Keys()
	keys := make([]string, 0, m.Length())

	for _, key := range mappingKeys {
		keys = append(keys, key.GetValue())
	}

	return keys
}

func (m *Mapping) Get(jsRuntime *js.JsRuntime, key string) (any, bool) {
	mapping := m.GetMapping(key)
	if mapping != nil {
		return mapping, true
	}

	sequence := m.GetSequence(key)
	if sequence != nil {
		return sequence, true
	}

	scalar := m.GetScalar(key)
	if scalar != nil {
		return scalar.GetValue(), true
	}

	return nil, false
}

func (m *Mapping) Set(jsRuntime *js.JsRuntime, key string, value sobek.Value) bool {
	mapping, err := jsToMapping(jsRuntime, value)
	if err == nil {
		m.SetMapping(key, mapping)
		return true
	}

	sequence, err := jsToSequence(jsRuntime, value)
	if err == nil {
		m.SetSequence(key, sequence)
		return true
	}

	if scalar := tryGetScalar(jsRuntime, value); scalar != nil {
		m.SetScalar(key, scalar)
		return true
	}

	return false
}

func (mapping *Mapping) ToJSON(jsRuntime *js.JsRuntime, dummy string) sobek.Value {
	object := jsRuntime.Runtime.NewObject()

	for _, keyScalar := range mapping.Keys() {
		key := keyScalar.GetValue()

		if childMapping := mapping.GetMapping(key); childMapping != nil {
			object.Set(key, childMapping.ToJSON(jsRuntime, dummy))
			continue
		}

		if childSequence := mapping.GetSequence(key); childSequence != nil {
			object.Set(key, childSequence.ToJSON(jsRuntime, dummy))
			continue
		}

		if childScalar := mapping.GetScalar(key); childScalar != nil {
			object.Set(key, childScalar.ToJSON(jsRuntime, dummy))
			continue
		}

		js.Throw(fmt.Errorf("key %s does not correspond to a mapping, sequence or scalar", key))
	}

	return object
}

func registerYamlMapping(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Mapping]()).Methods(
		js.Method("Clear"),
		js.Method("Clone"),
		js.Method("ContainsKey"),
		js.Method("EnsureMapping"),
		js.Method("EnsureMappingChain").JsName("ensureMapping"),
		js.Method("EnsureSequence"),
		js.Method("EnsureSequenceChain").JsName("ensureSequence"),
		js.Method("GetMapping"),
		js.Method("GetMappingChain").JsName("getMapping"),
		js.Method("GetScalar"),
		js.Method("GetScalarChain").JsName("getScalar"),
		js.Method("GetSequence"),
		js.Method("GetSequenceChain").JsName("getSequence"),
		js.Method("GetValue"),
		js.Method("GetValueChain").JsName("getValue"),
		js.Method("IndexOfKey"),
		js.Method("InsertMapping").JsName("insert"),
		js.Method("InsertScalar").JsName("insert"),
		js.Method("InsertSequence").JsName("insert"),
		js.Method("InsertValue").JsName("insert"),
		js.Method("KeyAt"),
		js.Method("Keys"),
		js.Method("Length"),
		js.Method("Move"),
		js.Method("Remove"),
		js.Method("SetMapping").JsName("set"),
		js.Method("SetScalar").JsName("set"),
		js.Method("SetSequence").JsName("set"),
		js.Method("SetValue").JsName("set"),
		js.Method("SetValueInt").JsName("set"),
		js.Method("SetValueFloat").JsName("set"),
		js.Method("SetValueBool").JsName("set"),
		js.Method("SortByKey"),
		js.Method("ToJSON"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewEmptyMapping)),
		js.Constructor(reflect.ValueOf(jsToMapping)),
	).TypeConversion(
		reflect.ValueOf(jsToMapping),
	).DisableObjectMapping()
}
