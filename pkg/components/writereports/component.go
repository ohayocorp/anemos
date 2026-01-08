package writereports

import (
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
	"github.com/ohayocorp/anemos/pkg/util"
)

//go:embed css/github-markdown.css
var CssGithubMarkdown string

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
	component.AddAction(core.StepOutput, component.output)

	return component.Component
}

func (component *component) sanitizeOptions(context *core.BuildContext) {
	options := component.options

	if options == nil {
		options = NewOptions()
		component.options = options
	}

	if len(options.OutputTypes) == 0 {
		options.OutputTypes = []ReportOutputType{
			ReportOutputTypeMarkdown,
			ReportOutputTypeHtml,
		}
	}
}

func (component *component) output(context *core.BuildContext) {
	outputConfiguration := context.BuilderOptions.OutputConfiguration
	reportsDirectory := filepath.Join(outputConfiguration.OutputPath, "reports")

	slog.Info("Writing reports to ${directory}", slog.String("directory", reportsDirectory))

	reports := context.GetAllReports()

	for _, report := range reports {
		if report.Metadata == nil {
			js.Throw(fmt.Errorf("report metadata is nil"))
		}
	}

	for _, report := range reports {
		for _, outputType := range component.options.OutputTypes {
			if outputType == ReportOutputTypeMarkdown {
				outputFile := component.createFile(changeFileExtension(report.Metadata.FilePath, ".md"), reportsDirectory)
				defer outputFile.Close()

				outputFile.WriteString(report.MarkdownContent)
			}

			if outputType == ReportOutputTypeHtml {
				htmlFile := component.createFile(changeFileExtension(report.Metadata.FilePath, ".html"), reportsDirectory)
				defer htmlFile.Close()

				htmlText := component.renderMarkdown(report.MarkdownContent, "")
				htmlFile.WriteString(htmlText)
			}
		}
	}
}

func changeFileExtension(fileName string, extension string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + extension
}

func (component *component) createFile(fileName, directory string) *os.File {
	filePath := filepath.Join(directory, fileName)
	directory = filepath.Dir(filePath)

	if directory != "" {
		if err := os.MkdirAll(directory, os.ModePerm); err != nil {
			js.Throw(fmt.Errorf("can't create directory %s, %v", directory, err))
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		js.Throw(fmt.Errorf("can't create report file %s, %v", filePath, err))
	}

	return file
}

func (component *component) renderMarkdown(mdText, title string) string {
	head := util.ParseTemplate(`
		<meta name="viewport" content="width=device-width, initial-scale=1, minimal-ui">
		<style>
		  body {
		    box-sizing: border-box;
		    min-width: 200px;
		    max-width: 1280px;
		    margin: 0 auto;
		    padding: 45px;
		    background-color: #0d1117;
		  }
		</style>
		<style>
		  {{ .Css }}
		</style>
		`,
		map[string]interface{}{
			"Css": util.Indent(CssGithubMarkdown, 2),
		})

	head = fmt.Sprintf("  %s\n", util.Indent(head, 2))

	p := parser.NewWithExtensions(parser.CommonExtensions)
	r := html.NewRenderer(html.RendererOptions{
		Flags:          html.CommonFlags | html.CompletePage,
		RenderNodeHook: renderHook,
		Head:           []byte(head),
		Title:          title,
	})

	b := markdown.ToHTML([]byte(mdText), p, r)

	return string(b)
}

func renderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if _, ok := node.(*ast.Document); ok {
		if entering {
			io.WriteString(w, `<article class="markdown-body">`)
			io.WriteString(w, "\n")
		} else {
			io.WriteString(w, "\n")
			io.WriteString(w, `</article>`)
		}

		return ast.GoToNext, false
	}
	return ast.GoToNext, false
}
