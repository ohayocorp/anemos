package cmd

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/ohayocorp/anemos/pkg"
	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
	"github.com/ohayocorp/anemos/pkg/util"
	"github.com/spf13/cobra"
)

func getWriteDeclarationsCommand(program *AnemosProgram) *cobra.Command {
	command := &cobra.Command{
		Use:   "declarations [directory]",
		Short: "Writes the type declarations to the specified directory.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return writeDeclarations(program, args[0])
		},
		Args: cobra.ExactArgs(1),
	}

	return command
}

func writeDeclarations(program *AnemosProgram, output string) error {
	slog.Info("Writing type declarations to ${directory}", slog.String("directory", output))

	output, err := filepath.Abs(output)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %s, %w", output, err)
	}

	err = os.MkdirAll(output, 0777)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %s, %w", output, err)
	}

	packageJsonFile := filepath.Join(output, "package.json")
	packageJson := core.ParseTemplate(`
		{
		  "name": "{{ .PackageName }}",
		  "version": "{{ .AppVersion }}",
		  "scripts": {}
		}
		`,
		map[string]interface{}{
			"AppVersion":  util.AppVersion,
			"PackageName": js.PackageName,
		})

	err = os.WriteFile(packageJsonFile, []byte(packageJson), 0666)
	if err != nil {
		return fmt.Errorf("failed to write package.json: %s, %w", packageJsonFile, err)
	}

	return createTypeDeclarations(program, output)
}

func createTypeDeclarations(program *AnemosProgram, outputDir string) error {
	declarations, err := fs.Sub(pkg.TypeDeclarations, "jsdeclarations")
	if err != nil {
		return err
	}

	err = os.CopyFS(outputDir, declarations)
	if err != nil {
		return fmt.Errorf("failed to copy type declarations to %s: %w", outputDir, err)
	}

	for _, extraDeclarations := range program.ExtraJsDeclarations {
		err = os.CopyFS(outputDir, extraDeclarations)
		if err != nil {
			return fmt.Errorf("failed to copy extra declarations to %s: %w", outputDir, err)
		}
	}

	indexContents := ""

	fs.WalkDir(os.DirFS(outputDir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.Name() == "nodejs.d.ts" {
			indexContents += fmt.Sprintf("import './%s';\n", path)
			return nil
		}

		if d.IsDir() {
			if path == "." {
				return nil
			}

			indexContents += fmt.Sprintf("export * from './%s';\n", path)
			return fs.SkipDir
		}

		if strings.HasSuffix(path, ".d.ts") {
			indexContents += fmt.Sprintf("export * from './%s';\n", strings.TrimSuffix(path, ".d.ts"))
		}

		return nil
	})

	indexFile := filepath.Join(outputDir, "index.d.ts")
	err = os.WriteFile(indexFile, []byte(indexContents), 0666)
	if err != nil {
		return fmt.Errorf("failed to write index file: %s, %w", indexFile, err)
	}

	return nil
}
