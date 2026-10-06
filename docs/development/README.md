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
| `v44.1.0` | 2026-10-06 | Minor |
| `v44.0.0` | 2026-10-06 | Major |
| `v43.0.1` | 2026-10-06 | Patch |
| `v43.0.0` | 2026-10-06 | Major |
| `v42.0.1` | 2026-10-04 | Patch |
| `v42.0.0` | 2026-10-02 | Major |

## v44.1.0 (2026-10-06)

Add `AuthMethodPIN` with wire value `pin` for reusable staff PIN authentication
through Identity's authorized POS flow. Existing authentication-method values,
`LoginSession` fields, identity providers and assurance levels are unchanged.
The value grants no MFA or assurance upgrade. No barcode authentication-method
value is introduced.

Add optional `UserProfile.employee_id` as a string, matching the existing
`AdminAccountProfile.employee_id`. It represents the same Identity-assigned
eight-digit workforce identifier, preserves leading zeros and is omitted for
customers and legacy records until assigned. Identity owns generation,
uniqueness, backfill, persistence and authorized projection. The identifier is
not a PIN or hash, and authorized profile exposure does not establish secrecy.

This additive model release retains the `/v44` module path. JSON fixtures cover
PIN and existing session methods, unknown strings, omitted legacy methods,
unchanged session evidence and employee-ID string/omission semantics. The
module remains dependency-free and JSON-model-only; BSON mapping and refresh
preservation tests belong to Identity. No routes, credentials, validators,
token claims, persistence rules or runtime authentication are added here.

Read-only source inventory on 2026-10-06 found Identity and Customers pinned to
`/v43 v43.0.0`; Orders, Payments and Pricing to `/v44 v44.0.0`; Supply to
`/v42 v42.0.0`; Insights to `/v39 v39.0.0`; and Notification to `/v40 v40.0.0`.
Current service CI declarations have no separate hardcoded contract-version
input. No unkeyed `UserProfile` construction was found in these services.

Before PIN emission, Identity must deliberately adopt the released module,
implement its service-owned credential/session behavior and employee-ID
projection, and retain the original sign-in method through refresh. Admin Web
must widen its closed authentication-method union, add the profile field and
localized PIN label, and provide an unknown-method display fallback. Its
generated browser snapshot must follow Identity's updated API schema. POS
currently strips unknown profile fields and must explicitly retain employee-ID
strings when adopting that projection. Customers' ordinary JSON decoding
tolerates additive fields, but intended employee-ID propagation needs explicit
mapping updates. Retail Web's closed session union needs review if workforce
sessions can reach it; retail iOS has an unknown-method fallback and Android
ignores unknown JSON fields. Wholesale Web has no affected projection, and
wholesale native repositories remain empty or planning-only.

Consumer repositories were inspected without edits or consumer gate execution.
These findings establish the scoped schema-publication order, not deployed
compatibility or PIN readiness. Consumer qualification and Identity admission
controls must precede new emission. Admin sign-in remains email and password;
PIN management and POS authentication are separate owning-service/client work.
Adopting `/v44` from an older major also requires the independent campaign and
other migration review already described in earlier releases.

## v44.0.0 (2026-10-06)

Replace campaign placement vocabulary with four canonical surfaces:
`announcement_bar` is the thin top header strip; `home_banner` is the web home
hero or mobile home banner above the collection rail; `account_banner` remains
the mobile Account banner below the order-status strip; and `home_modal` is the
initial home-page popup.

Remove the constants and validity support for `top_banner`, `home_hero`,
`modal`, `checkout_notice` and `product_notice`. There are no aliases, automatic
conversions, normalization helpers or default placement. Ordinary JSON decoding
preserves unrecognized strings, including retired values, but `IsValid` rejects
them. Decoding tolerance does not grant semantic or rendering support.

This breaking release uses
`github.com/Potato-Mart/Backend-Shared-Contract/v44`. Campaign fields and wire
names, CTA text/href/typed destinations, benefit references, localization, media,
targeting, scheduling, lifecycle and audit evidence retain their shapes. The
forecasting placement field remains a string-valued placement. The
`CampaignChangedEvent` payload, event type/topic and event version are unchanged;
the event contains no placement or authored-content field. Fixtures cover all
four surfaces, invalid legacy/future values and unchanged campaign/event evidence.

This is a Shared-Contract-only cutover. Customers and existing clients still use
legacy queries and authoring rules. Their adoption must update module and CI
pins, DTOs, generated OpenAPI, authoring, queries and rendering. Customers must
review migration of the three renamed surfaces and archive or otherwise review
checkout/product notices; this release performs no record migration. Pricing
retains eligibility and benefit authority. Publishing v44 establishes the schema
without activating runtime adoption or proving live rendering compatibility.

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
