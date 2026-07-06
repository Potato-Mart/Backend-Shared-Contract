[CmdletBinding()]
param()

$ErrorActionPreference = 'Stop'

$repoRoot = Resolve-Path (Join-Path $PSScriptRoot '..')
Push-Location $repoRoot
try {
    $env:GOWORK = 'off'
    go test ./...
    if ($LASTEXITCODE -ne 0) {
        throw "go test ./... failed"
    }
}
finally {
    Pop-Location
}
