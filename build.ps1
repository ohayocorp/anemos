param (
    [string]$outputPath = "bin\anemos.exe"
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest
$PSNativeCommandUseErrorActionPreference = $true

if (Test-Path pkg\jslib\dist) {
    Remove-Item -Recurse -Force pkg\jslib\dist
}
tsc -p pkg\jslib\tsconfig.json
go build -C cmd\anemos -o $outputPath