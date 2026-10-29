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
| 18 | `StockRestockAlertTriggeredEvent` | `stock.restock_alert_triggered` | `stock-events` |
| 19 | `InventoryPOSSaleCommittedEvent` | `inventory.pos_sale_committed` | `stock-events` |

## Release index

| Version | Release date | Type |
| --- | --- | --- |
| `v47.2.0` | 2026-10-29 | Minor |
| `v47.1.0` | 2026-10-29 | Minor |
| `v47.0.0` | 2026-10-29 | Major |
| `v46.0.1` | 2026-10-27 | Patch |
| `v46.0.0` | 2026-10-24 | Major |
| `v45.5.0` | 2026-10-22 | Minor |
| `v45.4.0` | 2026-10-21 | Minor |
| `v45.3.0` | 2026-10-20 | Minor |
| `v45.2.0` | 2026-10-14 | Minor |
| `v45.1.0` | 2026-10-14 | Minor |
| `v45.0.0` | 2026-10-14 | Major |
| `v44.1.0` | 2026-10-06 | Minor |
| `v44.0.0` | 2026-10-06 | Major |
| `v43.0.1` | 2026-10-06 | Patch |
| `v43.0.0` | 2026-10-06 | Major |
| `v42.0.1` | 2026-10-04 | Patch |
| `v42.0.0` | 2026-10-02 | Major |

## v47.2.0 (2026-10-29)

- Add an unreserved POS inventory movement event that distinguishes physical issues from deficit issues without reservation identities.

## v47.1.0 (2026-10-29)

Order status history includes the optional verified actor ID, allowing physical task authorization to use the same durable staff identity as the order audit.

## v47.0.0 (2026-10-29)

Breaking picking and packing cutover: retire picking lists and item-level picking progress. Orders record picking start and completion; packing records retain order generation, quantities, containers, discrepancies, actors and timestamps. Consumers must upgrade in coordination. The changed `fulfilment.packing_updated` payload uses event version v5; Supply and Orders must cut over together.

## v46.0.1 (2026-10-27)

Correct both `campaign.changed` test registry entries to v3, matching the existing payload and Customers publisher. This patch distributes metadata/test corrections only: no event version, payload shape, Go API, module major, or consumer dependency pin changes.

A separate patch release is required because the release guard treats tests and release documentation as immutable release contents. Keep the published v46.0.0 tag unchanged. Existing Customers deployments may continue using v46.0.0; their actual publisher already emits v3.

Verification: independent-module Go tests and release alignment.

## v46.0.0 (2026-10-24)

Campaign and coupon v2 direct cutover (change `campaign-coupon-v2`). The module is now `github.com/Potato-Mart/Backend-Shared-Contract/v46`.

Campaign records support incomplete actor-owned drafts, localized description, canonical owning market and currency, retail/wholesale audience sets, explicit schedule/timezone, multiple independently ordered placements, and typed promotion/coupon identity links. Publication and effective state, revisions and timestamps are server-owned. Singular placement, severity/message style, authored navigation/CTA, series selection and generic benefit references are retired. Campaign audience is independent from offer audience and checkout authentication.

Coupon definitions separate visibility, acquisition mode and automatic cohort policy; receiving eligibility, redemption eligibility, independent/inherited scope and windows; typed benefits, currency-aware minimum spend, entitlement issuance cap and redemption limits. Permanent customer/coupon entitlement identity is separate from usage/reservations. Customers owns authority and orchestration; Pricing owns issuance and redemption policy. HTTP DTOs and routes remain in their services.

`campaign.changed` payload v3 replaces campaign lifecycle fields. Its test-only version registry records v3; correcting the former v2 registry entry does not introduce a new payload version or change the wire shape. Unchanged `coupon.changed` v1 and `promotion.changed` v3 retain their versions. Event versions are independent from this module and HTTP v2.

Release this contract before updating backend module/CI pins. Customers, Pricing, Notification, Supply and Orders deploy as a coordinated replacement, followed by Admin Web and affected POS checkout contracts. New paths use separate storage and do not read or transform pre-cutover records. No migration, backfill, adapters, dual operation or rollback is included. Historical coupon handling, including pre-cutover refund restoration, is outside this release.

All retail/wholesale storefront repositories are excluded. Clients still using retired campaign/coupon APIs may lose those features until separately updated. Pause affected admission and workers, require drained old operations/holds/events, deploy production only, smoke test under held admission, and resume only on success. Failed checks remain paused for correction.

Verification: JSON serialization and enum coverage; campaign/coupon model boundaries; unrelated contract-preservation gates; full `GOWORK=off go test -count=1 ./...` and release alignment.

## v45.5.0 (2026-10-22)

Add the additive `revision` field to `product.Product`. Supply advances it on
every successful write to the stored product document and admin clients echo
it back as an `expected_revision` precondition; `0` only appears on documents
written before Supply backfills its own store. Add the additive leaf enum
`measurement_enums.NetContentUnit` (`g`, `kg`, `mL`, `L`, `m`, `ea`) that
publishes the closed net-content unit vocabulary Supply enforces; both measures
of one net content declaration must share a dimension (mass, volume, length
or count). `measurement.MeasurementUnit` remains an open string. Existing
models are otherwise unchanged; no consumer migration is required.

## v45.4.0 (2026-10-21)

