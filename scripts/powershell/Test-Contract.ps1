[CmdletBinding()]
param()

$ErrorActionPreference = 'Stop'

$repoRoot = Resolve-Path (Join-Path $PSScriptRoot '..\..')
Push-Location $repoRoot
try {
    $env:GOWORK = 'off'
    go test -count=1 ./...
    if ($LASTEXITCODE -ne 0) {
        throw "go test -count=1 ./... failed"
    }
}
finally {
    Pop-Location
}
