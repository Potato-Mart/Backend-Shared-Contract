# Version Log / 版本紀錄

Backend-Shared-Contract is the shared contract layer for the Potato Mart backend ecosystem. This file records public contract changes, migration impact, and consumer actions for backend services, web clients, mobile clients, and future service-to-service integrations.

Backend-Shared-Contract 是土豆商城後端生態系的共用契約層。本文件記錄公開契約變更、遷移影響，以及後端服務、前端、行動端與未來服務間整合需要採取的動作。

## Governance / 治理原則

- This module contains shared enums, DTOs, response envelopes, error codes, event payloads, and cross-service request/response contracts only.
- It must not depend on database drivers, web frameworks, authentication middleware, or runtime service implementations.
- Semantic versioning is enforced. Any removal, rename, JSON shape change, module path change, or incompatible exported type change requires a major version.
- Consumers should pin a released module tag and review the "Consumer Action / 使用方動作" section before upgrading.
- Remote release history was reconciled from GitHub tags in `Potato-Mart/Backend-Shared-Contract` on 2026-06-18.

- 本模組只包含共用列舉、DTO、回應信封、錯誤碼、事件負載，以及跨服務 request/response 契約。
- 本模組不得依賴資料庫驅動、Web 框架、身份驗證 middleware，或任何服務執行期實作。
- 本模組遵循 semantic versioning。任何移除、改名、JSON shape 改變、module path 改變，或不相容的 exported type 變更，都必須升 major version。
- 使用方應固定依賴已發布 tag，並在升級前閱讀 "Consumer Action / 使用方動作"。

## Release Index / 發布索引

| Version | Release date | Type | Impact |
| --- | --- | --- | --- |
| `v6.0.0` | 2026-06-18 | Major | Staged breaking release: V6 module path, identity/access model, retail/wholesale split, grouped support fields |
| `v5.6.0` | 2026-06-17 | Minor | Contract history, stock movement, loyalty expiry models |
| `v5.5.2` | 2026-06-16 | Patch | Payment method correction/extension |
| `v5.5.1` | 2026-06-15 | Patch | Payment method extension in sales contracts |
| `v5.5.0` | 2026-06-15 | Minor | Device tracking and customer segment refinement |
| `v5.4.0` | 2026-06-15 | Minor | Common party reference reuse for company/customer/supplier models |
| `v5.3.0` | 2026-06-15 | Minor | Company/customer shared detail, device detection, collections, security logs |
| `v5.2.0` | 2026-06-12 | Minor | Promotions, category tags, product lifecycle, effective promotion resolver |
| `v5.1.2` | 2026-06-12 | Patch | Product/SKU field refinement |
| `v5.1.1` | 2026-06-12 | Patch | BSON inline tag corrections and integration audit docs |
| `v5.1.0` | 2026-06-12 | Minor | Service-authenticated stock/pricing endpoints and API envelope clarification |
| `v5.0.0` | 2026-06-11 | Major | V5 module path, contract reroute, performance-oriented model cleanup |
| `v4.2.0` | 2026-06-11 | Patch | Version metadata bump |
| `v4.1.0` | 2026-06-11 | Minor | Field grouping, common contact/address/date/party references |
| `v4.0.0` | 2026-06-11 | Major | Generalized payment interfaces and user notification preferences |
| `v3.10.0` | 2026-06-05 | Minor | Customer `is_active` removal after status migration |
| `v3.9.0` | 2026-06-05 | Minor | Customer active flag replaced by status |
| `v3.8.0` | 2026-06-04 | Minor | Customer profile status enum added |
| `v3.7.0` | 2026-06-04 | Minor | Customer record field expansion |
| `v3.6.0` | 2026-06-02 | Minor | Warehouse damage report module |
| `v3.5.1` | 2026-06-01 | Patch | Customer type wording changed from company to wholesaler |
| `v3.5.0` | 2026-06-01 | Minor | ISO27001-aligned audit, security, media, data protection fields |
| `v3.3.0` | 2026-05-18 | Minor | Additional identity roles and shared media/security fields |
| `v3.2.0` | 2026-05-09 | Minor | MX51 payment terminal alignment |
| `v3.1.0` | 2026-05-09 | Minor | Warehouse 3D geometry and layout contracts |
| `v3.0.0` | 2026-05-02 | Major | V3 module path and automated release workflow |
| `v2.1.1` | 2026-04-27 | Patch | Product freshness represented as string |
| `v2.1.0` | 2026-04-27 | Minor | Product freshness/status field expansion |
| `v2.0.5` | 2026-04-27 | Patch | Product expiry field |
| `v2.0.4` | 2026-04-27 | Patch | Product/order JSON naming refinements |
| `v2.0.3` | 2026-04-27 | Patch | Product code retention follow-up |
| `v2.0.2` | 2026-04-27 | Patch | Product code retention |
| `v2.0.1` | 2026-04-26 | Patch | V2 module path correction |
| `v2.0.0` | 2026-04-26 | Major | V2 module path and package import migration |
| `v1.3.0` | 2026-04-25 | Minor | Product and placing area contracts |
| `v1.2.0` | 2026-04-25 | Minor | Purchase, supplier, sales, SKU, payment status model changes |
| `v1.1.0` | 2026-04-24 | Minor | Initial complete contract/model set |
| `v1.0.0` | 2026-04-21 | Major | Initial module baseline |
| `v0.1.0` | 2026-04-21 | Pre-release | Initial repository seed |

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

## v5.1.1 - BSON Inline Tag Fixes / BSON Inline 標籤修正

Release date: 2026-06-12

- Corrected BSON inline tags for embedded structs across category, customer, identity, loyalty, marketing, payment, product, promotion, purchase, sales, shared, shipping, subscription, warehouse, and wholesale-related contracts.
- Added integration audit documentation dated 2026-06-12.
- Improved persistence compatibility for consumers using document databases.

- 修正多個 category、customer、identity、loyalty、marketing、payment、product、promotion、purchase、sales、shared、shipping、subscription、warehouse、wholesale-related contract 的 embedded struct BSON inline tags。
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
- Added Adyen terminal readiness documentation and retained MX51 readiness documentation.
- Updated README and version metadata.

- 新增或擴充 audit 與 data protection contracts。
- identity device/session/user 新增 security fields。
- 新增 shared audit、access log、cloud security、security event、security policy、media、media upload contracts。
- 擴充 payment terminal constants 與 payment amount/settlement/transaction contracts。
- 新增 Adyen terminal readiness 文件並保留 MX51 readiness 文件。
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
