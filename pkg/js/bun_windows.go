//go:build windows

package js

import (
	_ "embed"
)

//go:embed bun/bun-windows-x64.zip
var bunZip []byte
var bunFileName = "bun.exe"
