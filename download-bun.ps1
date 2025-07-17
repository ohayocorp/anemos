$ErrorActionPreference = "Stop"
$ProgressPreference = 'SilentlyContinue'

$thisScriptDirectory=$PSScriptRoot

$platforms = @(
  "windows-x64",
  "linux-x64",
  "linux-aarch64",
  "darwin-x64",
  "darwin-aarch64"
)

foreach ($platform in $platforms) {
  mkdir "$thisScriptDirectory\pkg\js\bun" -Force | Out-Null

  Invoke-WebRequest `
    -Uri "https://github.com/oven-sh/bun/releases/download/bun-v1.2.18/bun-$platform.zip" `
    -OutFile "$thisScriptDirectory\pkg\js\bun\bun-$platform.zip"
}