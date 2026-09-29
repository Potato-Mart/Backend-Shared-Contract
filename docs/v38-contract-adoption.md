# v38.0.0 contract adoption

Import `github.com/Potato-Mart/Backend-Shared-Contract/v38` at released tag `v38.0.0`.
This is a breaking, model-only release. No compatibility aliases or database migrations are supplied.
All eight backend services must align their module imports, `go.mod` pin and CI `contract_version`; owning services then regenerate their OpenAPI and client projections. Backend and Frontend sessions own that work.

## Geographic scope and authority

`GeographicScope` and `StaffGeoScope` use `mode` (`GLOBAL` or `TARGETED`) and optional `targets`. GLOBAL has no targets. TARGETED requires at least one explicit independent parent path. The UI may describe multiple selections; the wire enum remains `TARGETED`.

```json
{"mode":"TARGETED","targets":[{"country_code":"AU"},{"country_code":"AU","market_code":"AU-RETAIL","state_code":"AU-NSW"},{"country_code":"NZ","market_code":"NZ-RETAIL","state_code":"NZ-AUK","depot_code":"NZ-AUK-01"}]}
```

Country is required in every path. Market, official state/subdivision and depot are optional suffixes: a deeper value requires all parents. Services validate actual parent relationships. Selections do not create a Cartesian product. `state_code` uses the existing `SubdivisionCode` value type; operational `DepotRegion` is separate. Depth vocabulary is COUNTRY/MARKET/STATE/DEPOT; access scope levels additionally include `state`.

`GeographicContext` carries optional `path` and `matched_path`, plus existing `source`, authoritative owning `market_code`, scope/rule revisions and evaluation timezone. Flat country/subdivision/depot-region/depot and matched target kind/code fields are removed from that context. Resource owning market/currency and target paths are separate; preserve existing authoritative ownership fields and validate their consistency.

Staff path selection intersects role-derived grants. GLOBAL does not elevate authority; global super-admin behavior remains role-derived. Shared models contain no authorization implementation.

## Packages and barcodes

```json
{"package_option":{"sku_code":"A00001","code":"CASE12","version":2}}
```

`PackageOptionRef` has three required fields: string `sku_code`, string `code`, positive int64 `version`. Package identity is the complete tuple; a code or version alone is insufficient. `EACH` means one base unit; `CASE{units}` encodes the integer case size. Version is immutable physical specification identity, scoped to SKU+code. A physical-spec change creates a future version; stock/price/approval revision is separate. SKU fields retained on enclosing records must agree with the reference.

Former `package_option_code` becomes `package_option`; source/destination/requested-case/replacement-each variants keep their semantic prefix and drop `_code`. Previously optional fields remain optional pointer objects. Promotion scope `package_option_codes` becomes `package_options`, a list of full references. `ProductPackageOption` and `SellingProductPackageOption` expose required `sku_code`, `code`, `version`. Stock, reservation, receipt, procurement, price, promotion allocation, picking, fulfilment and order snapshots retain full references, including package-composition components and analytics facts.

Product barcodes use only `CODE_128`. Values remain strings, e.g. `"000012345678"`; never convert to numbers. Remove stale `display_selling_count` from consumer DTOs, fixtures and forms; the canonical product contract already excludes it.

Existing price channels, package types and approval rules remain service-owned. A 90% POS price suggestion is an explicit staff-save action in the online edit modal. Saved normal POS prices are independent, with no automatic link. Checkout custom-price overrides retain a reason. Backend must apply eligible membership, quantity and bundle promotions equally online/POS for the same customer and basket.

## Membership tiers

```json
{"label":[{"language":"en","text":"Gold"},{"language":"zh-TW","text":"金卡"},{"language":"zh-CN","text":"金卡"}],"metadata":[{"key":"campaign","value":"spring"}],"tier_card":{"code":"MEDIA-TIER-GOLD"}}
```

