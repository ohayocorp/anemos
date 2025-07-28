package main

import (
	"log/slog"
	"os"

	"github.com/ohayocorp/anemos/pkg/cmd"
	"github.com/ohayocorp/anemos/pkg/util"
	"github.com/spf13/cobra"
)

func main() {
	program := &cmd.AnemosProgram{
		RootCommand: &cobra.Command{
			Use:     "anemos",
			Short:   "Anemos is a Kubernetes package manager.",
			Version: util.AppVersion,
		},
	}

	if err := cmd.Run(program); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