Add the additive `pkg/contracts/supply/catalogue/selection` package for curated,
ordered product lists. `ProductSelection` carries an immutable `code`, a
public `slug` editable only while `draft`, a `kind` (`badge` or `section`), a
`status` (`draft`, `published` or `archived`), localized `name` and optional
`description`, code-only `icon` and `cover` media references, an admin-set
`display_order` and a `revision` advanced by every successful write.
`ProductSelectionItem` records membership as `selection_code`, `sku_code` and
a zero-based `position`. Membership references products by SKU code and is
independent of collections and category tags. Supply owns persistence,
validation and APIs. Existing models are unchanged; no migration is required.

## v45.3.0 (2026-10-20)

Add the additive `stock.restock_alert_triggered` v1 event on the existing
`stock-events` topic, carried by `StockRestockAlertTriggeredEvent`. Supply owns
restock subscriptions and produces one event per subscriber when a SKU returns
to stock in a market; `alert_id` is the aggregate and idempotency key and the
subscription clears when it fires. The payload carries `user_id`,
`retail_customer_number` and localized `product_name` entries so Notification
can consume it and send push/inbox copy without a lookup. Orders acknowledges
it as a peer event on its stock-events subscription. Existing event types and
payloads are unchanged; no migration is required.

## v45.2.0 (2026-10-14)

Add `held`, `released` and `dropped` point-award states and the optional
`CustomerPaymentSummary.points_held` integer. Held points are excluded from
spendable wallet balances; a present zero is distinct from unavailable evidence.
Existing states remain readable for stored history. Pricing owns preparation,
calculation and wallet release; Payments owns completion/refund reconciliation.
Their HTTP DTOs and persistence policy remain service-owned.

Publish before Pricing and Payments adopt the pin. Admin, Orders, storefront
and POS projections must preserve unknown states and must not treat `held` as
spendable or `pending` as a known zero. Existing JSON decoding retains unknown
string states; `IsValid` intentionally rejects unknown values. Owner follow-up
issues track coordinated consumers; this release does not prove adoption.
No migration, backfill or event payload change is included.

## v45.1.0 (2026-10-14)

Add `coupon_enums.DistributionReferralAssigned` with wire value
`REFERRAL_ASSIGNED` for coupons issued only through qualified referral awards.
`DistributionAssigned` retains its existing `ASSIGNED` wire value and cohort
distribution meaning. `SELF_CLAIM_BY_CODE` and `AUTO_MEMBER_CLAIM` remain
unchanged. This additive enum release retains the `/v45` module path and all
existing coupon fields.

Pricing owns referral-only issuance, distribution selection, eligibility and
assignment persistence. Customers owns referral qualification and configured
coupon selection. Their service behavior and HTTP DTOs remain outside this
model-only module. The new value must not fall back to cohort assignment,
self-claim or automatic member distribution.

Publish the immutable release before Pricing and Customers adopt its module
pin. Pricing API schemas and Admin, storefront and POS enum consumers must
handle the new value before operators configure referral coupons. Older enum
validators reject it; ordinary JSON decoding preserves unknown strings without
validating or normalizing them. Consumer adoption and deployment require their
own checks and are not established by this release. There is no migration,
data reset, default coupon or new event payload.

Contract tests lock each existing and new distribution wire value, distinct
coupon round trips, `IsValid`/`String` behavior and unknown-value preservation.

## v45.0.0 (2026-10-14)

Remove vouchers as a supported instrument. Coupons retain their existing
definitions, assignments and permanent entitlement identity; gift cards, points
and other reward families retain their models. This breaking release uses
`github.com/Potato-Mart/Backend-Shared-Contract/v45` with no aliases or conversion
adapters for the removed instrument.

Remove `special.Voucher`, `reservation.VoucherBenefitReservation`,
`wallet_enums.VoucherStatus`, and `order.VoucherRedemptionSnapshot`. Remove
`RewardTypeVoucher`, `WalletInstrumentTypeVoucher` and
`CustomerPaymentAllocationKindVoucher`; their former wire values are no longer
valid. Ordinary JSON decoding of string enums does not itself validate them.

Remove `RewardBenefit.voucher_code_prefix` and `voucher_value`,
`RewardRedemptionOutcome.voucher_code`, `RewardRedemptionSnapshot.voucher_code`,
`CheckoutBenefitReservation.voucher`, `CustomerWalletSummary.active_vouchers`
and `Order.voucher_redemption`. Rename referral-profile
`RewardVouchersIssued` / `reward_vouchers_issued` to
`RewardCouponsIssued` / `reward_coupons_issued`, and `RewardVoucherCodes` /
`reward_voucher_codes` to `RewardCouponCodes` / `reward_coupon_codes`.
The coupon receipt fields retain their integer count and string-list shapes.

Remove `Reward.TriggerTierKey` / `trigger_tier_key` and
`Reward.IssueOnTierAchievement` / `issue_on_tier_achievement`. Tier gifts use
automatic coupon distribution and receiving-tier eligibility; they no longer
use the retired reward-driven issuance configuration. `MinimumTierKey` remains
the eligibility requirement for a points-redeemable reward.

Pricing and Customers must adopt the released module and implement the matching
coupon-only service behavior and referral issuance interface. Orders, Payments,
Admin Web, storefront and POS consumers must remove the retired fields and
instrument selections before claiming compatibility with the new provider APIs.
Consumers pinned to an older immutable module remain on that contract until
they deliberately adopt the new major. Publication alone does not qualify
runtime rollout or consumer adoption.

This is a fresh-data contract cutover: no data migration, deletion, reset,
runtime configuration, HTTP DTO, service policy or event workflow is added.
Issuance eligibility, referral limits, coupon selection, idempotency and
assignment persistence remain with their authoritative services. Tests cover
coupon referral JSON, supported and retired enum values, absence of voucher
production models/fields, retired reward tier-issuance fields, and the deliberate
historical-manifest removals.

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
