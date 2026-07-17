# Contract Versioning

This module follows semantic versioning for shared backend contracts.

## Version Rules

- Patch, for example `v2.1.2`: comments, docs, non-breaking helper fixes, or corrections that do not change JSON shape, enum values, field names, or behavior expected by consumers.
- Minor, for example `v2.2.0`: additive contract changes such as new optional fields, new contracts, new enum values, or new helper types.
- Major, for example `v3.0.0`: breaking changes such as removed fields, renamed JSON keys, changed enum wire values, changed primitive shapes, changed package paths, or serialized model shape changes.

## Release Flow

1. Make contract changes.
2. Prepare the version metadata and release notes on a feature branch, then
   merge it to protected `main` through an approved pull request.

   From PowerShell:

   ```powershell
   ./scripts/Publish-ContractVersion.ps1 -Bump minor -Push
   ```

   From Git Bash, Linux, macOS, or GitHub Actions-style shell:

   ```bash
   bash scripts/publish-contract-version.sh --bump minor --push
   ```

   Or use an explicit version.

   ```powershell
   ./scripts/Publish-ContractVersion.ps1 -Version v2.2.0 -CommitMessage "chore(release): v2.2.0" -Push
   ```

   ```bash
   bash scripts/publish-contract-version.sh --version v2.2.0 --message "chore(release): v2.2.0" --push
   ```

   To use AI-polished release notes locally, set `OPENAI_API_KEY` and run one of these:

   ```powershell
   ./scripts/Publish-ContractVersion.ps1 -Version v3.1.0 -UseAIReleaseNotes -Push
   ```

   ```bash
   bash scripts/publish-contract-version.sh --version v3.1.0 --use-ai-release-notes --push
   ```

3. Merging to `main` triggers `.github/workflows/release.yml`. The workflow
   reruns the contract tests, verifies source-version alignment, creates the
   immutable annotated tag, and creates or repairs the matching GitHub Release.
   A manual workflow dispatch from `main` may repair an aligned release; a
   tag-only push does not trigger the workflow.

## Useful Script Options

- `-Bump major|minor|patch`: calculate the next version from `pkg/versioning/version.go`.
- `-Version vX.Y.Z`: use an exact version.
- `-NoCommit`: update files and notes without committing.
- `-NoTag`: commit but do not create a tag.
- `-Push`: push the commit and tag.
- `-DryRun`: validate the requested version without changing files.
- `-UseAIReleaseNotes`: polish deterministic release notes through OpenAI. Requires `OPENAI_API_KEY`.

Bash equivalents:

- `--bump major|minor|patch`
- `--version vX.Y.Z`
- `--no-commit`
- `--no-tag`
- `--push`
- `--dry-run`
- `--use-ai-release-notes`

## AI Release Notes

The release workflow always creates deterministic release notes first. If the repository has an `OPENAI_API_KEY` secret, the workflow can polish those notes with AI before publishing the GitHub Release.

Recommended setup:

- Add a repository secret named `OPENAI_API_KEY`.
- Optionally add a repository variable named `OPENAI_RELEASE_NOTES_MODEL`. If omitted, the scripts use `gpt-5`.

The AI step is a polish layer, not the source of truth. The input is limited to git commits, changed contract files, diff summary, and deterministic notes.

## Release Notes

Release notes are generated from git history since the previous tag and grouped into:

- Breaking Contract Changes
- Added
- Fixed
- Other Changes
- Contract Files Changed
- Compatibility Notes
