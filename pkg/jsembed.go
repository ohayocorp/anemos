package pkg

import (
	"embed"
	"fmt"
	"io/fs"
)

var (
	//go:embed jslib/*.d.ts jslib/k8s
	libNativeDeclarations embed.FS
	LibNativeDeclarations = mustSub(libNativeDeclarations, "jslib")

	//go:embed jslib/dist/*.d.ts jslib/dist/**/*.d.ts
	libTypeDeclarations embed.FS
	LibTypeDeclarations = mustSub(libTypeDeclarations, "jslib/dist")

	//go:embed jslib/dist/*.js jslib/dist/**/*.js
	libJavaScript embed.FS
	LibJavaScript = mustSub(libJavaScript, "jslib/dist")
)

func mustSub(files embed.FS, subpath string) fs.FS {
	sub, err := fs.Sub(files, subpath)
	if err != nil {
		panic(fmt.Errorf("failed to create sub filesystem: %w", err))
	}
	return sub
}
