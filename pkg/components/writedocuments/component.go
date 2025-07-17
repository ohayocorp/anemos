package writedocuments

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
)

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
		options = &Options{}
		component.options = options
	}

	component.SetIdentifier("write-documents")
}

func (component *component) output(context *core.BuildContext) {
	outputConfiguration := context.BuilderOptions.OutputConfiguration
	outputDirectory := filepath.Join(outputConfiguration.OutputPath, core.DocumentsDir)
	outputDirectory, err := filepath.Abs(outputDirectory)
	if err != nil {
		js.Throw(fmt.Errorf("can't get absolute path for %s, %v", outputDirectory, err))
	}

	slog.Info("Writing documents to ${outputDirectory}", slog.String("outputDirectory", outputDirectory))

	if outputDirectory != "" {
		if err := os.MkdirAll(outputDirectory, os.ModePerm); err != nil {
			js.Throw(fmt.Errorf("can't create directory %s, %v", outputDirectory, err))
		}
	}

	paths := map[string]int{}

	for _, documentGroup := range context.GetDocumentGroups() {
		for _, document := range documentGroup.Documents {
			path := document.FullPath()
			paths[path]++
		}

		for _, additionalFile := range documentGroup.AdditionalFiles {
			path := fmt.Sprintf("%s/%s", documentGroup.Name, additionalFile.Path)
			paths[path]++
		}
	}

	duplicates := []string{}
	for path, count := range paths {
		if count > 1 {
			duplicates = append(duplicates, path)
		}
	}

	if len(duplicates) > 0 {
		sort.Strings(duplicates)
		message := ""

		for _, path := range duplicates {
			message += fmt.Sprintf("  %s -> %d times\n", path, paths[path])
		}

		js.Throw(fmt.Errorf("duplicate document paths found:\n%s", message))
	}

	for _, documentGroup := range context.GetDocumentGroups() {
		for _, document := range documentGroup.Documents {
			component.writeDocument(document, outputDirectory)
		}

		for _, additionalFile := range documentGroup.AdditionalFiles {
			outputDirectory := filepath.Join(outputDirectory, documentGroup.Name)
			filePath := filepath.Join(outputDirectory, additionalFile.Path)
			filePath = strings.ReplaceAll(filePath, "\\", "/")
			fileDirectory := filepath.Dir(filePath)

			if fileDirectory != "" {
				if err := os.MkdirAll(fileDirectory, os.ModePerm); err != nil {
					js.Throw(fmt.Errorf("can't create directory %s, %v", fileDirectory, err))
				}
			}

			err = os.WriteFile(filePath, []byte(additionalFile.Content), os.ModePerm)
			if err != nil {
				js.Throw(fmt.Errorf("can't write file %s, %v", filePath, err))
			}
		}
	}
}

func (component *component) writeDocument(document *core.Document, outputDirectory string) {
	documentPath := document.FullPath()
	filePath := filepath.Join(outputDirectory, documentPath)
	documentDirectory := filepath.Dir(filePath)

	slog.Debug("Writing document ${path}", slog.String("path", documentPath))

	if documentDirectory != "" {
		if err := os.MkdirAll(documentDirectory, os.ModePerm); err != nil {
			js.Throw(fmt.Errorf("can't create directory %s, %v", documentDirectory, err))
		}
	}

	yaml := core.SerializeToYaml(document)

	err := os.WriteFile(filePath, []byte(yaml), os.ModePerm)
	if err != nil {
		js.Throw(fmt.Errorf("can't write file %s, %v", filePath, err))
	}
}
