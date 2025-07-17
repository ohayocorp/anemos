#!/bin/bash

set -e
set -o pipefail

thisScriptFile=$(readlink -f "${BASH_SOURCE[0]:-$0}")
thisScriptDirectory=$(dirname "${thisScriptFile}")

cd "${thisScriptDirectory}/docs-code/js"
anemos package install
anemos build --verbose index.js

cd "${thisScriptDirectory}/docs-code/ts"

# Find all tsconfig.json files in the current directory and subdirectories ignoring node_modules
tsconfigFiles=$(find . -name "tsconfig.json" -not -path "*/node_modules/*")

# For each tsconfig.json file, run the anemos build command
for tsconfigFile in $tsconfigFiles
do
  tsconfigDir=$(dirname "$tsconfigFile")
  pushd "$tsconfigDir" > /dev/null || exit 1

  # Check if the package.json file exists in the same directory as tsconfig.json
  if [[ -f "package.json" ]]
  then
    # Run anemos package install in the directory of the tsconfig.json
    anemos package install

    # Check if copy-assets script exists in package.json
    if grep -q '"copy-assets":' package.json
    then
      anemos package run copy-assets
    fi
  fi

  anemos build --tsc . --verbose dist/index.js

  popd > /dev/null || exit 1
done