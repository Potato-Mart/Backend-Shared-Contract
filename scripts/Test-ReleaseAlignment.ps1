[CmdletBinding()]
param(
    [ValidatePattern('^v\d+\.\d+\.\d+(-[0-9A-Za-z.-]+)?$')]
    [string]$ExpectedVersion,

    [switch]$RequireTag,
    [switch]$PassThru
)

$ErrorActionPreference = 'Stop'

function Assert-Match {
    param(
        [bool]$Condition,
        [string]$Message
    )

    if (-not $Condition) {
        throw $Message
    }
}

$repoRoot = Resolve-Path (Join-Path $PSScriptRoot '..')
Push-Location $repoRoot
try {
    $versionPath = Join-Path $repoRoot 'pkg/versioning/version.go'
    $modulePath = Join-Path $repoRoot 'go.mod'
    $readmePath = Join-Path $repoRoot 'README.md'
    $releaseNotesPath = Join-Path $repoRoot 'RELEASE_NOTES.md'

    $versionSource = Get-Content -LiteralPath $versionPath -Raw
    $moduleSource = Get-Content -LiteralPath $modulePath -Raw
    $readmeSource = Get-Content -LiteralPath $readmePath -Raw
    $releaseNotesSource = Get-Content -LiteralPath $releaseNotesPath -Raw

    $versionMatch = [regex]::Match(
        $versionSource,
        'ModuleVersion\s*=\s*"(?<version>v\d+\.\d+\.\d+(?:-[0-9A-Za-z.-]+)?)"'
    )
    Assert-Match $versionMatch.Success 'pkg/versioning/version.go does not contain a valid ModuleVersion.'
    $version = $versionMatch.Groups['version'].Value

    $semverMatch = [regex]::Match(
        $version,
        '^v(?<major>0|[1-9]\d*)\.(?<minor>0|[1-9]\d*)\.(?<patch>0|[1-9]\d*)(?:-(?<prerelease>[0-9A-Za-z.-]+))?$'
    )
    Assert-Match $semverMatch.Success "ModuleVersion '$version' is not valid release semver."
    $versionMajor = "v$($semverMatch.Groups['major'].Value)"

    $majorMatch = [regex]::Match($versionSource, 'MajorVersion\s*=\s*"(?<major>v\d+)"')
    Assert-Match $majorMatch.Success 'pkg/versioning/version.go does not contain a valid MajorVersion.'
    Assert-Match (
        $majorMatch.Groups['major'].Value -eq $versionMajor
    ) "MajorVersion '$($majorMatch.Groups['major'].Value)' does not match ModuleVersion '$version'."

    $moduleMatch = [regex]::Match(
        $moduleSource,
        '(?m)^module\s+(?<path>\S+)/v(?<major>0|[1-9]\d*)\s*$'
    )
    Assert-Match $moduleMatch.Success 'go.mod must declare a versioned module path ending in /vMAJOR.'
    $moduleImportPath = "$($moduleMatch.Groups['path'].Value)/v$($moduleMatch.Groups['major'].Value)"
    Assert-Match (
        "v$($moduleMatch.Groups['major'].Value)" -eq $versionMajor
    ) "go.mod module path '$moduleImportPath' does not match ModuleVersion '$version'."

    if ($ExpectedVersion) {
        Assert-Match (
            $ExpectedVersion -eq $version
        ) "Requested release '$ExpectedVersion' does not match source ModuleVersion '$version'."
    }

    $versionLinePattern = "(?m)^$([regex]::Escape($version))\s*$"
    Assert-Match (
        [regex]::IsMatch($readmeSource, $versionLinePattern)
    ) "README.md does not advertise source ModuleVersion '$version'."

    $consumerRequirement = "require $moduleImportPath $version"
    Assert-Match (
        $readmeSource.Contains($consumerRequirement)
    ) "README.md does not contain the aligned consumer requirement '$consumerRequirement'."

    $releaseIndexEntry = '| `' + $version + '` |'
    Assert-Match (
        $releaseNotesSource.Contains($releaseIndexEntry)
    ) "RELEASE_NOTES.md does not contain a release-index entry for '$version'."

    $releaseHeadingPattern = "(?m)^##\s+$([regex]::Escape($version))(?:\s|\()"
    Assert-Match (
        [regex]::IsMatch($releaseNotesSource, $releaseHeadingPattern)
    ) "RELEASE_NOTES.md does not contain a detailed heading for '$version'."

    & git rev-parse --is-inside-work-tree *> $null
    Assert-Match ($LASTEXITCODE -eq 0) 'Release alignment must run inside the contract Git repository.'

    & git rev-parse --verify --quiet "refs/tags/$version^{commit}" *> $null
    $tagExists = $LASTEXITCODE -eq 0

    if ($tagExists) {
        $tagType = (& git cat-file -t "refs/tags/$version").Trim()
        Assert-Match (
            $LASTEXITCODE -eq 0 -and $tagType -eq 'tag'
        ) "Existing release tag '$version' must be an annotated tag."

        & git diff --quiet --exit-code "refs/tags/$version" HEAD -- go.mod go.sum README.md RELEASE_NOTES.md pkg
        $diffExitCode = $LASTEXITCODE
        if ($diffExitCode -eq 1) {
            throw "Release payload differs from immutable tag '$version'. Bump ModuleVersion and release notes before changing contract payload."
        }
        Assert-Match (
            $diffExitCode -eq 0
        ) "Could not compare the current release payload with tag '$version'."
    }
    elseif ($RequireTag) {
        throw "Required release tag '$version' does not exist."
    }

    Write-Host "Contract release metadata is aligned at $version ($moduleImportPath)."
    if ($PassThru) {
        Write-Output $version
    }
}
finally {
    Pop-Location
}
