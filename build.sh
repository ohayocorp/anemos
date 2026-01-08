set -e

outputPath="${1:-bin/anemos}"

tsc -p ../anemos/pkg/jslib/tsconfig.json
tsc -p pkg/jslib/tsconfig.json
go build -C cmd/anemos -o $outputPath