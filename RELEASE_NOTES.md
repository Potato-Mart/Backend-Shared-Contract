# Version Log / 版本紀錄

Backend-Shared-Contract is the shared contract layer for the Potato Mart backend ecosystem. This file records public contract changes, upgrade impact, and consumer actions for backend services, web clients, mobile clients, and future service-to-service integrations.

Backend-Shared-Contract 是土豆商城後端生態系的共用契約層。本文件記錄公開契約變更、遷移影響，以及後端服務、前端、行動端與未來服務間整合需要採取的動作。

## Governance / 治理原則

- This module contains reusable JSON domain entities, records, snapshots, events, value objects, typed enums, required model-identity constants, and single-value `String`/`IsValid` enum methods only.
- HTTP/API wire DTOs, response envelopes, command payloads, and backend-specific request/response structs belong in the owning backend service.
- It must not depend on web frameworks, authentication middleware, runtime service implementations, non-JSON struct tags, or custom codecs.
- Semantic versioning is enforced. Any removal, rename, JSON shape change, module path change, or incompatible exported type change requires a major version.
- Consumers should pin a released module tag and review the "Consumer Action / 使用方動作" section before upgrading.
- Remote release history was reconciled from GitHub tags in `Potato-Mart/Backend-Shared-Contract` on 2026-06-18.

- 本模組只包含可重用的 JSON domain entity、record、snapshot、event、value object、typed enum、必要的模型識別常數，以及單一值 `String`/`IsValid` enum 方法。
- HTTP/API wire DTO、回應信封、command payload，以及後端專屬 request/response struct 應由各自擁有的後端服務維護。
- 本模組不得依賴 Web 框架、身份驗證 middleware、服務執行期實作、非 JSON struct tag 或自訂 codec。
- 本模組遵循 semantic versioning。任何移除、改名、JSON shape 改變、module path 改變，或不相容的 exported type 變更，都必須升 major version。
- 使用方應固定依賴已發布 tag，並在升級前閱讀 "Consumer Action / 使用方動作"。

## Release Index / 發布索引

| Version | Release date | Type | Impact |
| --- |--------------| --- | --- |
| `v23.0.0` | 2026-08-07 | Major | Domain-oriented `pkg` layout hard cutover: moves contracts and enums into cohesive domain packages, centralizes routed Pub/Sub events, changes the module path to `/v23`, and requires downstream consumers to migrate explicitly. No JSON fields, tags, enum values, event versions, or runtime behavior change. |
| `v22.2.0` | 2026-08-07 | Minor | Repository cleanup and requested warehouse date-mark cutover: removes `InventoryDateMarkUseBy` / `USE_BY`, reorganizes repository and enum tests, groups scripts by language, and adds filename and Git workflow rules. Keeps the `/v22` module path; consumers must complete the documented breaking migration. |
| `v22.1.0` | 2026-08-06 | Minor | Adds the customer-safe public AU commercial projection and expiry display fields for price, canonical `EACH` package, aggregate stock state, market/freshness, soon-expiry thresholds, localized labels, and optional exact date. Keeps the `/v22` module path. |
| `v22.0.1` | 2026-08-05 | Patch | Locks the existing V22 geography, buyer, `EACH` offer, availability, revision, and stock/dependency error primitives required by the backend gate-closure wave. No exported model, JSON shape, enum value, or module-path change. |
| `v22.0.0` | 2026-08-04 | Major | JSON-only geography, depot, package-aware inventory, consolidated group fulfilment, and geographically scoped campaign/promotion cut-over. Replaces superseded fields and requires every consumer to adopt `/v22`. |
| `v21.1.0` | 2026-07-30 | Minor | Additive backend-gap models for persistent notification read state and consent provenance, refund-linked coupon usage, redemption/wallet timestamps, point award and debt summaries, secret-free voucher claim delivery, and versioned gift-card denomination policy. Keeps the `/v21` module path. |
| `v21.0.0` | 2026-07-29 | Major | Delivery, campaign, native notification, and wallet-policy cutover: adds revisioned delivery schedules/rates/preferences, campaign-promotion/media/typed-CTA linkage, safe storefront events, typed campaign notification references and push values; removes quiet-hours and requires every consumer to adopt `/v21`. |
| `v20.0.0` | 2026-07-29 | Major | Catalog brand identity and SKU relationship hard cut: canonical brands use Mongo-backed IDs, immutable slugs, localized names, and optional logo URLs; brand keys, summaries, counters, and embedded SKU product lists are removed. Requires Supply to adopt `/v20` and consumers to migrate explicitly. |
| `v19.0.0` | 2026-07-27 | Major | Admin Portal hard cutover: retail-only membership keyed by customer number, non-membership benefit ownership, wholesale applications/freight presets, qualified product placement, managed media visibility, analytical facts, and complete removal of deprecated shapes and `sort_order`. Requires every consumer to adopt `/v19`. |
| `v18.6.1` | 2026-07-22 | Patch | Adds optional buyer-identity fields (`retail_customer_number`, `organisation_access_id`) to `payments.PaymentFailedEvent` and `payments.RefundCompletedEvent` so notification consumers need no local buyer lookup. Additive optional fields only; no exported-type, enum, or module-path change. |
| `v18.6.0` | 2026-07-22 | Minor | Completes the eventing model surface (stock, fulfilment, customer, catalog, engagement, product-stats payloads + enriched order/refund events and invoice-issued), adds customer payment summary/allocation, invoice resend, membership tier progress + typed benefits, storefront origin/physical weight, and the reuse-first POS surface (registers/shifts/cash movements/receipt snapshots, cashier role, Stripe terminal provider, POS attribution). Additive only; keeps the `/v18` module path. |
| `v18.5.0` | 2026-07-21 | Minor | Adds the Pub/Sub event envelope (`contracts/events.EventEnvelope`), `EventTopic`/`EventType` enums, order lifecycle events (`contracts/sales`), and payment/refund events (`contracts/payments`) for the seven-service migration's eventing backbone. Additive only; keeps the `/v18` module path. |
| `v18.4.1` | 2026-07-21 | Patch | Consolidates version history in this file, streamlines the README, and documents the protected-main release workflow. Documentation and version metadata only; no exported model, JSON shape, enum value, or module-path change. |
| `v18.4.0` | 2026-07-20 | Minor | Adds customer-safe product supply/manufacturing provenance, ordered detail imagery, nullable display selling counts, audience-specific brand counts, rating/review models, and Make a Wish proposal/ballot models. Additive only; retains `supplier_code` and the `/v18` module path. |
| `v18.3.0` | 2026-07-19 | Minor | Adds the customer-safe brand catalogue summary and optional immutable `brand_key` fields on brand masters, references, and storefront products. Additive only; keeps existing brand JSON and the `/v18` module path. |
| `v18.2.0` | 2026-07-19 | Minor | Adds canonical localized product brand masters, lightweight brand references, and optional product/snapshot/storefront `brand_ref` fields. Additive only; keeps legacy `brand` arrays and the `/v18` module path. |
| `v18.1.0` | 2026-07-18 | Minor | Adds typed customer notification topics for order, payment, packing, dispatch, delivery, and invoice lifecycle milestones. Additive only; keeps the `/v18` module path. |
| `v18.0.0` | 2026-07-18 | Major | Customer-commerce hard cut: adds unified favourite-list, notification lifecycle, customer-safe product projection, promotion badge, and historical sales-ranking models; removes legacy editable product velocity fields and Commerce wholesale-list permissions; changes the module path to `/v18`. |
| `v17.4.0` | 2026-07-17 | Minor | Account security and wallet-pass models: adds session-bound access-token claims, exact device last-login IP, provider-neutral membership-pass content/barcode enums, and wholesale fixed/on-request price modes. Additive only; keeps the `/v17` module path. |
| `v17.3.0` | 2026-07-16 | Minor | Storefront customer support: adds optional stable IDs to persisted contact-address book entries and a customer-safe promotion catalogue projection that omits rule-engine internals. Additive only; keeps the `/v17` module path. |
| `v17.2.1` | 2026-07-16 | Patch | Release-alignment publication of the existing V17.2 import-compliance model surface. No JSON shape, enum value, exported contract, or module-path change. |
| `v17.2.0` | 2026-07-16 | Minor | Import-compliance model foundation: adds revisioned settings, manufacturer declarations, label masters, tariff profiles/assessments, trademark evidence, RFI records, immutable source snapshots, cited evidence/catalogue references, and generated-artifact references with fixed-point monetary, rate, exchange, weight, and volume fields. Additive only; keeps the `/v17` module path. |
| `v17.1.0` | 2026-07-16 | Minor | Receipt-safe promotion messaging: adds explicit customer-facing localized receipt copy and an opt-in print flag to promotions, plus a buyer/POS-safe `ReceiptOffer` projection that omits internal rules, discount configuration, counters, metadata, and authoring copy. Additive only; keeps the `/v17` module path. |
| `v17.0.0` | 2026-07-15 | Major | Retail wallet and checkout-benefit hard cut: changes the module path to `/v17`; removes all wallet-export contracts and enums; generalizes coupon ownership; adds reservation-aware vouchers and gift cards, explicit gift-card balances, order redemption snapshots, and the `gift_card` payment method/provider reference. |
| `v16.0.0` | 2026-07-13 | Major | Model-only hard cut: changes the module path to `/v16`; removes endpoint DTOs, paths, service scopes, HTTP envelopes, business functions/mappings/transitions, forbidden packages, standalone preorder models, and compatibility aliases; adds v16 identity, order-owned preorder, campaign prediction, packing/stock-arrival, and persisted discount decision models plus an AST/manifest boundary gate. |
| `v15.2.0` | 2026-07-12 | Minor | Durable customer notifications and idempotent stock operations: adds customer notification topics/channels/delivery states, preorder availability and delivery receipt contracts, customer/internal notification paths, the `notification:send` scope, and atomic reservation/packing-settlement commands and results. Additive only; keeps the `/v15` module path. |
| `v15.1.0` | 2026-07-11 | Minor | Outbound delivery completion: adds the `delivered` outbound-shipment status and optional `delivered_at` timestamp. Additive only; keeps the `/v15` module path. |
| `v15.0.0` | 2026-07-09 | Major | Retail account and portal cleanup: changes the module path to `/v15`, renames the general customer account persona to `retailCustomer`, renames storefront/partner portal wire values to `retail`/`wholesale`, removes legacy contact and activity timeline fields, adds retail profile completion/gender/billing/referral fields, and embeds order packing progress on sales orders. |
| `v14.1.0` | 2026-07-09 | Minor | Storefront profile and catalogue slugs: adds optional product collection/category slug fields for public storefront URLs and optional avatar fields on `identity.UserProfile`. Additive only; keeps the `/v14` module path. |
| `v14.0.0` | 2026-07-08 | Major | Wholesale customer consolidation: makes wholesale customer a compatibility name for the wholesale organisation/business account, moves stable sign-in identity references onto the organisation principal, makes organisation access the people/team record, removes separate wholesale customer number/customer id response fields, and requires `/v13` → `/v14` module-path migration. |
| `v13.3.0` | 2026-07-08 | Minor | Back-in-stock notifications: adds account-only SKU subscription contracts, email/SMS channel and lifecycle enums, consent snapshot and delivery-error metadata, a restock event payload, and the `restock:notify` service-auth scope for Operations to trigger Management-owned notification processing. Additive only; keeps the `/v13` module path. |
| `v13.2.0` | 2026-07-08 | Minor | Paid preorder checkout support: adds optional `sales.CartItem.properties` metadata so Commerce can carry server-validated preorder fulfilment markers from cart lines into order lines and skip immediate stock reservation for preorder-open products. Additive only; keeps the `/v13` module path. |
| `v13.1.0` | 2026-07-07 | Minor | Group-order buyer discount and manager permissions: adds the shared per-group-order discount application/approval domain model (`promotion.GroupOrderDiscountApplication`/`GroupOrderDiscountProposal`), its lifecycle enum (`promotionenum.GroupOrderDiscountState`), the internal discount-read endpoint constant `promotion.PathGroupOrderDiscountInternal`, the `promotion:grant` service-auth scope, and four wholesale group-order manager permissions. Additive only; keeps the `/v13` module path. |
| `v13.0.0` | 2026-07-06 | Major | V13 module path for reward tier-trigger fields and release engineering cleanup: adds membership reward tier-achievement issue fields, splits enum tests into small domain files, adds standalone `GOWORK=off` test scripts/CI, and requires `/v12` → `/v13` module-path migration. |
| `v12.0.0` | 2026-07-06 | Major | Breaking enum package cleanup: splits the flat enum package into domain enum subpackages, hard-moves remaining contract/API/service-auth typed enums, removes legacy wallet enum aliases, and removes provider-specific terminal adapter constants from the contract repo. Requires `/v11` → `/v12` module-path migration. |
| `v11.2.0` | 2026-07-06 | Minor | Sales delivery scheduling and payment processor fees: adds order `expected_delivery_date`/`expected_delivery_time`, `delivery_method` (new `DeliveryMethod` enum) with `outsourced_carrier`, derived `delivery_region` (new `DeliveryRegion` enum), shipping zone `is_local`, payment `processor_fee`/`net_amount`, Stripe `balance_transaction_id`, and the canonical MAMA order-number rule in `pkg/logic/sales`. Additive only; no module path change. |
| `v11.1.0` | 2026-07-06 | Minor | Retail preorder and expiry merchandising contracts: adds preorder interest/status contracts, product preorder policy fields, admin-configurable soon-expiry merchandising policy, storefront-safe preorder/expiry display fields, and wholesale storefront projection support. Additive only; no module path change. |
| `v11.0.0` | 2026-07-05 | Major | Unified coupon and wallet export contracts: removes public `CustomerCoupon`; adds Coupon-owned assignment/detail/issue/recipient-preview contracts; adds wallet export request/status/result models with `schema_version: "wallet_export_v1"` covering membership points, coupons, vouchers, gift cards, rewards, normalized history, filters, row counts, checksum, and status. Requires `/v10` → `/v11` module-path migration. |
| `v10.1.1` | 2026-07-02 | Patch | Identity claim alignment: adds optional `retail_customer_number` to `identity.AccessTokenClaims` so services can carry the retail customer business number in access-token claims. Additive only; no module path change, migration, or breaking JSON change. |
| `v10.1.0` | 2026-06-30 | Minor | Product category/catalogue cleanup: removes `category_key`, `category_path`, `catalogue`, and nested `merchandising`; adds product/wholesale `collection` object and top-level localized `category_tags`; renames product-list service auth and wholesale permissions from catalog/catalogue to products; replaces category promotion/coupon targeting with category-tag ID/localized-name contracts. Breaking wire-shape change shipped under requested `v10.1.0`. |
| `v10.0.0` | 2026-06-30  | Major | Breaking `/v9`→`/v10` module path. Completes the canonical reference migration: removes every deprecated id/legacy alias and converts remaining cross-struct references to code/number business keys (product→`sku_code`, order→`order_number`, depot→`depot_code`, supplier→`supplier_code`, coupon→`coupon_code`, reward→`reward_code`, retail customer→`customer_number`, wholesale org→`organisation_code`, device→`device_key`). Removes storage-driver struct tags, flattens customer identity keys and product secondary keys, adds `common.PartyRef.Code`, makes `Reward.Code` required, and removes legacy product/payment/user compatibility fields. Hard cutover — no legacy decode/fallback. |
| `v9.4.0` | 2026-06-27   | Minor | Canonical code/number reference migration: product refs add `product_sku_code`, order refs add `order_number`, depot/supplier refs add code fields, product display fields use localized description/brand arrays, and product vendor is renamed to supplier with legacy decode/writeback compatibility |
| `v9.3.0` | 2026-06-27   | Minor | Added required SKU `primary_name` as a singular `common.LocalizedName`; SKU `other_names` remains as optional alternate localized names |
| `v9.0.0` | 2026-06-25   | Major | Contract hygiene: inline enums relocated to `pkg/enums`; non-struct logic (promotion resolver, product/payments/promotion/membership/campaign helpers) moved to `pkg/logic`; contract files are structs-only. Requires the `/v9` module path migration. Also adds (additive) shared buyer/commercial context — `BuyerType`/`PriceAudience`/`PriceVisibility`/`FulfilmentIntent` enums, `sales.BuyerContext`/`PricingContext`, `Cart.Channel`/`buyer`, item `pricing`, and `product.Selling` — with POS treated as a channel, not a buyer type |
| `v8.1.0` | 2026-06-24   | Minor | Added product `taxed` field for GST/FRE invoice rendering (additive) |
| `v8.0.0` | 2026-06-23   | Major | Global membership consolidation; loyalty/subscription contracts moved into membership; wholesale membership renamed to organisation access; points reservation and reward redemption contracts |
| `v7.0.0` | 2026-06-23   | Major | V7 module path; Product enterprise redesign; contract DTO cleanup removing shared wire/action structs and moving API envelopes, pagination, service-token, pricing, stockops, media, payment terminal/settlement, identity/access, and packing settlement payloads into owning backends |
| `v6.0.2` | 2026-06-18   | Patch | Version metadata correction (`ModuleVersion = "v6.0.2"`) |
| `v6.0.1` | 2026-06-18   | Minor | Added product `description` field (additive) |
| `v6.0.0` | 2026-06-18   | Major | Staged breaking release: V6 module path, identity/access model, retail/wholesale split, grouped support fields |
| `v5.6.0` | 2026-06-17   | Minor | Contract history, stock movement, loyalty expiry models |
| `v5.5.2` | 2026-06-16   | Patch | Payment method correction/extension |
| `v5.5.1` | 2026-06-15   | Patch | Payment method extension in sales contracts |
| `v5.5.0` | 2026-06-15   | Minor | Device tracking and customer segment refinement |
| `v5.4.0` | 2026-06-15   | Minor | Common party reference reuse for company/customer/supplier models |
| `v5.3.0` | 2026-06-15   | Minor | Company/customer shared detail, device detection, collections, security logs |
| `v5.2.0` | 2026-06-12   | Minor | Promotions, category tags, product lifecycle, effective promotion resolver |
| `v5.1.2` | 2026-06-12   | Patch | Product/SKU field refinement |
| `v5.1.1` | 2026-06-12   | Patch | Embedded struct tag corrections and integration audit docs |
| `v5.1.0` | 2026-06-12   | Minor | Service-authenticated stock/pricing endpoints and API envelope clarification |
| `v5.0.0` | 2026-06-11   | Major | V5 module path, contract reroute, performance-oriented model cleanup |
| `v4.2.0` | 2026-06-11   | Patch | Version metadata bump |
| `v4.1.0` | 2026-06-11   | Minor | Field grouping, common contact/address/date/party references |
| `v4.0.0` | 2026-06-11   | Major | Generalized payment interfaces and user notification preferences |
| `v3.10.0` | 2026-06-05   | Minor | Customer `is_active` removal after status migration |
| `v3.9.0` | 2026-06-05   | Minor | Customer active flag replaced by status |
| `v3.8.0` | 2026-06-04   | Minor | Customer profile status enum added |
| `v3.7.0` | 2026-06-04   | Minor | Customer record field expansion |
| `v3.6.0` | 2026-06-02   | Minor | Warehouse damage report module |
| `v3.5.1` | 2026-06-01   | Patch | Customer type wording changed from company to wholesaler |
| `v3.5.0` | 2026-06-01   | Minor | ISO27001-aligned audit, security, media, data protection fields |
| `v3.3.0` | 2026-05-18   | Minor | Additional identity roles and shared media/security fields |
| `v3.2.0` | 2026-05-09   | Minor | MX51 payment terminal alignment |
| `v3.1.0` | 2026-05-09   | Minor | Warehouse 3D geometry and layout contracts |
| `v3.0.0` | 2026-05-02   | Major | V3 module path and automated release workflow |
| `v2.1.1` | 2026-04-27   | Patch | Product freshness represented as string |
| `v2.1.0` | 2026-04-27   | Minor | Product freshness/status field expansion |
| `v2.0.5` | 2026-04-27   | Patch | Product expiry field |
| `v2.0.4` | 2026-04-27   | Patch | Product/order JSON naming refinements |
| `v2.0.3` | 2026-04-27   | Patch | Product code retention follow-up |
| `v2.0.2` | 2026-04-27   | Patch | Product code retention |
| `v2.0.1` | 2026-04-26   | Patch | V2 module path correction |
| `v2.0.0` | 2026-04-26   | Major | V2 module path and package import migration |
| `v1.3.0` | 2026-04-25   | Minor | Product and placing area contracts |
| `v1.2.0` | 2026-04-25   | Minor | Purchase, supplier, sales, SKU, payment status model changes |
| `v1.1.0` | 2026-04-24   | Minor | Initial complete contract/model set |
| `v1.0.0` | 2026-04-21   | Major | Initial module baseline |
| `v0.1.0` | 2026-04-21   | Pre-release | Initial repository seed |

## v23.0.0 (2026-08-07) - Contract Domain Layout Hard Cutover

### Breaking Contract Changes / 破壞性契約變更

- Changes the module path from `/v22` to `/v23` and replaces the legacy split
  between `pkg/contracts/<legacy-domain>` and `pkg/enums/<domain>` with merged
  domain packages under `pkg/contracts/<domain>`, without compatibility aliases.
- Restructures the shared model surface under the `pkg/contracts` domain root,
  including the requested identity, pricing, supply, orders, payments,
  insights, customers, notifications, and `pubsub/envelop` / `pubsub/event`
  layout.

### Other Changes / 其他變更

- Merges each domain's cohesive contract and enum files into the same Go
  package, using descriptive filenames where names would otherwise collide.
- Keeps routed Pub/Sub payloads centralized in `pkg/contracts/pubsub/event` and leaves
  identity lifecycle, device, audit, and security records with their owning
  domains.
- Removes the four workflow SHA pins in favor of the stable
  `actions/checkout@v7.0.1` and `actions/setup-go@v7.0.0` tags.

### Consumer Action / 使用方動作

- The seven consumer repositories and the parent `go.work` are intentionally
  untouched. Each downstream repository must migrate its module requirement
  and imports from `/v22` legacy domain paths to `/v23/pkg/contracts/<domain>`
  as a separate change.
- Regenerate or update downstream package references only after choosing the
  matching V23 contract domain package; no compatibility aliases are provided.

## v22.2.0 (2026-08-07) - Repository Cleanup And Warehouse Date-Mark Cutover

### Breaking Contract Changes / 破壞性契約變更

- Removes `warehouseenum.InventoryDateMarkUseBy` and the `USE_BY` wire value.
  Consumers must stop constructing or advertising that value and coordinate
  their route, test, and generated-documentation updates before upgrading.
- This requested release remains on the `/v22` module path even though enum
  value removal is normally classified as a major-version change.

