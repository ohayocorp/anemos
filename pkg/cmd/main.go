package cmd

import (
	"io/fs"
	"log/slog"
	"os"

	"github.com/ohayocorp/anemos/pkg/js"
	"github.com/spf13/cobra"
)

type AnemosProgram struct {
	RootCommand               *cobra.Command
	InitializeRuntimeCallback func(runtime *js.JsRuntime) error
	ExtraJsDeclarations       []fs.FS
}

func Run(program *AnemosProgram) error {
	rootCmd := program.RootCommand

	logLevelVar := &slog.LevelVar{}
	logHandlerOptions := &slog.HandlerOptions{
		Level: logLevelVar,
	}
	slog.SetDefault(slog.New(NewCliSlogHandler(os.Stdout, logHandlerOptions)))

	var isVerbose bool

	rootCmd.PersistentFlags().BoolVarP(&isVerbose, "verbose", "v", false, "enable verbose logging")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		rootCmd.SilenceUsage = true
		rootCmd.SilenceErrors = true

		if isVerbose {
			logLevelVar.Set(slog.LevelDebug)
		}
	}

	rootCmd.AddCommand(
		getNewProjectCommand(program),
		getWriteDeclarationsCommand(program),
		getBuildCommand(program),
		getPackageCommand(program),
		getApplyCommand(program),
		getDeleteCommand(program),
		getListCommand(program),
	)

	return rootCmd.Execute()
}
