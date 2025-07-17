package core

import (
	"fmt"
	"path"
	"reflect"

	"github.com/ohayocorp/anemos/pkg/js"
	"gopkg.in/yaml.v3"
)

// Document corresponds to a single YAML document. Note that even though a YAML file can contain multiple documents,
// each one of these documents is represented by a separate Document object.
//
// Document objects can be created with the methods below. These methods ensure that the kind of the yaml node is [yaml.DocumentNode].
//   - [NewDocument]
//   - [NewEmptyDocument]
//   - [ParseDocument]
//   - [ParseTemplateAsDocument]
//
// Although the root of the document can be any kind of node, only [Mapping] is supported.
type Document struct {
	Path  string
	Group *DocumentGroup

	YamlNode *yaml.Node
	root     *Mapping
}

// Retuns a new [Document] with the root set to an empty mapping.
func NewEmptyDocument(path string) *Document {
	yamlNode := NewYamlDocumentNode()
	root := NewEmptyMapping()

	document := &Document{
		Path:     path,
		YamlNode: yamlNode,
		root:     root,
	}

	document.YamlNode.Content = []*yaml.Node{
		root.YamlNode,
	}

	return document
}

// Retuns a new [Document] with the root set to the given mapping.
func NewDocumentWithRoot(path string, root *Mapping) *Document {
	yamlNode := NewYamlDocumentNode()

	document := &Document{
		Path:     path,
		YamlNode: yamlNode,
		root:     root,
	}

	document.YamlNode.Content = []*yaml.Node{
		root.YamlNode,
	}

	return document
}

// Retuns a new [Document] by parsing the given YAML.
func NewDocumentWithYaml(path string, yamlContent string) *Document {
	return NewDocumentWithRoot(path, ParseMapping(yamlContent))
}

// Retuns a new [Document] with the content set to the given yaml node. Panics if the
// yaml node is nil or its type is not [yaml.DocumentNode].
func NewDocument(path string, yamlNode *yaml.Node) *Document {
	if yamlNode == nil {
		panic(fmt.Errorf("passed yaml node is nil"))
	}

	if yamlNode.Kind != yaml.DocumentNode {
		panic(fmt.Errorf("yaml node is not document, but %s", getYamlNodeKind(yamlNode)))
	}

	root := NewMapping(yamlNode.Content[0])

	return &Document{
		Path:     path,
		YamlNode: yamlNode,
		root:     root,
	}
}

// Retuns a new [Document] with the content set to an empty yaml node with type [yaml.DocumentNode].
func NewYamlDocumentNode() *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.DocumentNode,
		Style: yaml.TaggedStyle,
	}
}

// Returns the path to write the document. Adds group name as base directory if it is not nil.
func (document *Document) FullPath() string {
	if document.Group == nil {
		return document.Path
	}

	return path.Join(document.Group.Name, document.Path)
}

// Returns a clone of the document.
func (document *Document) Clone() *Document {
	clone := cloneYamlNode(document.YamlNode)
	return NewDocument(document.Path, clone)
}

// Return the root of the document as a [Mapping].
func (document *Document) GetRoot() *Mapping {
	return document.root
}

func registerYamlDocument(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Document]()).Fields(
		js.Field("Path"),
		js.Field("Group"),
	).Methods(
		js.Method("Clone"),
		js.Method("FullPath"),
		js.Method("GetRoot"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewEmptyDocument)),
		js.Constructor(reflect.ValueOf(NewDocumentWithRoot)),
		js.Constructor(reflect.ValueOf(NewDocumentWithYaml)),
	).DisableObjectMapping()
}
