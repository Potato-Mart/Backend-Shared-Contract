# Git Workflow

These rules apply to every commit and push in this repository.

## Branches

- Start from the latest protected `main`.
- Use `feature/<description>` for feature branches unless a release task
  specifies another branch.
- Do not push directly to `main`.

## Commits and Pull Requests

- Use a concise Conventional Commit message: `type(scope): summary`.
- Mark breaking changes with `!` or a `BREAKING CHANGE` footer.
- Run the release-alignment gate, contract gate, and `git diff --check` before
  pushing:

  ```powershell
  .\scripts\powershell\Test-ReleaseAlignment.ps1 -ExpectedVersion vX.Y.Z
  .\scripts\powershell\Test-Contract.ps1
  git diff --check
  ```
- Push only the feature branch and open a pull request targeting `main`.
- Resolve review comments and wait for the required `Go tests` check before
  merging.

## Release Maintainers

The visible `@Potato-Mart/release-maintainers` team owns the release-control
paths listed in [`.github/CODEOWNERS`](../.github/CODEOWNERS), including release
workflows, version metadata, validation, and release governance documents.

The `main` ruleset must require a pull request, code-owner review, stale-review
dismissal, and the `Go tests` check. For these paths, approval from either
release maintainer satisfies the required code-owner review.

## Releases

- Do not create or push release tags from a feature branch.
- Merging to `main` is the release boundary.
- The release workflow validates the aligned version, creates the immutable
  annotated tag, and publishes the matching GitHub release.
- Do not use the mutating publish scripts for repository releases.
