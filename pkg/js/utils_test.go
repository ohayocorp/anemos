package js_test

import (
	"os"
	"testing"

	"github.com/ohayocorp/anemos/pkg/js"
)

func ReadScript(t *testing.T, path string) *js.JsScript {
	t.Helper()

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read script %s: %v", path, err)
	}

	return &js.JsScript{
		FilePath: path,
		Contents: string(content),
	}
}
