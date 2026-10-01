# Backend-Shared-Contract

Shared Go models and enums for Potato Mart. Backend services own routes, request/response DTOs, validation, authorization, persistence and business workflows.

## Latest version

```text
v41.0.0
github.com/Potato-Mart/Backend-Shared-Contract/v41
```

```go
require github.com/Potato-Mart/Backend-Shared-Contract/v41 v41.0.0
```

## Contract rules

- Production source contains reusable domain records, snapshots, value objects, typed enums and event payloads with ordinary JSON tags.
- Common models live under `pkg/contracts/common`; domain models live under their owning service packages. Finite enums use leaf `_enums` packages.
- Pricing owns commercial listings, markets, tax categories and price books. Supply owns catalogue and physical inventory evidence.
- Money uses integer minor units with an explicit currency and currency exponent. Business references use their canonical immutable keys.
- Event envelopes and typed payloads live under `pkg/contracts/pubsub`. Runtime event versions and delivery behavior remain service-owned.
- Repository tests enforce the model-only boundary, package inventory, enum coverage, wire shapes and single model ownership.

## Model references

- [Pricing listings](docs/pricing-listing-model.md)
- [Procurement costs](docs/procurement-cost-model.md)
- [Commercial evidence](docs/commercial-evidence-model.md)
- [Delivery companies](docs/delivery-company-model.md)
- [Identity access](docs/identity-access-model.md)
- [Permission catalogue](docs/permission-catalogue.md)
- [Notification templates](docs/notification-template-model.md)
- [Capture timing](docs/capture-timing-model.md)
- [Gift-card issuance](docs/gift-card-issuance-model.md)
- [Net tier-refund proof](docs/net-tier-refund-proof.md)
- [Event payloads](docs/event-model.md)

## Verification and releases

```powershell
./scripts/powershell/Test-ReleaseAlignment.ps1 -ExpectedVersion v41.0.0
./scripts/powershell/Test-Contract.ps1
git diff --check
```

The contract gate runs tests with `GOWORK=off` to verify the declared module. Release metadata is declared in `go.mod`; published tags are immutable.

See [Git workflow](docs/git-workflow.md), [versioning](docs/contract-versioning.md) and [release notes](docs/release-notes.md).
