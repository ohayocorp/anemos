package js

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
)

var bunTargetDirectory = filepath.Join(os.TempDir(), "anemos", "bun")
var bunPath = filepath.Join(bunTargetDirectory, bunFileName)

type BunCommand struct {
	Description string
	Args        []string
	Cwd         *string
}

func RunBunCommand(bunCommand BunCommand) error {
	if err := extractBun(); err != nil {
		return fmt.Errorf("failed to extract bun: %w", err)
	}

	cmd := exec.Command(bunPath, bunCommand.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if bunCommand.Cwd != nil {
		cmd.Dir = *bunCommand.Cwd
	}

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("%s failed with exit code %d", bunCommand.Description, cmd.ProcessState.ExitCode())
	}

	return nil
}

func extractBun() error {
	checkPath := filepath.Join(bunTargetDirectory, bunFileName)

	if _, err := os.Stat(checkPath); err == nil {
		slog.Debug("bun executable already extracted to ${path}", "path", checkPath)
		return nil
	}

	err := os.MkdirAll(bunTargetDirectory, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create target directory %s: %w", bunTargetDirectory, err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(bunZip), int64(len(bunZip)))
	if err != nil {
		return fmt.Errorf("failed to create zip reader: %w", err)
	}

	for _, file := range zipReader.File {
		if filepath.Base(file.Name) != bunFileName {
			continue
		}

		outFile, err := os.OpenFile(bunPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", bunPath, err)
		}
		defer outFile.Close()

		rc, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open zip file %s: %w", file.Name, err)
		}
		defer rc.Close()

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return fmt.Errorf("failed to write file %s: %w", bunPath, err)
		}
	}

	return nil
}
