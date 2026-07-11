# Version Log / 版本紀錄

Backend-Shared-Contract is the shared contract layer for the Potato Mart backend ecosystem. This file records public contract changes, migration impact, and consumer actions for backend services, web clients, mobile clients, and future service-to-service integrations.

Backend-Shared-Contract 是土豆商城後端生態系的共用契約層。本文件記錄公開契約變更、遷移影響，以及後端服務、前端、行動端與未來服務間整合需要採取的動作。

## Governance / 治理原則

- This module contains reusable domain contracts, value structs, enums, constants, validation helpers, error codes, and durable event/domain payloads only.
- HTTP/API wire DTOs, response envelopes, command payloads, and backend-specific request/response structs belong in the owning backend service.
- It must not depend on database drivers, web frameworks, authentication middleware, or runtime service implementations.
- Semantic versioning is enforced. Any removal, rename, JSON shape change, module path change, or incompatible exported type change requires a major version.
- Consumers should pin a released module tag and review the "Consumer Action / 使用方動作" section before upgrading.
- Remote release history was reconciled from GitHub tags in `Potato-Mart/Backend-Shared-Contract` on 2026-06-18.

- 本模組只包含可重用的 domain contract、value struct、列舉、常數、驗證 helper、錯誤碼，以及可持久化的事件/domain payload。
- HTTP/API wire DTO、回應信封、command payload，以及後端專屬 request/response struct 應由各自擁有的後端服務維護。
- 本模組不得依賴資料庫驅動、Web 框架、身份驗證 middleware，或任何服務執行期實作。
- 本模組遵循 semantic versioning。任何移除、改名、JSON shape 改變、module path 改變，或不相容的 exported type 變更，都必須升 major version。
- 使用方應固定依賴已發布 tag，並在升級前閱讀 "Consumer Action / 使用方動作"。

## Release Index / 發布索引

| Version | Release date | Type | Impact |
| --- |--------------| --- | --- |
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
- Contract tests can be run with `scripts/Test-Contract.ps1` or
  `scripts/test-contract.sh`, both using `GOWORK=off go test ./...` so the
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
