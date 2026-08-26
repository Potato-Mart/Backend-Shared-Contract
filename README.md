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
v32.0.0
github.com/Potato-Mart/Backend-Shared-Contract/v32
```

See [release notes](docs/release-notes.md) for the release history,
breaking changes and consumer actions.

## Usage

Pin the latest release in the consuming service's `go.mod`:

```go
require github.com/Potato-Mart/Backend-Shared-Contract/v32 v32.0.0
```

Import packages from the same `/v32` module path.

## Package Layout

- Common models are grouped by concern under `pkg/contracts/common`, such as
  `geography`, `party`, `packaging`, `security`, `temporal`, `measurement`, and
  `money`. The legacy `common/shared` package does not exist.
- Finite enum types live in a leaf `<domain>_enums` package beside the models
  that use them. For example, product models import
  `supply/product/product_enums`, while common security models import
  `common/security/security_enums`.
- Supply classifications, import compliance, operations, market listings, and
  cost use `supply/classification`, `supply/import_compliance`,
  `supply/operations`, `supply/listing`, and `supply/cost`. The Go package name
  for import compliance is `import_compliance`.
- Commercial markets, price books, and immutable quote evidence live under
  `pricing/market`, `pricing/pricebook`, and `pricing/quote`. Promotion
  mechanics live under `pricing/promotion`; wallet models under
  `pricing/wallet`; shipment models under `orders/shipping`.
- Customer market is resolved only from the frozen
  `orders/shipping.FulfilmentLocationSnapshot`; retail analytics are standalone
  `insights/analytics` records. Notification topics, preferences, protected
  deliveries, and customer-safe in-app inbox projections live only under
  `notifications`.
- Marketing owns canonical `marketing/campaign` and `marketing/message`
  authoring models. General customer reviews live under `customers/review`,
  with separate protected, public, and owner-safe projections.
- `supply/product.Product` is keyed by the immutable `sku_code`, owns the
  authoritative storage requirement, localized content, classification,
  provenance, package options, barcode assignments, net content, status, and
  optional administration. `supply/product.SellingProduct` is the customer-safe
  market-priced projection; raw product records carry no price, tax,
  channel sellability, or market-dependent metric.
- Authoritative prices live in `pricing/pricebook.PriceEntry` inside a
  market-scoped `PriceBook`; commercial availability lives in
  `supply/listing.MarketListing`. `MarketCode` and `CountryCode` are separate
  concepts, and one country may carry several markets.
- Cross-domain catalogue and commercial links use immutable business codes:
  `sku_code`, `market_code`, `price_book_code`, `tax_category_code`, brand,
  collection, category-tag, supplier, package-option, and media codes. Root
  masters retain API `id`; references never use IDs or slugs.
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
or wire values require a new major module path. A cut-over removes the replaced
field, type, enum value, event, and test in the same change; deprecated aliases
and fallback JSON shapes are not retained. `docs/release-notes.md` is the source of
truth for release-specific JSON changes and consumer actions.

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
Go command is `GOWORK=off go test ./...`.

## Change and Release Workflow

Follow the [Git workflow](docs/git-workflow.md) for required branch, commit, push,
pull request, merge, and release-tag rules.

See [Contract Versioning](docs/contract-versioning.md) for the complete
version rules and release flow.
