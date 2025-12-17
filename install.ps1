# Build and install godo to your Go bin, then remove the local binary
# Usage: .\install.ps1

$projectDir = Split-Path -Path $MyInvocation.MyCommand.Definition -Parent
Push-Location $projectDir

Write-Output "Building godo.exe..."
go build -o godo.exe

$gobin = (go env GOBIN)
if ($gobin -eq '') {
    $gobin = Join-Path $env:USERPROFILE 'go\bin'
}

if (-not (Test-Path $gobin)) {
    New-Item -ItemType Directory -Path $gobin | Out-Null
}

Write-Output "Installing to $gobin"
Copy-Item -Force .\godo.exe (Join-Path $gobin 'godo.exe')

Write-Output "Removing local godo.exe from project directory"
Remove-Item -Force .\godo.exe -ErrorAction SilentlyContinue

Write-Output "Done. Ensure $gobin is in your PATH, then run 'godo -list' from any terminal."
Pop-Location
