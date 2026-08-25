# Contract Versioning

This module follows semantic versioning for shared backend contracts.

## Version Rules

- Patch, for example `v2.1.2`: comments, docs, non-breaking helper fixes, or corrections that do not change JSON shape, enum values, field names, or behavior expected by consumers.
- Minor, for example `v2.2.0`: additive contract changes such as new optional fields, new contracts, new enum values, or new helper types.
- Major, for example `v3.0.0`: breaking changes such as removed fields, renamed JSON keys, changed enum wire values, changed primitive shapes, changed package paths, or serialized model shape changes.

## Release Metadata

The exact source release is the `// contract-release: vX.Y.Z` declaration at
the top of `go.mod`. The major in that declaration must match the versioned Go
module path. This keeps release metadata outside the production Go package
tree, whose contents are restricted to contract models and enums.

## Release Flow

1. Create a feature branch from the latest protected `main` and make the
   additive or breaking contract changes there.
2. In the same feature-branch pull request, align the `go.mod`
   `contract-release` declaration, `README.md`, the model manifest, and the
   existing canonical [release notes](release-notes.md). The release notes must contain both
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

Release-control changes additionally require the review policy in the
[Git workflow](git-workflow.md#release-maintainers).

## Release Notes

The canonical [release notes](release-notes.md) are the source of truth. Each release records:

- Breaking Contract Changes
- Added
- Fixed
- Other Changes
- Contract Files Changed
- Compatibility Notes
