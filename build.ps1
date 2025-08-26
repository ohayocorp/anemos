param (
    [string]$outputPath = "bin\anemos.exe"
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest
$PSNativeCommandUseErrorActionPreference = $true

tsc -p pkg\jslib\tsconfig.json
go build -C cmd\anemos -o $outputPath