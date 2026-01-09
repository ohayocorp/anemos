//go:build darwin

package js

import (
	_ "embed"
)

//go:embed bun/bun-darwin-x64.zip
var bunZip []byte
var bunFileName = "bun"

//go:embed ts/ts-darwin-x64.tgz
var tsTgz []byte
var tsFileName = "tsgo"