The localized list is extensible; services require en, zh-TW and zh-CN. `metadata` is optional inert string key/value data. `tier_card` is optional `ObjectMediaRef`, containing only `code`; public/current/next tier projections share that shape. Customers owns the existing `membership-tier-card` managed media folder and upload/lifecycle; Pricing owns the tier association and compare-and-swap behavior. No new infrastructure is required by this model.

Remove tier `is_active`, `discount_percent`, `free_shipping_threshold`, `birthday_bonus_points` and those three benefit-kind enum values. The existing authoritative minimum-spend ladder must have unique thresholds with the same metric/currency. Services derive one-based read rank and ladder revision; this release adds no writable/persisted rank.

## Coupon policy data

```json
{"distribution":"SELF_CLAIM_BY_CODE","visibility":"UNLISTED","receiving_tier":{"specification":"ALL_TIERS"},"tier_restriction":{"active":false},"profile_completion_restriction":{"active":false}}
```

Required policy fields are `distribution`, `visibility`, `receiving_tier`, `tier_restriction`, `profile_completion_restriction`. Distribution accepts SELF_CLAIM_BY_CODE, ASSIGNED, AUTO_MEMBER_CLAIM. Visibility accepts PUBLIC or UNLISTED. Receiving specification accepts ALL_TIERS or SELECTED_TIER; selected policies carry `tier_key` and `match` EXACT/AT_OR_ABOVE. Active redemption restrictions independently carry `tier_key` and `match`; inactive restrictions do not constrain redemption. Services normalize default receiving ALL_TIERS and visibility UNLISTED; Go zero values are not valid defaults.

Receiving eligibility and redemption current-tier checks are separate. Profile completion requires 100% when active. Entitlement is permanently once per customer and immutable coupon ID, never reset by use, refund, revoke or edit. Assigned issuance snapshots eligible recipients at activation; automatic issuance scans existing/newly eligible customers within 30 minutes, idempotently. Guests browse but authenticate to claim. Only available active public self-claim coupons enter discovery; unlisted coupons can be claimed by an exact known code without leaking counts or autocomplete results. Pricing owns all validation, scheduling, entitlement and enforcement.

## Gift cards and purchase completion

GiftCard `code` is the single customer-facing claim/connected-POS identifier: GC followed by exactly 12 decimal digits. Internal IDs remain separate. Optional `customer_number` records customer binding and must agree with the authoritative owner; the service binds a POS redemption to the order customer. Owned online use needs no PIN.

PINs are system-generated four-digit strings, preserving leading zeroes. Plaintext PIN, verifier material and delivery credentials stay in protected service-owned data and privileged DTOs; they are absent from canonical/public/event models.

Orders generates the GC gift-purchase order prefix, expires unpaid purchases after 15 minutes using `SalesOrderStatusExpired` (`expired`), disallows customer/staff cancellation and permits paid refund only. Fulfilment completes only after issuance AND successful recipient delivery, independently of claim.

```json
{"issuance_id":"ISSUANCE-01","delivered_at":"2026-09-29T00:00:00Z"}
```

`pubsub/notification.GiftCardDeliveredEvent` and `notification.gift_card_delivered` v1 carry only that safe correlation/time. Notification produces the completed delivery fact; Orders correlates it to the purchase. Backend/DevOps must establish authenticated transport/routing and durable dedupe before use; this release does not deploy a topic/subscription or claim operational readiness. Existing `wallet.gift_card_issued` v1 remains unchanged and carries no delivery material.

## Event compatibility and rollout

The [v38 release note](release-notes.md#v3800-2026-09-29) lists every changed event version. Package-composition changes are transitive through order/refund analytical item facts and packing projections. Other registered event versions remain unchanged. Deploy compatible consumers before new-version producers; fail closed on missing required references, rather than inferring version 1 or upgrading historical snapshots.

Release Shared-Contract first, then Backend service contracts/logic and checks, then Frontend/POS consumers. In this approved empty-database workflow there is no migration, backfill, reset, deletion or legacy conversion. Consumer compile/API/event/behavior checks remain with those sessions; a passing Shared-Contract suite proves model consistency only.
