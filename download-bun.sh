#!/bin/bash

set -e
set -o pipefail

thisScriptFile=$(readlink -f "${BASH_SOURCE[0]:-$0}")
thisScriptDirectory=$(dirname "${thisScriptFile}")

platforms=("windows-x64" "linux-x64" "linux-aarch64" "darwin-x64" "darwin-aarch64")

for platform in "${platforms[@]}"
do
  echo "Downloading Bun for platform: ${platform}"
  mkdir -p "${thisScriptDirectory}/pkg/js/bun"

  wget -q \
    -O "${thisScriptDirectory}/pkg/js/bun/bun-${platform}.zip" \
    https://github.com/oven-sh/bun/releases/download/bun-v1.2.18/bun-${platform}.zip
done