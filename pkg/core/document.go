package core

import (
	"fmt"
	"path"
	"reflect"
	"slices"
	"strings"

	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg/js"
	"github.com/ohayocorp/anemos/pkg/util"
)

type NewDocumentOptions struct {
	Yaml          *string
	Object        *sobek.Object
	Path          *string
	DocumentGroup *string
}

type Document struct {
	*sobek.Object

	path         *string
	Group        *DocumentGroup
	Dependencies *Dependencies[*Document]
}

func NewNewDocumentOptions() *NewDocumentOptions {
	return &NewDocumentOptions{}
}

func NewDocument(jsRuntime *js.JsRuntime) *Document {
	content := jsRuntime.Runtime.NewObject()
	return NewDocumentWithContent(content)
}

func NewDocumentWithYaml(jsRuntime *js.JsRuntime, yaml string) (*Document, error) {
	content, err := Parse(jsRuntime, yaml)
	if err != nil {
		return nil, err
	}

	return NewDocumentWithContent(content), nil
}

func NewDocumentWithContent(content *sobek.Object) *Document {
	document := &Document{
		Object:       content,
		Dependencies: NewDependencies[*Document](),
	}

	return document
}

func NewDocumentWithOptions(jsRuntime *js.JsRuntime, options *NewDocumentOptions) (*Document, error) {
	if options == nil {
		return nil, fmt.Errorf("options cannot be nil")
	}

	if options.Yaml == nil && options.Object == nil {
		return nil, fmt.Errorf("content must be specified")
	}

	var document *Document
	var err error

	if options.Yaml != nil {
		document, err = NewDocumentWithYaml(jsRuntime, *options.Yaml)
		if err != nil {
			return nil, err
		}
	} else if options.Object != nil {
		document = NewDocumentWithContent(options.Object)
	}

	document.SetPath(options.Path)

	return document, nil
}

func SobekObjectGetString(object *sobek.Object, key string) *string {
	value := object.Get(key)
	if value == nil {
		return nil
	}

	result := value.String()

	return &result
}

func SobekObjectGetStringChain(object *sobek.Object, keys ...string) *string {
	var property *sobek.Object
	var ok bool

	for i, key := range keys {
		if i == len(keys)-1 {
			break
		}

		property, ok = object.Get(key).(*sobek.Object)
		if !ok || property == nil {
			return nil
		}
	}

	return SobekObjectGetString(property, keys[len(keys)-1])
}

// Returns the file path of the document. May contain multiple segments separated by slashes.
func (document *Document) GetPath() string {
	path := document.path
	if path == nil {
		kind := SobekObjectGetString(document.Object, "kind")
		name := SobekObjectGetStringChain(document.Object, "metadata", "name")

		if kind != nil && name != nil {
			p := fmt.Sprintf("%s-%s.yaml", strings.ToLower(*kind), util.ToKubernetesIdentifier(*name))
			path = &p
		}
	}

	if path == nil {
		p := "document.yaml"

		if document.Group != nil {
			index := slices.Index(document.Group.Documents, document)
			p = fmt.Sprintf("document-%d.yaml", index+1)
		}

		path = &p
	}

	return *path
}

// Sets the file path of the document. May contain multiple segments separated by slashes.
func (document *Document) SetPath(path *string) {
	document.path = path
}

// Returns the path to write the document. Adds group path as base directory if it is not nil.
func (document *Document) FullPath() string {
	documentPath := document.GetPath()

	if document.Group == nil {
		return documentPath
	}

	return path.Join(document.Group.Path, documentPath)
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

func (document *Document) ToJSON(jsRuntime *js.JsRuntime, dummy string) sobek.Value {
	object := jsRuntime.Runtime.NewObject()

	object.Set("path", document.GetPath())
	object.Set("content", document.Object)

	return jsRuntime.Runtime.ToValue(object)
}

func registerDocument(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Document]()).JsModule(
		"document",
	).Fields(
		js.Field("Group"),
	).Methods(
		js.Method("GetPath"),
		js.Method("SetPath"),
		js.Method("FullPath"),
		js.Method("ProvisionAfter"),
		js.Method("ProvisionBefore"),
		js.Method("ToJSON"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewDocument)),
		js.Constructor(reflect.ValueOf(NewDocumentWithOptions)),
		js.Constructor(reflect.ValueOf(NewDocumentWithContent)),
		js.Constructor(reflect.ValueOf(NewDocumentWithYaml)),
	)

	jsRuntime.Type(reflect.TypeFor[NewDocumentOptions]()).JsModule(
		"document",
	).Fields(
		js.Field("DocumentGroup"),
		js.Field("Path"),
		js.Field("Yaml").JsName("content"),
		js.Field("Object").JsName("content"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewNewDocumentOptions)),
	)
}
