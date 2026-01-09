//go:build windows

package js

import (
	_ "embed"
)

//go:embed bun/bun-windows-x64.zip
var bunZip []byte
var bunFileName = "bun.exe"

//go:embed ts/ts-win32-x64.tgz
var tsTgz []byte
var tsFileName = "tsgo.exe"
