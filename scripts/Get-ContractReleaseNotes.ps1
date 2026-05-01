param(
    [Parameter(Mandatory = $true)]
    [ValidatePattern('^v\d+\.\d+\.\d+(-[0-9A-Za-z.-]+)?$')]
    [string]$Version,

    [string]$PreviousTag,

    [string]$OutputPath
)

$ErrorActionPreference = 'Stop'

function Invoke-Git {
    param([Parameter(ValueFromRemainingArguments = $true)][string[]]$Args)
    $output = & git @Args 2>$null
    if ($LASTEXITCODE -ne 0) {
        return $null
    }
    return $output
}

if (-not $PreviousTag) {
    $PreviousTag = Invoke-Git describe --tags --abbrev=0 "$Version^"
    if (-not $PreviousTag) {
        $PreviousTag = Invoke-Git describe --tags --abbrev=0
    }
}

$range = if ($PreviousTag) { "$PreviousTag..HEAD" } else { "HEAD" }
$commits = Invoke-Git log --no-merges '--pretty=format:- %s (%h)' $range
$changedFiles = Invoke-Git diff --name-only $range

$breaking = @()
$features = @()
$fixes = @()
$other = @()

foreach ($commit in @($commits)) {
    if (-not $commit) { continue }
    switch -Regex ($commit) {
        'BREAKING CHANGE|!:' { $breaking += $commit; break }
        '^- feat(\(.+\))?:' { $features += $commit; break }
        '^- fix(\(.+\))?:' { $fixes += $commit; break }
        default { $other += $commit }
    }
}

$notes = New-Object System.Collections.Generic.List[string]
$notes.Add("# $Version")
$notes.Add("")
if ($PreviousTag) {
    $notes.Add("Changes since `$PreviousTag`.")
} else {
    $notes.Add("Initial tracked release notes for this contract version.")
}
$notes.Add("")

if ($breaking.Count -gt 0) {
    $notes.Add("## Breaking Contract Changes")
    $breaking | ForEach-Object { $notes.Add($_) }
    $notes.Add("")
}

if ($features.Count -gt 0) {
    $notes.Add("## Added")
    $features | ForEach-Object { $notes.Add($_) }
    $notes.Add("")
}

if ($fixes.Count -gt 0) {
    $notes.Add("## Fixed")
    $fixes | ForEach-Object { $notes.Add($_) }
    $notes.Add("")
}

if ($other.Count -gt 0) {
    $notes.Add("## Other Changes")
    $other | ForEach-Object { $notes.Add($_) }
    $notes.Add("")
}

if ($changedFiles) {
    $contractFiles = @($changedFiles | Where-Object { $_ -like 'pkg/*' -or $_ -like 'VERSIONING.md' -or $_ -like 'go.mod' })
    if ($contractFiles.Count -gt 0) {
        $notes.Add("## Contract Files Changed")
        $contractFiles | Sort-Object | ForEach-Object { $notes.Add("- ``$_``") }
        $notes.Add("")
    }
}

$notes.Add("## Compatibility Notes")
$notes.Add("- Consumers should upgrade deliberately and run contract serialization/deserialization tests.")
$notes.Add("- Major versions may include removed fields, renamed JSON keys, enum value changes, or changed primitive shapes.")
$notes.Add("- Minor versions may add fields, contracts, enum values, or helper types in a backward-compatible way.")
$notes.Add("- Patch versions should contain documentation, comments, helper fixes, or non-breaking corrections.")

$text = $notes -join [Environment]::NewLine

if ($OutputPath) {
    $directory = Split-Path -Parent $OutputPath
    if ($directory) {
        New-Item -ItemType Directory -Path $directory -Force | Out-Null
    }
    Set-Content -Path $OutputPath -Value $text -Encoding utf8
} else {
    Write-Output $text
}