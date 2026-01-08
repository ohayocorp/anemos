package core

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/js"
)

type ReportMetadata struct {
	FilePath string
}

// A Report analyzes the output documents and writes some information into a file.
type Report struct {
	Metadata        *ReportMetadata
	MarkdownContent string

	component *Component
}

func NewReport(metadata *ReportMetadata, markdownContent string) *Report {
	return &Report{
		Metadata:        metadata,
		MarkdownContent: markdownContent,
	}
}

func NewReportMetadata(filePath string) *ReportMetadata {
	return &ReportMetadata{
		FilePath: filePath,
	}
}

func registerReport(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[Report]()).JsModule(
		"report",
	).Fields(
		js.Field("MarkdownContent"),
		js.Field("Metadata"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewReport)),
	)

	jsRuntime.Type(reflect.TypeFor[ReportMetadata]()).JsModule(
		"report",
	).Fields(
		js.Field("FilePath"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewReportMetadata)),
	)
}
