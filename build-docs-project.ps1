$thisScriptDirectory=$PSScriptRoot

Push-Location "${thisScriptDirectory}\docs-code\js"

try {
    anemos.exe package install
    if ($LastExitCode -ne 0) {
        exit $LastExitCode
    }

    anemos.exe build --verbose index.js
    if ($LastExitCode -ne 0) {
        exit $LastExitCode
    }
} finally {
    Pop-Location
}

Push-Location "${thisScriptDirectory}\docs-code\ts"
try {
    # Find all tsconfig.json files in the current directory and subdirectories ignoring node_modules.
    $tsconfigFiles = Get-ChildItem -Path . -Filter "tsconfig.json" -Recurse -File | Where-Object { $_.FullName -notlike "*node_modules*" }

    # For each tsconfig.json file, run the anemos build command
    foreach ($tsconfigFile in $tsconfigFiles) {
        $tsconfigPath = $tsconfigFile.FullName
        
        Push-Location $tsconfigFile.DirectoryName

        try {
            # Check if the package.json file exists in the same directory as tsconfig.json
            $packageJsonPath = Join-Path -Path $tsconfigFile.DirectoryName -ChildPath "package.json"
            if (Test-Path -Path $packageJsonPath) {
                # Run anemos package install in the directory of the tsconfig.json
                anemos.exe package install
                if ($LastExitCode -ne 0) {
                    exit $LastExitCode
                }

                # Run anemos package run copy-assets if it exists in the package.json
                $packageJson = Get-Content -Path $packageJsonPath | ConvertFrom-Json
                if ($packageJson.scripts -and $packageJson.scripts."copy-assets") {
                    anemos.exe package run copy-assets
                    if ($LastExitCode -ne 0) {
                        exit $LastExitCode
                    }
                }
            }

            anemos.exe build --tsc . --verbose "dist/index.js"

            if ($LastExitCode -ne 0) {
                exit $LastExitCode
            }
        } finally {
            Pop-Location
        }
    }
} finally {
    Pop-Location
}