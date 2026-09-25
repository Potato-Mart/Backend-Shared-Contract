# Backend-Shared-Contract

## Responsibility Summary

Backend-Shared-Contract is Potato Mart's shared Go module for reusable
cross-service data models. It contains domain entities, records, snapshots,
events, value objects, typed enums, and build metadata declared in `go.mod`.

The module is deliberately model-only. Backend services remain responsible for
their routes, request and response DTOs, HTTP envelopes, validation,
authorization, state transitions, calculations, normalization, and business
workflows.

## Latest Version

```text
v36.0.0
github.com/Potato-Mart/Backend-Shared-Contract/v36
```

See [release notes](docs/release-notes.md) for the release history,
breaking changes and consumer actions.

## Usage

Pin the latest release in the consuming service's `go.mod`:

```go
require github.com/Potato-Mart/Backend-Shared-Contract/v36 v36.0.0
```

Import packages from the same `/v36` module path.

The v36.0.0 release removes courier service-area `routing_priority` and adds
`market_code` while retaining explicit shipping-zone references. Selecting all
country zones means materializing the existing active zones, never a wildcard.
Legacy markets require explicit service-owned mapping. Derived credential
requirements may advertise `authentication_methods` with method identifiers and
required field names, never credential values or implied account-login support.

Promotion and coupon controls now share optional `audience` customer-type and
client-platform restrictions. Missing dimensions remain unrestricted; platform
is independent of order channel. The offer holds an editable snapshot, not a
live campaign link. Pricing owns validation and enforcement. Packing models and
event schemas are unchanged. See the [v36 migration](docs/v36-migration.md) for
exact shapes, compatibility, and consumer actions.

The additive v35.2.0 release adds `FulfillmentStatusCancelled` with wire value
`cancelled`. It means no further fulfilment work is scheduled; it does not
erase recorded picking, packing, shipment, or item facts and is distinct from
`fulfilled`. Consumers should accept this value before backend services emit
it, and should preserve existing physical records when applying the status.

The additive v35.1.0 contract exposes the current order `delivery_address`
override without changing the frozen checkout address in
`fulfilment_location`, adds `supply/fulfilment.OrderAllocationEvidence` for
linking an order line to its reservation, allocation, picking, and staging
facts, and adds the optional `PickingList.allocation_fingerprint` used to bind
ensure/reuse retries to the same frozen allocation set. Routes, readiness
decisions, edit validation, authorization, and audit persistence remain owned by
Orders and Supply. Previously issued invoices and shipments keep their captured
addresses.

The v35.0.0 breaking change removes the remaining API/manual `integration`
classification from `DeliveryCompany` and `DeliveryCompanyRef`. The legacy
`adapter` field is absent from both shapes; `code` is the sole, open company
identity and selects any verified Supply provider implementation.
`DeliveryCompany` retains derived, non-sensitive `credential_requirements`
metadata, which is null when the exact company code has no verified registered
provider implementation.
See the [delivery company model](docs/delivery-company-model.md) for the wire
shapes and migration guidance.

The v33.8.0 additions add optional Orders-owned shipping-zone identity and ISO
subdivision state codes to courier service areas, plus a standalone privileged
`DeliveryProviderCredentials` value model. Its optional JSON fields are
`sign_in_account`, `password`, `api_base_url`, and `api_token`; use it only on
authorized credential write and writer-only detail operations. Never embed it
in company, connection, list, or customer models. Encryption, authorization,
credential versioning, rotation, and redaction remain Supply-owned. Legacy
configured courier schedules remain readable for compatibility and are
deprecated for new availability decisions.

The v33.7.0 additions add trusted publication context to `promotion.changed`
v3 and an identity-only `coupon.changed` v1 invalidation event. Consumers must
read promotion.changed v2 and v3 before Pricing emits v3. V3 requires a known
publication context; v2 payloads omit it and retain their existing handling.
Campaign-linked benefits use opaque Promotion.ID or Coupon.ID references, not
redeemable coupon codes. Service validation, storage, publication workflows and
notification behavior remain backend-owned.

The v33.6.0 release added optional provider-discriminator fields while preserving
the then-existing `integration` API/manual mode. v34 removed the discriminator;
v35 removes the integration classification.

The v33.5.0 additions describe Supply-owned delivery companies, explicit postal
coverage filters, safe connection status, configured service windows, and frozen
delivery selections carried from Orders to Supply. See the
[delivery company model](docs/delivery-company-model.md) for migration and
validation requirements. Provider registration, credentials and booking
remain backend-owned.

