package core

import (
	"fmt"
	"reflect"
	"slices"
	"sort"

	"github.com/ohayocorp/anemos/pkg/js"
)

// BuildContext provides all the necessary objects and options to generate documents when [Builder.Build] is called.
type BuildContext struct {
	BuilderOptions         *BuilderOptions
	KubernetesResourceInfo *KubernetesResourceInfo
	CustomData             map[string]any

	builder          *Builder
	documentGroups   map[*Component][]*DocumentGroup
	currentComponent *Component
}

// Adds given document to the document group named "". Creates the document group if it doesn't exist.
func (context *BuildContext) AddDocument(document *Document) {
	if document == nil {
		js.Throw(fmt.Errorf("document cannot be nil"))
	}

	documentGroup := context.GetDocumentGroupWithName("")
	if documentGroup == nil {
		documentGroup = NewDocumentGroup("")
		context.AddDocumentGroup(documentGroup)
	}

	documentGroup.AddDocument(document)
}

func (context *BuildContext) AddDocumentParse(path string, yamlContent string) {
	document := NewDocumentWithYaml(path, yamlContent)
	context.AddDocument(document)
}

func (context *BuildContext) AddDocumentMapping(path string, root *Mapping) {
	document := NewDocumentWithRoot(path, root)
	context.AddDocument(document)
}

// Adds given group to the document groups list.
func (context *BuildContext) AddDocumentGroup(group *DocumentGroup) {
	context.documentGroups[context.currentComponent] = append(context.documentGroups[context.currentComponent], group)
	group.component = context.currentComponent
}

// Adds given additional file to the document group named "". Creates the document group if it doesn't exist.
func (context *BuildContext) AddAdditionalFile(additionalFile *AdditionalFile) {
	if additionalFile == nil {
		js.Throw(fmt.Errorf("additionalFile cannot be nil"))
	}

	documentGroup := context.GetDocumentGroupWithName("")
	if documentGroup == nil {
		documentGroup = NewDocumentGroup("")
		context.AddDocumentGroup(documentGroup)
	}

	documentGroup.AddAdditionalFile(additionalFile)
}

// Removes given group from the document groups list.
func (context *BuildContext) RemoveDocumentGroup(group *DocumentGroup) {
	for component, groups := range context.documentGroups {
		context.documentGroups[component] = slices.DeleteFunc(groups, func(dg *DocumentGroup) bool {
			return dg == group
		})
	}
}

// Returns all documents inside all document groups as a slice.
func (context *BuildContext) GetAllDocuments() []*Document {
	var documents []*Document

	for _, documentGroups := range context.documentGroups {
		for _, documentGroup := range documentGroups {
			documents = append(documents, documentGroup.Documents...)
		}
	}

	return documents
}

// Returns all documents inside all document groups sorted by their file path as a slice.
func (context *BuildContext) GetAllDocumentsSorted() []*Document {
	allDocuments := context.GetAllDocuments()

	sort.SliceStable(allDocuments, func(i, j int) bool {
		return allDocuments[i].FullPath() < allDocuments[j].FullPath()
	})

	return allDocuments
}

// Returns the first document that satisfies the given predicate. Returns nil if no document is found.
func (context *BuildContext) GetDocument(predicate func(*Document) bool) *Document {
	for _, document := range context.GetAllDocuments() {
		if predicate(document) {
			return document
		}
	}

	return nil
}

// Returns the first document that has the given path. Returns nil if no document is found.
func (context *BuildContext) GetDocumentWithPath(path string) *Document {
	for _, document := range context.GetAllDocuments() {
		if document.FullPath() == path {
			return document
		}
	}

	return nil
}

func (context *BuildContext) GetDocumentGroups() []*DocumentGroup {
	var documentGroups []*DocumentGroup

	for _, r := range context.documentGroups {
		documentGroups = append(documentGroups, r...)
	}

	return documentGroups
}

func (context *BuildContext) GetDocumentGroupWithName(name string) *DocumentGroup {
	for _, r := range context.documentGroups {
		for _, documentGroup := range r {
			if documentGroup.Name == name {
				return documentGroup
			}
		}
	}

	return nil
}

func (context *BuildContext) GetDocumentGroupsForComponent(component *Component) []*DocumentGroup {
	return context.documentGroups[component]
}

func (context *BuildContext) GetCurrentComponent() *Component {
	return context.currentComponent
}

func (context *BuildContext) GetAllComponents() []*Component {
	// Return a copy of the components slice to avoid modification
	// by the caller.
	components := make([]*Component, len(context.builder.Components))
	copy(components, context.builder.Components)

	return components
}

func (context *BuildContext) GetComponentWithIdentifier(identifier string) *Component {
	for _, component := range context.builder.Components {
		componentIdentifier := component.GetIdentifier()
		if componentIdentifier != nil && *componentIdentifier == identifier {
			return component
		}
	}

	return nil
}

func registerBuildContext(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[BuildContext]()).Fields(
		js.Field("BuilderOptions"),
		js.Field("KubernetesResourceInfo"),
		js.Field("CustomData"),
	).Methods(
		js.Method("AddDocument"),
		js.Method("AddDocumentParse").JsName("addDocument"),
		js.Method("AddDocumentMapping").JsName("addDocument"),
		js.Method("AddDocumentGroup"),
		js.Method("AddAdditionalFile"),
		js.Method("GetAllDocuments"),
		js.Method("GetAllDocumentsSorted"),
		js.Method("GetDocumentGroups"),
		js.Method("GetDocumentGroupWithName").JsName("getDocumentGroup"),
		js.Method("GetDocumentGroupsForComponent").JsName("getDocumentGroups"),
		js.Method("GetDocument"),
		js.Method("GetDocumentWithPath").JsName("getDocument"),
		js.Method("RemoveDocumentGroup"),
	)
}
