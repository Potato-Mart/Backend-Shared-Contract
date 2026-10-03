# Backend-Shared-Contract

Shared Go models and enums for Potato Mart. Backend services own routes, request/response DTOs, validation, authorization, persistence and business workflows.

## Latest version

```text
v42.0.1
github.com/Potato-Mart/Backend-Shared-Contract/v42
```

```go
require github.com/Potato-Mart/Backend-Shared-Contract/v42 v42.0.1
```

## Contract rules

- Production source contains reusable domain records, snapshots, value objects, typed enums and event payloads with ordinary JSON tags.
- Common models live under `pkg/contracts/common`; domain models live under their owning service packages. Finite enums use leaf `_enums` packages.
- Pricing owns commercial listings, markets, tax categories and price books. Supply owns catalogue and physical inventory evidence.
- Money uses integer minor units with an explicit currency and currency exponent. Business references use their canonical immutable keys.
- Event envelopes and typed payloads live under `pkg/contracts/pubsub`. Runtime event versions and delivery behavior remain service-owned.
- Repository tests enforce the model-only boundary, package inventory, enum coverage, wire shapes and single model ownership.

## Model references

Current model and enum definitions live under [pkg/contracts](pkg/contracts/).
The checked event payload table is in [Development Notes](docs/development/README.md#event-payload-reference).

## Verification and releases

```powershell
./scripts/powershell/Test-ReleaseAlignment.ps1 -ExpectedVersion v42.0.1
./scripts/powershell/Test-Contract.ps1
git diff --check
```

The contract gate runs tests with `GOWORK=off` to verify the declared module. Release metadata is declared in `go.mod`; published tags are immutable.

See [Git Workflow](docs/development/git-workflow.md) and [Development Notes](docs/development/README.md) for versioning and the current release notes.

## Repository documentation

See [Documentation](docs/README.md) for OpenAPI usage, [Git Workflow](docs/development/git-workflow.md), [Package Boundaries](docs/development/package-boundaries.md), and [Development Notes](docs/development/README.md).