### Other Changes / 其他變更

- Moves repository gates to `pkg/test` and aggregate enum tests to
  `pkg/enums/enums_test`, while keeping package boundary tests beside their
  packages.
- Groups PowerShell scripts under `scripts/powershell` and Bash scripts under
  `scripts/bash`, updating all repository consumers and script-relative paths.
- Adds stable filename rules to `README.md` and commit/push/release rules to
  `GIT_WORKFLOW.md`.

### Contract Files Changed / 契約檔案變更

- `pkg/enums/warehouse/inventory.go`
- `pkg/enums/enums_test/warehouse_test.go`
- `pkg/contracts/warehouse/geo_regional_json_test.go`
- `pkg/versioning/version.go`
- `pkg/versioning/version_test.go`
- `README.md`
- `GIT_WORKFLOW.md`
- `RELEASE_NOTES.md`

### Compatibility Notes / 相容性說明

- The seven backend services currently pin `v22.1.0`; upgrade consumers only
  after replacing `InventoryDateMarkUseBy` / `USE_BY` references.
- Known downstream references remain in Supply source/tests and generated
  Swagger documentation. This repository release does not modify those
  service repositories.

## v22.1.0 (2026-08-06) - Public AU Storefront Commercial And Expiry Projection

### Added / 新增

- Adds `product.StorefrontCommercial` to the public retail product projection.
  It contains the canonical `EACH` package snapshot, optional retail `AUD`
  price, aggregate `stock_state`, `AU` market, and `as_of` freshness timestamp.
  Depot, lot, offer identity, geographic-resolution, and raw stock quantities
  remain outside the public shape.
- Restores the customer-safe `storefront_display.expiry` projection and its
  merchandising policy. It carries `soon_expiry`, the existing broad status,
  warning/critical `alert_level`, remaining days, the 30-day display window,
  localized labels/descriptions, and an optional exact date controlled by
  `show_exact_expiry_date`.
- Adds JSON, enum, compatibility-surface, and version-alignment coverage for
  the new public fields.

### Consumer Action / 使用方動作

- Consumers remain on the `github.com/Potato-Mart/Backend-Shared-Contract/v22`
  module path and must upgrade the required version to `v22.1.0`.
- Supply owns population and filtering of the public AU projection. Orders
  remains the authenticated authority for cart offer resolution.
- Frontends must treat `stock_state=unknown` as unknown, not as out of stock,
  and must never reconstruct price or stock from raw offers or availability.

## v22.0.1 (2026-08-05) - Backend Gate Model Lock

### Breaking Contract Changes / 破壞性契約變更

- None. The `/v22` module path and every exported JSON model and enum value are
  unchanged.

### Added / 新增

- Adds a focused contract test that locks the existing canonical address,
  geographic context, buyer context, `EACH` handling unit, accepted-offer
  revision, product availability, and shared stock/dependency error primitives
  used by the V22 backend gate-closure implementation.

### Fixed / 修正

- Makes the contract release consumed by the backend gate wave explicitly prove
  that the required reusable model surface is present without moving resolver
  DTOs, stock commands, HTTP semantics, authorization, or business rules into
  the shared module.

### Other Changes / 其他變更

- Aligns module metadata and consumer documentation at `v22.0.1`.

### Contract Files Changed / 契約檔案變更

- `pkg/test/backend_gate_lock_test.go`
- `pkg/versioning/version.go`
- `pkg/versioning/version_test.go`
- `README.md`
- `RELEASE_NOTES.md`

### Compatibility Notes / 相容性說明

- Consumers remain on `github.com/Potato-Mart/Backend-Shared-Contract/v22` and
  may update the required version from `v22.0.0` to `v22.0.1` without source or
  data migration.
- Provider and consumer route DTOs, validation, authorization, persistence,
  idempotency, and orchestration remain owned by their service repositories.

## v22.0.0 (2026-08-04) - Geography, Depot, Packaging, and Inventory JSON Cut-over

### Breaking JSON Changes / 破壞性 JSON 變更

- Changes the module path from `/v21` to `/v22` and removes each replaced
  field, type, enum value, event, test fixture, and compatibility alias.
- Replaces free-text address and delivery-region geography with typed country,
  administrative-area, depot-region, depot, shipping-zone, and coverage JSON.
- Replaces product placing-area and depot-product shapes with qualified stock
  location assignments and per-location product availability.
- Replaces singular product barcodes and scalar carton fields with stable
  package options, barcode assignments, and priced CASE/EACH components.
- Replaces scalar product, storefront, POS, and forecast stock/expiry fields
  with lot, bucket, package, offer, and structured availability snapshots.
- Replaces ambiguous stock-location and packing-box fields with
  `location_code`, `storage_type`, product CASE, and outbound CONTAINER models.
- Replaces free-text campaign region and unscoped promotion/coupon models with
  explicit GLOBAL/TARGETED scope and resolved profile geography.
- Replaces competing stock-arrival/restock payloads with revisioned,
  depot-qualified inventory and location-availability events.

The cut-over removes the following v21 surface and uses only the listed v22
replacement JSON:

| Removed v21 surface | v22 replacement JSON |
| --- | --- |
| Address `city`, `state`, `postcode`, and string `country` | `locality`, typed `administrative_area`, `postal_code`, and typed `country` |
| `PostcodeRule`, `PostcodeRules`, and Melbourne `DeliveryRegion` values | typed `country_code`, `administrative_area_codes`, `postal_codes`, `DepotRegion`, and `DepotCoverageRule` |
| Shipping `states`, `postcodes`, and `is_local` | typed `country_code`, `administrative_area_codes`, and `postal_codes` |
| Depot postcode rules and unqualified depot codes | globally qualified `code`, `region_code`, required `address`, and required `timezone` |
| `Product.PlacingArea`, `ProductPlacement`, `DepotProduct`, stock-location `code`, and `zone` | `StockLocationRef`, `location_code`, `storage_type`, location policy, `StockLocationAssignment`, and `StockLocationProductBalance` |
| Product `sku`, singular `barcode`, `physical`, and ambiguous `storage` | `category_sku_code`, `package_options`, `barcode_assignments`, and `storage_type`; `sku_code` remains the product identity |
| Product/storefront/POS `current_stock`, `restocked_at`, stock-derived `display_status`, scalar price, and product-wide expiry | `ProductStockSummary`, `SellableOffer`, and immutable `SellableOfferSnapshot` |
| Purchase/WMS scalar expiry and location fields, including `expired_at`, `expire_at`, `expiry_ym`, and singular `location_code` | `InventoryLot`, `InventoryDateMark`, `InventoryStockBucket`, bucket/location references, and package compositions |
| Legacy stock-adjust, reserve/release movement values, and competing arrival/restock events | physical `StockMovement`, logical `StockReservation`, `StockReservationAllocation`, `StockStagingRecord`, package conversion, quality, availability, and offer events |
| Cart/order `quantity`, `unit_price`, `carton_qty`, and `carton_size` | `components`, `requested_package_count`, `requested_base_units`, accepted offer/package snapshots, and package compositions |
| Campaign planning `*_units`, `minimum_order_quantity`, and supplier `total_units` | explicit `int64` `*_base_units` fields and `suggested_composition` |
| `PackingBoxPlan`, `PackingBoxContent`, scalar picking/packing quantities, and duplicated damage balances | `OutboundContainerPlan`, `OutboundContainerContent`, package-aware picking/packing compositions, substitution snapshots, and quality-assessment references |
| Participant-owned group inventory fields | `GroupOrderContext` and parent-owned `GroupOrderFulfilmentPlan` aggregate lines with participant shares |
| Campaign audience `region` and unscoped campaign/promotion/coupon projections | required `geographic_scope`, resolved `geographic_context`, promotion `series_key`, revisions, and schedule timezone |

Operational instants use standard `time.Time`/`*time.Time` JSON values. Values
are normalized to UTC before encoding; local-calendar meaning carries an
explicit IANA timezone, and global schedules use `Etc/UTC`.

### Added / 新增

- Typed country/subdivision references, administrative-area kinds,
  `DepotRegion`, `DepotCoverageRule`, and required depot IANA timezone.
- Location purpose, handling, customer access, collection eligibility,
  system-location identity, SKU assignment, and electronic-shelf-label fields.
- Product package options, package-qualified barcode assignments, package
  composition snapshots, inventory lots/date marks/buckets/units, physical
  condition, stock disposition, quality assessment, and package conversion.
- Logical reservation/allocation, physical staging, canonical movement,
  sellable offer, product stock summary, and zero-crossing availability models.
- Mixed CASE/EACH cart and order components, explicit loose-case substitution,
  consolidated group fulfilment, participant shares, and outbound containers.
- Reusable geographic scope/context on campaigns, marketing campaigns,
  promotions, coupons, pricing snapshots, offers, and order records.

### Contract Boundary / 契約邊界

- Ordinary `json` tags and standard `encoding/json` are the only wire-shape
  mechanism. Production non-JSON tags and custom codecs are rejected.
- Production code remains structs and enums only. Receiver methods are limited
  to `String()` and single-value `IsValid()` enum helpers.
- Deprecated declarations, type aliases, fallback fields, and dual JSON shapes
  are rejected by the v22 hard-cut tests.

### Consumer Action / 使用方動作

- Adopt `github.com/Potato-Mart/Backend-Shared-Contract/v22 v22.0.0` and replace
  every `/v21` import in one consumer release.
- Update JSON mappings for typed geography, qualified depot locations,
  package-aware quantities, lots/buckets/offers, group fulfilment, and resolved
  campaign/promotion geography before publishing or consuming v22 events.

## v21.1.0 (2026-07-30) - Backend Gap Models and Gift-Card Policy V2

### Breaking Contract Changes / 破壞性契約變更

- None. This is an additive minor release and retains the `/v21` module path.

### Added / 新增

- Notification status value `read`; read notifications remain distinct from
  dismissed notifications and continue to use the existing optional `read_at`.
- Retail marketing `push_opt_in` plus optional email, SMS, Line, and push
  consent timestamps and source provenance. The customer-consent event also
  carries the push decision.
- Optional `refund_id` and `refunded_at` on coupon usage records.
- Optional authoritative `occurred_at` timestamps on point, reward, voucher,
  and gift-card order redemption snapshots.
- Optional issued, activated, and single-use redeemed timestamps on wallet
  instruments. Re-spendable gift cards do not use a singular redemption time.
- Point-award status values `ineligible`, `disabled`, `pending`, `awarded`, and
  `failed`; customer payment summaries now distinguish gross points earned,
  debt repaid, net points credited, and remaining debt.
- Point-debt totals on membership and customer wallet projections, optional
  debt transitions on ledger entries, and `DEBT_INCURRED` / `DEBT_REPAID`
  ledger reasons.
- `voucher.claim_issued` and its reusable event payload containing only an
  issuance ID and a Customers-bound delivery handle. Claim tokens, recipient
  addresses, and other claim material are deliberately absent.
- Reusable `wallet.GiftCardDenominationPolicy` with a policy version, currency,
  and ordered minor-unit amounts. Gift-card issuance events can carry an
  optional denomination-policy version for V1/V2 coexistence.

### Fixed / 修正

- Makes read and dismissed notification lifecycle states independently
  representable.
- Makes refund reversal, point debt, and redemption commit timing explicit
  without fabricating values for historical records.

### Other Changes / 其他變更

- Adds JSON round-trip and omission tests for all new fields and events,
  including known-zero debt, claim-material exclusion, and the five V2 policy
  values `50000`, `80000`, `100000`, `150000`, and `200000` AUD minor units.
- Updates the reviewed exported-type manifest for the three new model/enum
  types while preserving the model-only boundary.

### Contract Files Changed / 契約檔案變更

- `pkg/contracts/{customers,membership,notification,promotion,sales,wallet}`
- `pkg/enums/{events,membership,notification}`
- Module/version metadata, release notes, the exported-type manifest, and JSON
  shape tests.

### Compatibility Notes / 相容性

- Existing V21 consumers remain source-compatible. New scalar fields decode to
  zero values and new optional timestamps/debt transitions may be absent on
  historical records and V1 events.
- This module does not define gift-card routes, cache behavior, validation,
  authorization, cutover timestamps, the rolling purchase cap, or persistence.
  Pricing remains the runtime policy owner; service-local APIs must return this
  model and enforce business rules.
- Existing issued gift cards and delayed V1 events remain valid. The V2 allowed
  amounts apply only when owning services authorize a new purchase-backed
  issuance.

### Consumer Action / 使用方動作

- Pin `github.com/Potato-Mart/Backend-Shared-Contract/v21 v21.1.0` exactly and
  regenerate or update service-local mappings after the release is published.
- Customers should support both V1 and V2 issuance events before Pricing and
  Orders start publishing V2. Pricing should return policy version `2`, currency
  `AUD`, and exactly `50000`, `80000`, `100000`, `150000`, `200000`; Orders must
  validate new quote/start flows against that policy.
- Preserve missing historical timestamps rather than substituting unrelated
  order times, and never put voucher claim material into Pub/Sub, outboxes,
  logs, traces, or DLQs.

## v21.0.0 (2026-07-29) - Delivery, Campaign, Notification, and Wallet Policy Cutover

### Breaking Contract Changes / 破壞性契約變更

- The module path changes from `/v20` to `/v21`; consumers must update every
  contract import together and must not mix types from the two major lines.
- `identity.UserNotificationPreferences.quiet_hours` and the exported
  `identity.NotificationQuietHours` type are removed without a compatibility
  alias. Services and clients must remove quiet-hours validation, storage,
  scheduling, and UI before adopting V21.
- V20's catalog hard cut remains intact: `product.BrandRef.ID`,
  `analytics.OrderItemFact.BrandID`, and `analytics.RefundItemFact.BrandID`
  remain the only canonical shared brand identity fields; `brand_key` does not
  return.

### Added / 新增

- Shipping models `DeliverySlot`, `DeliveryDateGroup`, `DeliverySchedule`,
  `DeliveryAreaRate`, and `PreferredDeliverySlot`. Area rates expose
  `postcode`, optional `suburb`, `delivery_region`, depot code/name,
  `shipping_fee`, and `free_shipping_threshold`; monetary values use
  `common.Money` integer minor units and services emit AUD.
- Retail profile fields `billing_same_as_delivery` and
  `preferred_delivery_slot`. A preferred slot is display-only metadata and
  remains subject to authoritative checkout revalidation.
- Campaign-promotion linkage, stable managed-media identity, typed CTA
  destinations, content revision, activation revision, and deactivation
  timestamp on `campaign.Campaign`; the existing `cta_href` remains available
  to current storefront clients.
- Customer-safe `promotion.PromotionChangedEvent` and
  `campaign.CampaignChangedEvent`, the `storefront-events` topic, and
  `promotion.changed` / `campaign.changed` event-type values. These payloads
  carry identifiers, lifecycle/revision state, and timestamps only so
  storefront consumers can refetch authoritative content.
- Typed `notification.CampaignReference` on durable customer notifications,
  the `push` notification channel, and the `order_placed`,
  `promotion_available`, and `announcement` notification topics.
- Stable API error codes `CART_NOT_ACTIVE` and
  `EMAIL_VERIFICATION_REQUIRED` for deterministic native recovery flows.
- Reusable `membership.PointsPolicy`, surfaced optionally from the customer
  wallet summary. Services publish `points_per_minor_unit=2`,
  `minimum_eligible_balance=1000`, `redemption_step_points=200`, and the
  server-calculated `maximum_redemption_points`; clients do not reproduce the
  calculation.

### Fixed / 修正

- Makes campaign activation and content changes independently revisioned so
  one push can be deduplicated per activation revision while content updates
  remain safe refetch signals.
- Gives retail clients a shared cart-independent delivery schedule/preference
  shape while preserving checkout as the authority for availability and fees.

### Model-Only Boundary / 僅模型邊界

- FCM/FID destinations, encryption and destination hashes remain Customers
  service-local; V21 intentionally exports no destination/token model.
- Coupon previews, request/response DTOs, routes, query parameters, validation,
  authorization, state transitions, activation/push logic, and points
  calculations remain in their owning services.
- The model-only AST boundary, reviewed exported-type manifest, and hard-cut
  searches guard these exclusions and reject reintroduction of quiet-hours,
  FCM destination, or coupon-preview fields.

### Contract Files Changed / 契約檔案變更

- `pkg/contracts/shipping/delivery_slots.go`
- `pkg/contracts/customers/retail_customer.go`
- `pkg/contracts/campaign/campaign.go` and `events.go`
- `pkg/contracts/promotion/events.go`
- `pkg/contracts/notification/customer_notification.go`
- `pkg/contracts/identity/notification_preferences.go`
- `pkg/contracts/membership/points.go` and `pkg/contracts/wallet/wallet.go`
- `pkg/enums/{apiresponse,campaign,events,notification}`
- Module/version metadata, release notes, boundary/manifest gates, and JSON
  shape tests across the V21 module path.

### Compatibility Notes / 相容性

- This is a hard major-version cutover. Existing `/v20` consumers continue to
  compile against `v20.0.0`, but no `/v20` value can be passed as a `/v21`
  value and there is no dual-import compatibility layer.
- Existing stored quiet-hours values must be unset by the owning service;
  changing the shared model does not mutate persistence.
- Delivery availability strings, route DTOs, and service persistence are not
  defined by this module. Consumers should use the owning service's OpenAPI
  contract and treat checkout responses as authoritative.

### Consumer Action / 使用方動作

- Pin `github.com/Potato-Mart/Backend-Shared-Contract/v21 v21.0.0` and update
  every import to `/v21` after this tag is published.
- Remove quiet-hours reads/writes/UI, migrate delivery/campaign/notification
  projections to the new shared shapes, and refetch on safe storefront events.
- Keep FCM destination and coupon-preview DTOs inside Customers and Orders,
  respectively. Do not add them to the shared module.

## v20.0.0 (2026-07-29) - Catalog Brand Identity and SKU Relationship Hard Cut

### Breaking Contract Changes / 破壞性契約變更

- `product.Brand` and embedded `product.BrandRef` now use exactly `id`,
  immutable `slug`, localized `name`, and optional `logo_url`.
- Removed `product.BrandSummary`, `brand_key`, brand audience, featured state,
  active-product counters, and public brand audit fields. Storefront product
  projections expose only `brand_ref`.
- `analytics.OrderItemFact` and `analytics.RefundItemFact` now carry
  `brand_id` instead of `brand_key`.
- `category.SKU` no longer embeds a derived `products` list. Canonical product
  records own SKU-family membership through their `sku` field.
- The module path changes from `/v19` to `/v20`; there is no compatibility
  alias, fallback decoder, or dual-read period.

### Added / 新增

- Optional HTTPS-capable `logo_url` on canonical brand masters and embedded
  brand references.
- Stable canonical brand identity in `BrandRef.id`, matching the master Mongo
  ObjectId represented as a 24-character hexadecimal string.

### Fixed / 修正

- Removes stale SKU-product snapshot modelling from the shared contract so
  SKU membership has one authoritative source.

### Other Changes / 其他變更

- Supply-owned brand activity, validation, API routes, persistence indexes,
  and product-browse endpoints remain implementation concerns in
  Backend-Supply rather than shared DTO fields.

### Contract Files Changed / 契約檔案變更

- `pkg/contracts/product/brand.go`
- `pkg/contracts/product/storefront.go`
- `pkg/contracts/category/sku.go`
- `pkg/contracts/analytics/facts.go`
- Module/version metadata, model-manifest, and JSON-shape tests.

### Compatibility Notes / 相容性

- This is a hard major-version cutover. Existing `/v19` consumers continue to
  work against `v19.0.0`; they must not mix `/v19` types with `/v20` types.
- Existing data must be imported in the v20 clean shape. This release does not
  provide a live-data migration or compatibility read path.

### Consumer Action / 使用方動作

- Pin `github.com/Potato-Mart/Backend-Shared-Contract/v20 v20.0.0` and update
  imports to `/v20`.
- Update catalog consumers to link brands by `brand_ref.id`, use slugs for
  public navigation, and stop reading/writing `brand_key` and brand summaries.
- Fetch related products through Backend-Supply relationship endpoints instead
  of a `SKU.products` response field.

## v19.0.0 (2026-07-27) - Admin Portal Hard Cutover

This breaking release establishes the final model surface for the fully wired
Admin Portal. Membership is retail-only and uses the customer number as its
aggregate key. Wholesale commercial benefits use a separate owner reference,
wholesale applications are durable records, freight presets have no tier
semantics, product placements are qualified by depot and location, media
records carry bucket/visibility/reference metadata, and analytical facts are
typed for Admin reporting.

All deprecated models, enum values, fallback fields, and every serialized
`sort_order` field are removed. Consumers must update imports to `/v19`, use
domain-derived ordering, migrate affected data, and regenerate service-owned
OpenAPI and client types before deployment. There is no dual-read or alias
period.

### Consumer Action / 使用方動作

- Pin `github.com/Potato-Mart/Backend-Shared-Contract/v19 v19.0.0`.
- Re-key retail membership references to customer number and remove wholesale
  membership/tier fields.
- Migrate product brand, placement, media, freight, and ordering fields to the
  v19 models.
- Regenerate OpenAPI and every generated/manual client model.

## v18.6.1 (2026-07-22) - Buyer Identity on Payment Failure and Refund Completion Events

This patch adds optional `retail_customer_number` and `organisation_access_id`
fields to `payments.PaymentFailedEvent` and `payments.RefundCompletedEvent`,
matching the buyer-identity enrichment the sales events already carry, so
notification consumers can resolve the recipient from the payload alone.

本修補版本為 `payments.PaymentFailedEvent` 與 `payments.RefundCompletedEvent` 增加
選填的 `retail_customer_number` 與 `organisation_access_id` 欄位，使通知消費端可
僅憑 payload 解析收件人。

### Compatibility / 相容性

- Additive optional fields only; no exported-type, JSON-shape, enum, or
  module-path change. Consumers tolerating unknown fields need no action.
- 純追加選填欄位；使用方無需調整。

### Consumer Action / 使用方動作

- Publishers should populate the new fields where buyer identity is known;
  consumers should prefer them over local lookups (keeping fallbacks for old
  envelopes).
- 發布端在已知買家身分時填入新欄位；消費端優先使用事件欄位，並保留舊信封的
  後備解析。

## v18.6.0 (2026-07-22) - Complete Event Taxonomy, Customer Payment Models, Membership Progress, POS Surface

This additive release completes the event payload surface for all nine Pub/Sub
topic families and delivers the storefront-experience and POS model
foundations: customer-safe payment summaries, invoice redelivery, membership
tier progress and typed benefits, product origin/weight display fields, and a
reuse-first point-of-sale surface.

