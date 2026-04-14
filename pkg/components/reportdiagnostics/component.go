package reportdiagnostics

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ohayocorp/anemos/pkg/core"
)

const componentType = "report-diagnostics"

type component struct {
	*core.Component
	options *Options
}

func NewComponent(options *Options) *core.Component {
	component := &component{
		Component: core.NewComponent(),
		options:   options,
	}

	component.AddAction(core.StepSanitize, component.sanitizeOptions)
	component.AddAction(core.StepReport, component.report)

	component.SetComponentType(componentType)
	component.SetIdentifier(componentType)

	return component.Component
}

func (component *component) sanitizeOptions(context *core.BuildContext) {
	options := component.options

	if options == nil {
		options = &Options{}
		component.options = options
	}
}

func (component *component) report(context *core.BuildContext) {
	report := core.NewReport(
		core.NewReportMetadata("diagnostics.md"),
		"# Diagnostics\n\n",
	)

	severitiesToCheck := []core.DiagnosticSeverity{
		core.DiagnosticSeverityError,
		core.DiagnosticSeverityWarning,
	}

	severities := map[core.DiagnosticSeverity][]*core.Diagnostic{}

	for _, diagnostic := range context.GetAllDiagnostics() {
		severity := diagnostic.Metadata.Severity
		if !slices.Contains(severitiesToCheck, severity) {
			continue
		}

		severities[diagnostic.Metadata.Severity] = append(severities[diagnostic.Metadata.Severity], diagnostic)
	}

	for _, diagnostics := range severities {
		slices.SortStableFunc(diagnostics, func(a, b *core.Diagnostic) int {
			return strings.Compare(a.Metadata.Id, b.Metadata.Id)
		})
	}

	for _, severity := range severitiesToCheck {
		diagnostics := severities[severity]
		if len(diagnostics) == 0 {
			continue
		}

		ids := map[string][]*core.Diagnostic{}
		for _, diagnostic := range diagnostics {
			ids[diagnostic.Metadata.Id] = append(ids[diagnostic.Metadata.Id], diagnostic)
		}

		for _, id := range core.SortedKeys(ids) {
			diagnostics := ids[id]
			if len(diagnostics) == 0 {
				continue
			}

			report.MarkdownContent += fmt.Sprintf("## %s (%s)\n", diagnostics[0].Metadata.Name, severity)
			report.MarkdownContent += fmt.Sprintf("%s\n\n", diagnostics[0].Metadata.Description)

			diagnosticsByDocument := map[string][]*core.Diagnostic{}
			for _, diagnostic := range diagnostics {
				key := ""
				if diagnostic.Document != nil {
					key = diagnostic.Document.FullPath()
				}

				diagnosticsByDocument[key] = append(diagnosticsByDocument[key], diagnostic)
			}

			for _, key := range core.SortedKeys(diagnosticsByDocument) {
				report.MarkdownContent += fmt.Sprintf("#### `%s`\n", key)

				for _, diagnostic := range diagnosticsByDocument[key] {
					if diagnostic.Message != "" {
						report.MarkdownContent += fmt.Sprintf("- %s\n", diagnostic.Message)
					}
				}

				report.MarkdownContent += "\n"
			}
		}
	}

	if len(severities) > 0 {
		context.AddReport(report)
	}
}
