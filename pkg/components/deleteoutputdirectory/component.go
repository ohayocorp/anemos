package deleteoutputdirectory

import (
	"errors"
	"fmt"
	"os"

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
	// This component should be run just before the output step otherwise all output files will be lost.
	component.AddAction(core.NewStep("Delete outputs", append(core.StepOutput.Numbers, -1)...), component.output)

	return component.Component
}

func (component *component) sanitizeOptions(context *core.BuildContext) {
	options := component.options

	if options == nil {
		options = &Options{}
		component.options = options
	}

	component.SetIdentifier("delete-output-directory")
}

func (component *component) output(context *core.BuildContext) {
	outputDirectory := context.BuilderOptions.OutputConfiguration.OutputPath
	errs := []error{}

	// Try multiple times to delete the output directory. Git commands run by VS Code
	// can sometimes lock files, preventing deletion.
	for {
		err := os.RemoveAll(outputDirectory)
		if err == nil {
			break
		}

		errs = append(errs, err)
		if len(errs) > 5 {
			js.Throw(fmt.Errorf("failed to delete directory %s after multiple attempts: %v", outputDirectory, errors.Join(errs...)))
		}
	}
}
