# Development Notes

Read [Git Workflow](git-workflow.md) for branches, PRs and checks, and [Package Boundaries](package-boundaries.md) for code ownership and dependency direction.

## Verification and versioning

Run `./scripts/powershell/Test-ReleaseAlignment.ps1` and `./scripts/powershell/Test-Contract.ps1` from the repository root. Tests run independently with `GOWORK=off`. Release metadata is declared in `go.mod`; module majors follow semantic versioning and published tags remain immutable. Consumers pin released versions in their own repositories.

## Event payload reference

This reference covers the catalogue, inventory, order and analytics payload set checked by the repository. Other payloads are defined alongside these under `pkg/contracts/pubsub`.

| # | Payload | Event type | Topic |
| --- | --- | --- | --- |
| 1 | `StockLocationAvailabilityChangedEvent` | `stock.location_availability_changed` | `stock-events` |
| 2 | `InventoryLotReceivedEvent` | `inventory.lot_received` | `stock-events` |
| 3 | `InventoryStockBucketChangedEvent` | `inventory.stock_bucket_changed` | `stock-events` |
| 4 | `InventoryPackageConvertedEvent` | `inventory.package_converted` | `stock-events` |
| 5 | `InventoryQualityAssessedEvent` | `inventory.quality_assessed` | `stock-events` |
| 6 | `InventoryReservationChangedEvent` | `inventory.reservation_changed` | `stock-events` |
| 7 | `StockStagingChangedEvent` | `inventory.staged` | `stock-events` |
| 8 | `InventorySaleCommittedEvent` | `inventory.sold` | `stock-events` |
| 9 | `InventoryDateMarkThresholdEvent` | `inventory.date_mark_threshold_reached` | `stock-events` |
| 10 | `OrderPaidEvent` | `order.paid` | `order-events` |
| 11 | `RefundCompletedEvent` | `refund.completed` | `refund-events` |
| 12 | `OrderFact` | `analytics.order_fact` | `order-events` |
| 13 | `RefundFact` | `analytics.refund_fact` | `refund-events` |
| 14 | `ProductSalesRollup` | `product.sales_performance_updated` | `product-stats` |
| 15 | `CatalogBaseCostChangedEvent` | `catalog.base_cost_changed` | `catalog-events` |
| 16 | `CatalogListingChangedEvent` | `catalog.listing_changed` | `catalog-events` |
| 17 | `OrderPackingProjection` | `fulfilment.packing_updated` | `fulfilment-events` |

## Release index

| Version | Release date | Type |
| --- | --- | --- |
| `v43.0.1` | 2026-10-06 | Patch |
| `v43.0.0` | 2026-10-06 | Major |
| `v42.0.1` | 2026-10-04 | Patch |
| `v42.0.0` | 2026-10-02 | Major |

## v43.0.1 (2026-10-06)

Add shared-wire acceptance fixtures for existing gift-card, checkout benefit
reservation, order, payment, customer allocation, receipt and refund-event
records. Cover two-card mixed funding, gift-only and external-only funding,
partial mixed refunds, ordered wallet restoration and absent, known-zero and
positive refund amounts. Integer minor units, currencies, payment identities,
wallet references and original receipt funding rows survive JSON round trips.
Gift funding remains a tender rather than a promotion discount or a reduction
of the commercial order total.

Production models, enums, JSON names and the `/v43` module path are unchanged.
Orders, Pricing and Payments retain allocation and refund authority. These
fixtures establish shared-wire acceptance, not live integration evidence;
migrations, service DTOs and pins, OpenAPI and client adoption remain outside
this release. Existing money, currency-exponent, enum and gift-card
secret-exclusion coverage continues to run in the contract gate.

## v43.0.0 (2026-10-06)

Replace `MembershipTier.MarketCode` and `MembershipAccount.MarketCode` with
`MarketCodes []string`, serialized as `market_codes` without `omitempty`. The
legacy `market_code` field is removed from these two models; no compatibility
alias, automatic conversion or data migration is included. Nil lists serialize
as `null` and empty lists as `[]`; neither grants global availability. Unrelated
owning and transaction market fields remain singular.

Membership is selected-market scoped within one tier group per country and
currency. All tiers in the group share its selected markets. Existing
`country_code` and tier threshold currency describe the group; tier keys remain
globally unique and account IDs remain retail customer numbers. The market list
also describes staff geographic access, subject to Pricing-owned role
permissions. Other tier and account fields retain their existing shapes.

This is a breaking schema release under
`github.com/Potato-Mart/Backend-Shared-Contract/v43`. Consumers adopt the module
major and their CI contract-version pins deliberately. Pricing must implement
nonempty, unique market validation, country/currency and group consistency,
account participation, selected-market availability and staff scope enforcement.
It must also establish annual qualification calendar behavior because the
account no longer carries a singular market timezone anchor. Service DTOs, BSON
persistence, migrations, OpenAPI and client projections remain consumer-owned;
publishing this contract does not activate those runtime behaviors.

## v42.0.1 (2026-10-04)

Consolidate repository documentation into OpenAPI status and three development files. Update documentation-dependent tests and release tooling to the surviving paths. Shared production model files and their JSON shapes are unchanged from v42.0.0. The patch metadata preserves immutable release payload checks without rewriting the existing tag.

Historical source and complete release records remain available through published Git tags and GitHub releases.
