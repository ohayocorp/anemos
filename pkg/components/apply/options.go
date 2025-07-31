package apply

import "github.com/ohayocorp/anemos/pkg/core"

type Options struct {
	Documents               []*core.Document
	ApplySetParentName      string
	ApplySetParentNamespace string
	SkipConfirmation        bool
}

func NewOptions() *Options {
	return &Options{}
}

func NewOptionsWithDocumentGroup(documentGroup *core.DocumentGroup) *Options {
	return &Options{
		Documents:          documentGroup.Documents,
		ApplySetParentName: core.ToKubernetesIdentifier(documentGroup.Path),
	}
}

func NewOptionsWithDocumentGroupAndNamespace(documentGroup *core.DocumentGroup, namespace string) *Options {
	return &Options{
		Documents:               documentGroup.Documents,
		ApplySetParentName:      core.ToKubernetesIdentifier(documentGroup.Path),
		ApplySetParentNamespace: namespace,
	}
}

func NewOptionsWithDocumentsAndName(documents []*core.Document, name string) *Options {
	return &Options{
		Documents:          documents,
		ApplySetParentName: name,
	}
}

func NewOptionsWithDocumentsNameAndNamespace(documents []*core.Document, name string, namespace string) *Options {
	return &Options{
		Documents:               documents,
		ApplySetParentName:      name,
		ApplySetParentNamespace: namespace,
	}
}
