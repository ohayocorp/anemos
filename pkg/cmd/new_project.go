package cmd

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/ohayocorp/anemos/pkg/js"
	"github.com/ohayocorp/anemos/pkg/util"
	"github.com/spf13/cobra"
)

var (
	//go:embed all:templates
	templates embed.FS
)

type language string
type projectType string

const (
	javascript language = "javascript"
	typescript language = "typescript"

	projectTypePackage projectType = "package"
	projectTypeApp     projectType = "app"
)

func (l *language) String() string {
	return string(*l)
}

func (l *language) Type() string {
	return fmt.Sprintf("[%s|%s]", javascript, typescript)
}

func (l *language) Set(value string) error {
	switch value {
	case "javascript", "typescript":
		*l = language(value)
		return nil
	default:
		return fmt.Errorf("unsupported language: %s, only javascript and typescript are supported", value)
	}
}

func (p *projectType) String() string {
	return string(*p)
}
func (p *projectType) Type() string {
	return fmt.Sprintf("[%s|%s]", projectTypeApp, projectTypePackage)
}
func (p *projectType) Set(value string) error {
	switch value {
	case "package", "app":
		*p = projectType(value)
		return nil
	default:
		return fmt.Errorf("unsupported project type: %s, only package and app are supported", value)
	}
}

func getNewProjectCommand(program *AnemosProgram) *cobra.Command {
	var language language
	var projectType projectType

	command := &cobra.Command{
		Use:   "new [directory]",
		Short: "Initializes a new project.",
		RunE: func(ctx *cobra.Command, args []string) error {
			return newProject(program, args, language, projectType)
		},
		Args: cobra.MaximumNArgs(1),
	}

	command.Flags().VarPF(&language, "language", "l", "The language to use for the project").DefValue = string(javascript)
	command.Flags().VarPF(&projectType, "project-type", "p", "The type of project to create").DefValue = string(projectTypeApp)

	return command
}

func newProject(program *AnemosProgram, args []string, language language, projectType projectType) error {
	var output string

	if len(args) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current working directory: %w", err)
		}

		output = cwd
	} else {
		output = args[0]
	}

	output, err := filepath.Abs(output)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %s, %w", output, err)
	}

	entries, err := os.ReadDir(output)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read directory: %s, %w", output, err)
	}

	if len(entries) > 0 {
		//return fmt.Errorf("output directory is not empty: %s", output)
	}

	if language == "" {
		language = javascript
	}

	if projectType == "" {
		projectType = projectTypeApp
	}

	projectName := filepath.Base(output)

	slog.Info(
		"Creating new project ${projectName} in ${directory} with language ${language} and project type ${projectType}",
		slog.String("projectName", projectName),
		slog.String("directory", output),
		slog.String("language", language.String()),
		slog.String("projectType", projectType.String()))

	copyTemplates := func(filesystem fs.FS, subfs string) error {
		templates, err := fs.Sub(filesystem, subfs)
		if err != nil {
			return fmt.Errorf("failed to get templates: %s, %w", subfs, err)
		}

		err = copyFS(output, templates, projectName)
		if err != nil {
			return fmt.Errorf("failed to copy templates: %s, %w", subfs, err)
		}

		return nil
	}

	err = copyTemplates(templates, fmt.Sprintf("templates/%s/%s", language, projectType))
	if err != nil {
		return err
	}

	anemosTypesPath := filepath.Join(output, ".anemos", "types")

	err = writeDeclarations(program, anemosTypesPath)
	if err != nil {
		return fmt.Errorf("failed to write declarations: %w", err)
	}

	slog.Info("Initializing packages in ${directory}", slog.String("directory", output))

	return js.RunBunCommand(js.BunCommand{
		Description: "Initialize packages",
		Args:        []string{"install"},
		Cwd:         &output,
		Stdout:      os.Stdout,
		Stderr:      os.Stderr,
		Stdin:       os.Stdin,
	})
}

func copyFS(dir string, fsys fs.FS, projectName string) error {
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fpath, err := filepath.Localize(path)
		if err != nil {
			return err
		}
		newPath := filepath.Join(dir, fpath)
		if d.IsDir() {
			return os.MkdirAll(newPath, 0777)
		}

		if !d.Type().IsRegular() {
			return &os.PathError{Op: "CopyFS", Path: path, Err: os.ErrInvalid}
		}

		r, err := fsys.Open(path)
		if err != nil {
			return err
		}
		defer r.Close()
		info, err := r.Stat()
		if err != nil {
			return err
		}

		content, err := io.ReadAll(r)
		if err != nil {
			return err
		}

		modifiedContent := bytes.ReplaceAll(content, []byte("ANEMOS_VERSION"), []byte(util.AppVersion))
		modifiedContent = bytes.ReplaceAll(modifiedContent, []byte("PACKAGE_NAME"), []byte(projectName))

		w, err := os.OpenFile(newPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666|info.Mode()&0777)
		if err != nil {
			return err
		}

		if _, err := w.Write(modifiedContent); err != nil {
			w.Close()
			return &os.PathError{Op: "Copy", Path: newPath, Err: err}
		}
		return w.Close()
	})
}
