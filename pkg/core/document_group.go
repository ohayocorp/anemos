package core

import (
	"fmt"
	"reflect"
	"slices"
	"sort"

	"github.com/ohayocorp/anemos/pkg/js"
)

// DocumentGroup is a named container for multiple [Document] instances.
type DocumentGroup struct {
	Path             string
	Documents        []*Document
	AdditionalFiles  []*AdditionalFile
	ApplyProvisioner *Provisioner
	WaitProvisioner  *Provisioner

	component *Component
}

type AdditionalFile struct {
	Path    string
	Content string
}

// Creates a new [DocumentGroup] with given path.
func NewDocumentGroup(path string) *DocumentGroup {
	documentGroup := &DocumentGroup{
		Path:      path,
		Documents: make([]*Document, 0),
	}

	documentGroup.ApplyProvisioner = ApplyDocuments(documentGroup)
	documentGroup.WaitProvisioner = WaitDocuments(documentGroup)

	documentGroup.WaitProvisioner.RunAfter(documentGroup.ApplyProvisioner)

	return documentGroup
}

// Creates a new [AdditionalFile] with given path and content.
func NewAdditionalFile(path, content string) *AdditionalFile {
	return &AdditionalFile{
		Path:    path,
		Content: content,
	}
}

// Adds the given document to this group and sets its Group field to this group.
func (group *DocumentGroup) AddDocument(document *Document) {
	if document == nil {
		js.Throw(fmt.Errorf("document cannot be nil"))
		return
	}

	group.Documents = append(group.Documents, document)
	document.Group = group
}

// Adds the given documents to this group and sets their Group field to this group.
func (group *DocumentGroup) AddDocuments(documents []*Document) {
	for _, document := range documents {
		group.AddDocument(document)
	}
}

// Adds the given additional file to this group.
func (group *DocumentGroup) AddAdditionalFile(additionalFile *AdditionalFile) {
	group.AdditionalFiles = append(group.AdditionalFiles, additionalFile)
}

// Returns the component that created this document group. Component is set when the
// document group is added to the builder context.
func (group *DocumentGroup) GetComponent() *Component {
	return group.component
}

// Returns the first document that has the given path. Returns nil if no document is found.
func (group *DocumentGroup) GetDocument(path string) *Document {
	return group.GetDocumentFunc(func(document *Document) bool {
		return document.GetPath() == path
	})
}

// Returns the first document that satisfies the given predicate. Returns nil if no document is found.
func (group *DocumentGroup) GetDocumentFunc(predicate func(*Document) bool) *Document {
	for _, document := range group.Documents {
		if predicate(document) {
			return document
		}
	}

	return nil
}

// Returns the documents in this group sorted by their file path.
func (group *DocumentGroup) SortedDocuments() []*Document {
	sorted := make([]*Document, len(group.Documents))
	copy(sorted, group.Documents)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].GetPath() < sorted[j].GetPath()
	})

	return sorted
}

// Removes the given document from this group and sets its Group field to nil.
func (group *DocumentGroup) RemoveDocument(document *Document) {
	if document == nil {
		js.Throw(fmt.Errorf("document cannot be nil"))
		return
	}

	// Using slices.DeleteFunc modifies the original slice that can lead to unexpected behavior since
	// JavaScript objects may reference it. Create a new slice and copy the elements over.
	newDocuments := make([]*Document, 0, len(group.Documents))
	for _, d := range group.Documents {
		if d != document {
			newDocuments = append(newDocuments, d)
		}
	}
	group.Documents = newDocuments

	document.Group = nil
}

// Removes the given additional file from this group.
func (group *DocumentGroup) RemoveAdditionalFile(additionalFile *AdditionalFile) {
	group.AdditionalFiles = slices.DeleteFunc(group.AdditionalFiles, func(f *AdditionalFile) bool {
		return f == additionalFile
	})
}

// Removes all documents and additional files from this group and adds them to the given group.
func (group *DocumentGroup) MoveTo(other *DocumentGroup) {
	for _, document := range group.Documents {
		other.AddDocument(document)
	}

	for _, additionalFile := range group.AdditionalFiles {
		other.AddAdditionalFile(additionalFile)
	}

	group.Documents = nil
	group.AdditionalFiles = nil
}

func (documentGroup *DocumentGroup) ProvisionAfter(other *DocumentGroup) {
	if documentGroup.ApplyProvisioner != nil {
		if other.ApplyProvisioner != nil {
			documentGroup.ApplyProvisioner.RunAfter(other.ApplyProvisioner)
		}

		if documentGroup.WaitProvisioner != nil {
			documentGroup.ApplyProvisioner.RunAfter(other.WaitProvisioner)
		}
	}

	if documentGroup.WaitProvisioner != nil {
		if other.ApplyProvisioner != nil {
			documentGroup.WaitProvisioner.RunAfter(other.ApplyProvisioner)
		}

		if documentGroup.WaitProvisioner != nil {
			documentGroup.WaitProvisioner.RunAfter(other.WaitProvisioner)
		}
	}
}

func (documentGroup *DocumentGroup) ProvisionBefore(other *DocumentGroup) {
	if documentGroup.ApplyProvisioner != nil {
		if other.ApplyProvisioner != nil {
			documentGroup.ApplyProvisioner.RunBefore(other.ApplyProvisioner)
		}

		if documentGroup.WaitProvisioner != nil {
			documentGroup.ApplyProvisioner.RunBefore(other.WaitProvisioner)
		}
	}

	if documentGroup.WaitProvisioner != nil {
		if other.ApplyProvisioner != nil {
			documentGroup.WaitProvisioner.RunBefore(other.ApplyProvisioner)
		}

		if documentGroup.WaitProvisioner != nil {
			documentGroup.WaitProvisioner.RunBefore(other.WaitProvisioner)
		}
	}
}

func registerDocumentGroup(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[DocumentGroup]()).JsModule(
		"documentGroup",
	).Fields(
		js.Field("Path"),
		js.Field("Documents"),
		js.Field("AdditionalFiles"),
		js.Field("ApplyProvisioner"),
		js.Field("WaitProvisioner"),
	).Methods(
		js.Method("AddDocument"),
		js.Method("AddDocuments"),
		js.Method("AddAdditionalFile"),
		js.Method("GetComponent"),
		js.Method("GetDocument"),
		js.Method("GetDocumentFunc").JsName("getDocument"),
		js.Method("SortedDocuments"),
		js.Method("MoveTo"),
		js.Method("RemoveDocument"),
		js.Method("RemoveAdditionalFile"),
		js.Method("ProvisionAfter"),
		js.Method("ProvisionBefore"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewDocumentGroup)),
	)

	jsRuntime.Type(reflect.TypeFor[AdditionalFile]()).JsModule(
		"documentGroup",
	).Fields(
		js.Field("Path"),
		js.Field("Content"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewAdditionalFile)),
	)
}
