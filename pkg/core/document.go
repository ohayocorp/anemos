package core

import (
	"fmt"
	"path"
	"reflect"

	"github.com/grafana/sobek"
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
	Path         string
	Group        *DocumentGroup
	Dependencies *Dependencies[*Document]

	YamlNode *yaml.Node
	root     *Mapping
}

// Retuns a new [Document] with the root set to an empty mapping.
func NewEmptyDocument(path string) *Document {
	root := NewEmptyMapping()
	return NewDocumentWithRoot(path, root)
}

// Retuns a new [Document] with the root set to the given mapping.
func NewDocumentWithRoot(path string, root *Mapping) *Document {
	yamlNode := NewYamlDocumentNode()

	document := &Document{
		Path:         path,
		YamlNode:     yamlNode,
		Dependencies: NewDependencies[*Document](),
		root:         root,
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

// Retuns a new [Document] with the given content. Doesn't act on document group, it just creates the document.
func NewDocumentWithOptions(options *AddDocumentOptions) *Document {
	if options == nil {
		js.Throw(fmt.Errorf("options cannot be nil"))
	}

	if options.Path == "" {
		js.Throw(fmt.Errorf("path cannot be empty"))
	}

	if options.Root == nil && options.Yaml == nil && options.Object == nil {
		js.Throw(fmt.Errorf("content must be specified"))
	}

	var document *Document

	if options.Root != nil {
		document = NewDocumentWithRoot(options.Path, options.Root)
	} else if options.Yaml != nil {
		document = NewDocumentWithYaml(options.Path, *options.Yaml)
	} else if options.Object != nil {
		yaml := SerializeToYaml(options.Object)
		document = NewDocumentWithYaml(options.Path, yaml)
	}

	return document
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

	return NewDocumentWithRoot(path, root)
}

// Retuns a new [Document] with the content set to an empty yaml node with type [yaml.DocumentNode].
func NewYamlDocumentNode() *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.DocumentNode,
		Style: yaml.TaggedStyle,
	}
}

// Returns the path to write the document. Adds group path as base directory if it is not nil.
func (document *Document) FullPath() string {
	if document.Group == nil {
		return document.Path
	}

	return path.Join(document.Group.Path, document.Path)
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

func (document *Document) ProvisionAfter(other *Document) {
	if document.Group == nil || other.Group != document.Group {
		js.Throw(fmt.Errorf("cannot set provision order for documents that are not in the same group"))
	}

	document.Dependencies.RunAfter(other)
}

func (document *Document) ProvisionBefore(other *Document) {
	if document.Group == nil || other.Group != document.Group {
		js.Throw(fmt.Errorf("cannot set provision order for documents that are not in the same group"))
	}

	document.Dependencies.RunBefore(other)
}

func (document *Document) Get(jsRuntime *js.JsRuntime, key string) any {
	return document.root.Get(jsRuntime, key)
}

func (document *Document) Set(jsRuntime *js.JsRuntime, key string, value sobek.Value) bool {
	return document.root.Set(jsRuntime, key, value)
}

func registerYamlDocument(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Document]()).Fields(
		js.Field("Path"),
		js.Field("Group"),
	).Methods(
		js.Method("Clone"),
		js.Method("FullPath"),
		js.Method("GetRoot"),
		js.Method("ProvisionAfter"),
		js.Method("ProvisionBefore"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewEmptyDocument)),
		js.Constructor(reflect.ValueOf(NewDocumentWithRoot)),
		js.Constructor(reflect.ValueOf(NewDocumentWithYaml)),
		js.Constructor(reflect.ValueOf(NewDocumentWithOptions)),
	).DisableObjectMapping()
}