本次追加版本補齊九個 Pub/Sub topic 家族的事件 payload 模型，並提供 storefront 體驗
與 POS 模型基礎：客戶安全的付款摘要、發票重寄、會員等級進度與型別化權益、商品
產地/重量顯示欄位，以及以重用為先的 POS 模型。

### Other Changes / 其他變更

- Events: `warehouse.StockAdjustedEvent`; `sales` fulfilment events
  (`FulfilmentShippedEvent`/`Delivered`/`Completed`/`Tracking`) +
  `PreorderAvailabilityEvent`; `customers` registered/profile/consent events;
  `product` catalog-changed events + `ProductSalesRollup`;
  `notification.NotificationEngagementEvent` + `EngagementAction`;
  `payments.InvoiceIssuedEvent`; enriched `sales` order events (buyer/tracking/
  invoice refs) and `payments.RefundCompletedEvent` (benefit/points restoration
  fields). `stock.arrived` reuses `sales.PreorderStockArrivalEvent`;
  `fulfilment.packing_updated` reuses `sales.OrderPackingProjection`.
  19 new `EventType` constants.
- Customer payments: `sales.CustomerPaymentSummary`/`CustomerPaymentAllocation`
  (also the invoice payment-row shape), `sales.InvoiceEmailDelivery`, and
  payment enums (allocation kind, completeness, summary component,
  invoice-resend status).
- Membership: `CustomerTierProgress`/`TierProgressTier`/`QualificationWindow`
  + `TierProgressReason`; typed `TierBenefit`/`TierBenefitValue` +
  `TierBenefitKind`; `MembershipTier.Benefits` (untyped `Perks` deprecated).
- Product: `StorefrontOrigin` + `country_of_origin`/`physical_weight`
  (reusing `common.Weight`) on the storefront projection and product master.
- POS (reuse-first): `contracts/pos` (`Register`, `RegisterShift`,
  `CashMovement`, `ShiftTotalsSnapshot`, `MethodTotal`, `ReceiptSnapshot`
  reusing order items/payment rows/receipt offers), `enums/pos`
  (`ShiftStatus`, `CashMovementKind`), `sales.POSAttribution` on
  `SourceDevice`, `identityenum.UserRoleCashier`,
  `paymentenum.TerminalProviderStripe`.
- 事件、客戶付款、會員、商品與 POS 模型如上；模型清單分類、digest 與版本中繼資料
  更新至 `v18.6.0`。

### Compatibility / 相容性

- Additive only: no existing exported model, field, JSON/BSON shape, enum wire
  value, package path, or Go module path changes.
- 純追加變更；現有 v18 使用方不需遷移程式碼或資料。

### Consumer Action / 使用方動作

- Services wiring the remaining topics pin `v18.6.0`; POS-facing services adopt
  the `contracts/pos` + cashier/Stripe enums as their façades land.
- 佈線其餘 topic 的服務請固定使用 `v18.6.0`；POS 相關服務於 façade 上線時採用
  `contracts/pos` 與收銀員/Stripe enum。

## v18.5.0 (2026-07-21) - Event Envelope and Domain Event Models

This additive release supplies the shared model truth for the seven-service
migration's GCP Pub/Sub eventing backbone: a transport-neutral event envelope,
typed topic and event-type enums, and the first wired payload models for order,
payment, and refund lifecycle facts. Payload models for the remaining topics
(stock, fulfilment, customer, catalog, engagement, product-stats) will follow
in later additive minors as those flows are wired.

本次追加版本為七服務遷移的 GCP Pub/Sub 事件骨幹提供共用模型：與傳輸層無關的事件
信封、型別化的 topic 與事件類型 enum，以及訂單、付款、退款生命週期事實的首批
payload 模型。其餘 topic 的 payload 模型將於後續追加 minor 版本補齊。

### Other Changes / 其他變更

- Added `pkg/contracts/events` with `EventEnvelope` (`event_id`, `event_type`,
  `event_version`, `occurred_at`, `aggregate_id`, raw JSON `payload`); the
  envelope carries delivery metadata only — typed payloads stay with their
  owning domain packages.
- Added `pkg/enums/events` with `EventTopic` (the nine aggregate-family topics)
  and `EventType` (the wired order/payment/refund event names).
- Added `contracts/sales` order lifecycle events: `OrderCreatedEvent`,
  `OrderPaidEvent`, `OrderStatusChangedEvent`, `OrderCancelledEvent`.
- Added `contracts/payments` money-path events: `PaymentCapturedEvent`,
  `PaymentFailedEvent`, `RefundRequestedEvent`, `RefundCompletedEvent`,
  `RefundFailedEvent`.
- Updated the model manifest classifications/digest and
  `versioning.ModuleVersion` to `v18.5.0`.
- 新增事件信封、topic/事件類型 enum，以及訂單/付款/退款事件模型；更新模型清單
  分類、digest 與版本中繼資料至 `v18.5.0`。

### Compatibility / 相容性

- Additive only: no existing exported model, field, JSON/BSON shape, enum wire
  value, package path, or Go module path changes.
- Existing v18 consumers require no code or data migration.
- 純追加變更；現有 v18 使用方不需遷移程式碼或資料。

### Consumer Action / 使用方動作

- Services adopting the Pub/Sub backbone should pin `v18.5.0` and build their
  publisher/receiver runtime around `events.EventEnvelope`, deduping on
  `event_id` (push delivery is at-least-once).
- 採用 Pub/Sub 骨幹的服務應固定使用 `v18.5.0`，以 `events.EventEnvelope` 建置
  發布/接收執行期，並以 `event_id` 去重（push 傳遞為 at-least-once）。

## v18.4.1 (2026-07-21) - Documentation and Release Guidance

This patch consolidates release history and migration guidance in the canonical
`RELEASE_NOTES.md`, keeps the README focused on module responsibility and
consumer usage, and documents the protected-main contribution and release
workflow. The `/v18` module path and all public contract behavior remain
unchanged.

本修補版本將發布歷史與遷移指引集中於標準 `RELEASE_NOTES.md`，並讓 README 專注於
模組責任與使用方式，同時記錄受保護 `main` 分支的貢獻與發布流程。`/v18` 模組路徑及
所有公開契約行為均維持不變。

### Other Changes / 其他變更

- Reorganized the README into responsibility, latest version, usage, boundary
  governance, verification, and change policy sections.
- Removed duplicated v18 release summaries from the README; the existing
  detailed entries in this file remain the release-history source of truth.
- Documented that contributors must use a feature branch and pull
  request, align the next semantic version across release metadata, and allow
  the post-merge workflow to publish the tag and GitHub release.
- Updated `versioning.ModuleVersion` and its metadata test to `v18.4.1`.
- 重新整理 README，移除重複的 v18 發布摘要，並記錄功能分支、pull request、版本
  對齊及合併後自動發布的流程。

### Compatibility / 相容性

- No exported model, field, JSON/BSON shape, enum wire value, package path, or
  Go module path changes.
- Existing v18 consumers require no code or data migration.
- 公開模型、欄位、JSON/BSON shape、enum wire value、package path 與 Go module
  path 均無變更；現有 v18 使用方不需遷移程式碼或資料。

### Consumer Action / 使用方動作

- Consumers may pin `v18.4.1` after publication to use the aligned metadata;
  no runtime behavior changes are required.
- 使用方可在發布後固定使用 `v18.4.1` 以取得一致的版本中繼資料，無需變更執行期行為。

## v18.4.0 (2026-07-20) - Product Story, Reviews, and Make a Wish

This additive release supplies shared model truth for customer-safe product
provenance and detail imagery, rating and review projections, and Make a Wish
proposals and ballots. It retains the `/v18` module path, preserves the legacy
`supplier_code` field, and does not define HTTP routes, request DTOs, response
envelopes, pagination, authorization, validation, ETags, or OpenAPI. Those
remain owned by the implementing backend. All removals are deferred to v19.

此加法版本提供顧客安全的商品供應來源、製造資訊與詳細圖片，以及評分、評論和
「許願商品」提案與投票的共用模型真相。模組路徑維持 `/v18`，並保留既有
`supplier_code` 欄位。本模組不定義 HTTP 路由、request DTO、response envelope、
分頁、授權、驗證、ETag 或 OpenAPI；這些仍由實作後端擁有。所有移除項目延後至
v19。

### Added / 新增

- `common.LocalizedText` with required `language` and `text` JSON fields for
  localized customer-facing copy that is not a name or description.
- `product.ProductSupplierRef` with customer-safe `code` and `name`;
  `product.ProductManufacturing` with optional `company_name` and `location`;
  and `product.ProductSupply` with independently optional `supplier` and
  `manufacturing` sections. Supplier contacts, addresses, legal identifiers,
  and operational fields are deliberately absent.
- `product.DetailImage` with required `url` and optional localized `alt_text`
  and `caption`. Array position is the canonical display order.
- Optional `supply` and nullable/omittable `display_selling_count` on
  `product.Product`; optional `supply` and `display_selling_count`, plus ordered
  top-level `detail_images`, on `product.StorefrontProduct`; ordered
  `detail_images` on `product.Media`; and optional `supply` on
  `product.Snapshot` while retaining `supplier_code`.
- Optional `audience` on `product.BrandSummary`, using
  `productenum.PriceAudience`, so retail and wholesale
  `active_product_count` values cannot be confused.
- `review.RatingDistributionBucket` with `score` and `count`, and
  `review.RatingSummary` with `average_score`, `rating_count`,
  `published_text_review_count`, and an exactly five-bucket ordered
  `distribution` for scores 1 through 5.
- Public `review.ProductReview`, customer-owned `review.MyProductReview`, and
  PII-free moderation `review.ProductReviewModeration` projections. Their JSON
  fields cover review ID, `product_sku_code`, score, approved and/or original
  title/body as appropriate, locale, verified-purchase truth, moderation
  status/reason or note as appropriate, and explicit safe timestamps. None
  embeds audit actors or customer identity.
- `reviewenum.ReviewModerationStatus` values `not_required`, `pending`,
  `approved`, `rejected`, and `suppressed`; `ReviewRejectionReason` values
  `spam`, `off_topic`, `inappropriate`, `personal_information`,
  `unsupported_language`, and `other`; and stable `ReviewErrorCode` values
  `REVIEW_NOT_FOUND` and `REVIEW_PURCHASE_REQUIRED`.
- `wish.WishProposal` with `id`, `product_name`, optional `description` and
  `reference_url`, `state`, optional conversion/product references, and safe
  timestamps; `wish.WishCandidate` with localized `name` and `description`,
  ordered `image_urls`, state, publication/fulfilment facts, and safe
  timestamps; revisioned `wish.WishBallot` with ordered `candidate_ids`,
  opening/closing times, and `as_of`; `wish.WishRankingEntry` with
  `candidate_id`, `rank`, and `vote_count`; and identity-free
  `wish.WishSelection` with `ballot_id`, ordered `candidate_ids`, and
  `updated_at`.
- `wishenum.WishProposalState` values `pending`, `converted`, and `rejected`;
  `WishCandidateState` values `draft`, `published`, `retired`, and `fulfilled`;
  `WishBallotState` values `scheduled`, `open`, and `closed`; and stable
  `WishErrorCode` values `WISH_NO_ACTIVE_BALLOT`, `WISH_BALLOT_CLOSED`, and
  `WISH_CANDIDATE_UNAVAILABLE`.
- 新增 `common.LocalizedText`、顧客安全的供應商／製造／供應資訊、依陣列順序顯示
  的本地化詳細圖片，以及可接受缺少／`null` 並保留明確零值的手動銷售顯示數量。
- 新增固定五個有序分數桶的評分摘要、公開／顧客本人／無 PII 的審核評論投影，
  以及評論狀態、拒絕原因與穩定錯誤碼。
- 新增許願提案、候選商品、修訂式投票、排序結果與無身份選擇模型，以及提案、
  候選、投票狀態與穩定錯誤碼；所有公開時間欄位直接宣告，且不嵌入審計人員或
  顧客身份。

### Compatibility / 相容性

- Every new product/storefront/snapshot field is additive and optional.
  Pointer-backed `display_selling_count` preserves an explicit zero while an
  absent value remains omitted. Existing v18.3 payloads decode unchanged.
- Legacy `supplier_code`, localized `brand`, `brand_ref`, `image_urls`, and all
  other existing v18 JSON keys remain present. Consumers may migrate from
  `supplier_code` to `supply` during v18; no legacy removal is permitted before
  v19.
- Product supply may contain supplier-only, manufacturing-only, both sections,
  or neither. Backends must expose only the customer-safe fields declared here.
- Rating distributions are always five ordered buckets even when counts are
  zero. Rating summaries are retail-only and must not be added to wholesale
  responses.
- Review and wish enums and error codes are stable wire values. Backend-owned
  APIs may add transport-specific status and envelope details without changing
  these shared domain values.
- 所有新商品、店面與快照欄位均為選用加法欄位；既有 v18.3 payload 可原樣解析。
  `display_selling_count` 可保留明確的零值；未提供時則省略。
- `supplier_code`、既有品牌與圖片欄位全部保留。v18 期間可逐步改用 `supply`，
  v19 前不得移除相容欄位。評分摘要僅供零售，不得出現在批發回應。

### Consumer Action / 使用方動作

- Operations must pin `v18.4.0`, adopt the released product/review/wish structs
  at domain, persistence, and projection boundaries, migrate/hydrate supply
  data, implement brands/reviews/wishes and their indexes/ETags/errors, publish
  capability flags, and freeze its provider OpenAPI before Commerce regenerates
  the catalogue consumer. Backend-local write DTOs must preserve omitted/set/null
  validation semantics, and wishes must not reuse favourite-list persistence.
- Commerce must pin `v18.4.0`, regenerate its Operations catalogue wire from the
  frozen provider OpenAPI, explicitly preserve `Snapshot.Supply` through cart,
  checkout, and durable order items, and add the least-privilege authenticated
  review purchase-eligibility endpoint. It must not own review/wish moderation
  or persistence.
- Management must pin `v18.4.0`, add review read/moderate and wish read/manage
  permissions, and issue only the frozen Operations-to-Commerce eligibility
  service grant. Identity, RBAC, tokens, service scopes, and HTTP DTOs remain
  Management-owned; product/review/wish persistence does not.
- Retail web must consume the backend-owned OpenAPI for brands, product story,
  rating/review, wish, conditional caching, errors, and runtime capabilities,
  while excluding identity, moderator notes, supplier contacts, and operational
  data. Wholesale web adopts only audience-correct brands and shared safe
  product fields; it must not add ratings, reviews, proposals, ballots, or
  voting.
- Operations、Commerce 與 Management 必須在各自功能 branch 固定使用已發布的
  `v18.4.0`，並依上述邊界採用結構、OpenAPI、快照相容性、權限與服務授權。
  零售網頁採用品牌、商品故事、評論與許願 API；批發網頁只採用受眾正確的品牌
  與顧客安全商品欄位，不得加入評分、評論、提案或投票。

## v18.3.0 (2026-07-19) - Storefront Brand Catalogue

This additive release defines the customer-safe brand catalogue projection and
introduces an optional stable `brand_key` for brand masters, lightweight brand
references, and storefront products. The module path remains `/v18`; all
existing brand fields and payloads remain compatible.

此加法版本定義顧客安全的品牌目錄投影，並在品牌主檔、輕量品牌參照及店面商品上新增
選用且穩定的 `brand_key`。模組路徑維持 `/v18`；所有既有品牌欄位與 payload
保持相容。

### Added / 新增

- `product.BrandSummary` with required `brand_key`, localized `names`,
  `featured`, `sort_order`, and `active_product_count`, plus optional
  `logo_url`.
- Optional `brand_key` on `product.Brand`, `product.BrandRef`, and
  `product.StorefrontProduct`.
- `brand_key` is the lowercase URL-safe canonical filter/navigation key and is
  immutable after assignment by the owning catalogue.
- 新增 `product.BrandSummary`，包含必要的 `brand_key`、本地化 `names`、
  `featured`、`sort_order`、`active_product_count`，以及選用的 `logo_url`。
- `product.Brand`、`product.BrandRef` 與 `product.StorefrontProduct` 新增選用的
  `brand_key`；此欄位為小寫、URL 安全且一經指定即不可變更的正規鍵值。

### Compatibility / 相容性

- Empty `brand_key` values are omitted from compatibility-bearing models.
- Existing payloads without `brand_key` decode unchanged, and existing `slug`,
  localized `brand`, and `brand_ref` fields remain present.
- 不含 `brand_key` 的既有 payload 可繼續解析；空值不會輸出，既有 `slug`、
  本地化 `brand` 及 `brand_ref` 欄位全部保留。

### Consumer Action / 使用方動作

- Operations should pin `v18.3.0`, populate `brand_key`, and expose
  `BrandSummary` only after its brand backfill is validated. Storefront clients
  may continue using existing brand fields during rollout.
- Operations 應固定使用 `v18.3.0`，完成 `brand_key` 回填驗證後才公開
  `BrandSummary`；店面客戶端在推出期間仍可使用既有品牌欄位。

## v18.2.0 (2026-07-19) - Canonical Product Brand Models

This additive release defines reusable canonical product brand masters and
lightweight embedded brand references. The module path remains `/v18`; existing
localized `brand` arrays remain unchanged and `brand_ref` is optional.

此加法版本定義可重用的正規商品品牌主檔與輕量嵌入式品牌參照。模組路徑維持
`/v18`；既有本地化 `brand` 陣列保持不變，`brand_ref` 為選用欄位。

### Added / 新增

- `product.Brand` with stable ID and slug, localized canonical names,
  localized aliases, and shared audit fields.
- Lightweight `product.BrandRef` with stable ID, slug, and localized names.
- Optional `brand_ref` on `product.Product`, `product.Snapshot`, and
  `product.StorefrontProduct`.
- `product.Brand` 新增穩定 ID、slug、本地化正規名稱、本地化別名與共用審計欄位。
- `product.BrandRef` 提供穩定 ID、slug 與本地化名稱的輕量參照。

### Compatibility / 相容性

- Existing payloads without `brand_ref` decode unchanged.
- Existing localized `brand` arrays remain present for compatible consumers;
  this release adds no renamed or removed JSON keys.
- 不含 `brand_ref` 的既有 payload 可繼續解析；本地化 `brand` 陣列仍保留，
  本版本沒有改名或移除任何 JSON key。

### Consumer Action / 使用方動作

- Consumers that need canonical brand identity should pin `v18.2.0` and prefer
  `brand_ref`; consumers may continue using the legacy localized `brand` array
  during migration.
- 需要正規品牌識別的使用方應固定 `v18.2.0` 並優先使用 `brand_ref`；
  遷移期間仍可繼續使用既有本地化 `brand` 陣列。

## v18.1.0 (2026-07-18) - Customer Order Lifecycle Notification Topics

This additive release defines the Management-owned customer notification topics
needed for retail and wholesale order lifecycle messaging. The module path
remains `/v18` and no existing JSON field or enum value changes.

此加法版本定義由 Management 擁有的顧客通知主題，支援零售與批發訂單生命週期訊息。
模組路徑維持 `/v18`，既有 JSON 欄位及 enum 值皆不變。

### Added / 新增

- Order confirmation and cancellation topics.
- Payment received, failed, and refunded topics.
- Packing started, order packed, dispatched, delivered, and invoice available topics.
- 訂單確認與取消、付款成功／失敗／退款、開始包裝、完成包裝、出貨、送達及發票可用通知主題。

### Consumer Action / 使用方動作

- Management, Commerce, and Operations must pin `v18.1.0` before publishing
  lifecycle notification behavior. Frontends must consume the corresponding
  backend OpenAPI topic values.
- Management、Commerce 與 Operations 在發布生命週期通知功能前必須固定使用
  `v18.1.0`；前端須採用對應後端 OpenAPI 的通知主題值。

## v18.0.0 (2026-07-18) - Customer Commerce And Storefront Models

This breaking release moves the module path to `/v18` and establishes shared
model truth for customer lists, notification dismissal, customer-safe catalogue
details, promotion ribbons, and computed product sales rankings.

### Breaking Contract Changes

- `product.Product.sales_performance` is now structured, computed history;
  `avg_weekly_sales` and the `hot|normal|slow` enum are removed.
- Legacy Commerce wholesale favourite/order-list permissions are replaced by
  organisation-scoped favourite-list view/write permissions.
- All consumers must update imports to `/v18` and pin `v18.0.0`.

### Added

- Retail-user and wholesale-organisation favourite-list entities, ownership,
  stable limit/name error codes, and product-membership records.
- Customer notification `unread|dismissed` status, read/dismiss timestamps,
  and explicit expiry.
- 7/30/90-day and lifetime order/unit statistics, refund-aware net units,
  timestamps, timezone, and per-category rank/population.
- A customer-safe storefront product model with audience-filtered pricing,
  exact policy-approved catalogue details, promotion badges, and rankings.

### Consumer Action / 使用方動作

- Management, Operations, and Commerce must adopt `/v18` before publishing
  dependent API changes. Storefront clients must consume backend projections,
  not the raw product master or merchandising policy.

## v17.4.0 (2026-07-17) - Session, Membership Pass, And Wholesale Price Models

This additive minor release provides reusable model truth for immediate
session revocation, exact device sign-in reporting, external membership-pass
content, and wholesale price-on-request merchandising. The module path remains
`/v17`.

### Added

- `identity.AccessTokenClaims.session_id` binds an access token to its durable
  login-session record.
- `identity.UserDevice.last_login_ip` records the IP used for the most recent
  successful sign-in independently from the latest observed device IP.
- `wallet.MembershipPassContent` and `wallet.MembershipPassBarcode` define the
  provider-neutral membership, tier, points, and Code 128 barcode snapshot.
- `walletenum.WalletPassPlatform` and
  `walletenum.WalletPassBarcodeFormat` classify supported pass targets and the
  canonical barcode format.
- `productenum.WholesalePriceMode` distinguishes fixed prices from
  approved-buyer price-on-request catalogue entries.

### Safety And Ownership

- No endpoint paths, service scopes, provider signing payloads, HTTP envelopes,
  authorization policy, or Apple/Google media definitions are introduced.
- Backend services remain responsible for session introspection, pass signing,
  wholesale enquiry behavior, and transport validation.

### Compatibility And Consumer Action

- All fields and enums are additive and keep the
  `github.com/Potato-Mart/Backend-Shared-Contract/v17` module path.
- Management, Operations, and Commerce should upgrade to `v17.4.0`, run
  `go mod tidy`, and update their backend-owned auth/API behavior together.

## v17.3.0 (2026-07-16) - Storefront Address And Promotion Models

This minor release adds the model surface required for customer-owned saved
address CRUD and a safe public promotions feed. The module path remains
`/v17`.

