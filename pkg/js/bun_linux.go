//go:build linux

package js

import (
	_ "embed"
)

//go:embed bun/bun-linux-x64.zip
var bunZip []byte
var bunFileName = "bun"
