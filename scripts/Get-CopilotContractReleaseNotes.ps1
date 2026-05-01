param(
    [Parameter(Mandatory = $true)]
    [ValidatePattern('^v\d+\.\d+\.\d+(-[0-9A-Za-z.-]+)?$')]
    [string]$Version,

    [Parameter(Mandatory = $true)]
    [string]$InputPath,

    [Parameter(Mandatory = $true)]
    [string]$OutputPath,

    [string]$PreviousTag
)

$ErrorActionPreference = 'Stop'

# Copilot requires a token (usually a PAT with Copilot access)
if (-not $env:GH_COPILOT_TOKEN) {
    throw 'GH_COPILOT_TOKEN is required for Copilot release notes.'
}

function Invoke-Git {
    param([Parameter(ValueFromRemainingArguments = $true)][string[]]$Args)
    $output = & git @Args 2>$null
    if ($LASTEXITCODE -ne 0) {
        return $null
    }
    return $output
}

if (-not $PreviousTag) {
    # Find the tag before the current one
    $PreviousTag = Invoke-Git describe --tags --abbrev=0 "${Version}^" 2>$null
    if (-not $PreviousTag) {
        $PreviousTag = Invoke-Git describe --tags --abbrev=0 HEAD
    }
}

$range = if ($PreviousTag) { "$PreviousTag..HEAD" } else { "HEAD" }
$deterministicNotes = Get-Content -Path $InputPath -Raw
$commits = (Invoke-Git log --no-merges '--pretty=format:%h %s' $range) -join "`n"
$changedFiles = (Invoke-Git diff --name-only $range) -join "`n"
$diff = (Invoke-Git diff --stat $range) -join "`n"

$prompt = @"
You are writing release notes for a Go module that is a shared backend contract package.

Version: $Version
Previous tag: $PreviousTag

Use the deterministic notes, commit subjects, changed files, and diff summary below.
Produce Markdown release notes for engineers consuming the contract.

Rules:
- Keep it factual. Do not invent changes.
- Call out breaking contract changes prominently.
- Explain compatibility impact in practical terms.
- Mention changed JSON field names, enum wire values, and new exported types.
- Do not include marketing language.

Deterministic notes:
$deterministicNotes

Commits:
$commits

Changed files:
$changedFiles

Diff summary:
$diff
"@

# Prepare the API Call
$headers = @{
    "Authorization" = "Bearer $env:GH_COPILOT_TOKEN"
    "Content-Type"  = "application/json"
    # Essential for GitHub Copilot API endpoints
    "Editor-Version" = "vscode/1.85.0" 
}

$body = @{
    # Copilot usually maps 'gpt-4' or 'claude-3.5-sonnet' depending on your Org settings
    model = "gpt-4" 
    messages = @(
        @{ 
            role = "system" 
            content = "You are an expert backend engineer. Write accurate, concise contract release notes." 
        },
        @{ 
            role = "user" 
            content = $prompt 
        }
    )
    temperature = 0.2
} | ConvertTo-Json -Depth 10

Write-Host "Requesting release notes from GitHub Copilot..."

$response = Invoke-RestMethod -Method Post -Uri "https://api.githubcopilot.com/chat/completions" -Headers $headers -Body $body

# Extract the content from the response
$text = $response.choices[0].message.content

if (-not $text) {
    throw 'GitHub Copilot response did not contain release note text.'
}

# Ensure directory exists and write file
$directory = Split-Path -Parent $OutputPath
if ($directory -and -not (Test-Path $directory)) {
    New-Item -ItemType Directory -Path $directory -Force | Out-Null
}

Set-Content -Path $OutputPath -Value $text.Trim() -Encoding utf8
Write-Host "Successfully generated release notes to $OutputPath"