### Added

- `common.ContactAddress.id` is an optional stable identifier for persisted
  address-book entries. Inline checkout, order, billing, and fulfilment
  snapshots may continue to omit it.
- `promotion.StorefrontPromotion` is a customer-safe promotion catalogue
  record with display copy, public target references, active window, and
  computed active state.

### Safety

- `promotion.StorefrontPromotion` deliberately excludes discount rules,
  pricing configuration, usage limits and counters, priority, stacking
  settings, source metadata, audit history, and timestamps.
- Backend-specific routes, request bodies, pagination, and authorization stay
  in the owning services.

### Compatibility

- Additive JSON shape only; the Go module remains
  `github.com/Potato-Mart/Backend-Shared-Contract/v17`.
- Existing decoders may ignore `ContactAddress.id`.
- Consumers that expose saved-address CRUD or storefront promotion lists should
  upgrade to `v17.3.0` and run `go mod tidy`.

## v17.2.1 (2026-07-16) - Release Alignment Publication

This patch release republishes the existing V17.2 contract payload after the
repository release-alignment workflow was strengthened. The module path remains
`/v17`.

### Compatibility

- No JSON property, requiredness rule, enum wire value, exported contract type,
  package path, or fixed-point representation changes from `v17.2.0`.
- Existing V17.2 consumers can upgrade without model, serialization, or
  migration changes.

### Consumer Action

- Upgrade the `/v17` dependency to `v17.2.1` and run `go mod tidy`.
- Continue to keep API validation, persistence, authorization, workflows, and
  OpenAPI metadata in the owning backend services.

## v17.2.0 (2026-07-16) - Import-Compliance Model Foundation

This additive minor release introduces reusable, model-only records for the
administrative import-compliance workflow. It does not define HTTP DTOs,
authorization, validation, lifecycle transitions, calculations, automation,
external-provider behavior, or regulatory decisions. The module path remains
`/v17`.

### Added

- `contracts/importcompliance.ImportSettings` and its nested air, ambient-sea,
  frozen-sea, and ingredient-declaration models preserve the import settings
  surface. Charges use `common.Money`; the exchange rate uses micros; margin
  and Taiwan tax use basis points; and weight/volume inputs use grams and cubic
  centimetres.
- `ManufacturerDeclaration`, `LabelMaster`, `TariffAssessment`,
  `TariffProfile`, `TrademarkEvidence`, and `RFIRecord` provide revisioned
  records for the five administrative import pages. Declaration/order/product
  source data is frozen in snapshots, while signature, package-photo,
  attachment, and generated-file bytes are referenced through managed media.
- `RevisionMetadata`, `EvidenceReference`, `CatalogueReference`, and
  `ArtifactReference` preserve review history, cited provenance, official
  catalogue versions/checksums, and deterministic artifact identity without
  embedding transport or storage behavior.
- `enums/importcompliance` defines stable review states, AU/TW jurisdictions,
  import modes, evidence kinds, RFI channels/submission states/requested times,
  label sizes/orientations, and generated-artifact kinds.
- JSON-shape, enum-validation, model-boundary, and exported-model manifest
  coverage protects the new wire values and fixed-point representations.

### Compatibility

- No existing package, exported type, field, or enum value is removed or
  renamed. Existing V17 consumers can upgrade without changing unrelated
  model usage.
- The new models are deliberately persistence-neutral. Owning services must
  continue to publish API DTOs and workflows through their own OpenAPI
  documents.
- Tariff rates retain authoritative raw text and expose basis points only when
  a percentage representation is faithful. These models do not assert tariff
  classification or trademark clearance.
- RFI submission events record explicit external evidence; the model does not
  treat opening a portal or downloading a package as submission.

### Consumer Action

- Upgrade the `/v17` dependency to `v17.2.0` after the tag is published and run
  `go mod tidy`.
- Use the fixed-point units documented on the models and preserve catalogue,
  evidence, snapshot, revision, and artifact references when mapping service
  persistence or API DTOs.
- Keep review transitions, ETags/idempotency, permissions, AI/OCR proposals,
  tariff ingestion, RFI handoff, and artifact rendering in the owning backend;
  do not infer those behaviors from the shared records.

## v17.1.0 (2026-07-16) - Receipt-Safe Promotion Messaging

This additive minor release gives promotion authors an explicit, localized copy
surface for POS receipts and tax invoices while keeping internal promotion
configuration out of buyer-facing payloads. The module path remains `/v17`.

### Added

- `promotion.Promotion.receipt_enabled` explicitly approves a promotion for
  receipt and tax-invoice printing.
- `promotion.Promotion.receipt_messages` carries customer-facing copy as
  `common.LocalizedName` values instead of reusing internal campaign names or
  descriptions.
- `promotion.ReceiptOffer` provides the buyer/POS-safe projection: promotion
  id, localized receipt messages, optional active-window timestamps, and
  priority.
- JSON and exported-model manifest coverage verifies the new fields and guards
  the projection from leaking rule details, discount configuration, usage
  counters, source metadata, or internal authoring copy.

### Compatibility

- No existing fields, packages, or wire values are removed or renamed, and the
  module continues to use
  `github.com/Potato-Mart/Backend-Shared-Contract/v17`.
- Promotion JSON now includes `receipt_enabled` and `receipt_messages`;
  consumers with strict schemas must allow these additive fields.
- `ReceiptOffer.starts_at` and `expires_at` remain absent when unset because
  both fields use `omitempty`.

### Consumer Action

- Upgrade the `/v17` dependency to `v17.1.0` after the tag is published and run
  `go mod tidy`.
- Author and approve localized `receipt_messages`, and print them only when
  `receipt_enabled` is true. Never substitute `Promotion.name` or
  `Promotion.description` when approved receipt copy is unavailable.
- Backends should expose the minimal `ReceiptOffer` projection to receipt/POS
  consumers instead of serializing the full promotion model.

## v17.0.0 (2026-07-15) - Retail Wallet And Checkout Benefits

### Breaking removals and ownership changes

- Changes the module path to
  `github.com/Potato-Mart/Backend-Shared-Contract/v17` and reports exactly
  `v17.0.0` from `versioning.ModuleVersion`.
- Removes `wallet.WalletExport*`, `walletenum.WalletExportFormat`, and
  `walletenum.WalletExportStatus` with no compatibility aliases. Runtime
  wallet-export routes, storage, and generated API documentation are removed by
  their owning services.
- Replaces coupon assignment `customer_number` with the required generalized
  `owner` reference. Coupon usage uses the same owner reference when the
  redemption has an authenticated retail customer or wholesale organisation.

### Wallet, reservation, and order models

- Replaces ambiguous gift-card `balance` fields with committed, reserved, and
  available balances. Wallet instruments expose the same three balances;
  wallet summaries aggregate all three and use available value for storefront
  display.
- Adds typed voucher lifecycle state including `reserved`, reservation linkage
  and expiry timestamps, exact source reward-redemption linkage for idempotent
  backfill, and a voucher face-value projection on wallet instruments.
- Adds the idempotency-keyed `CheckoutBenefitReservation` record with distinct
  coupon, voucher, and ordered gift-card allocation state, including refunded
  gift-card amounts for deterministic partial-refund replay.
- Adds voucher redemption snapshots and reshapes the existing gift-card order
  snapshot around applied amount, reservation id, and committed wallet
  transaction id.
- Adds the `gift_card` payment method and a wallet provider reference containing
  the gift-card code and authoritative wallet transaction id.
- Adds `replaces_gift_card_code` so refund-issued replacements remain linked to
  expired or voided source cards.
- Adds membership point-ledger reason `REFUND` so original-source point
  restoration is distinguishable from manual adjustments and new earnings.

### Consumer action

- Pin `github.com/Potato-Mart/Backend-Shared-Contract/v17 v17.0.0` and replace
  all `/v16` imports.
- Remove wallet-export APIs and persisted export payloads rather than decoding
  them into v17.
- Migrate coupon ownership before serving v17 traffic and initialize every gift
  card's reserved balance to zero, with available balance equal to committed
  balance.
- Treat checkout quote/start/commit transports as provider-owned API DTOs; the
  shared module remains model-only.

## v16.0.0 (2026-07-13) - Model-Only Contract Hard Cut

### Breaking boundary changes

- Changes the module path to
  `github.com/Potato-Mart/Backend-Shared-Contract/v16` and reports exactly
  `v16.0.0` from `versioning.ModuleVersion`.
- Removes every endpoint request/response/query/command/acknowledgement and
  pagination DTO, route/path inventory, service scope catalogue, token endpoint
  type, and HTTP response envelope/helper.
- Removes portal/account mappings, permission resolution, state transitions,
  calculations, normalizers, and every `pkg/logic` package. Backend services
  own these policies and publish their transports through provider OpenAPI.
- Removes standalone preorder records/statuses and wholesale compatibility
  aliases. Preorder state is stored on cart/order models.
- Keeps error codes as typed data enums while each backend owns HTTP status and
  envelope mapping.

### Shared model additions and changes

- Adds identity-domain/account/portal isolation fields to claims, sessions,
  refresh/security records, plus group-order-manager application records.
- Adds order-owned preorder snapshots/allocation state and fulfillment
  readiness.
- Adds storefront campaign planning targets, immutable prediction revisions and
  evidence, packing projections, stock-arrival events, and persisted
  group-order discount decisions.
- Renames the preorder limit field to `max_quantity_per_order`.

### Boundary governance

- Adds an exported-model classification manifest and digest.
- Adds AST tests that reject endpoint DTO naming, transport tags, paths, scopes,
  free functions, non-intrinsic methods, exported aliases, and forbidden
  packages.

### Consumer action

- Pin `github.com/Potato-Mart/Backend-Shared-Contract/v16 v16.0.0`.
- Remove v15 imports and all local compatibility decoding.
- Own endpoint DTOs, paths, scopes, validation, authorization, and workflows in
  the provider backend; generate consumer clients from provider OpenAPI.

## v15.2.0 (2026-07-12) - Durable Notifications And Idempotent Stock Operations

Release date: 2026-07-12

This additive minor release provides shared domain/wire contracts for durable
customer notifications, paid-preorder availability, atomic stock reservation,
and packing-time stock settlement without changing the `/v15` module path.

### Added

- Customer notification topic, channel, and delivery-status enums.
- Customer-safe portal notification and delivery receipt contracts.
- A preorder availability command that carries stable business identifiers and
  intentionally excludes recipient addresses, subjects, and message bodies.
- Customer list/mark-read and internal preorder-availability path constants.
- The least-privilege `notification:send` service-auth scope.
- Idempotent atomic reservation and packing-settlement commands/results plus the
  internal packing-settlement path.
- An optional paid-preorder readiness projection carrying reservation ids and
  the durable customer-notification delivery receipt.

### Compatibility

- All additions are optional/new exported symbols; existing `/v15` payloads and
  paths are unchanged.
- Existing reserve, commit, and release path constants remain available.

### Consumer Action

- Upgrade `github.com/Potato-Mart/Backend-Shared-Contract/v15` to `v15.2.0`
  after the tag is published and run `go mod tidy`.
- Management should own message copy, verified-recipient resolution, portal
  persistence, and email delivery state.
- Operations and Commerce should use stable idempotency keys and reject reuse
  with a different request fingerprint.

## v15.1.0 (2026-07-11) - Outbound Delivery Completion

Release date: 2026-07-11

This additive minor release lets warehouse and delivery consumers represent the
final delivery milestone on an outbound shipment without changing the `/v15`
module path or any existing JSON field.

### Added

- `warehouseenum.OutboundShipmentStatusDelivered` with wire value `delivered`.
- Optional `warehouse.OutboundShipment.delivered_at` timestamp.
- JSON and enum validation coverage for the delivered state and timestamp.

### Compatibility

- Existing `packed` and `dispatched` statuses remain unchanged.
- `delivered_at` uses `omitempty`; existing payloads and consumers remain valid.
- No package or module path changes are required.

### Consumer Action

- Upgrade `github.com/Potato-Mart/Backend-Shared-Contract/v15` to `v15.1.0`
  and run `go mod tidy`.
- Consumers that complete deliveries should persist the `delivered` status and
  stamp `delivered_at`; read-only consumers may continue ignoring the new
  optional field.

## v15.0.0 (2026-07-09) - Retail Account And Portal Cleanup

Release date: 2026-07-09

This major release aligns the shared identity and retail customer contracts with
the retail/wholesale portal model, removes legacy contact and customer activity
shapes, and adds the shared fields needed by retail profile completion,
referral rewards, billing, gender, and order packing progress.

### Breaking Contract Changes

- `go.mod` module path changes to `github.com/Potato-Mart/Backend-Shared-Contract/v15`.
- All in-repository imports now use `/v15`; consumers must update imports from
  `/v14` to `/v15` before upgrading.
- `accountenum.AccountTypeGeneralCustomer` is replaced by
  `accountenum.AccountTypeRetailCustomer`, with the wire value changing from
  `generalCustomer` to `retailCustomer`.
- `accountenum.PortalStore` and `accountenum.PortalPartner` are replaced by
  `accountenum.PortalRetail` and `accountenum.PortalWholesale`, with wire values
  changing from `store`/`partner` to `retail`/`wholesale`.
- `common.ContactChannels.mobile` and `common.ContactChannels.line_id` are
  removed. Use `phone` for the primary phone number and `external_handles` for
  channel-specific handles such as LINE.
- `customers.RetailCustomer.recent_orders` is removed from the shared customer
  profile; consumers should fetch order history/projections from Commerce.
- `customers.RetailCustomerActivity` and `customerenum.CustomerActivityType` are
  removed from the shared contract.
- `customers.RetailCustomerReferralProfile.credited` is removed and replaced by
  explicit referral confirmation, usage, reward count, voucher code, and reward
  timestamp fields.
- `pkg/versioning.ModuleVersion` now reports `v15.0.0`.

### Added

- `customerenum.CustomerGender` and `customers.RetailCustomerBasicInfo.gender`.
- `customers.RetailCustomerProfileCompletion` and
  `customers.RetailCustomer.profile_completion`.
- `customers.RetailCustomer.default_billing`.
- Expanded `customers.RetailCustomerReferralProfile` reward tracking fields.
- `sales.Order.packing` with `sales.OrderPackingProgress`, reusing warehouse
  packing lines, box plan, damage, and discrepancy contracts.
- JSON-shape tests for the retail customer profile and sales order packing
  progress changes.

### Documentation And Release Notes

- README usage examples now point to `v15.0.0` and the `/v15` module path.
- Identity access documentation now uses the retail/wholesale portal names.
- The generated lowercase `release-notes.md` content has been consolidated into
  this canonical `RELEASE_NOTES.md` history.

### Consumer Action

- Upgrade imports and dependencies from `/v14` to `/v15`, then run
  `go mod tidy`.
- Update any persisted or serialized account/portal values from
  `generalCustomer`/`store`/`partner` to `retailCustomer`/`retail`/`wholesale`.
- Move LINE and other channel-specific contact handles into
  `contacts.external_handles`.
- Stop reading `recent_orders` from retail customer payloads; fetch order
  summaries from Commerce-owned order projections.
- Replace shared customer activity timeline usage with the owning service's
  history or CRM activity model.

## v14.1.0 (2026-07-09) - Storefront Slugs And Account Avatars / 零售網址 Slug 與帳號頭像

Release date: 2026-07-09

This minor release adds the shared contract fields needed by the retail
storefront profile and product navigation work. Product collection and category
records can now expose stable public slugs while preserving their canonical ids,
and the shared user profile can project the caller's avatar media reference and
public avatar URL.

本 minor release 新增零售網站個人資料與商品導覽所需的共用欄位。商品集合與分類可在保留
canonical id 的同時提供穩定公開 slug；共用使用者資料也可投影登入者的頭像媒體參照與公開頭像網址。

### Added / 新增

- `product.CollectionRef.slug` and `product.Collection.slug` are optional JSON
  fields for public product collection URLs.
- `product.CategoryTag.slug` is an optional JSON field for public category tag
  URLs.
- `identity.UserProfile.avatar_media_id` and `identity.UserProfile.avatar_url`
  are optional JSON fields for account avatar projection.
- JSON-shape tests cover slug inclusion/omission and avatar field
  inclusion/omission.

### Contract Files Changed

- `pkg/contracts/identity/user.go`
- `pkg/contracts/identity/user_json_test.go`
- `pkg/contracts/product/category_tag.go`
- `pkg/contracts/product/collection.go`
- `pkg/contracts/product/json_shape_test.go`
- `pkg/versioning/version.go`

### Compatibility / 相容性

- No module path change: consumers remain on
  `github.com/Potato-Mart/Backend-Shared-Contract/v14`.
- All new fields use `omitempty`; existing JSON consumers that ignore unknown
  fields remain compatible.
- Canonical ids remain unchanged. Slugs are public routing/display helpers, not
  primary persistence keys.

### Consumer Action / 使用方動作

- Upgrade `github.com/Potato-Mart/Backend-Shared-Contract/v14` to `v14.1.0` and
  run `go mod tidy`.
- Backends that project storefront product collections/categories should emit
  slug fields when available or derive/backfill them from English names.
- Backends that expose `/users/me` should project avatar fields where supported.

## v14.0.0 (2026-07-08) - Wholesale Customer Organisation Consolidation / 批發客戶組織化整併

Release date: 2026-07-08

This major release makes the wholesale organisation the canonical wholesale
customer/business account. `wholesale.WholesaleCustomer` and
`WholesaleCustomerSummary` remain exported as compatibility names for the
organisation contracts, while people inside the business are represented by
`OrganisationAccess` records. The stable sign-in identity now belongs to the
organisation principal, so changing the main person-in-charge changes
organisation/contact/access metadata without changing the principal
`AuthIdentity`.

### Breaking Contract Changes / 破壞性契約變更

- `go.mod` module path changes to `github.com/Potato-Mart/Backend-Shared-Contract/v14`.
- All in-repository imports now use `/v14`; consumers must update imports from
  `/v13` to `/v14` before upgrading.
- `wholesale.WholesaleCustomer` is now an alias of
  `wholesale.WholesaleOrganisation`; the old person-profile fields
  `customer_number`, `basic_info`, `commercial`, `account_profile`,
  `organisation_access_id`, and `primary_wholesale_customer_id` are removed
  from the wholesale customer shape.
- `WholesaleOrganisation` adds organisation-principal identity references:
  `principal_user_id`, `principal_account_id`, `primary_auth_identity_id`,
  `auth_identity_ids`, and `primary_organisation_access_id`.
- `OrganisationAccess` is the canonical organisation-person/team-member
  contract and adds optional `name`, `contacts`, and `department`.
- Wholesale application responses no longer expose
  `wholesale_customer_number` or `customer_status`.
- `WholesaleAccountTerms` no longer exposes a separate `customer_id`; the
  customer/business key is `organisation_code`.
- `pkg/versioning.ModuleVersion` now reports `v14.0.0`.

### Consumer Action / 使用方動作

- Upgrade imports and dependencies from `/v13` to `/v14`, then run
  `go mod tidy`.
- Treat `organisation_code` / `wholesale_organisation_code` as the wholesale
  customer/business account key.
- Store people, roles, status, and the main person-in-charge through
  `OrganisationAccess`; enforce one active primary access per organisation in
  the owning backend.
- Stop reading or writing separate wholesale customer numbers for B2B accounts;
  migrate any per-person wholesale customer profile data into organisation
  records or organisation access records.

## v13.3.0 (2026-07-08) - Back-In-Stock Notification Contracts / 到貨通知契約

Release date: 2026-07-08

This minor release adds the shared contract surface for account-only
back-in-stock notifications. Management remains the owning API and persistence
service, while Operations can emit a service-authenticated restock event when a
storefront-visible, sellable SKU moves from unavailable to available. The
subscription model is one-shot: each pending channel subscription can be marked
`notified` once, and a customer can subscribe again after notification.

本 minor release 新增登入帳號專用的到貨通知共用契約。Management 仍是 API 與持久化資料
的擁有服務；Operations 可在 storefront 可見、可銷售的 SKU 從不可購買變成可購買時，
送出 service-authenticated restock event。訂閱模型為一次性通知：每個 pending channel
subscription 只能被標記為 `notified` 一次，通知後客戶可再次訂閱。

### Additive Changes / 新增相容變更

- `notification.BackInStockSubscription` captures account-owned SKU
  subscriptions with `product_sku_code`, `user_id`, customer type, selected
  channel, locale, lifecycle status, consent snapshot, requested/notified/
  cancelled timestamps, and delivery-error metadata.
- `notification.BackInStockConsentSnapshot` records the email/SMS consent and
  contact availability observed at subscription time.
- `notification.BackInStockDeliveryError` stores the last provider/channel
  delivery failure code and message without making provider adapters part of the
  shared contract.
- `notification.BackInStockRestockEvent` is the durable payload Operations sends
  when a restock should be evaluated by Management.
- `notificationenum.BackInStockChannel` adds `email` and `sms`.
- `notificationenum.BackInStockStatus` adds `pending`, `notified`, and
  `cancelled`.
- `serviceauth.ScopeRestockNotify` adds the `restock:notify` internal
  service-auth scope.
- `pkg/versioning.ModuleVersion` now reports `v13.3.0`.

### Consumer Action / 使用方動作

- Backend services that implement back-in-stock notifications should upgrade to
  `github.com/Potato-Mart/Backend-Shared-Contract/v13 v13.3.0` after this
  contract tag is pushed.
- Backend-Management should expose the account-owned subscription endpoints,
  dedupe `user_id + product_sku_code + channel + pending`, enforce account
  contact/consent checks for SMS, and mark subscriptions `notified` after a
  successful one-shot trigger.
- Backend-Operations should request/use the `restock:notify` scope and send
  `BackInStockRestockEvent` only for storefront-visible, sellable SKU restocks
  that transition from unavailable to available.
- Frontends should treat these as account-only subscriptions, redirect signed-out
  users to sign-in, and call the owning Management-backed APIs rather than
  persisting local notification promises.

## v13.2.0 (2026-07-08) - Paid Preorder Checkout Metadata / 付費預購結帳中繼資料

Release date: 2026-07-08

This minor release adds optional cart-line metadata for prepaid preorder
checkout. Commerce can now stamp a cart line with a server-validated fulfilment
mode such as `{"fulfilment_mode":"preorder"}`, carry it into the resulting order
line, and decide whether immediate stock reservation is required. The field is
additive, optional, and omitted for ordinary cart lines, so existing consumers
that ignore unknown JSON fields remain compatible.

