$ErrorActionPreference = "Stop"
$ProgressPreference = 'SilentlyContinue'

$thisScriptDirectory=$PSScriptRoot

$platforms = @(
  "win32-x64",
  "linux-x64",
  "linux-arm64",
  "darwin-x64",
  "darwin-arm64"
)

foreach ($platform in $platforms) {
  mkdir "$thisScriptDirectory\pkg\js\ts" -Force | Out-Null

  Invoke-WebRequest `
    -Uri "https://registry.npmjs.org/@typescript/native-preview-${platform}/-/native-preview-${platform}-7.0.0-dev.20260108.1.tgz" `
    -OutFile "$thisScriptDirectory\pkg\js\ts\ts-${platform}.tgz"
}