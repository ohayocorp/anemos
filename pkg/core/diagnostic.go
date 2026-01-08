package core

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/js"
)

type DiagnosticSeverity string
type DiagnosticCategory string

const (
	DiagnosticSeverityInfo    DiagnosticSeverity = "info"
	DiagnosticSeverityWarning DiagnosticSeverity = "warning"
	DiagnosticSeverityError   DiagnosticSeverity = "error"

	DiagnosticCategoryLinting  DiagnosticCategory = "linting"
	DiagnosticCategorySecurity DiagnosticCategory = "security"
	DiagnosticCategorySpecs    DiagnosticCategory = "specs"
)

type DiagnosticMetadata struct {
	Id          string
	Name        string
	Description string
	Severity    DiagnosticSeverity
	Categories  []DiagnosticCategory
}

type Diagnostic struct {
	Metadata DiagnosticMetadata
	Message  string
	Document *Document

	component *Component
}

func NewDiagnosticMetadata(id string, name string, description string, severity DiagnosticSeverity, categories []DiagnosticCategory) *DiagnosticMetadata {
	return &DiagnosticMetadata{
		Id:          id,
		Name:        name,
		Description: description,
		Severity:    severity,
		Categories:  categories,
	}
}

func NewDiagnostic(metadata *DiagnosticMetadata, message string) *Diagnostic {
	return &Diagnostic{
		Metadata: *metadata,
		Message:  message,
	}
}

func NewDiagnosticWithDocument(metadata *DiagnosticMetadata, message string, document *Document) *Diagnostic {
	return &Diagnostic{
		Metadata: *metadata,
		Message:  message,
		Document: document,
	}
}

func registerDiagnostic(jsRuntime *js.JsRuntime) {
	jsRuntime.Variable("diagnostic", "info", reflect.ValueOf(DiagnosticSeverityInfo))
	jsRuntime.Variable("diagnostic", "warning", reflect.ValueOf(DiagnosticSeverityWarning))
	jsRuntime.Variable("diagnostic", "error", reflect.ValueOf(DiagnosticSeverityError))

	jsRuntime.Variable("diagnostic", "linting", reflect.ValueOf(DiagnosticCategoryLinting))
	jsRuntime.Variable("diagnostic", "security", reflect.ValueOf(DiagnosticCategorySecurity))
	jsRuntime.Variable("diagnostic", "specs", reflect.ValueOf(DiagnosticCategorySpecs))

	jsRuntime.Type(reflect.TypeFor[DiagnosticMetadata]()).JsModule(
		"diagnostic",
	).Fields(
		js.Field("Id"),
		js.Field("Name"),
		js.Field("Description"),
		js.Field("Severity"),
		js.Field("Categories"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewDiagnosticMetadata)),
	)

	jsRuntime.Type(reflect.TypeFor[Diagnostic]()).JsModule(
		"diagnostic",
	).Fields(
		js.Field("Metadata"),
		js.Field("Message"),
		js.Field("Document"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewDiagnostic)),
		js.Constructor(reflect.ValueOf(NewDiagnosticWithDocument)),
	)
}