本 minor release 新增可選的購物車明細中繼資料，用於付費預購結帳。Commerce 可在通過
伺服器端驗證後，將明細標記為例如 `{"fulfilment_mode":"preorder"}`，並把該標記帶入
後續訂單明細，以決定是否需要立即預留庫存。此欄位為相容新增、可省略；一般購物車明細
不會輸出，忽略未知 JSON 欄位的既有使用方可保持相容。

### Additive Changes / 新增相容變更

- `sales.CartItem.Properties` (`json:"properties,omitempty"`) carries optional
  metadata for cart lines before checkout turns them into `sales.OrderItem`
  records.
- `pkg/versioning.ModuleVersion` now reports `v13.2.0`.

### Consumer Action / 使用方動作

- Backend services that need prepaid preorder checkout should upgrade to
  `github.com/Potato-Mart/Backend-Shared-Contract/v13 v13.2.0` after this
  contract tag is pushed.
- Consumers that only read ordinary carts do not need code changes unless they
  want to display or preserve cart-line properties.

## v13.1.0 (2026-07-07) - Group-Order Buyer Discount And Manager Permissions / 團購買方折扣與團購管理權限

Release date: 2026-07-07

This minor release adds the shared contract surface for the buyer-facing group
order feature: a per-group-order discount application/approval domain model, its
lifecycle enum, an internal endpoint constant plus a service-auth scope for the
discount grant read, and the wholesale-buyer permission keys a group-order
manager needs. It keeps the `/v13` Go module path and does not remove, rename, or
narrow any exported type or JSON field. Group-order discount money is carried as
`common.Money` minor units (fixed amount) or an integer basis-point value
(percentage), never as a stringified major-unit value.

本 minor release 新增團購（group order）買方功能所需的共用契約：每個團購的折扣
申請／核准 domain 模型、其生命週期列舉、折扣讀取用的 internal endpoint 常數與
service-auth scope，以及團購管理者所需的批發買方權限。此版本維持 `/v13` Go module
path，未移除、改名或縮窄任何 exported type 或 JSON 欄位。團購折扣金額一律以
`common.Money`（最小貨幣單位，定額）或整數 basis points（百分比）表示，不使用
字串型主單位數值。

### Additive Changes / 新增相容變更

- `promotionenum.GroupOrderDiscountState` (`pending` / `approved` / `rejected`)
  captures the lifecycle of a per-group-order discount application submitted by a
  wholesale group-order manager and decided by a staff approver.
- `promotion.GroupOrderDiscountApplication` is the shared application record
  produced by Backend-Management and read by Backend-Commerce and the admin
  console: group order code, organisation/access, state, the selected or approved
  promotion id, the requested proposal, and requester/decider audit fields.
- `promotion.GroupOrderDiscountProposal` describes a newly requested benefit with
  a `common.Money` `amount` (fixed amount, minor units), an optional `max_discount`
  cap (minor units), and integer `percent_basis_points` (percentage, 1000 =
  10.00%). It deliberately avoids `DiscountSpec`'s stringified major-unit
  `discount_value`, so no float or major-unit money crosses the wire.
- `promotion.PathGroupOrderDiscountInternal`
  (`/v1/internal/promotions/group-order-discount`) is the internal endpoint
  Backend-Commerce uses to resolve an application's approved promotion for a group
  order.
- `serviceauthenum.ScopePromotionGrant` (`promotion:grant`) authorises that
  internal group-order-discount read.
- `wholesaleenum` gains four group-order manager permissions —
  `group_orders.manage_org`, `group_orders.invite`, `group_orders.submit`, and
  `group_order_discount.apply` — all granted to the `owner` buyer role by
  `PermissionsForWholesaleBuyerRole` and validated by `IsValid`.
- `pkg/versioning.ModuleVersion` now reports `v13.1.0`.
- Consistent with the contract's DTO-ownership guard, the manager-apply and
  admin approve/reject wire request envelopes are NOT added to the shared
  contract; they remain owned by Backend-Management. The shared module keeps only
  the domain record, the proposal, the enum, the endpoint constant, the scope, and
  the permission keys.

### 新增相容變更

- `promotionenum.GroupOrderDiscountState`（`pending`／`approved`／`rejected`）
  表示由批發團購管理者提交、由後台核准人員裁決的每團購折扣申請生命週期。
- `promotion.GroupOrderDiscountApplication` 是由 Backend-Management 產生、供
  Backend-Commerce 與後台管理台讀取的共用申請紀錄：團購代碼、組織／存取、狀態、
  選定或核准的 promotion id、申請提案，以及申請者／裁決者稽核欄位。
- `promotion.GroupOrderDiscountProposal` 描述新申請的折扣：`common.Money` 的
  `amount`（定額，最小單位）、選用的 `max_discount` 上限（最小單位），以及整數
  `percent_basis_points`（百分比，1000 = 10.00%）。刻意不重用 `DiscountSpec` 的
  字串型主單位 `discount_value`，避免任何浮點或主單位金額進入 wire。
- `promotion.PathGroupOrderDiscountInternal`
  （`/v1/internal/promotions/group-order-discount`）為 Backend-Commerce 解析某團購
  已核准 promotion 所使用的 internal endpoint。
- `serviceauthenum.ScopePromotionGrant`（`promotion:grant`）授權該 internal
  團購折扣讀取。
- `wholesaleenum` 新增四個團購管理權限——`group_orders.manage_org`、
  `group_orders.invite`、`group_orders.submit`、`group_order_discount.apply`——
  由 `PermissionsForWholesaleBuyerRole` 全數授予 `owner` 買方角色並納入 `IsValid`。
- `pkg/versioning.ModuleVersion` 現在回報 `v13.1.0`。
- 依照契約的 DTO 歸屬守則，管理者申請與後台核准／退回的 wire request 信封不加入
  共用契約，仍由 Backend-Management 擁有；共用模組只保留 domain record、proposal、
  列舉、endpoint 常數、scope 與權限鍵。

### Consumer Action / 使用方動作

- Management / Operations / Commerce: upgrade the pin to
  `github.com/Potato-Mart/Backend-Shared-Contract/v13 v13.1.0` and run
  `go mod tidy`. No `/v14` module-path migration is required; all changes are
  additive.
- Backend-Management: own the manager-apply and admin approve/reject request DTOs
  locally, persist the `GroupOrderDiscountApplication` record, grant the new
  wholesale permissions to group-order managers, and gate the internal
  group-order-discount read with the `promotion:grant` scope.
- Backend-Commerce: resolve the approved promotion for a group order through
  `PathGroupOrderDiscountInternal` (or by storing the approved promotion id on the
  group order) and apply it at pricing/submit; carry all discount money as
  `common.Money` minor units.
- Frontends: treat the group-order discount fields as optional; render the
  proposal amount from the `common.Money` field, never from the percentage
  integer as money.
- All consumers: run `go mod tidy` and contract JSON round-trip tests after
  upgrading.

- Management／Operations／Commerce：將相依固定升級為
  `github.com/Potato-Mart/Backend-Shared-Contract/v13 v13.1.0` 並執行
  `go mod tidy`。全部為 additive，不需要 `/v14` module-path 遷移。
- Backend-Management：於後端自行擁有管理者申請與後台核准／退回的 request DTO、
  持久化 `GroupOrderDiscountApplication` 紀錄、將新批發權限授予團購管理者，並以
  `promotion:grant` scope 保護 internal 團購折扣讀取。
- Backend-Commerce：透過 `PathGroupOrderDiscountInternal`（或將已核准 promotion id
  存於團購上）解析某團購的已核准 promotion，並於定價／送出時套用；所有折扣金額以
  `common.Money` 最小單位表示。
- 前端：將團購折扣欄位視為選用；折扣金額一律取自 `common.Money` 欄位，切勿將
  百分比整數當作金額顯示。
- 所有使用方：升級後執行 `go mod tidy` 與契約 JSON round-trip 測試。

## v13.0.0 (2026-07-06) - Reward Tier Triggers And Test Infrastructure / 獎勵等級觸發與測試基礎設施

Release date: 2026-07-06

This major release moves the Go module path to `/v13`, adds tier-achievement
trigger fields to membership rewards, and hardens the contract test setup. The
enum test suite is split into small domain-owned files so enum cases and helper
behaviour can be maintained without growing a single large test file.

本 major release 將 Go module path 升至 `/v13`，為會員獎勵新增等級達成觸發欄位，
並強化 shared contract 的測試設定。enum 測試已拆成小型領域檔案，方便後續新增、
修改或刪除 enum case 與 helper 行為，而不再擴大單一大型測試檔。

### Breaking Changes / 破壞性變更

- `go.mod` module path changes to `github.com/Potato-Mart/Backend-Shared-Contract/v13`.
- All in-repository imports now use `/v13`; consumers must update imports from
  `/v12` to `/v13` before upgrading.
- No enum JSON wire values are changed by this release.

### Additive Contract Changes / 新增相容合約變更

- `membership.Reward` gains `trigger_tier_key` for rewards issued when a member
  reaches a specific tier.
- `membership.Reward` gains `issue_on_tier_achievement` to mark rewards that
  should be issued automatically on tier achievement.
- Both new reward fields are `omitempty`, so existing serialized reward records
  remain readable.

### Release Engineering / 發布工程

- The previous large `pkg/enums/enums_test.go` file is split into small
  domain-focused enum test files under `pkg/enums`.
- Shared enum assertions now live in `pkg/enums/enum_assertions_test.go`.
- Contract tests can be run with `scripts/powershell/Test-Contract.ps1` or
  `scripts/bash/test-contract.sh`, both using `GOWORK=off go test ./...` so the
  parent Potato Mart workspace does not hide this standalone module.
- GitHub test and release workflows run with `GOWORK=off`.

### Consumer Action / 使用方動作

- Update dependencies from `/v12` to `/v13`, then run `go mod tidy`.
- Treat the new reward tier-trigger fields as optional for existing data.
- Run contract serialization/deserialization tests after upgrading.

## v12.0.0 (2026-07-06) - Domain Enum Subpackages / 領域列舉子套件

Release date: 2026-07-06

This major release replaces the flat enum package with domain enum subpackages
under `pkg/enums/*`. It also hard-moves the remaining typed enum definitions
that still lived in contract, API response, and service-auth packages. No
compatibility aliases are kept, so consumers must update imports and enum
qualifiers while JSON wire values remain unchanged.

本 major release 將原本扁平的列舉套件拆分為 `pkg/enums/*` 下的領域列舉子套件，
並將仍留在 contract、API response、service-auth 套件中的 typed enum 全部硬搬遷。
本版本不保留相容別名，因此使用方必須更新 import 與列舉限定詞；JSON wire 值不變。

### Breaking Changes / 破壞性變更

- `go.mod` module path changes to `github.com/Potato-Mart/Backend-Shared-Contract/v12`.
- The former flat enum package is split into domain packages such as
  `pkg/enums/sales`, `pkg/enums/payment`, `pkg/enums/product`, and
  `pkg/enums/wallet`.
- `apiresponse.Code` moves to `pkg/enums/apiresponse`; `apiresponse.Error`
  remains in `pkg/apiresponse` and now uses the moved code type.
- `serviceauth.Scope` and scope helpers move to `pkg/enums/serviceauth`;
  service-auth endpoint constants and `ServiceClaims` remain in `pkg/serviceauth`.
- Preorder status, storefront preorder/expiry status, wholesale application
  state, and wallet export status are now enum-subpackage types instead of
  contract-package types.
- Provider-specific terminal adapter constants were removed from shared
  contracts; provider adapter code belongs in the owning backend.

### 破壞性變更

- `go.mod` module path 改為 `github.com/Potato-Mart/Backend-Shared-Contract/v12`。
- 原本扁平的 enum package 拆分為 `pkg/enums/sales`、`pkg/enums/payment`、
  `pkg/enums/product`、`pkg/enums/wallet` 等領域套件。
- `apiresponse.Code` 移至 `pkg/enums/apiresponse`；`apiresponse.Error` 留在
  `pkg/apiresponse`，並改用搬遷後的 code type。
- `serviceauth.Scope` 與 scope helpers 移至 `pkg/enums/serviceauth`；
  service-auth endpoint constants 與 `ServiceClaims` 留在 `pkg/serviceauth`。
- Preorder status、storefront preorder/expiry status、wholesale application
  state、wallet export status 改由 enum 子套件定義，不再由 contract 套件定義。
- shared contract 移除 provider-specific terminal adapter constants；provider
  adapter code 應留在擁有該整合的 backend。

### Consumer Action / 使用方動作

- Update dependencies from `/v11` to `/v12`, then replace flat enum imports with
  the matching domain enum package. Example: `salesenum.OrderTypeOnline`,
  `paymentenum.PaymentStatusPaid`, `walletenum.WalletExportFormatJSON`, and
  `serviceauthenum.ScopePricingQuote`.
- Replace old contract-local enum qualifiers such as `wallet.WalletExportFormat`,
  `sales.PreorderStatusRequested`, `product.StorefrontPreorderStatusOpen`, and
  `wholesale.WholesaleApplicationStatePending` with their new enum-subpackage
  qualifiers.

- 將依賴從 `/v11` 更新到 `/v12`，並把原本扁平的 enum import 改為對應領域的 enum
  子套件。例如：`salesenum.OrderTypeOnline`、`paymentenum.PaymentStatusPaid`、
  `walletenum.WalletExportFormatJSON`、`serviceauthenum.ScopePricingQuote`。
- 將舊的 contract-local enum 限定詞（例如 `wallet.WalletExportFormat`、
  `sales.PreorderStatusRequested`、`product.StorefrontPreorderStatusOpen`、
  `wholesale.WholesaleApplicationStatePending`）改為新的 enum 子套件限定詞。

## v11.2.0 (2026-07-06) - Sales Delivery Scheduling And Payment Processor Fees / 銷售配送排程與支付手續費

Release date: 2026-07-06

This minor release adds additive delivery-scheduling fields to sales orders,
payment processor fee fields, and the canonical MAMA order-number rule. It
keeps the `/v11` Go module path and does not remove, rename, or narrow any
exported type or JSON field.

本 minor release 為銷售訂單新增相容的配送排程欄位、支付手續費欄位，以及標準
MAMA 訂單編號規則。此版本維持 `/v11` Go module path，沒有移除、重新命名或縮窄
任何 exported type 或 JSON 欄位。

### Additive Changes / 新增相容變更

- `sales.Order` gains `expected_delivery_date` (`common.Date`) and
  `expected_delivery_time` (`common.TimeOfDay`) for the promised delivery
  slot, distinct from the after-the-fact `delivered_at` lifecycle timestamp.
- `enums.DeliveryMethod` (`delivery`, `pickup`, `outsourced`) and
  `sales.Order.delivery_method` record the authoritative fulfilment method;
  `sales.Order.outsourced_carrier` names the third-party company when the
  method is `outsourced`.
- `enums.DeliveryRegion` (`local_melbourne`, `regional_vic`, `interstate`)
  and `sales.Order.delivery_region` classify the shipping destination for
  packing and dispatch planning; `shipping.Zone.is_local` marks the metro
  zone used to derive it.
- `sales.Payment` gains `processor_fee` and `net_amount` (`common.Money`,
  minor units) and `payments.StripePaymentReference` gains
  `balance_transaction_id` so processor-reported transaction fees can be
  stored and aggregated.
- `pkg/logic/sales` adds `OrderNumberPrefix`, `OrderNumberPattern`
  (`^MAMA\d{6}[A-Z0-9]{6}$`), and `IsValidOrderNumber` as the canonical
  order-number rule shared by services.
- `pkg/versioning.ModuleVersion` now reports `v11.2.0`.

### Consumer Action / 使用方動作

- Commerce: upgrade to
  `github.com/Potato-Mart/Backend-Shared-Contract/v11 v11.2.0`, mint order
  numbers in the MAMA shape, validate supplied numbers with
  `saleslogic.IsValidOrderNumber`, populate the new delivery fields, and
  store Stripe balance-transaction fees on payments.
- Management / Operations: upgrade the pin to `v11.2.0` and `go mod tidy`;
  no behavioural change required.
- Frontends: treat the new order/payment fields as optional (absent on
  pre-existing records) and mirror the order-number pattern literal only
  with a comment pointing at `pkg/logic/sales/order_number.go`.
- All consumers: run `go mod tidy` and contract JSON round-trip tests after
  upgrading. No `/v12` module-path migration is required.

## v11.1.0 (2026-07-06) - Retail Preorder And Expiry Merchandising Contracts / 零售預購與效期銷售契約

Release date: 2026-07-06

This minor release adds additive retail merchandising contracts for preorder
interest and soon-expiry product display. It keeps the `/v11` Go module path and
does not remove, rename, or narrow any exported type or JSON field.

本 minor release 新增零售預購與即期商品銷售展示所需的相容契約。此版本維持 `/v11`
Go module path，沒有移除、重新命名或縮窄任何 exported type 或 JSON 欄位。

### Additive Changes / 新增相容變更

- `sales.PreorderStatus`, `sales.Preorder`, and `sales.PreorderSummary` define
  durable preorder interest/status contracts for customer preorder history,
  customer preorder checks, and downstream order conversion visibility.
- `product.StorefrontMerchandising` adds optional product-level preorder policy
  fields through `product.PreorderPolicy`, including active windows, expected
  availability, quantity limits, and localized customer labels/descriptions.
- `product.SoonExpiryMerchandisingPolicy` adds an admin-configurable soon-expiry
  merchandising policy with active windows, display window days, exact-date
  visibility control, and localized customer labels/descriptions.
- `product.StorefrontDisplay`, `product.StorefrontPreorderDisplay`, and
  `product.StorefrontExpiryDisplay` provide backend-computed, storefront-safe
  display fields for product cards and detail pages without exposing stock lots,
  internal service names, or operational implementation details.
- `wholesale.ApprovedStorefrontProduct` now includes optional
  `storefront_display` so wholesale storefront responses can reuse the same
  customer-safe preorder and expiry display state.
- `pkg/versioning.ModuleVersion` and README usage examples now point to
  `v11.1.0`.

### Consumer Action / 使用方動作

- Management: upgrade to
  `github.com/Potato-Mart/Backend-Shared-Contract/v11 v11.1.0`, then map any
  admin-configurable preorder and soon-expiry settings into the new product
  policy structs. Keep endpoint-specific command/request DTOs in Management if
  Management owns those APIs.
- Operations: upgrade to `v11.1.0`, then map catalog expiry/stock-lot decisions
  into `product.StorefrontExpiryDisplay` only after filtering out lot IDs,
  backend module names, and other operational details that storefronts should
  not see.
- Commerce: upgrade to `v11.1.0`, use `sales.Preorder` /
  `sales.PreorderSummary` for preorder persistence or read projections, and
  continue owning checkout/order-conversion command payloads locally.
- Retail and wholesale storefronts: consume backend-provided
  `storefront_display.preorder` and `storefront_display.expiry` fields for
  cards, detail pages, comparison views, profile preorder checks, and home
  listing merchandising. Do not infer preorder eligibility or expiry truth from
  frontend-only state.
- All consumers: run `go mod tidy`, contract JSON round-trip tests, and the
  affected API/frontend integration tests after upgrading. No `/v12` module-path
  migration is required.

## v11.0.0 (2026-07-05) - Unified Coupons and Wallet Export Contracts / 統一優惠券與錢包匯出契約

Release date: 2026-07-05

This major release moves coupon ownership to the Coupon aggregate and adds the shared
wallet export contracts used by Management-owned wallet export APIs. It upgrades the Go
module path from `/v10` to `/v11`. The database is empty for this migration window, so no
old customer-coupon data migration or backward-compatible read path is required.

本 major release 將優惠券發放歸屬移至 Coupon aggregate，並新增 Management 錢包匯出 API
使用的共用錢包匯出契約。Go module path 從 `/v10` 升級為 `/v11`。本次遷移期間資料庫為空，
因此不需要舊 customer-coupon 資料遷移，也不保留向後相容讀取路徑。

### Breaking Changes / 破壞性變更

- Removed public `promotion.CustomerCoupon`. Consumers must use Coupon-owned assignment
  and history contracts instead.
- Module path changes from `github.com/Potato-Mart/Backend-Shared-Contract/v10` to
  `github.com/Potato-Mart/Backend-Shared-Contract/v11`.

### Added / 新增

- `promotion.CouponAssignment`, `promotion.CouponAssignmentSummary`,
  `promotion.CouponDetail`, `promotion.CouponIssueSpec`, and
  `promotion.CouponRecipientPreview`.
- Wallet export contracts with `schema_version: "wallet_export_v1"` for async export
  request/status/result flows.
- Wallet export item/history models for membership points, coupon assignments/usage,
  vouchers, gift cards, rewards, normalized history, filters, record counts, checksum,
  status, requester, and generated/export metadata.

### Consumer Action / 使用方動作

- Update imports and `go.mod` from `/v10` to `/v11` and require
  `github.com/Potato-Mart/Backend-Shared-Contract/v11 v11.0.0`.
- Remove local `/v10` replace directives and run `go mod tidy`.
- Run contract JSON round-trip tests and backend integration tests for coupon assignment,
  coupon redemption, wallet export, checkout, and any consumer package importing shared
  promotion or wallet contracts.

## v10.1.1 (2026-07-02) - Retail Customer Number Identity Claim / 零售顧客編號身份聲明

Release date: 2026-07-02

This patch release adds an optional retail customer business number claim to the shared
access-token contract. It keeps the `/v10` module path and does not require a data migration
or compatibility fallback.

本 patch release 在共用 access-token 契約中加入可選的零售顧客業務編號聲明。此版本維持
`/v10` module path，不需要資料遷移，也不需要相容性 fallback。

### Additive Changes / 新增相容變更

- `identity.AccessTokenClaims` now includes optional `retail_customer_number`.
- Services can use this claim to mirror the retail customer number that Management issues
  for accounts linked to a retail-customer profile.

### Consumer Action / 使用方動作

- Upgrade backend services and other consumers to
  `github.com/Potato-Mart/Backend-Shared-Contract/v10 v10.1.1`.
- Ensure any service-local JWT/access-token claim mirrors accept `retail_customer_number`.
- No module path change, schema migration, or breaking JSON handling change is required.

## v10.1.0 (2026-06-30) - Product Category/Collection Contract Cleanup / 商品分類與集合契約整理

Release date: 2026-06-30

This release replaces the legacy product category/catalogue shape with collection and
category-tag contracts. The database is empty, so no legacy decode, fallback, or migration
compatibility fields are retained. Although these are breaking JSON and enum wire changes,
the release is published as the requested `v10.1.0` metadata version and keeps the `/v10`
module path.

### Breaking Changes / 破壞性變更

- `product.Product` removes `category_key`, `category_path`, `catalogue`, and the nested
  `merchandising` group.
- Product and wholesale product projections now expose category/collection data as a lightweight
  `collection` object plus top-level `category_tags`.
