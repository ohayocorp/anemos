package cmd

import (
	"os"

	"github.com/ohayocorp/anemos/pkg/js"
	"github.com/spf13/cobra"
)

type packageCommand struct {
}

func getPackageCommand(program *AnemosProgram) *cobra.Command {
	command := &cobra.Command{
		Use:   "package",
		Short: "Manages NPM packages.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	p := &packageCommand{}

	command.AddCommand(
		p.getRunCommand(),
		p.getInstallCommand(program),
		p.getAddCommand(program),
		p.getUpdateCommand(program),
		p.getRemoveCommand(program),
		p.getPackCommand(),
		p.getPublishCommand(),
		p.getLoginCommand(),
		p.getConfigCommand(),
		p.getLinkCommand(program),
		p.getUnlinkCommand(program),
	)

	return command
}

func (p *packageCommand) getRunCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "run",
		Short:              "Runs a package.json script",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return js.RunBunCommand(js.BunCommand{
				Description: "Running script",
				Args:        append([]string{"run"}, args...),
			})
		},
	}
}

func (p *packageCommand) getPackCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "pack",
		Short:              "Packages the application for distribution",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return js.RunBunCommand(js.BunCommand{
				Description: "Packaging application",
				Args:        append([]string{"pm", "pack"}, args...),
			})
		},
	}
}

func (p *packageCommand) getPublishCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "publish",
		Short:              "Publish the package to the registry",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return js.RunBunCommand(js.BunCommand{
				Description: "Publishing application",
				Args:        append([]string{"publish"}, args...),
			})
		},
	}
}

func (p *packageCommand) getLoginCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "login",
		Short:              "Log in to the package registry",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return js.RunBunCommand(js.BunCommand{
				Description: "Logging in to the package registry",
				Args:        append([]string{"x", "npm", "login"}, args...),
			})
		},
	}
}

func (p *packageCommand) getConfigCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "config",
		Short:              "Manage NPM configuration",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return js.RunBunCommand(js.BunCommand{
				Description: "Managing NPM configuration",
				Args:        append([]string{"x", "npm", "config"}, args...),
			})
		},
	}
}

func (p *packageCommand) getInstallCommand(program *AnemosProgram) *cobra.Command {
	return &cobra.Command{
		Use:                "install",
		Short:              "Installs the dependencies defined in the package.json",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			p.writeAnemosTypeDeclarations(program)

			return js.RunBunCommand(js.BunCommand{
				Description: "Installing dependencies",
				Args:        append([]string{"install"}, args...),
			})
		},
	}
}

func (p *packageCommand) getAddCommand(program *AnemosProgram) *cobra.Command {
	return &cobra.Command{
		Use:                "add",
		Short:              "Add a dependency to package.json",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			p.writeAnemosTypeDeclarations(program)

			return js.RunBunCommand(js.BunCommand{
				Description: "Adding dependencies",
				Args:        append([]string{"add"}, args...),
			})
		},
	}
}

func (p *packageCommand) getUpdateCommand(program *AnemosProgram) *cobra.Command {
	return &cobra.Command{
		Use:                "update",
		Short:              "Update dependencies in package.json",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			p.writeAnemosTypeDeclarations(program)

			return js.RunBunCommand(js.BunCommand{
				Description: "Updating dependencies",
				Args:        append([]string{"update"}, args...),
			})
		},
	}
}

func (p *packageCommand) getRemoveCommand(program *AnemosProgram) *cobra.Command {
	return &cobra.Command{
		Use:                "remove",
		Short:              "Remove a dependency from package.json",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			p.writeAnemosTypeDeclarations(program)

			return js.RunBunCommand(js.BunCommand{
				Description: "Removing dependencies",
				Args:        append([]string{"remove"}, args...),
			})
		},
	}
}

func (p *packageCommand) getLinkCommand(program *AnemosProgram) *cobra.Command {
	return &cobra.Command{
		Use:                "link",
		Short:              "Link a local package",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			p.writeAnemosTypeDeclarations(program)

			return js.RunBunCommand(js.BunCommand{
				Description: "Linking dependencies",
				Args:        append([]string{"link"}, args...),
			})
		},
	}
}

func (p *packageCommand) getUnlinkCommand(program *AnemosProgram) *cobra.Command {
	return &cobra.Command{
		Use:                "unlink",
		Short:              "Unlink a local package",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			p.writeAnemosTypeDeclarations(program)

			return js.RunBunCommand(js.BunCommand{
				Description: "Unlinking dependencies",
				Args:        append([]string{"unlink"}, args...),
			})
		},
	}
}

func (p *packageCommand) writeAnemosTypeDeclarations(program *AnemosProgram) {
	pwd, err := os.Getwd()
	if err == nil {
		writeTypeDeclarations(program, pwd)
	}
}
