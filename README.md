# Backend-Shared-Contract

## Responsibility Summary

Backend-Shared-Contract is Potato Mart's shared Go module for reusable
cross-service data models. It contains domain entities, records, snapshots,
events, value objects, typed enums, required model-identity constants, and
model-version metadata.

The module is deliberately model-only. Backend services remain responsible for
their routes, request and response DTOs, HTTP envelopes, validation,
authorization, state transitions, calculations, normalization, and business
workflows.

## Latest Version

```text
v22.0.0
github.com/Potato-Mart/Backend-Shared-Contract/v22
```

See [RELEASE_NOTES.md](RELEASE_NOTES.md) for the release history,
breaking JSON changes and consumer actions.

## Usage

Pin the latest release in the consuming service's `go.mod`:

```go
require github.com/Potato-Mart/Backend-Shared-Contract/v22 v22.0.0
```

Import packages from the same `/v22` module path.

## Boundary Governance

Ordinary `json` struct tags and standard `encoding/json` define every shared
wire shape. The package manifest and AST boundary tests reject unclassified
exported models, non-JSON tags, custom codecs, endpoint DTO naming, paths,
scopes, free business functions, type aliases, deprecated declarations, and
non-intrinsic receiver methods. Approved receiver behavior is limited to
single-value `String` or `IsValid` enum methods.

Published model changes follow semantic versioning. Breaking exported shapes
or wire values require a new major module path. A cut-over removes the replaced
field, type, enum value, event, and test in the same change; deprecated aliases
and fallback JSON shapes are not retained. `RELEASE_NOTES.md` is the source of
truth for release-specific JSON changes and consumer actions.

## Verification

Run the standalone contract gate so a parent Go workspace cannot alter
dependency resolution:

```powershell
./scripts/Test-Contract.ps1
```

On Bash-based systems, run `bash scripts/test-contract.sh`. The equivalent Go
command is `GOWORK=off go test ./...`.

## Change Policy

Direct pushes to the protected `main` branch are prohibited. Create a feature
branch from the latest `main`, make the change, and merge it through a pull
request after the required checks pass.

Before opening the pull request, choose the next semantic version after the
latest published tag and keep `pkg/versioning/version.go`, its tests, this
README's latest version and usage example, and `RELEASE_NOTES.md` aligned. Run:

```powershell
./scripts/Test-ReleaseAlignment.ps1 -ExpectedVersion vX.Y.Z
./scripts/Test-Contract.ps1
git diff --check
```

Do not create or push the release tag from the feature branch. Merging to
`main` triggers the release workflow, which verifies the aligned version and
publishes the immutable tag and matching GitHub release.

See [Contract Versioning](pkg/versioning/VERSIONING_RULE.md) for the complete
version rules and release flow.
