# Contract Versioning

This module follows semantic versioning for shared backend contracts.

## Version Rules

- Patch, for example `v2.1.2`: comments, docs, non-breaking helper fixes, or corrections that do not change JSON shape, enum values, field names, or behavior expected by consumers.
- Minor, for example `v2.2.0`: additive contract changes such as new optional fields, new contracts, new enum values, or new helper types.
- Major, for example `v3.0.0`: breaking changes such as removed fields, renamed JSON keys, changed enum wire values, changed primitive shapes, changed package paths, or serialized model shape changes.

## Release Flow

1. Create a feature branch from the latest protected `main` and make the
   additive or breaking contract changes there.
2. In the same feature-branch pull request, align `pkg/versioning/version.go`,
   `pkg/versioning/version_test.go`, `README.md`, the model manifest, and the
   existing canonical `RELEASE_NOTES.md`. The release notes must contain both
   a release-index row and a detailed `## vX.Y.Z` section.
3. Run the release-alignment and contract gates before pushing the branch:

   ```powershell
   ./scripts/powershell/Test-ReleaseAlignment.ps1 -ExpectedVersion vX.Y.Z
   ./scripts/powershell/Test-Contract.ps1
   git diff --check
   ```

4. Push only the feature branch, open an approved pull request, resolve review
   threads, and wait for the required `Go tests` check. Do not create or push
   the version tag from the feature branch.
5. Merging to `main` triggers `.github/workflows/release.yml`. The workflow
   reruns the contract tests, verifies source-version alignment, creates the
   immutable annotated tag, and creates or repairs the matching GitHub Release.
   A manual workflow dispatch from `main` may repair an aligned release; a
   tag-only push does not trigger the workflow.

The mutating `scripts/powershell/Publish-ContractVersion.ps1` and
`scripts/bash/publish-contract-version.sh` scripts predate this protected-main workflow.
They generate a standalone `release-notes.md` and may create a tag before a
pull request is merged. Do not use their mutating or push modes for repository
releases. `-DryRun` / `--dry-run` remains safe for version calculation only.

## Safe Local Version Calculation

The legacy preparation scripts may be used only in non-mutating dry-run mode:

```powershell
./scripts/powershell/Publish-ContractVersion.ps1 -Bump minor -DryRun
```

```bash
bash scripts/bash/publish-contract-version.sh --bump minor --dry-run
```

Prepare the actual version metadata and canonical release notes explicitly in
the protected-main pull request as described above.

## Release Notes

`RELEASE_NOTES.md` is the source of truth. Each release records:

- Breaking Contract Changes
- Added
- Fixed
- Other Changes
- Contract Files Changed
- Compatibility Notes
