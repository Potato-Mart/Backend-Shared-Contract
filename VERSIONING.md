# Contract Versioning

This module follows semantic versioning for shared backend contracts.

## Version Rules

- Patch, for example `v2.1.2`: comments, docs, non-breaking helper fixes, or corrections that do not change JSON shape, enum values, field names, or behavior expected by consumers.
- Minor, for example `v2.2.0`: additive contract changes such as new optional fields, new contracts, new enum values, or new helper types.
- Major, for example `v3.0.0`: breaking changes such as removed fields, renamed JSON keys, changed enum wire values, changed primitive shapes, changed package paths, or response shape changes.

## Release Flow

1. Make contract changes.
2. Run tests:

   ```powershell
   $env:GOCACHE="$PWD\.gocache"
   go test ./...
   go vet ./...
   ```

3. Prepare the version, commit, and tag:

   ```powershell
   ./scripts/Publish-ContractVersion.ps1 -Bump minor -Push
   ```

   Or use an explicit version:

   ```powershell
   ./scripts/Publish-ContractVersion.ps1 -Version v2.2.0 -CommitMessage "chore(release): v2.2.0" -Push
   ```

   To use AI-polished release notes locally, set `OPENAI_API_KEY` and run:

   ```powershell
   ./scripts/Publish-ContractVersion.ps1 -Version v3.1.0 -UseAIReleaseNotes -Push
   ```

4. Pushing the tag triggers `.github/workflows/release.yml`, which runs tests, generates release notes, and creates or updates the GitHub Release.

## Useful Script Options

- `-Bump major|minor|patch`: calculate the next version from `pkg/versioning/version.go`.
- `-Version vX.Y.Z`: use an exact version.
- `-NoCommit`: update files and notes without committing.
- `-NoTag`: commit but do not create a tag.
- `-Push`: push the commit and tag.
- `-DryRun`: validate the requested version without changing files.
- `-UseAIReleaseNotes`: polish deterministic release notes through OpenAI. Requires `OPENAI_API_KEY`.

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
