//go:build linux

package js

import (
	_ "embed"
)

//go:embed bun/bun-linux-x64.zip
var bunZip []byte
var bunFileName = "bun"

//go:embed ts/ts-linux-x64.tgz
var tsTgz []byte
var tsFileName = "tsgo"