- `product.Collection.name`, `product.CollectionRef.name`, `product.CategoryTag.name`, and
  `product.CategoryTag.collection_name` are `[]common.LocalizedName`.
- `product.CategoryTag` now carries `id`, localized `name`, `collection_id`, and localized
  `collection_name`.
- Product-list auth naming changed from catalog/catalogue to products:
  `serviceauth.ScopeProductsRead` uses `products:read`, and
  `enums.WholesalePermissionProductsView` uses `products.view`.
- The internal wholesale product resolver path changed from `/v1/internal/catalog/wholesale/products`
  to `/v1/internal/wholesale/products`.
- Promotion targeting changes from category paths to category tags:
  `DiscountScopeCategoryTag` uses wire value `category_tag`,
  `target_category_tag_id` and localized `target_category_tag_name` replace
  `target_category_key`, and descendant matching is removed.
- Coupon targeting changes from `specific_categories` / `category_ids` to
  `specific_category_tags` / `category_tags`.

### Consumer Action / 使用方動作

- Update all services and clients that read product category/catalogue fields to use the
  product `collection` object and localized category tag names.
- Update service-token scope checks from `catalog:read` to `products:read`.
- Update wholesale storefront permission checks from `catalogue.view` to `products.view`.
- Update promotion and coupon rule management to store category-tag IDs and names, not
  category keys, category paths, or descendant flags.

## v10.0.0 (2026-06-30) - ID Reference Removal / 移除 ID 參照（改用代碼與單號）

Release date: 2026-06-30

Breaking `/v9` → `/v10` module path migration. This release completes the v9.4.0 canonical
code/number migration by removing every deprecated id/legacy alias and converting the
remaining cross-struct references to their code/number business key. Each struct keeps its
own `id`; only references whose target has no code/number stay id-based. HARD CUTOVER — all
legacy decode/backfill/fallback is removed (no backward compatibility; empty datastore).

### Breaking Changes / 破壞性變更

- Module path `github.com/Potato-Mart/Backend-Shared-Contract/v9` → `/v10`.
- Removed all deprecated `omitempty` id aliases (canonical siblings already existed): product
  `product_id`/`product_ids`, depot `depot_id`, order `order_id`/`purchase_order_id`/`sales_order_id`,
  supplier `supplier_id`, coupon `redeemed_order_id`, membership `related_order_id` (where a
  `related_order_number` sibling existed), and the inbound legacy `supplier`.
- Renamed remaining id references to code/number: `customer_profile_id` / `retail_customer_id` /
  cart `customer_id` / `referrer_id` → `*customer_number`; `coupon_id` → `coupon_code`;
  `reward_id` / `related_reward_id` → `reward_code` / `related_reward_code`;
  `wholesale_organisation_id` / `organisation_id` → `*organisation_code`; `wholesale_customer_id`
  → `wholesale_customer_number`; `device_id` (identity session, request context, order source) →
  `device_key`; analytics forecast `sku` and WMS draft/discrepancy `sku` → `sku_code` /
  `product_sku_code`; product/snapshot `supplier` → `supplier_code`.
- Removed deprecated non-id fields: product `vendor` and `brand_key`, snapshot `vendor`,
  storefront `vendor`/`brand_key`, `Payment.currency` (use `amount.currency`),
  `UserProfile.user_role`, and enum `UserRoleClient` plus `UserRole.IsStaff()`/`IsAdmin()`.
- Removed the product legacy-decode compat (`product/json_compat.go`) and the promotion
  resolver legacy `ProductID` fallback.
- `customers.RetailCustomerActivity.amount` changed from a bare number to `common.Money`.
- Removed storage-driver struct tags from shared contracts so persistence adapters own all
  storage-specific field mapping.
- `customers.RetailCustomer` and `wholesale.WholesaleCustomer` now expose `customer_number`,
  `user_id`, `account_id`, `primary_auth_identity_id`, and `auth_identity_ids` at the top level;
  the former nested `identity` object and nested `basic_info.customer_number` are removed.
- `product.Product` now exposes `catalogue`, `supplier_code`, and `placing_area_code` at the
  top level; the former nested `identifiers` object is removed.
- Flattened the supplier reference on `purchase.Order`: removed the nested `supplier` object
  (the `SupplierSnapshot` type is deleted); a purchase order references its supplier by the flat
  `supplier_code` (required) plus an optional flat `supplier_name`. `purchase.Receipt` was
  already flat. Supplier remains organisation-only (no login persona).

### Additive Changes / 新增

- Added `common.PartyRef.Code` (`code`) as the canonical organisation/supplier business key
  (surfaces on `Supplier`, `WholesaleOrganisation`, and their snapshots/summaries).
- `membership.Reward.Code` is now required (was optional) so `reward_code` is a reliable key.
- OAuth / social sign-in: added the `line`, `discord`, `microsoft` (consumer MSA, distinct from
  enterprise `azureAD`), and `oidc` (catch-all) `AuthIdentityProvider` values, plus an
  `identity.UserConnectedIdentities` projection for the connected-accounts screen. The
  one-user→many-`AuthIdentity` model already supports account linking; the backend implements
  each provider's flow.
- Order snapshot: added `sales.OrderSummary` + `sales.OrderLineSummary` (slim, customer-facing,
  no audit fields), and embedded a bounded `recent_orders` strip on `customers.RetailCustomer`
  and `wholesale.WholesaleCustomer` (hard-capped display projection, never the full history).
- Customer wallet: new `pkg/contracts/wallet` package — `CustomerWallet` (owner-keyed read
  aggregate linking points / gift cards / vouchers / coupons / rewards by business key),
  `WalletInstrument`, `CustomerWalletSummary`, a stored-value `GiftCard` + `GiftCardTransaction`
  balance ledger, and first-class `Voucher`. New enums `WalletInstrumentType`, `GiftCardStatus`,
  `GiftCardTransactionReason`. Order-side gift-card spend is captured via
  `sales.GiftCardRedemptionSnapshot` (+ `Order.gift_card_redemptions`). The existing points-only
  `membership.MembershipWalletSummary` is unchanged (linked, not widened).

- Update the `require` to `/v10` and rewrite imports; there is no `/v9` compatibility.
- Send/store all cross-entity references as the code/number key; backends must populate
  `PartyRef.Code` (supplier/organisation code) and `customer_number` on write so the new
  references resolve.
- Persistence adapters must map storage primary keys and any provider-specific storage field
  names locally; shared contracts expose only database-neutral JSON/domain fields.
- The 3 backends and frontends must be re-synced (separate task): fix imports, adapt
  repositories/handlers/DTOs to the renamed/removed fields, and move persistence lookup/index keys
  off the removed `*_id` to the canonical code/number (notably analytics `sku`→`sku_code` and
  request-context `device_id`→`device_key`).

## v9.4.0 (2026-06-27) - Canonical Code/Number References / 代碼與單號參照標準化

Release date: 2026-06-27

### Additive Changes / 新增

- Product display contracts now expose localized `description []common.LocalizedDescription` and `brand []common.LocalizedName` arrays.
- Product snapshots now expose `sku_code`, localized `description`, localized `brand`, and `supplier`.
- Product, warehouse, purchase, sales, wholesale, payment terminal, promotion, coupon, inbound, and layout contracts gained canonical business reference fields such as `product_sku_code`, `order_number`, `purchase_order_number`, `sales_order_number`, `depot_code`, and `supplier_code`.
- Promotion targeting and promotion resolver inputs now support product SKU code matching through `target_product_sku_code` and `ResolveTarget.ProductSKUCode`.

### Migration Compatibility / 遷移相容

- Legacy product flat `description` / `brand` JSON values decode into localized arrays.
- `identifiers.vendor` and snapshot `vendor` still decode and backfill canonical `supplier` during migration.
- Legacy `product_id`, `order_id`, `depot_id`, and `supplier_id` fields remain present as deprecated `omitempty` aliases where required for dual-read/write migration.
- Promotion resolver falls back to legacy product ID matching only when `target_product_sku_code` is absent.

### Consumer Action / 使用方動作

- Backends should write canonical fields as the primary lookup/index keys and keep legacy aliases only for transitional dual-write/backfill windows.
- API consumers should send product references as `product_sku_code`, order references as `order_number`, depot references as `depot_code`, and supplier references as `supplier_code`.
- Product UIs should render localized `description` and `brand`, and rename product `vendor` copy/fields to `supplier`.
- Before removing legacy fields/indexes, run duplicate/missing-code preflight checks for `sku_code`, `order_number`, `depot_code`, and `supplier_code`.

## v9.3.0 (2026-06-27) - SKU Primary Localized Name / SKU 主要本地化名稱

Release date: 2026-06-27

### Additive Changes / 新增

- `category.SKU` now includes required `primary_name` as a singular `common.LocalizedName` for the canonical display/listing name.
- `category.SKU.other_names` remains available for optional alternate localized names.

### Consumer Action / 使用方動作

- SKU-owning backends must require and persist `primary_name` on create/update paths.
- Frontends should use `primary_name` as the default SKU label and fall back to `other_names` only as alternates.

## v9.0.0 (Unreleased) - Contract Hygiene & Buyer/Commercial Context / 契約整理與買方／商業情境

Release date: Unreleased

### Executive Summary / 摘要

V9 makes the `pkg/contracts/**` files pure data shapes. Every inline enum type that was defined inside a contract file moves to the enum area `pkg/enums`, and every non-struct function/method — the promotion resolver and the derived-property helpers for product, payments, promotion, membership, and campaign — moves to a new `pkg/logic/<domain>` tree. Contract files now contain only structs and their tests. Because exported types change packages, this is a major release; consumers update import qualifiers and call the moved methods as free functions.

V9 讓 `pkg/contracts/**` 內的檔案成為純資料結構。所有原本定義在契約檔內的列舉型別都移至列舉區 `pkg/enums`，所有非 struct 的函式/方法（promotion resolver 以及 product、payments、promotion、membership、campaign 的衍生屬性 helper）都移至新的 `pkg/logic/<domain>`。契約檔現在只包含 struct 與其測試。由於匯出型別更換套件，此為 major 版本；使用方需更新 import 限定詞，並將已遷移的方法改以自由函式呼叫。

V9 also additively introduces a shared buyer/commercial vocabulary so retail, wholesale, and POS sales share one model — without renaming any existing field. POS is treated strictly as an order channel (`enums.OrderType`), never a buyer type: a wholesale organisation buying in the shop is `channel=pos`, `buyer.type=wholesale_organisation`. New `enums.BuyerType`, `enums.PriceAudience`, `enums.PriceVisibility`, and `enums.FulfilmentIntent`, the `sales.BuyerContext`/`sales.PricingContext` structs, an optional `Channel`/`Buyer` on `Cart`, an optional `Buyer` on `Order`, an optional `Pricing` on `CartItem`/`OrderItem`, and a `product.Selling` group describe who is buying and under what commercial pricing context. Every addition is an `omitempty` field (pointer or string), so existing JSON round-trips unchanged.

V9 同時以新增方式導入共用的買方／商業詞彙，讓零售、批發與 POS 銷售共用同一套模型，且不改名任何既有欄位。POS 一律視為訂單通路（`enums.OrderType`），絕非買方類型：批發組織到店購買即為 `channel=pos`、`buyer.type=wholesale_organisation`。新增 `enums.BuyerType`、`enums.PriceAudience`、`enums.PriceVisibility`、`enums.FulfilmentIntent`，以及 `sales.BuyerContext`／`sales.PricingContext`、`Cart` 的選用 `Channel`／`Buyer`、`Order` 的選用 `Buyer`、`CartItem`／`OrderItem` 的選用 `Pricing`，與 `product.Selling` 群組，用以描述買方身分與商業定價情境。所有新增皆為 `omitempty` 欄位（指標或字串），既有 JSON 序列化維持不變。

### Breaking Changes / 破壞性變更

- Relocated inline contract enums to `pkg/enums`: `CampaignPlacement`, `CampaignSeverity`, `CampaignCustomerType`, `CampaignPlatform` (from `campaign`), `PackingSessionStatus`, `PackingDamageHandling` (from `warehouse`), `AlertLevel` (from `analytics`), `AuditOutcome`, `MediaStatus` (from `shared`). The four campaign enums gained a `Campaign` prefix. JSON values are unchanged.
- Relocated the promotion resolver and derived-property helpers out of `pkg/contracts` into `pkg/logic/<domain>`: `promotionlogic` (`ResolveEffective`, `Matches`, `DiscountedUnitPriceMinor`, `ResolveTarget`, `EffectiveClass`, `IncludesDescendants`, `IsTargeted`, `OverrideReasonSpecialCampaign`), `productlogic` (`IsNew`, `NewExpiresAt`, `IsRestocked`, `RestockedExpiresAt`, `DisplayStatus`, `LifecycleTags`, and the `DisplayStatus*`/`LifecycleTag*` constants), `paymentslogic` (`TotalMinor`), `membershiplogic` (`RequiresOrganisationAccess`, `CanReserve`, `CanCommit`, `HasRequiredSpendAccess`, `PermissionMembershipPointsSpend`), `campaignlogic` (`LiveAt`). Former struct methods became free functions taking the struct as the first argument.
- Requires the Go major-version migration: module path `…/Backend-Shared-Contract/v8` → `…/v9`; update all consumer imports accordingly when the tag is cut.

### Additive Changes / 新增

- Added the `pkg/logic/{promotion,product,payments,membership,campaign}` packages.
- Added `IsValid()`/`String()` to the relocated `AlertLevel`, `AuditOutcome`, and `MediaStatus` enums to match the enum-area convention; all relocated enums are covered by `pkg/enums/enums_test.go`.
- `promotion.EffectivePromotion` remains a contract struct, relocated to `pkg/contracts/promotion/effective_promotion.go`.
- Added shared buyer/commercial enums in `pkg/enums`: `BuyerType` (`guest_retail`/`retail_customer`/`wholesale_organisation`), `PriceAudience` (`retail`/`wholesale`), `PriceVisibility` (`public`/`login_required`/`wholesale_approved_only`/`hidden`), and `FulfilmentIntent` (`delivery`/`pickup`/`in_store_carry`); all covered by `pkg/enums/enums_test.go`.
- Added the `sales.BuyerContext` (buyer type, optional retail-customer / wholesale-organisation / organisation-access references, and fulfilment intent) and `sales.PricingContext` (price audience + visibility) structs.
- Added optional `channel` and `buyer` to `sales.Cart`, optional `buyer` to `sales.Order`, and optional `pricing` to `sales.CartItem` and `sales.OrderItem` — all `omitempty` pointers/strings, so legacy JSON is unchanged. `Order.Channel`, `Cart.CustomerID`, and `Product.Pricing` are untouched.
- Added the `product.Selling` group (`channels []enums.OrderType`, `buyer_types []enums.BuyerType`, `visibility enums.PriceVisibility`) for per-product channel/buyer sellability rules.
- POS is modelled as a channel, not a buyer type — e.g. a wholesale organisation buying in store is `channel=pos`, `buyer.type=wholesale_organisation`, item `pricing.audience=wholesale`; a walk-in is `channel=pos`, `buyer.type=guest_retail`.

- 在 `pkg/enums` 新增共用買方／商業列舉：`BuyerType`、`PriceAudience`、`PriceVisibility`、`FulfilmentIntent`，皆納入 `pkg/enums/enums_test.go` 驗證。
- 新增 `sales.BuyerContext`（買方類型、選用的零售客戶／批發組織／組織存取參照與取貨意圖）與 `sales.PricingContext`（定價對象＋可見性）struct。
- `sales.Cart` 新增選用 `channel`／`buyer`，`sales.Order` 新增選用 `buyer`，`sales.CartItem`／`sales.OrderItem` 新增選用 `pricing`（皆為 `omitempty`，既有 JSON 不變）。`Order.Channel`、`Cart.CustomerID`、`Product.Pricing` 維持不變。
- 新增 `product.Selling` 群組，描述每個商品的通路／買方可售規則。
- POS 視為通路而非買方類型：批發組織到店購買為 `channel=pos`、`buyer.type=wholesale_organisation`、品項 `pricing.audience=wholesale`；一般散客為 `channel=pos`、`buyer.type=guest_retail`。

### Consumer Action / 使用方動作

- Update enum import qualifiers, e.g. `campaign.PlacementTopBanner` → `enums.CampaignPlacementTopBanner`, `warehouse.PackingSessionStatusPending` → `enums.PackingSessionStatusPending`, `analytics.AlertLevelOK` → `enums.AlertLevelOK`, `shared.AuditOutcomeSuccess` → `enums.AuditOutcomeSuccess`, `shared.MediaStatusActive` → `enums.MediaStatusActive`.
- Replace moved method calls with the logic-package functions, e.g. `p.DisplayStatus(now)` → `productlogic.DisplayStatus(p, now)`, `amounts.TotalMinor()` → `paymentslogic.TotalMinor(amounts)`, `p.EffectiveClass()` → `promotionlogic.EffectiveClass(p)`, `promotion.ResolveEffective(...)` → `promotionlogic.ResolveEffective(...)`.
- Update the module path from `/v8` to `/v9` when the release is tagged, and re-pin all backend services.
- The buyer/commercial additions are backward compatible: set `cart.channel`/`buyer`, `order.buyer`, item `pricing`, and `product.selling` where relevant; leaving them unset preserves prior behaviour and JSON shape. Model POS as a channel — never add a "pos" buyer type.
- 買方／商業新增為向後相容：可在需要時設定 `cart.channel`／`buyer`、`order.buyer`、品項 `pricing` 與 `product.selling`；未設定則維持既有行為與 JSON 形狀。POS 一律以通路表示，切勿新增「pos」買方類型。

## v8.1.0 - Product Tax Flag / 商品稅務旗標

Release date: 2026-06-24

### Executive Summary / 摘要

V8.1.0 adds a `taxed` boolean to the product contracts so consumers can render GST/FRE tax treatment on invoices and receipts. The change is additive and backward-compatible.

V8.1.0 在商品契約新增 `taxed` 布林欄位，讓使用方能在發票與收據上呈現 GST/FRE 稅務處理。此變更為新增且向後相容。

### Additive Changes / 新增

- Added `taxed` to `product.Product` and `product.Snapshot` for GST/FRE invoice rendering.
- Added product JSON shape coverage for the `taxed` field.

### Consumer Action / 使用方動作

- Run contract serialization/deserialization tests after upgrading.
- Set `taxed` when creating or syncing products; the default `false` preserves prior behaviour.

## v8.0.0 - Global Membership Consolidation / 全域會員整合

Release date: 2026-06-23

### Executive Summary / 摘要

V8 upgrades the module path to `github.com/Potato-Mart/Backend-Shared-Contract/v8` and consolidates membership programme contracts into `pkg/contracts/membership`. The new membership domain owns retail and wholesale-organisation wallets, tiers, point ledgers, point reservations, reward catalog redemptions, point promotions, check-ins, and member subscriptions.

V8 also removes the overloaded wholesale "membership" naming. Wholesale portal access is now `OrganisationAccess`; JSON fields that identify that B2B access grant use `organisation_access_id`. Global membership now means points, tiers, rewards, and recurring member subscriptions.

V8 將 module path 升級為 `github.com/Potato-Mart/Backend-Shared-Contract/v8`，並將會員方案契約整合到 `pkg/contracts/membership`。新的 membership domain 負責零售與 wholesale organisation wallet、tier、點數 ledger、點數 reservation、reward catalog redemption、point promotion、check-in，以及 member subscription。

V8 同時移除 wholesale "membership" 的混淆命名。Wholesale portal access 改名為 `OrganisationAccess`；代表該 B2B access grant 的 JSON 欄位使用 `organisation_access_id`。全域 membership 現在專指點數、tier、reward，以及 recurring member subscription。

### Breaking Changes / 破壞性變更

- Changed Go module path to `github.com/Potato-Mart/Backend-Shared-Contract/v8`; updated metadata to `ModuleVersion = "v8.0.0"`, `MajorVersion = "v8"`.
- Removed old `pkg/contracts/loyalty` and `pkg/contracts/subscription` exported contracts. Use `pkg/contracts/membership`.
- Renamed `wholesale.WholesaleMembership` to `wholesale.OrganisationAccess`.
- Renamed `wholesale.WholesaleMembershipSummary` to `wholesale.OrganisationAccessSummary`.
- Replaced wholesale access JSON key `membership_id` with `organisation_access_id` across wholesale customer, identity session, claims, permissions, and events.
- Replaced loyalty/subscription enum names with membership enum names, including `MembershipPointReason`, `MembershipPromotionTarget`, and `MemberSubscriptionStatus`.
- Replaced `IDENTITY_WHOLESALE_MEMBERSHIP_REQUIRED` with `IDENTITY_ORGANISATION_ACCESS_REQUIRED`.

### Additive Changes / 新增

- Added `membership.MembershipAccount`, `MembershipOwnerRef`, `MembershipWalletSummary`, and `MembershipAccountSummary`.
- Added `membership.MembershipTier` with qualification metric and benefit fields.
- Added `membership.PointLedgerEntry`, `PointAllocation`, `PointBucket`, `PointBalanceBreakdown`, and `PointReservation`.
- Added `membership.Reward` and `RewardRedemption` for catalog rewards.
- Added `membership.SubscriptionPlan` and `MemberSubscription`.
- Added internal membership endpoint constants for point quote/reserve/commit/cancel and reward redeem.
- Added service-auth scopes `membership:read`, `membership:points`, and `membership:redeem`.
- Added membership API error codes for inactive membership, insufficient points, unavailable rewards, and expired point reservations.
- Added sales order snapshots for `point_redemption` and `reward_redemptions`.

### Consumer Action / 使用方動作

- Update imports from `/v7` to `/v8`.
- Migrate `loyalty.LoyaltyLedgerEntry` to `membership.PointLedgerEntry`.
- Migrate `loyalty.LoyaltyTier` to `membership.MembershipTier`.
- Migrate `subscription.CustomerSubscription` to `membership.MemberSubscription`.
- Migrate `wholesale.WholesaleMembership` to `wholesale.OrganisationAccess`.
- Migrate wholesale access fields from `membership_id` to `organisation_access_id`.
- Treat membership wallet balances as projections; use point ledger/reservation contracts as the source of truth.

## v7.0.0 - Product Enterprise Redesign And Contract DTO Cleanup / 商品企業級重新設計與契約 DTO 清理

Release date: 2026-06-23

### Executive Summary / 摘要

V7 upgrades the module path to `github.com/Potato-Mart/Backend-Shared-Contract/v7` and redesigns `product.Product` for international trading: tidy nested groups (`pricing`, `localization`, `media`, `physical`, `merchandising`, `identifiers`) over a small set of flat, indexable identity/filter/sort fields. It adds localized name/description/brand, a six-field pricing block, a proper enterprise status enum plus a computed display status, and a sales-performance indicator. `product.Snapshot` gains fields additively.

