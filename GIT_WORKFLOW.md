# Git Workflow

These rules apply to every commit and push in this repository.

## Branches

- Start from the latest protected `main`.
- Use `feature/<description>` for feature branches.
- For this release, use `feature/repo-cleanup-v22.2.0`.
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

## Releases

- Do not create or push release tags from a feature branch.
- Merging to `main` is the release boundary.
- The release workflow validates the aligned version, creates the immutable
  annotated tag, and publishes the matching GitHub release.
- Do not use the mutating publish scripts for repository releases.