The v33.4.0 additions describe the canonical Identity E.164 phone projection and
nullable phone verification timestamp. The v33.3.0 additions describe
country-scoped notification template definitions,
immutable publications and bindings, reviewed localized email/SMS/push content,
typed placeholders, and constrained email blocks and style tokens. Existing
marketing, localization and pricing shapes are unchanged.
See the [notification template model](docs/notification-template-model.md) for
wire semantics, reusable fixtures and consumer requirements. Validation,
country authorization, translation, rendering, persistence and send workflows
belong to Notification and are not implemented by this module.

## Package Layout

- Common models are grouped by concern under `pkg/contracts/common`, such as
  `geography`, `party`, `packaging`, `security`, `temporal`, `measurement`, and
  `money`. The legacy `common/shared` package does not exist.
- Finite enum types live in a leaf `<domain>_enums` package beside the models
  that use them. For example, product models import
  `supply/catalogue/product/product_enums`, while common security models import
  `common/security/security_enums`.
- Identity lives in `identity/{access,account,authorisation}`; customer records
  live in `customers/{retail,wholesale,group,preference}`. Orders, payment,
  pricing, notification, insight, marketing, and supply models use their named
  v35 domain subpackages. See the complete package map and migration table in
  [the v33 migration guide](docs/v33-migration.md).
- Catalogue models live in `supply/catalogue/{classification,product,listing,
  review,wish,favourite}`. Marketing retains `marketing/campaign`,
  `marketing/audience`, and the complete `marketing/message` package.
- Payment contracts split into `payments/payment`, `provider`, `merchant`,
  `receipt`, `register`, `settlement`, and `terminal`; there are no standalone
  invoice, refund, or finance packages. Physical outbound fulfilment is Supply
  owned, while delivery intent, rates, schedules, and order projections remain
  under Orders.
- Notification uses the singular `notification` root with `core`, `email`,
  `sms`, `push`, `preference`, `delivery`, and `template` packages. Pub/Sub uses
  `pubsub/envelope`, `pubsub/routing`, and producer-owned payload packages.
- Cross-domain catalogue and commercial links generally use immutable business
  keys: `sku_code`, `market_code`, `price_book_code`, `tax_category_code`,
  brand, collection, category-tag, supplier, package-option, and media codes.
  Marketing `BenefitRef.Code` is the explicit exception: it uses the opaque
  Promotion.ID or Coupon.ID and never a redeemable coupon code. Root masters
  retain API `id`; other references do not use IDs or slugs.
- Money is always `{amount_minor, currency}` in minor units with a typed
  `money.CurrencyCode`. `money.CurrencyExponent` is carried alongside because
  not every currency has two decimals.
- Promotion mechanics use open string kinds with reusable scopes, ALL/ANY
  groups, qualifier-to-target relations, typed terms, and frozen ordered
  applications. Only promotion lifecycle status and match mode are closed
  promotion enums.
- This module is contract-only. All eight backend services must explicitly pin
  a released contract major version before consuming it; the parent `go.work`
  uses local directory entries and does not carry a module-major import path.

## Boundary Governance

Ordinary `json` struct tags and standard `encoding/json` define every shared
wire shape. The package manifest and AST boundary tests reject unclassified
exported models, non-JSON tags, hidden persistence fields, raw provider
payloads, endpoint DTO and workflow naming, command-like Pub/Sub events,
paths, scopes, free business
functions, type aliases, deprecated declarations, and non-intrinsic receiver
methods. Approved receiver behavior is limited to single-value `String` or
`IsValid` enum methods. The sole raw JSON exception is the generic Pub/Sub
event envelope. Every production model source file contains one exported
struct, and every closed enum has its own source file in a leaf `_enums`
package.

Published model changes follow semantic versioning. Breaking exported shapes
or wire values require a new major module path. V33 preserves every exported
v32 contract and enum at one reviewed v33 destination; it does not retire
public contracts or enums. `docs/release-notes.md` is the source of truth for
release-specific JSON changes and consumer actions.

## Repository Layout and Naming

- Use stable, descriptive filenames. Do not include numeric release or version
  tokens in source, test, script, or documentation filenames.
- Keep unit tests beside the code they test, package boundary tests beside
  their packages, repository gates in `pkg/test`, and aggregate enum tests in
  `pkg/test/enums`.
- Keep Bash scripts in `scripts/bash` and PowerShell scripts in
  `scripts/powershell`.
- Version numbers remain allowed in release tags, module paths, the
  `go.mod` contract-release declaration, and historical release content.

## Verification

Run the standalone contract gate so a parent Go workspace cannot alter
dependency resolution:

```powershell
./scripts/powershell/Test-Contract.ps1
```

On Bash-based systems, run `bash scripts/bash/test-contract.sh`. The equivalent
Go command is `GOWORK=off go test -count=1 ./...`.

## Change and Release Workflow

Follow the [Git workflow](docs/git-workflow.md) for required branch, commit, push,
pull request, merge, and release-tag rules.

See [Contract Versioning](docs/contract-versioning.md) for the complete
version rules and release flow.
