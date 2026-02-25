package js

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func RunTsc(tsconfigPath string) error {
	return runTsGo(tsconfigPath)
}

func runTsGo(directory string) error {
	tsTargetDirectory := filepath.Join(directory, "tsgo")
	tsPath := filepath.Join(tsTargetDirectory, "lib", tsFileName)

	if err := extractTypeScriptPackage(tsTargetDirectory); err != nil {
		return fmt.Errorf("failed to extract TypeScript package: %w", err)
	}

	slog.Info(
		"Running tsgo to compile ${directory}",
		slog.String("directory", directory))

	startTime := time.Now()

	cmd := exec.Command(tsPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Dir = directory

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("TypeScript compilation failed with exit code %d", cmd.ProcessState.ExitCode())
	}

	endTime := time.Now()
	slog.Debug("tsgo execution time: ${duration}", slog.String("duration", endTime.Sub(startTime).String()))

	return nil
}

func extractTypeScriptPackage(tsTargetDirectory string) error {
	err := os.MkdirAll(tsTargetDirectory, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create target directory %s: %w", tsTargetDirectory, err)
	}

	tgzReader, err := gzip.NewReader(bytes.NewReader(tsTgz))
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}

	defer tgzReader.Close()

	tarReader := tar.NewReader(tgzReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		name := strings.TrimPrefix(header.Name, "package/")
		if name == "" {
			continue
		}

		destPath := filepath.Join(tsTargetDirectory, name)

		switch header.Typeflag {

		case tar.TypeDir:
			if err := os.MkdirAll(destPath, os.FileMode(header.Mode)); err != nil {
				return err
			}

		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
				return err
			}

			outFile, err := os.OpenFile(
				destPath,
				os.O_CREATE|os.O_RDWR|os.O_TRUNC,
				os.FileMode(header.Mode),
			)
			if err != nil {
				return err
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()

		case tar.TypeSymlink:
			if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
				return err
			}
			linkName := strings.TrimPrefix(header.Linkname, "package/")
			if err := os.Symlink(linkName, destPath); err != nil {
				return err
			}

		default:
			return fmt.Errorf("unknown type flag: %c", header.Typeflag)
		}
	}

	return nil
}
