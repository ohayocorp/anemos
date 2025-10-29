package cmd

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/ohayocorp/anemos/pkg"
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
	packageJson := util.ParseTemplate(`
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
	indexBuilder := &strings.Builder{}

	// Append native type declarations to index.d.ts file.
	err := copyDeclarations(pkg.LibNativeDeclarations, outputDir, indexBuilder)
	if err != nil {
		return err
	}

	// Append library type declarations to index.d.ts file.
	err = copyDeclarations(pkg.LibTypeDeclarations, outputDir, indexBuilder)
	if err != nil {
		return err
	}

	for _, extraDeclarations := range program.ExtraJsDeclarations {
		err = copyDeclarations(extraDeclarations, outputDir, indexBuilder)
		if err != nil {
			return fmt.Errorf("failed to copy extra declarations to %s: %w", outputDir, err)
		}
	}

	indexFile := filepath.Join(outputDir, "index.d.ts")

	err = os.WriteFile(indexFile, []byte(indexBuilder.String()), 0666)
	if err != nil {
		return fmt.Errorf("failed to write index file: %s, %w", indexFile, err)
	}

	return nil
}

func copyDeclarations(files fs.FS, outputDir string, indexBuilder *strings.Builder) error {
	// Walk the source FS and copy files, overwriting any existing ones.
	return fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		target := filepath.Join(outputDir, path)

		if d.IsDir() {
			// Ensure the directory exists.
			return os.MkdirAll(target, 0777)
		}

		// Ensure the parent directory exists.
		if err := os.MkdirAll(filepath.Dir(target), 0777); err != nil {
			return err
		}

		// Read the file from the source FS.
		b, err := fs.ReadFile(files, path)
		if err != nil {
			return err
		}

		// Don't write the index.d.ts file as multiple index files will be merged afterwards.
		if indexBuilder != nil && path == "index.d.ts" {
			// Libraries use @ohayocorp/anemos to import modules, but VS Code intellisense can't resolve these paths.
			// Use relative paths instead. Since the index.d.ts is a declaration file, it doesn't affect module resolution.
			indexContents := string(b)
			indexContents = strings.ReplaceAll(indexContents, "@ohayocorp/anemos/", "./")

			indexBuilder.WriteString(indexContents)
			indexBuilder.WriteString("\n")

			return nil
		}

		// Write the file, truncating if it already exists.
		return os.WriteFile(target, b, 0666)
	})
}
