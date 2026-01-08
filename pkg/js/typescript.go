package js

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

//go:embed ts/typescript-5-9-3.zip
var TypeScriptPackage []byte
var typeScriptTargetDirectory = filepath.Join(os.TempDir(), "anemos", "typescript", "5.9.3")

func RunTsc(tsconfigPath string) error {
	// Running tsc with Goja requires many more NodeJS module definitions to be implemented. A working implementation
	// is available in a private branch, but the performance is not acceptable for production use. So, it is not included
	// in the main branch.

	// Running tsc with Bun is about 10x faster than running it with Goja. Still, it can take seconds to compile
	// large projects.
	// When https://github.com/microsoft/typescript-go is released, we can switch to it.
	return runTscWithBun(tsconfigPath)
}

func runTscWithBun(directory string) error {
	if err := extractTypeScriptPackage(); err != nil {
		return fmt.Errorf("failed to extract typescript package: %w", err)
	}

	slog.Info(
		"Running tsc with Bun to compile ${directory}",
		slog.String("directory", directory))

	startTime := time.Now()

	err := RunBunCommand(BunCommand{
		Description: "TypeScript compilation",
		Args:        []string{"run", filepath.Join(typeScriptTargetDirectory, "bin", "tsc")},
		Cwd:         &directory,
		Stdout:      os.Stdout,
		Stderr:      os.Stderr,
		Stdin:       os.Stdin,
	})

	endTime := time.Now()
	slog.Debug("tsc execution time: ${duration}", slog.String("duration", endTime.Sub(startTime).String()))

	return err
}

func extractTypeScriptPackage() error {
	checkPath := filepath.Join(typeScriptTargetDirectory, "bin", "tsc")

	if _, err := os.Stat(checkPath); err == nil {
		slog.Debug("typescript package already extracted to ${path}", "path", checkPath)
		return nil
	}

	err := os.MkdirAll(typeScriptTargetDirectory, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create target directory %s: %w", bunTargetDirectory, err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(TypeScriptPackage), int64(len(TypeScriptPackage)))
	if err != nil {
		return fmt.Errorf("failed to create zip reader: %w", err)
	}

	for _, file := range zipReader.File {
		destPath := filepath.Join(typeScriptTargetDirectory, file.Name)
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(destPath, 0755)
			if err != nil {
				return fmt.Errorf("failed to create directory %s: %w", destPath, err)
			}
			continue
		}

		srcFile, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open zip file %s: %w", file.Name, err)
		}
		defer srcFile.Close()

		destFile, err := os.Create(destPath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", destPath, err)
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return fmt.Errorf("failed to extract file %s: %w", file.Name, err)
		}
	}

	return nil
}
