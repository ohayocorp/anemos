#!/bin/bash

set -e
set -o pipefail

thisScriptFile=$(readlink -f "${BASH_SOURCE[0]:-$0}")
thisScriptDirectory=$(dirname "${thisScriptFile}")

platforms=("win32-x64" "linux-x64" "linux-arm64" "darwin-x64" "darwin-arm64")

for platform in "${platforms[@]}"
do
  echo "Downloading TypeScript compiler for platform: ${platform}"
  mkdir -p "${thisScriptDirectory}/pkg/js/ts"

  wget -q \
    -O "${thisScriptDirectory}/pkg/js/ts/ts-${platform}.tgz" \
    https://registry.npmjs.org/@typescript/native-preview-${platform}/-/native-preview-${platform}-7.0.0-dev.20260108.1.tgz
done