V7 also cleans the shared-contract ownership boundary. Exported HTTP/API wire structs and action DTOs were removed from this module and moved into their owning backend services while preserving the intended public JSON shapes. The shared module now keeps domain/value contracts, enums, constants, validation helpers, error codes, and durable event/domain payloads only.

V7 將 module path 升級為 `github.com/Potato-Mart/Backend-Shared-Contract/v7`，並針對國際貿易重新設計 `product.Product`：在少量保持扁平、可索引的識別/篩選/排序欄位之上，採用整潔的巢狀群組（`pricing`、`localization`、`media`、`physical`、`merchandising`、`identifiers`）。新增本地化名稱/描述/品牌、六欄定價、正式的企業級狀態列舉與計算型顯示狀態，以及銷售表現指標。`product.Snapshot` 以 additive 方式新增欄位。

V7 同時清理 shared-contract 的 ownership boundary。HTTP/API wire struct 與 action DTO 已從本模組移除並移至各自擁有的後端服務，同時維持既有公開 JSON shape 的預期相容性。shared module 現在只保留 domain/value contract、列舉、常數、驗證 helper、錯誤碼，以及可持久化的事件/domain payload。

### Breaking Changes / 破壞性變更

- Changed Go module path to `github.com/Potato-Mart/Backend-Shared-Contract/v7`; updated metadata to `ModuleVersion = "v7.0.0"`, `MajorVersion = "v7"`.
- Removed exported shared wire/action structs with DTO-style ownership (`Request`, `Response`, `Input`, `Payload`, `DTO`, or `Dto` suffixes); backend services now own runtime API payloads.
- Removed shared API response envelopes, pagination wire structs, service-auth token request/response structs, payment terminal and settlement command structs, pricing quote/effective-promotion request/response structs, stockops reserve/ref request/response structs, media upload/finalize request/response structs, identity/access command request structs, and the packing stock settlement command line.
- `product.Product`: renamed `code` → `sku_code` (now the barcode payload / unique key); removed the flat `catalogue` (now `identifiers.catalogue`).
- `product.Product`: nested `price`/`pos_price` into `pricing` (`pricing.online` ← old `price`, `pricing.offline` ← old `pos_price`) plus new `original`, `tag`, `wholesale`, `cost`.
- `product.Product`: `brand` (string) → flat `brand_key` + localized `localization.brand_names`; nested `other_names`/`dimensions`/`weight`/`cover_url`/`image_urls`/`vendor`/`placing_area_code` into `localization`/`physical`/`media`/`identifiers`.
- `product.Product`: replaced `freshness_status` (string) with `sales_performance` (`enums.SalesPerformance`: `hot`/`normal`/`slow`).
- `enums.ProductStatus` values changed from `publish`/`draft`/`dismiss` to `draft`/`active`/`archived`/`discontinued`.

- Go module path 改為 `github.com/Potato-Mart/Backend-Shared-Contract/v7`；metadata 更新為 `ModuleVersion = "v7.0.0"`、`MajorVersion = "v7"`。
- 移除帶有 DTO-style ownership 的 exported shared wire/action struct（名稱以 `Request`、`Response`、`Input`、`Payload`、`DTO` 或 `Dto` 結尾）；執行期 API payload 現由各後端服務自行擁有。
- 移除 shared API response envelope、pagination wire struct、service-auth token request/response struct、payment terminal/settlement command struct、pricing quote/effective-promotion request/response struct、stockops reserve/ref request/response struct、media upload/finalize request/response struct、identity/access command request struct，以及 packing stock settlement command line。
- `product.Product`：`code` 改名為 `sku_code`（成為條碼內容/唯一鍵）；移除扁平 `catalogue`（改為 `identifiers.catalogue`）。
- `product.Product`：`price`/`pos_price` 巢狀化為 `pricing`，並新增 `original`、`tag`、`wholesale`、`cost`。
- `product.Product`：`brand` 改為扁平 `brand_key` + 本地化 `localization.brand_names`；其餘屬性巢狀化。
- `product.Product`：`freshness_status` 改為 `sales_performance`（`hot`/`normal`/`slow`）。
- `enums.ProductStatus` 值由 `publish`/`draft`/`dismiss` 改為 `draft`/`active`/`archived`/`discontinued`。

### Additive Changes / 新增（向後相容）

- `product.Snapshot` adds `sku_code` and `display_status` (and keeps the flat `name`). Safe for all embedders (`category.SKU`, `sales.Order`/`Cart`, `purchase.Order`, `subscription.Plan`).
- New `product.Product.DisplayStatus(now)` helper — read-time merge of status + recency + stock (`new`/`restocked`/`out_of_stock`/`active`), never stored.
- New `enums.SalesPerformance`.
- Added a contract-shape guard test that fails if new exported shared-contract types use forbidden DTO-style suffixes (`Request`, `Response`, `Input`, `Payload`, `DTO`, `Dto`).

### Contract Ownership And Backend Migration / 契約歸屬與後端遷移

- **Backend-Operations** now owns local paging, API envelope, media, stockops, packing settlement, effective-promotion, and service-token payloads.
- **Backend-Commerce** now owns local paging, API envelope, pricing wire, stock wire, terminal command, and service-token payloads.
- **Backend-Management** now owns local paging, API envelope, media, pricing, identity/access, and service-token payloads.
- **Frontend-Admin-Web** TypeScript API interfaces remain frontend-local.
- Public HTTP JSON shapes are intended to remain unchanged; this release changes Go package ownership, not endpoint wire names.

- **Backend-Operations** 現在自行擁有 paging、API envelope、media、stockops、packing settlement、effective-promotion 與 service-token payload。
- **Backend-Commerce** 現在自行擁有 paging、API envelope、pricing wire、stock wire、terminal command 與 service-token payload。
- **Backend-Management** 現在自行擁有 paging、API envelope、media、pricing、identity/access 與 service-token payload。
- **Frontend-Admin-Web** 的 TypeScript API interface 維持 frontend-local。
- 公開 HTTP JSON shape 預期維持不變；本次 release 調整的是 Go package ownership，而不是 endpoint wire name。

### Validation / 驗證

- `Backend-Shared-Contract`: `go test ./...`
- `Backend-Operations`: `go test ./...`
- `Backend-Commerce`: `go test ./...`
- `Backend-Management`: `go test ./...`
- `Frontend-Admin-Web`: `npm run test`, `npm run build`

### Consumer Action / 使用方動作

- **Backend-Operations & Frontend-Admin-Web** (this coordinated release): adopt the new field/enum names, the nested pricing, `sku_code`, and the computed `display_status`; migrate the `ops_products` collection (flat→nested, `code`→`sku_code`, status remap) and reconcile indexes.
- **Backend-Commerce & Backend-Management** (pin `v6.0.2`): unaffected until they upgrade. On upgrade they inherit the `ProductStatus` value change (any code validating `publish`/`dismiss` breaks), the additive Snapshot fields (safe), and the `Pricing` nesting (any read of `product.Price`/`POSPrice` breaks → move to `Pricing.Online`/`Pricing.Offline`).
- Consumers importing removed shared DTO/request/response types must switch to backend-local equivalents. Domain/value structs, enums, constants, validation helpers, and error code constants remain shared contract responsibilities.
- Pin the dependency to `v7.0.0` once published.

- **Backend-Operations 與 Frontend-Admin-Web**（本次協同發布）：採用新欄位/列舉名稱、巢狀定價、`sku_code` 與計算型 `display_status`；遷移 `ops_products` 集合並重整索引。
- **Backend-Commerce 與 Backend-Management**（固定 `v6.0.2`）：升級前不受影響；升級時需處理 `ProductStatus` 值變更、additive Snapshot 欄位與 `Pricing` 巢狀化。
- 任何仍 import 已移除 shared DTO/request/response type 的使用方，需改用後端本地 equivalent。Domain/value struct、列舉、常數、驗證 helper 與錯誤碼常數仍由 shared contract 負責。
- 發布後將相依版本固定為 `v7.0.0`。

## v6.0.2 - Version Metadata Correction / 版本中繼資料修正

Release date: 2026-06-18

- Corrected the exported module version metadata to `ModuleVersion = "v6.0.2"` and refreshed README version guidance.
- No contract type, enum, or JSON shape change.

- 修正 exported module version metadata 為 `ModuleVersion = "v6.0.2"`，並更新 README 版本說明。
- 無 contract type、enum 或 JSON shape 變更。

### Consumer Action / 使用方動作

- No code change required. Optionally pin the dependency to `v6.0.2`.
- 不需修改程式碼。可選擇將相依版本固定為 `v6.0.2`。

## v6.0.1 - Product Description Field / 商品描述欄位

Release date: 2026-06-18

- Added a proper `description` field to `product.Product`. Additive and backward-compatible — existing consumers are unaffected.

- 為 `product.Product` 新增正式的 `description` 欄位。屬 additive 且向後相容，既有使用方不受影響。

### Consumer Action / 使用方動作

- No breaking change. Consumers may now read/write `product.description`.
- 無破壞性變更。使用方現在可讀寫 `product.description`。

## v6.0.0 - V6 Shared Contract Model / V6 共用契約模型

Release date: 2026-06-18

### Executive Summary / 摘要

V6 upgrades the shared contract module to `github.com/Potato-Mart/Backend-Shared-Contract/v6` and introduces a cleaner domain model for identity, account personas, portal admission, RBAC, retail customers, wholesale organisations, wholesale memberships, and payment terminal provider metadata.

V6 將共用契約模組升級至 `github.com/Potato-Mart/Backend-Shared-Contract/v6`，並重新整理 identity、account/persona、portal admission、RBAC、retail customer、wholesale organisation、wholesale membership，以及 payment terminal provider metadata 的契約模型。

### Breaking Changes / 破壞性變更

- Changed Go module path to `github.com/Potato-Mart/Backend-Shared-Contract/v6`.
- Updated exported version metadata to `ModuleVersion = "v6.0.0"` and `MajorVersion = "v6"`.
- Removed legacy `common.CompanyDetail`; use `common.OrganisationDetail`.
- Removed legacy `customers.Customer`; use `customers.RetailCustomer`.
- Removed legacy `customers.CompanyCustomer`; use `wholesale.WholesaleCustomer` and related wholesale organisation/membership contracts.
- Split customer contracts into retail/general customer profiles and wholesale-specific profiles.
- Replaced company-prefixed organisation fields with organisation-neutral JSON keys.
- Moved payment provider metadata into nested payment support groups.
- Moved repeated lifecycle action metadata into `common.LifecycleAction` groups where appropriate.
- Moved identity command/request metadata into nested `context` fields for command DTOs.

- Go module path 改為 `github.com/Potato-Mart/Backend-Shared-Contract/v6`。
- exported version metadata 更新為 `ModuleVersion = "v6.0.0"` 與 `MajorVersion = "v6"`。
- 移除舊的 `common.CompanyDetail`，改用 `common.OrganisationDetail`。
- 移除舊的 `customers.Customer`，改用 `customers.RetailCustomer`。
- 移除舊的 `customers.CompanyCustomer`，改用 `wholesale.WholesaleCustomer` 以及相關 wholesale organisation/membership 契約。
- customer contract 拆分為 retail/general customer profile 與 wholesale-specific profile。
- company-prefixed organisation 欄位改為 organisation-neutral JSON key。
- payment provider metadata 改為 nested payment support group。
- 重複的 lifecycle action metadata 改為 `common.LifecycleAction` group。
- identity command/request metadata 改為 command DTO 內的 nested `context` 欄位。

### Common Contracts / 共用契約

- Added `common.OrganisationDetail` as the canonical organisation/company profile for suppliers, wholesale organisations, invoices, fulfilment documents, and future business-party contracts.
- Added `common.IdentityLink` for linking business profiles to canonical user/account/auth identity records.
- Added `common.PersonName` for reusable person-name fields.
- Added `common.ContactChannels` for reusable email, phone, mobile, LINE, and external handle fields.
- Added `common.LifecycleAction` for reusable by/at/reason lifecycle metadata.
- Kept `common.PartyRef`, `common.ContactAddress`, `common.AuditFields`, `common.DataProtectionFields`, `common.Money`, and measurement primitives as shared building blocks.
- Reused `OrganisationDetail` in supplier and wholesale organisation contracts.

- 新增 `common.OrganisationDetail`，作為 supplier、wholesale organisation、invoice、fulfilment document 與未來 business-party contract 的 canonical organisation/company profile。
- 新增 `common.IdentityLink`，用於 business profile 連結 canonical user/account/auth identity。
- 新增 `common.PersonName`，供 person-like profile 重用姓名欄位。
- 新增 `common.ContactChannels`，供 email、phone、mobile、LINE 與 external handles 重用。
- 新增 `common.LifecycleAction`，供 by/at/reason lifecycle metadata 重用。
- 保留並延伸使用 `common.PartyRef`、`common.ContactAddress`、`common.AuditFields`、`common.DataProtectionFields`、`common.Money` 與 measurement primitive。
- supplier 與 wholesale organisation 改用 `OrganisationDetail`。

### Identity And Access / 身分與權限

- Added account/persona contracts: `UserAccount`, `UserAccountSummary`, and account profile extension DTOs.
- Added account and portal enums: `AccountType`, `AccountStatus`, `Portal`, and `PortalAccessStatus`.
- Added auth identity contracts: `AuthIdentity`, `AuthIdentitySummary`, `LinkAuthIdentityRequest`, and `DisableAuthIdentityRequest`.
- Added auth identity enums for provider, domain, status, auth method, and auth assurance level.
- Added portal admission contracts: `PortalAccess`, `PortalAccessDecision`, `ResolvePortalAccessRequest`, `GrantPortalAccessRequest`, and `RevokePortalAccessRequest`.
- Added RBAC assignment contracts: `RoleAssignment`, `GrantRoleAssignmentRequest`, `RevokeRoleAssignmentRequest`, and `EffectivePermissionSet`.
- Added `AccessTokenClaims` as the framework-neutral token claim shape.
- Added identity event payloads for account creation/status changes, portal access grants/revocations, auth identity linking, role assignment grants/revocations, and wholesale membership changes.
- Expanded `LoginSession` with account/persona, portal, audience, roles, permissions, wholesale organisation, membership, role key, MFA, auth assurance, risk, and device context.
- Expanded `UserProfile` with account summaries, primary account information, MFA, notification preferences, user device, login, password, and access-review timestamps.
- Preserved legacy `UserRole` compatibility while documenting that new portal admission must use `AccountType` and `PortalAccess`.

- 新增 account/persona 契約：`UserAccount`、`UserAccountSummary` 與 account profile extension DTO。
- 新增 account 與 portal 列舉：`AccountType`、`AccountStatus`、`Portal`、`PortalAccessStatus`。
- 新增 auth identity 契約：`AuthIdentity`、`AuthIdentitySummary`、`LinkAuthIdentityRequest`、`DisableAuthIdentityRequest`。
- 新增 auth identity provider、domain、status、auth method、auth assurance level 等列舉。
- 新增 portal admission 契約：`PortalAccess`、`PortalAccessDecision`、`ResolvePortalAccessRequest`、`GrantPortalAccessRequest`、`RevokePortalAccessRequest`。
- 新增 RBAC assignment 契約：`RoleAssignment`、`GrantRoleAssignmentRequest`、`RevokeRoleAssignmentRequest`、`EffectivePermissionSet`。
- 新增 framework-neutral 的 `AccessTokenClaims`。
- 新增 identity event payload，涵蓋 account creation/status change、portal access grant/revoke、auth identity linked、role assignment grant/revoke、wholesale membership change。
- `LoginSession` 擴充 account/persona、portal、audience、roles、permissions、wholesale organisation、membership、role key、MFA、auth assurance、risk 與 device context。
- `UserProfile` 擴充 account summary、primary account、MFA、notification preferences、user device、login/password/access-review timestamp。
- 保留 legacy `UserRole` 相容性，但明確要求新的 portal admission 使用 `AccountType` 與 `PortalAccess`。

### Customer And Wholesale / 客戶與批發

- Added grouped retail/general customer contract: `customers.RetailCustomer`.
- Added retail customer projections and grouped profiles for basic info, lifecycle, management, loyalty, marketing, commerce, analytics, and referral data.
- Renamed/reworked customer activity contracts for retail customer usage.
- Added retail JSON shape tests.
- Added wholesale domain package docs and contracts for organisations, memberships, customers, summaries, commercial profile, account profile, and terms.
- Added `WholesaleOrganisation`, using `common.OrganisationDetail`.
- Added `WholesaleMembership` for organisation-scoped user/account access.
- Added `WholesaleCustomer` for wholesale customer contact/profile details.
- Added wholesale status enums for organisation and membership lifecycle.
- Added wholesale JSON shape and lifecycle grouping tests.
- Added migration guide at `docs/v6-customer-contract-migration.md`.

- 新增 grouped retail/general customer contract：`customers.RetailCustomer`。
- 新增 retail customer summary 與 basic info、lifecycle、management、loyalty、marketing、commerce、analytics、referral 等分組 profile。
- customer activity contract 重新命名並調整為 retail customer 使用。
- 新增 retail JSON shape 測試。
- 新增 wholesale domain package docs，以及 organisation、membership、customer、summary、commercial profile、account profile、terms 契約。
- 新增 `WholesaleOrganisation`，並使用 `common.OrganisationDetail`。
- 新增 `WholesaleMembership`，表示 user/account 在 wholesale organisation scope 下的 access。
- 新增 `WholesaleCustomer`，表示 wholesale customer contact/profile。
- 新增 wholesale organisation/membership lifecycle status enum。
- 新增 wholesale JSON shape 與 lifecycle grouping 測試。
- 新增 `docs/v6-customer-contract-migration.md` 遷移指南。

### Payments / 支付與終端機

- Added `payments.TerminalProviderDetails` for provider-side merchant, store, terminal, device, nickname, and base URL metadata.
- Added `payments.ProviderOperationContext` for provider request ID, merchant reference, and idempotency key.
- Added `payments.ProviderPayloads` for raw provider request, response, notification, and display notification payloads.
- Refactored `Terminal`, `RegisterTerminalRequest`, and `TerminalConnectionInfo` to use `ProviderDetails`.
- Refactored `TerminalTransaction` and `Settlement` to use provider details, operation context, and provider payload groups.
- Refactored create request DTOs to use `OperationContext`.
- Kept core contract fields top-level: terminal IDs, provider, connection mode, transaction type, settlement type, status, financial status, requested/result amounts, totals, and timestamps.
- Extended terminal transaction JSON tests for grouped provider metadata and top-level critical field assertions.

- 新增 `payments.TerminalProviderDetails`，集中 provider-side merchant、store、terminal、device、nickname、base URL metadata。
- 新增 `payments.ProviderOperationContext`，集中 provider request ID、merchant reference、idempotency key。
- 新增 `payments.ProviderPayloads`，集中 raw provider request、response、notification、display notification payload。
- `Terminal`、`RegisterTerminalRequest`、`TerminalConnectionInfo` 改用 `ProviderDetails`。
- `TerminalTransaction` 與 `Settlement` 改用 provider details、operation context、provider payload group。
- create request DTO 改用 `OperationContext`。
- core contract fields 保持 top-level：terminal ID、provider、connection mode、transaction type、settlement type、status、financial status、requested/result amount、total、timestamp。
- terminal transaction JSON 測試擴充 provider metadata grouping 與 critical top-level field assertion。

### API, Security, History, And Data Protection / API、安全、歷史與資料保護

- Expanded shared API error codes for identity account type checks, portal admission, revoked/suspended identity access, MFA requirements, wholesale organisation approval, wholesale membership requirements, and terminal outcomes.
- Clarified shared API response envelope usage for internal service-to-service endpoints.
- Preserved and updated shared history, audit, access log, security event, cloud security, media, and data protection contracts for the v6 module path.
- Continued `HistoryEntry` and `HistoryChange` support for high-risk operational timelines.
- Continued `DataProtectionFields` for classification, lawful basis, PII flags, retention, legal hold, and deletion lifecycle.

- 擴充 shared API error code，涵蓋 identity account type check、portal admission、revoked/suspended identity access、MFA requirement、wholesale organisation approval、wholesale membership requirement、terminal outcome。
- 明確化 internal service-to-service endpoint 也使用共用 API response envelope。
- shared history、audit、access log、security event、cloud security、media、data protection contract 更新為 v6 module path。
- 保留 `HistoryEntry` 與 `HistoryChange`，支援高風險流程 timeline。
- 保留 `DataProtectionFields`，支援 classification、lawful basis、PII flag、retention、legal hold、deletion lifecycle。

### Commerce, Catalogue, Fulfilment, And Operations / 商務、商品、履約與營運

- Updated pricing quote and effective promotion contracts for v6 module path compatibility.
- Preserved service-to-service stock operation request/response DTOs and endpoint path constants.
- Updated product, category, snapshot, lifecycle, collection, and category tag contracts for v6 compatibility.
- Preserved promotion and coupon domain controls including discount spec, usage limits, active window, targeting, resolver behaviour, and tests.
- Updated purchase order, receipt, supplier, and supplier snapshot contracts, with supplier now using organisation detail.
- Updated sales order, cart, payment, order history, volume discount, and order JSON tests.
- Updated shipping zone/rate/package/arrival rule contracts.
- Updated subscription plan and customer subscription contracts.
- Updated warehouse damage, depot, draft, inbound, layout, picking, shipment, stock movement contracts and tests.
- Updated analytics forecast and loyalty tier/ledger/check-in contracts and tests.
- Preserved domain-defining fields as top-level fields instead of over-grouping important business concepts.

- pricing quote 與 effective promotion contract 更新為 v6 module path。
- 保留 service-to-service stock operation request/response DTO 與 endpoint path constant。
- product、category、snapshot、lifecycle、collection、category tag contract 更新為 v6 相容。
- 保留 promotion/coupon domain controls，包括 discount spec、usage limits、active window、targeting、resolver behaviour 與測試。
- purchase order、receipt、supplier、supplier snapshot contract 更新；supplier 改用 organisation detail。
- sales order、cart、payment、order history、volume discount 與 order JSON 測試更新。
- shipping zone/rate/package/arrival rule contract 更新。
- subscription plan 與 customer subscription contract 更新。
- warehouse damage、depot、draft、inbound、layout、picking、shipment、stock movement contract 與測試更新。
- analytics forecast 與 loyalty tier/ledger/check-in contract 與測試更新。
- domain-defining fields 保持 top-level，避免過度 grouping 重要 business concept。

### Documentation And Tests / 文件與測試

