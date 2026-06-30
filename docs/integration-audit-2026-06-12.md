# Integration Consistency Audit & Fix — 2026-06-12

Scope: Backend-Shared-Contract (source of truth) → Backend-Management / Backend-Operations / Backend-Commerce → Frontend-Admin-Web.

## A. Summary

**Checked:** contract module (go.mod, all 40+ enum files, common types, contracts/*, serviceauth paths/scopes) · all three backends' go.mod pins, contract imports, full Gin route tables, status literals, error envelopes, internal S2S routes · all 27 frontend API client files (195 calls), enum mirrors, pages/buttons, the wiring contract test.

**Changed:** 7 frontend files (enum drift + stale page literals + i18n), 1 backend file (guard comment). No public API names changed; no endpoints added/removed.

**Could not verify:** build/typecheck/test/vitest/go-test runs — the local sandbox shell is down ("no space left on device" at the host level). All findings are from direct source comparison.

**Needs manual review:** orphan backend endpoints (below), the Operations PO-transition timestamp switch, free-form product `status` string.

## B. Contract source of truth

- Module `github.com/Potato-Mart/Backend-Shared-Contract/v5`, working tree at **v5.1.0**; all three backends pin exactly `v5.1.0`, no `replace`, no vendoring → backends and contract working tree are the same version. No generation tooling exists (plain Go types; no ABI/TypeChain — N/A for this stack).
- S2S paths are contract constants: `serviceauth.PathToken=/v1/internal/token`, `pricing.PathQuote=/v1/internal/pricing/quote`, `stockops.PathReserve|Commit|Release`. Management and Operations register routes **derived from these constants**, Commerce calls them via the same constants → consistent by construction.
- Key wire-value facts that drove fixes: `SalesOrderStatus` includes `paid`, `packed`, `completed` (no plain `packing`); `FulfillmentStatus` = unfulfilled, picking_printed, packing, packed, partial, fulfilled; `PaymentStatus` adds unknown/pending/partially_refunded; `PaymentRecordStatus` adds processing/cancelled/awaiting_action/unknown; `PaymentMethod` = card, cash, qr, bank_transfer, line_pay, ecpay, manual, eftpos, moto, cashout (no `adyen`); `PurchaseOrderStatus` is **UPPERCASE** (DRAFT…REFUNDED, incl. PARTIALLY_RECEIVED); `OrderType` = online, pos, b2b, relay, manual, import (no `phone`).

## C. Backend changes

**Backend-Management — no changes needed.** 83 files import the contract; no local enum duplicates; all enums validated via contract `IsValid()`; envelope = contract `apiresponse`; svctoken/pricinginternal/customersinternal mounted on contract paths.

**Backend-Operations — no code changes; flags:**
- `internal/modules/stock/types.go` MovementKind/ReservationStatus are deliberately service-private (documented in-file) — OK.
- `internal/modules/catalog/product_service.go:69` default product status `"active"` is a free string; no contract ProductStatus enum exists. Consistent across stack — flagged for future contract promotion, not changed.
- `internal/modules/purchasing/module.go:443-452` transition timestamp switch has no case for PARTIALLY_RECEIVED/REFUNDED (transition itself is allowed via contract `CanTransitionTo`; only the optional `*_at` stamp is skipped). Cosmetic — flagged.

**Backend-Commerce — 1 change:**
- Backend-Commerce invoice index registration — added guard comment tying the partial-index literals `"draft"/"issued"` to `invoices.InvoiceDraft/InvoiceIssued` (import would cycle; values verified identical).
- Local InvoiceStatus/CartStatus/refund Kind and checkout session extra states (`awaiting_payment`, `expired`) are intentional module-local vocabulary — verified the frontend mirrors them exactly.

## D. Frontend changes (Frontend-Admin-Web)

| File | Fix |
|---|---|
| `src/lib/api/orders.ts` | `OrderStatus`: +paid, +packed, +completed, −packing (invalid). `OrderPaymentStatus`: +unknown, +pending, +partially_refunded. `FulfillmentStatus`: −partially_fulfilled (invalid) → contract's picking_printed/packing/packed/partial. `OrderChannel`: −phone (invalid) → +b2b, +relay, +manual |
| `src/lib/api/payments.ts` | `PaymentMethod`: −adyen (invalid) → full contract set incl. qr/line_pay/ecpay/eftpos/moto/cashout. `PaymentRecordStatus`: +processing, +cancelled, +awaiting_action, +unknown |
| `src/lib/api/purchasing.ts` | PO statuses lowercase → **UPPERCASE** wire values; +PARTIALLY_RECEIVED, +REFUNDED in PO_STATUSES and PO_TRANSITION_TARGETS |
| `src/pages/import/PurchaseOrdersPage.tsx` | statusTone cases, default transition target, and row-disable checks → UPPERCASE; disable now also covers REFUNDED |
| `src/pages/import/ReceiptsPage.tsx` | RECEIPT_STATUSES → contract values actually used (DRAFT, CONFIRMED); isDraft/labels/tones → UPPERCASE |
| `src/contexts/LanguageContext.tsx` | Added zh-TW + zh-CN labels for every new enum value (b2b, relay, qr, line_pay, ecpay, eftpos, moto, cashout, failed, partially_refunded, unknown, awaiting_action, picking_printed, SUBMITTED, CONFIRMED, PARTIALLY_RECEIVED, RECEIVED, REFUNDED) |

Impact before fix: PO status filter/transition sent lowercase → backend `IsValid()` rejection or empty lists; order filters `packing`/`phone`/`partially_fulfilled` matched nothing or were rejected; manual-payment method `adyen` rejected; orders in paid/packed/completed states rendered with no label and couldn't be targeted by the transition dialog.

State/loading/error handling: verified the transition/cancel/confirm/dispatch buttons disable on `saving`, reload list on success, and surface the contract error envelope — no changes needed. Tests: the existing `_wiring.contract.test.ts` snapshot was re-verified by hand against all three live route tables — accurate, no update required.

## E. Consistency matrices

### 1. Contract → backends (drift-relevant rows; ✓ = uses contract type/value correctly)

| Contract item | Mgmt | Ops | Commerce | Fix made | Remaining concern |
|---|---|---|---|---|---|
| SalesOrderStatus / FulfillmentStatus / PaymentStatus | n/a | n/a | ✓ | — | — |
| PaymentRecordStatus / PaymentMethod | n/a | n/a | ✓ | — | — |
| PurchaseOrderStatus (UPPERCASE) | n/a | ✓ | n/a | — | missing `*_at` stamps for PARTIALLY_RECEIVED/REFUNDED |
| Picking/Packing/Outbound/Inbound/Storage enums | n/a | ✓ | n/a | — | — |
| UserRole / coupons / promotions / marketing / loyalty enums | ✓ | n/a | n/a | — | — |
| apiresponse envelope + error codes | ✓ | ✓ | ✓ | — | — |
| common.Money / PageRequest / PageResponse | ✓ | ✓ | ✓ | — | — |
| serviceauth.PathToken + scopes | ✓ provider | ✓ consumer | ✓ consumer | — | — |
| pricing.PathQuote | ✓ provider | n/a | ✓ consumer | — | — |
| stockops Reserve/Commit/Release | n/a | ✓ provider | ✓ consumer | — | — |
| Invoice partial-index literals | n/a | n/a | ⚠ local literals | guard comment | promote InvoiceStatus to contract someday |
| Product status | n/a | ⚠ free string "active" | n/a | flagged only | consider contract enum |

### 2. Backend → frontend (all 195 frontend calls path+method-verified; only deltas shown)

| Backend | Endpoint | Frontend caller | Status | Fix |
|---|---|---|---|---|
| All 3 | every path called by `src/lib/api/*` | 195 functions | ✓ exists, method matches | — |
| Commerce | POST /v1/payments/eftpos | none | orphan | flag — POS-terminal flow not in admin UI |
| Commerce | POST/GET /v1/orders/:id/fulfillments | none | orphan | flag |
| Commerce | terminals/settlements/subscriptions/shipping/volume-discounts groups | none | orphan | flag — likely POS app / store app consumers |
| Management | sessions, devices, security/*, wholesale-presets, customer-coupons, marketing recipients, loyalty promos/check-ins, notification-prefs | none | orphan | flag |
| Operations | inbound-receipts, damage-reports, wms-drafts, layouts, forecasts, depot extras | none | orphan | flag — WMS screens not built yet |

No class-1 defects (dead frontend calls) found; the repo's own wiring contract test enforces this and its snapshot is current.

### 3. End-to-end user flows

| Action/button | FE handler | Endpoint | Backend svc | Contract interaction | Status | Fix |
|---|---|---|---|---|---|---|
| Order status change | OrdersListPage transition dialog | POST /v1/orders/:id/transition | orders.Transition (CanTransitionTo) | SalesOrderStatus | **fixed** | valid targets incl. paid/packed/completed; −packing |
| Order filters | OrdersListPage selects | GET /v1/orders | orders.List | 4 enums | **fixed** | all filter values now valid |
| Cancel order | cancel button | POST /v1/orders/:id/cancel | orders.Transition→cancelled (+stock release via stockops) | ✓ | OK | — |
| Bulk order import | OrderUploadPage | POST /v1/orders/import | import validate (IsValid) | OrderType | OK | CSV passthrough; backend validates |
| Record manual payment | PaymentsListPage form | POST /v1/payments | payments.RecordManual | PaymentMethod | **fixed** | −adyen, full contract set |
| PO transition | PurchaseOrdersPage dialog | POST /v1/purchase-orders/:id/transition | purchasing.Transition | PurchaseOrderStatus | **fixed** | UPPERCASE + new targets |
| Receipt confirm | ReceiptsPage confirm | POST /v1/receipts/:id/confirm | receipts.Confirm (DRAFT→CONFIRMED) | PurchaseOrderStatus | **fixed** | UPPERCASE checks |
| Picking assign/start/complete/cancel, Packing mark-packed/discrepancy, Shipment dispatch, Stock receive/adjust/transfer/reserve/commit/release/import, Expiry run, Invoice issue/void, Refund create, Checkout, Cart ops, Media upload, Auth/Users/Customers/Roles/Settings/Loyalty/Marketing/Coupons/Promotions CRUD | respective pages | verified routes | verified handlers | contract enums | OK | — |

## F. Verification results

| Repo | Command | Result | Notes |
|---|---|---|---|
| all | any shell command (go build/vet/test, npm run typecheck/test/build, vitest) | **BLOCKED** | Sandbox VM failed: `useradd: /etc/passwd: No space left on device` — host disk full, not a repo problem. Rerun needs a working shell/disk space. |
| Frontend | `npm run typecheck && npm test` | not run | Run first — `_wiring.contract.test.ts` re-validates all 195 calls; tsc validates the enum-union changes. |
| Backends | `go build ./... && go test ./...` per repo | not run | No Go code changed except one comment — risk minimal. |
| Static cross-check (done) | source-level diff of contract enums vs all consumers; route-table vs all FE paths | **PASS** | Performed via direct file reads/greps documented above. |

## Diff summary

```
Backend-Commerce invoice index registration                    | +3 (comment)
Frontend-Admin-Web/src/lib/api/orders.ts                       | enum unions/arrays rewritten to contract v5
Frontend-Admin-Web/src/lib/api/payments.ts                     | enum unions/arrays rewritten to contract v5
Frontend-Admin-Web/src/lib/api/purchasing.ts                   | PO statuses → UPPERCASE, +2 values
Frontend-Admin-Web/src/pages/import/PurchaseOrdersPage.tsx     | 4 literal sites → UPPERCASE
Frontend-Admin-Web/src/pages/import/ReceiptsPage.tsx           | 5 literal sites → contract values
Frontend-Admin-Web/src/contexts/LanguageContext.tsx            | +36 i18n keys (2 locales)
```
