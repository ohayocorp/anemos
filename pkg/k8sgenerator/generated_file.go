package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type GeneratedFile struct {
	Path string
	file *os.File
}

func NewGeneratedFile(path string) (*GeneratedFile, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create output file: %w", err)
	}

	file := &GeneratedFile{
		Path: path,
		file: f,
	}

	return file, nil
}

func (file *GeneratedFile) Close() error {
	return file.file.Close()
}

func (file *GeneratedFile) Write(text string) {
	file.file.WriteString(text)
}
