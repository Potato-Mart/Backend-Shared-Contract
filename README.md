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
v27.3.0
github.com/Potato-Mart/Backend-Shared-Contract/v27
```

See [RELEASE_NOTES.md](RELEASE_NOTES.md) for the release history,
breaking changes and consumer actions.

## Usage

Pin the latest release in the consuming service's `go.mod`:

```go
require github.com/Potato-Mart/Backend-Shared-Contract/v27 v27.3.0
```

Import packages from the same `/v27` module path.

## Package Layout

- Common models are grouped by concern under `pkg/contracts/common`, such as
  `geography`, `party`, `packaging`, `security`, `temporal`, `measurement`, and
  `money`. The legacy `common/shared` package does not exist in v27.
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
- `supply/product.Product` is the global conceptual product — content,
  classification, provenance, status, and optional administration only.
  `supply/product.SKU` is the sellable and stockable identity that owns package
  options, barcodes, net content, and storage type. Neither carries a price, a
  tax flag, channel sellability, or a market-dependent metric.
- Authoritative prices live in `pricing/pricebook.PriceEntry` inside a
  market-scoped `PriceBook`; commercial availability lives in
  `supply/listing.MarketListing`. `MarketID` and `CountryCode` are separate
  concepts, and one country may carry several markets.
- Cross-domain links use the scalar `sku_id`. A frozen `sku_code` survives only
  on transaction evidence: cart items, order items, order line summaries, POS
  receipt lines, purchase order items, receipt items, and supplier invoice
  lines. Transaction lines own their frozen display facts.
- Money is always `{amount_minor, currency}` in minor units with a typed
  `money.CurrencyCode`. `money.CurrencyExponent` is carried alongside because
  not every currency has two decimals.
- Promotion mechanics use open string kinds with reusable scopes, ALL/ANY
  groups, qualifier-to-target relations, typed terms, and frozen ordered
  applications. Only promotion lifecycle status and match mode are closed
  promotion enums.
- This module is contract-only. The seven backend services and the parent
  `go.work` must migrate to `/v27` before consuming this release.

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

## Repository Layout and Naming

- Use stable, descriptive filenames. Do not include numeric release or version
  tokens in source, test, script, or documentation filenames.
- Keep unit tests beside the code they test, package boundary tests beside
  their packages, repository gates in `pkg/test`, and aggregate enum tests in
  `pkg/test/enums`.
- Keep Bash scripts in `scripts/bash` and PowerShell scripts in
  `scripts/powershell`.
- Version numbers remain allowed in release tags, module paths, version
  metadata, and historical release content.

## Verification

Run the standalone contract gate so a parent Go workspace cannot alter
dependency resolution:

```powershell
./scripts/powershell/Test-Contract.ps1
```

On Bash-based systems, run `bash scripts/bash/test-contract.sh`. The equivalent
Go command is `GOWORK=off go test ./...`.

## Change and Release Workflow

Follow [GIT_WORKFLOW.md](GIT_WORKFLOW.md) for required branch, commit, push,
pull request, merge, and release-tag rules.

See [Contract Versioning](pkg/versioning/VERSIONING_RULE.md) for the complete
version rules and release flow.