- Added `docs/identity-access-model.md` for canonical user, login identity, account/persona, portal admission, RBAC, session/claims, customer profile, and wholesale access guidance.
- Added `docs/v6-customer-contract-migration.md` for V5 to V6 customer/organisation migration mapping.
- Added JSON round-trip and JSON shape tests for retail customer, wholesale customer, wholesale organisation, wholesale membership lifecycle, supplier organisation detail, payment provider grouping, marketing contact channels, and identity access/context grouping.
- Updated enum tests for new account, portal, identity, wholesale, and role semantics.
- Verification command: `go test ./...`.

- 新增 `docs/identity-access-model.md`，說明 canonical user、login identity、account/persona、portal admission、RBAC、session/claims、customer profile、wholesale access。
- 新增 `docs/v6-customer-contract-migration.md`，記錄 V5 到 V6 customer/organisation 遷移對照。
- 新增 JSON round-trip 與 JSON shape 測試，涵蓋 retail customer、wholesale customer、wholesale organisation、wholesale membership lifecycle、supplier organisation detail、payment provider grouping、marketing contact channels、identity access/context grouping。
- enum 測試更新，涵蓋 account、portal、identity、wholesale、role 語意。
- 驗證指令：`go test ./...`。

### Consumer Action / 使用方動作

- Update `go.mod` dependency and imports to `github.com/Potato-Mart/Backend-Shared-Contract/v6`.
- Replace `common.CompanyDetail` usage with `common.OrganisationDetail`.
- Replace `customers.Customer` usage with `customers.RetailCustomer`.
- Replace `customers.CompanyCustomer` usage with `wholesale.WholesaleCustomer` plus `WholesaleOrganisation` and `WholesaleMembership` where organisation access is required.
- Update identity consumers to use `AccountType`, `Portal`, `PortalAccess`, and `RoleAssignment` rather than global role checks for portal admission.
- Update payment terminal integrations to read/write `provider_details`, `operation_context`, and `provider_payloads`.
- Update identity command clients to send request metadata under `context`.
- Run consumer-side JSON contract tests before deploying services that rely on customer, identity, wholesale, or payment terminal DTOs.

- 將 `go.mod` dependency 與 imports 更新為 `github.com/Potato-Mart/Backend-Shared-Contract/v6`。
- 將 `common.CompanyDetail` 改為 `common.OrganisationDetail`。
- 將 `customers.Customer` 改為 `customers.RetailCustomer`。
- 將 `customers.CompanyCustomer` 改為 `wholesale.WholesaleCustomer`，並在需要 organisation access 時搭配 `WholesaleOrganisation` 與 `WholesaleMembership`。
- identity 使用方應改用 `AccountType`、`Portal`、`PortalAccess`、`RoleAssignment`，不要再只靠 global role 判斷 portal admission。
- payment terminal integration 應改讀寫 `provider_details`、`operation_context`、`provider_payloads`。
- identity command client 應將 request metadata 放在 `context` 下。
- 依賴 customer、identity、wholesale、payment terminal DTO 的服務，部署前需跑 consumer-side JSON contract tests。

## Remote Release History / 遠端發布歷史

The following sections summarize the released GitHub tag history up to remote `v5.6.0`.

以下內容整理 GitHub 遠端已發布 tags，直到 remote `v5.6.0` 為止。

## v5.6.0 - Contract History, Stock Movement, Loyalty Expiry / 契約歷史、庫存異動與點數到期

Release date: 2026-06-17

- Added `contracts/shared.HistoryEntry` and `HistoryChange` for high-risk process visibility.
- Added `contracts/warehouse.StockMovement` as the shared read model for stock balance changes linked to purchase orders, receipts, sales orders, and damage reports.
- Added loyalty point allocation and expiry bucket summaries, allowing each earned point batch to expire and be redeemed independently.
- Added or updated JSON round-trip tests for history-bearing contracts, including terminal transactions, purchase orders, sales orders, warehouse stock movements, and loyalty ledgers.
- Added history fields across customer, identity device/session, marketing, payment, product, promotion, purchase, sales, subscription, warehouse, and wholesale-facing contracts.

- 新增 `contracts/shared.HistoryEntry` 與 `HistoryChange`，支援高風險流程可視化。
- 新增 `contracts/warehouse.StockMovement` 作為 inventory balance change 的 shared read model，可連結採購單、收貨單、銷售單與損耗紀錄。
- 新增 loyalty point allocation 與 expiry bucket summary，支援每批 earned points 獨立到期與兌換。
- 新增或更新 history-bearing contract 的 JSON round-trip 測試，涵蓋 terminal transaction、purchase order、sales order、warehouse stock movement、loyalty ledger。
- 多個 customer、identity device/session、marketing、payment、product、promotion、purchase、sales、subscription、warehouse、wholesale-facing contract 新增 history 欄位。

## v5.5.2 - Payment Method Follow-up / 支付方式修正

Release date: 2026-06-16

- Continued payment method updates from `v5.5.1`.
- Updated `enums.PaymentMethod` and release metadata.
- Refined README version guidance for the latest v5 patch.

- 延續 `v5.5.1` 的 payment method 更新。
- 更新 `enums.PaymentMethod` 與 release metadata。
- 調整 README 最新 v5 patch 版本資訊。

## v5.5.1 - Payment Method Expansion / 支付方式擴充

Release date: 2026-06-15

- Added additional payment method support used by sales order and payment flows.
- Updated sales order contract payment method references.
- Updated version metadata and README release notes.

- 新增 sales order 與 payment flow 使用的 payment method。
- 更新 sales order contract 內的 payment method references。
- 更新 version metadata 與 README release notes。

## v5.5.0 - Device Tracking Detail / 裝置追蹤細節

Release date: 2026-06-15

- Improved device-tracking detail for identity/session-related contracts.
- Updated customer segment enum support.
- Updated version metadata and README release notes.

- 強化 identity/session 相關 contract 的 device tracking detail。
- 更新 customer segment enum 支援。
- 更新 version metadata 與 README release notes。

## v5.4.0 - Reusable Party Reference / 可重用交易方參照

Release date: 2026-06-15

- Refactored `common.PartyRef` to make shared identity/contact references more reusable.
- Applied shared party references across company customer and supplier-facing contracts.
- Reduced duplicate party id/name/phone/email fields across purchasing and customer models.

- 重構 `common.PartyRef`，提升 shared identity/contact reference 的可重用性。
- company customer 與 supplier-facing contract 改用 shared party reference。
- 減少 purchasing 與 customer model 中重複的 party id/name/phone/email 欄位。

## v5.3.0 - Company Detail And Security Context / 公司細節與安全上下文

Release date: 2026-06-15

- Added or refined `common.CompanyDetail` and `common.PartyRef` usage.
- Enhanced customer and company-customer shared fields.
- Added identity device shared fields and improved user/device projections.
- Added product collection/category updates.
- Added shared access log, audit, cloud security, media, security event, and security policy contract changes.
- Updated purchase and sales contracts to use shared party/detail fields more consistently.

- 新增或調整 `common.CompanyDetail` 與 `common.PartyRef` 使用方式。
- 強化 customer 與 company-customer shared fields。
- 新增 identity device shared fields，並改善 user/device projections。
- 新增 product collection/category 相關更新。
- 新增 shared access log、audit、cloud security、media、security event、security policy contract 變更。
- purchase 與 sales contract 更一致地使用 shared party/detail fields。

## v5.2.0 - Promotion And Category Targeting / 促銷與分類目標

Release date: 2026-06-12

- Added effective promotion lookup contracts under `pkg/contracts/pricing`.
- Added product category tag contract and product lifecycle helpers/tests.
- Expanded promotion rules with class, target scope, product/category targeting, descendant targeting, quantity modes, and pricing tiers.
- Added promotion resolver and resolver tests for deterministic effective-promotion precedence.
- Updated sales order applied promotion support.
- Added promotion class enum and discount type updates.

- 新增 `pkg/contracts/pricing` 下的 effective promotion lookup contracts。
- 新增 product category tag contract 與 product lifecycle helper/tests。
- promotion rule 擴充 class、target scope、product/category targeting、descendant targeting、quantity mode、pricing tiers。
- 新增 promotion resolver 與測試，確保 effective-promotion precedence 一致。
- 更新 sales order applied promotion 支援。
- 新增 promotion class enum 與 discount type 更新。

## v5.1.2 - Product And SKU Fields / 商品與 SKU 欄位

Release date: 2026-06-12

- Added proper product/SKU fields to category contract.
- Updated README and version metadata for the patch release.

- category contract 新增更合適的 product/SKU 欄位。
- 更新 README 與 version metadata。

## v5.1.1 - Embedded Struct Tag Fixes / Embedded Struct 標籤修正

Release date: 2026-06-12

- Corrected embedded struct tags across category, customer, identity, loyalty, marketing, payment, product, promotion, purchase, sales, shared, shipping, subscription, warehouse, and wholesale-related contracts.
- Added integration audit documentation dated 2026-06-12.
- Improved persistence compatibility for consumers using document databases.

- 修正多個 category、customer、identity、loyalty、marketing、payment、product、promotion、purchase、sales、shared、shipping、subscription、warehouse、wholesale-related contract 的 embedded struct tags。
- 新增 2026-06-12 integration audit 文件。
- 改善使用 document database 的消費方持久化相容性。

## v5.1.0 - Service-Authenticated Internal APIs / 服務驗證內部 API

Release date: 2026-06-12

- Added internal stock operation endpoint path constants under `pkg/contracts/stockops`: `PathReserve`, `PathCommit`, and `PathRelease`.
- Added service token endpoint path constant `PathToken` under `pkg/serviceauth`.
- Added cross-service pricing quote contracts `QuoteRequest` and `QuoteResponse` with `PathQuote`.
- Clarified that internal service-to-service endpoints use the same `apiresponse.APIResponse` envelope.

- `pkg/contracts/stockops` 新增內部庫存端點路徑常數：`PathReserve`、`PathCommit`、`PathRelease`。
- `pkg/serviceauth` 新增服務權杖端點路徑常數 `PathToken`。
- 新增跨服務 pricing quote 契約 `QuoteRequest`、`QuoteResponse` 與 `PathQuote`。
- 明確化 internal service-to-service endpoint 同樣使用 `apiresponse.APIResponse` 回應信封。

## v5.0.0 - V5 Model Reroute / V5 模型重整

Release date: 2026-06-11

- Upgraded module path and imports to the v5 contract line.
- Rerouted and cleaned contract models across API responses, common fields, analytics, category, customers, identity, loyalty, marketing, payments, pricing, product, promotion, purchase, sales, shared, shipping, stock operations, subscription, warehouse, and wholesale packages.
- Added or refined service-to-service pricing/stock contract support.
- Expanded API errors for newer identity, discount, inventory, and terminal scenarios.
- Updated README and version metadata to the v5 release line.

- module path 與 imports 升級至 v5 contract line。
- API response、common fields、analytics、category、customers、identity、loyalty、marketing、payments、pricing、product、promotion、purchase、sales、shared、shipping、stock operations、subscription、warehouse、wholesale package 進行模型重整。
- 新增或調整 service-to-service pricing/stock contract 支援。
- 擴充 identity、discount、inventory、terminal 等情境的 API error。
- README 與 version metadata 更新至 v5 release line。

## v4.2.0 - Version Metadata Bump / 版本中繼資料更新

Release date: 2026-06-11

- Confirmed the v4.2.0 release number in README and version metadata.
- No broad contract model change was identified in the remote tag diff.

- README 與 version metadata 確認 v4.2.0 release number。
- 遠端 tag diff 未顯示大型 contract model 變更。

## v4.1.0 - Field Consolidation And Shared Common Models / 欄位整併與共用模型

Release date: 2026-06-11

- Refactored and split repeated fields across customer, identity, loyalty, marketing, payment, product, promotion, purchase, sales, shared, shipping, subscription, warehouse, and wholesale-facing contracts.
- Added or refined `common.ContactAddress`, `common.DataProtectionFields`, `common.Date`, and `common.PartyRef`.
- Moved customer profile concerns into shared/customer-specific files.
- Updated module metadata and API response envelope usage.

- 重構並拆分 customer、identity、loyalty、marketing、payment、product、promotion、purchase、sales、shared、shipping、subscription、warehouse、wholesale-facing contract 中的重複欄位。
- 新增或調整 `common.ContactAddress`、`common.DataProtectionFields`、`common.Date`、`common.PartyRef`。
- customer profile concerns 移至 shared/customer-specific files。
- 更新 module metadata 與 API response envelope 使用方式。

## v4.0.0 - Generalized Payment Interfaces And Notifications / 泛化支付介面與通知

Release date: 2026-06-11

- Generalized payment-facing interfaces so the backend can cooperate with different third-party providers.
- Added user notification preference contracts.
- Added payment reference contracts and expanded terminal/settlement/transaction amount models.
- Added terminal connection mode, refund type, provider, status, and transaction status enums.
- Added stock operation contract package.
- Updated sales payment/order contracts to align with generalized payment flows.

- 支付接口泛化，可與不同第三方 payment provider 合作。
- 新增 user notification preference contracts。
- 新增 payment reference contracts，並擴充 terminal/settlement/transaction amount models。
- 新增 terminal connection mode、refund type、provider、status、transaction status enums。
- 新增 stock operation contract package。
- sales payment/order contract 對齊 generalized payment flows。

## v3.10.0 - Customer Active Field Removal / 移除 Customer Active 欄位

Release date: 2026-06-05

- Removed the customer `is_active` field after status-based lifecycle handling was introduced.
- Updated README and version metadata.

- 在 status-based lifecycle 導入後移除 customer `is_active` 欄位。
- 更新 README 與 version metadata。

## v3.9.0 - Customer Status Migration / Customer 狀態遷移

Release date: 2026-06-05

- Replaced customer active-state handling with a status field.
- Updated customer contract and release metadata.

- 以 status 欄位取代 customer active-state handling。
- 更新 customer contract 與 release metadata。

## v3.8.0 - Customer Profile Status / Customer Profile 狀態

Release date: 2026-06-04

- Added customer profile status enum support.
- Updated README and version metadata.

- 新增 customer profile status enum 支援。
- 更新 README 與 version metadata。

## v3.7.0 - Customer Record Expansion / Customer 記錄擴充

Release date: 2026-06-04

- Added additional fields to the customer struct.
- Updated customer type enum and release metadata.

- customer struct 新增更多記錄欄位。
- 更新 customer type enum 與 release metadata。

## v3.6.0 - Warehouse Damage Module / 倉儲損耗模組

Release date: 2026-06-02

- Added warehouse damage report contract.
- Added damage stage enum.
- Updated warehouse packing enum support.
- Updated README and version metadata.

- 新增 warehouse damage report contract。
- 新增 damage stage enum。
- 更新 warehouse packing enum 支援。
- 更新 README 與 version metadata。

## v3.5.1 - Customer Type Rename / Customer Type 命名調整

Release date: 2026-06-01

- Changed customer type wording from `COMPANY` to `WHOLESALER`.
- Updated README and version metadata.

- customer type 命名由 `COMPANY` 調整為 `WHOLESALER`。
- 更新 README 與 version metadata。

## v3.5.0 - ISO27001-Aligned Contract Fields / ISO27001 對齊欄位

Release date: 2026-06-01

- Added or expanded audit and data protection contracts.
- Added identity device/session/user security fields.
- Added shared audit, access log, cloud security, security event, security policy, media, and media upload contracts.
- Expanded payment terminal constants and payment amount/settlement/transaction contracts.
- Added payment terminal readiness documentation.
- Updated README and version metadata.

- 新增或擴充 audit 與 data protection contracts。
- identity device/session/user 新增 security fields。
- 新增 shared audit、access log、cloud security、security event、security policy、media、media upload contracts。
- 擴充 payment terminal constants 與 payment amount/settlement/transaction contracts。
- 新增 payment terminal readiness 文件。
- 更新 README 與 version metadata。

## v3.3.0 - Identity Role And Shared Media Updates / Identity Role 與 Shared Media 更新

Release date: 2026-05-18

- Added new user role support.
- Updated identity role and session contracts.
- Updated shared audit, media, and media upload contracts.
- Updated version metadata.

- 新增 user role 支援。
- 更新 identity role 與 session contracts。
- 更新 shared audit、media、media upload contracts。
- 更新 version metadata。

## v3.2.0 - MX51 Payment Alignment / MX51 支付對齊

Release date: 2026-05-09

- Added MX51 integration readiness documentation.
- Added payment action framework, amount, settlement, terminal, and terminal transaction contracts.
- Updated sales payment/order contracts for terminal payment flows.
- Added payment, terminal provider/status/type, settlement type, and recovery decision enums.
- Expanded API errors for terminal payment outcomes.

- 新增 MX51 integration readiness 文件。
- 新增 payment action framework、amount、settlement、terminal、terminal transaction contracts。
- sales payment/order contracts 對齊 terminal payment flows。
- 新增 payment、terminal provider/status/type、settlement type、recovery decision enums。
- 擴充 terminal payment outcome 相關 API errors。

## v3.1.0 - Warehouse 3D Layout / 倉儲 3D Layout

Release date: 2026-05-09

- Added geometry primitives under `pkg/common`.
- Added warehouse layout contracts and depot layout references.
- Added warehouse layout enum support.
- Updated version metadata.

- `pkg/common` 新增 geometry primitives。
- 新增 warehouse layout contracts 與 depot layout references。
- 新增 warehouse layout enum 支援。
- 更新 version metadata。

## v3.0.0 - V3 Module Path And Release Automation / V3 Module Path 與發布自動化

Release date: 2026-05-02

- Upgraded module path and imports to the v3 contract line.
- Added release workflow automation under GitHub Actions.
- Added release note scripts and versioning documentation.
- Added date, localized name, measurement, analytics forecast, customer activity/profile, loyalty, marketing, payment, promotion, stock operation, subscription, and warehouse contract changes.
- Established automated release content generation from this major version onward.

- module path 與 imports 升級至 v3 contract line。
- GitHub Actions 新增 release workflow automation。
- 新增 release note scripts 與 versioning documentation。
- 新增 date、localized name、measurement、analytics forecast、customer activity/profile、loyalty、marketing、payment、promotion、stock operation、subscription、warehouse contract changes。
- 自此 major version 起建立 automated release content generation。

## v2.1.1 - Product Freshness String / 商品新鮮度字串化

Release date: 2026-04-27

- Changed product freshness representation to a string-backed field.
- Updated product freshness enum usage and version metadata.

- product freshness 改為 string-backed field。
- 更新 product freshness enum 使用方式與 version metadata。

## v2.1.0 - Product Freshness Field / 商品新鮮度欄位

Release date: 2026-04-27

- Added a new product freshness/status field.
- Added product freshness status enum support.
- Updated version metadata.

- 新增 product freshness/status 欄位。
- 新增 product freshness status enum 支援。
- 更新 version metadata。

## v2.0.5 - Product Expiry Field / 商品到期欄位

Release date: 2026-04-27

- Added product expiry support.
- Updated version metadata.

- 新增 product expiry 支援。
- 更新 version metadata。

## v2.0.4 - Product JSON Naming / 商品 JSON 命名

Release date: 2026-04-27

- Updated JSON naming in product-related contract fields.
- Updated purchase order, receipt, cart, and sales order JSON field alignment.
- Updated version metadata.

- 更新 product-related contract fields 的 JSON 命名。
- purchase order、receipt、cart、sales order JSON fields 對齊。
- 更新 version metadata。

## v2.0.3 - Product Code Retention Follow-up / 商品 Code 保留後續修正

Release date: 2026-04-27

- Follow-up tag for product code retention changes.
- Shares the same remote commit as `v2.0.2`.

- product code retention 變更的後續 tag。
- 與 `v2.0.2` 指向相同遠端 commit。

## v2.0.2 - Product Code Retention / 商品 Code 保留

Release date: 2026-04-27

- Kept product code support in the product contract.
- Updated version metadata.

- product contract 保留 product code 支援。
- 更新 version metadata。

## v2.0.1 - V2 Module Path Correction / V2 Module Path 修正

Release date: 2026-04-26

- Corrected the module path to the v2 contract line.
- Shares the same remote commit as `v2.0.0`.

- 修正 module path 至 v2 contract line。
- 與 `v2.0.0` 指向相同遠端 commit。

## v2.0.0 - V2 Module Path Migration / V2 Module Path 遷移

Release date: 2026-04-26

- Upgraded module path and imports to the v2 contract line.
- Updated README and package imports across category, customer, identity, product, promotion, purchase, sales, shipping, warehouse, and versioning contracts.
- Continued broad contract/package restructuring from the v1 line.

- module path 與 imports 升級至 v2 contract line。
- category、customer、identity、product、promotion、purchase、sales、shipping、warehouse、versioning contract 更新 README 與 package imports。
- 延續 v1 line 的大範圍 contract/package restructuring。

## v1.3.0 - Product And Placing Area / 商品與陳列區域

Release date: 2026-04-25

- Updated product contract fields.
- Added warehouse placing area contract.
- Added customer tier enum support.
- Updated sales status enum and version metadata.

- 更新 product contract fields。
- 新增 warehouse placing area contract。
- 新增 customer tier enum 支援。
- 更新 sales status enum 與 version metadata。

## v1.2.0 - Model Structure Refinement / 模型結構調整

Release date: 2026-04-25

- Refined SKU, purchase order, supplier, and sales order model structure.
- Updated payment method and sales status enums.

- 調整 SKU、purchase order、supplier、sales order model structure。
- 更新 payment method 與 sales status enums。

## v1.1.0 - Initial Complete Contracts / 初始完整契約

Release date: 2026-04-24

- Added initial complete contract/model set after repository bootstrap.
- Added API response and error contracts.
- Added common address, audit, metadata, money, pagination, and recipient models.
- Added category, customer, identity, product, promotion, purchase, sales, shipping, warehouse, enum, versioning, and documentation files.

- repository bootstrap 後新增初始完整 contract/model set。
- 新增 API response 與 error contracts。
- 新增 common address、audit、metadata、money、pagination、recipient models。
- 新增 category、customer、identity、product、promotion、purchase、sales、shipping、warehouse、enum、versioning、documentation files。

## v1.0.0 - Initial Module Baseline / 初始模組基線

Release date: 2026-04-21

- Established the repository baseline.
- Shares the same remote commit as `v0.1.0`.

- 建立 repository baseline。
- 與 `v0.1.0` 指向相同遠端 commit。

## v0.1.0 - Repository Seed / Repository 初始建立

Release date: 2026-04-21

- Added initial `.gitignore` and README.
- Created the first public repository seed for the shared contract module.

- 新增初始 `.gitignore` 與 README。
- 建立 shared contract module 的第一個 public repository seed。
