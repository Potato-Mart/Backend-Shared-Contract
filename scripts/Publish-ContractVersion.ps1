[CmdletBinding(SupportsShouldProcess = $true)]
param(
    [ValidatePattern('^v\d+\.\d+\.\d+(-[0-9A-Za-z.-]+)?$')]
    [string]$Version,

    [ValidateSet('major', 'minor', 'patch')]
    [string]$Bump,

    [string]$CommitMessage,

    [string]$ReleaseNotesPath = 'release-notes.md',

    [switch]$NoCommit,
    [switch]$NoTag,
    [switch]$Push,
    [switch]$DryRun
)

$ErrorActionPreference = 'Stop'

function Invoke-CheckedGit {
    param([Parameter(ValueFromRemainingArguments = $true)][string[]]$Args)
    & git @Args
    if ($LASTEXITCODE -ne 0) {
        throw "git $($Args -join ' ') failed"
    }
}

function Get-CurrentContractVersion {
    $versionFile = Join-Path $PSScriptRoot '..\pkg\versioning\version.go'
    $content = Get-Content $versionFile -Raw
    if ($content -notmatch 'ModuleVersion\s*=\s*"(?<version>v\d+\.\d+\.\d+(?:-[0-9A-Za-z.-]+)?)"') {
        throw "Could not find ModuleVersion in $versionFile"
    }
    return $Matches.version
}

function Get-BumpedVersion {
    param(
        [string]$Current,
        [string]$Kind
    )

    if ($Current -notmatch '^v(?<major>\d+)\.(?<minor>\d+)\.(?<patch>\d+)') {
        throw "Current version '$Current' is not valid semver."
    }

    $major = [int]$Matches.major
    $minor = [int]$Matches.minor
    $patch = [int]$Matches.patch

    switch ($Kind) {
        'major' { $major++; $minor = 0; $patch = 0 }
        'minor' { $minor++; $patch = 0 }
        'patch' { $patch++ }
        default { throw "Unknown bump '$Kind'" }
    }

    return "v$major.$minor.$patch"
}

function Set-ContractVersion {
    param([string]$NewVersion)

    $versionFile = Join-Path $PSScriptRoot '..\pkg\versioning\version.go'
    $majorVersion = ($NewVersion -replace '^v(\d+)\..*$', 'v$1')
    $content = Get-Content $versionFile -Raw
    $content = $content -replace 'ModuleVersion\s*=\s*"v\d+\.\d+\.\d+(?:-[0-9A-Za-z.-]+)?"', "ModuleVersion = `"$NewVersion`""
    $content = $content -replace 'MajorVersion\s*=\s*"v\d+"', "MajorVersion  = `"$majorVersion`""
    Set-Content -Path $versionFile -Value $content -Encoding utf8
}

$repoRoot = Resolve-Path (Join-Path $PSScriptRoot '..')
Set-Location $repoRoot

Invoke-CheckedGit rev-parse --is-inside-work-tree | Out-Null

$currentVersion = Get-CurrentContractVersion
if (-not $Version) {
    if (-not $Bump) {
        throw "Provide -Version vX.Y.Z or -Bump major|minor|patch."
    }
    $Version = Get-BumpedVersion -Current $currentVersion -Kind $Bump
}

if (-not $CommitMessage) {
    $CommitMessage = "chore(release): $Version"
}

Write-Host "Current contract version: $currentVersion"
Write-Host "Target contract version:  $Version"

if ($DryRun) {
    Write-Host "Dry run: no files, commits, tags, or pushes will be changed."
    exit 0
}

Set-ContractVersion -NewVersion $Version

$notesScript = Join-Path $PSScriptRoot 'Get-ContractReleaseNotes.ps1'
& $notesScript -Version $Version -OutputPath $ReleaseNotesPath
if ($LASTEXITCODE -ne 0) {
    throw "Release note generation failed."
}

$changes = & git status --porcelain
if ($changes) {
    Write-Host "Detected changes:"
    $changes | ForEach-Object { Write-Host $_ }

    if (-not $NoCommit) {
        Invoke-CheckedGit add pkg/versioning/version.go $ReleaseNotesPath
        Invoke-CheckedGit add -A
        Invoke-CheckedGit commit -m $CommitMessage
    } else {
        Write-Host "Skipping commit because -NoCommit was supplied."
    }
} else {
    Write-Host "No file changes to commit."
}

if (-not $NoTag) {
    $existingTag = & git tag --list $Version
    if ($existingTag) {
        throw "Tag '$Version' already exists."
    }

    $tagMessage = Get-Content $ReleaseNotesPath -Raw
    Invoke-CheckedGit tag -a $Version -m $tagMessage
}

if ($Push) {
    Invoke-CheckedGit push
    if (-not $NoTag) {
        Invoke-CheckedGit push origin $Version
    }
}

Write-Host "Release preparation complete for $Version."