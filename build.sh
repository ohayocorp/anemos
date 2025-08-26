outputPath="bin/anemos"

tsc -p pkg/jslib/tsconfig.json
go build -C cmd/anemos -o $outputPath