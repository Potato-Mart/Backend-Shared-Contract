# MX51 Integration Readiness — Shared Contract Audit

Date: 2026-05-09
Scope: `Backend-Shared-Contract` vs MX51 Developer Portal (https://developer.mx51.io/reference) and the SCI guides linked from it.

The MX51 stack offers three integration paths: SCI (cloud, terminal-paired), Spice (local adapter), and SPI (direct device protocol library). For a backend running headless against EFTPOS terminals — i.e. POS data flows through your services — SCI is the right choice and is the model used throughout this audit. A Spice/SPI integration would push the same data shapes into the contract; the gap list below is the same for either.

## TL;DR

The current shared contract models e-commerce orders and a single, generic `Payment` row. It does **not** yet model anything an MX51 SCI integration needs:

- no concept of a paired EFTPOS terminal (pairing_id, key_id, sci_api_base_url, tid, signing_secret_part_b),
- no per-transaction polling state (transaction_id, version, status PENDING/AWAITING_POS/FINALISED, result_financial_status APPROVED/DECLINED/CANCELLED/UNKNOWN),
- no split of purchase / tip / surcharge / cashout / moto amounts,
- no settlement / enquiry record,
- no override (recovery) decision,
- no merchant/customer receipt strings,
- no terminal-side error codes,
- `PaymentMethod` enum collapses card/eftpos/moto/cashout into a single `card`,
- `PaymentRecordStatus` has no `AWAITING_ACTION` or `UNKNOWN` value (both are mandatory MX51 outcomes).

In other words: the contract is fine for "we charged a card and got a yes/no back", but MX51's SCI flow is a long-running, polled, action-driven state machine and none of that surface area exists yet. The integration cannot be built end-to-end on top of the contract as it stands.

---

## What MX51 SCI demands the contract carry

References used throughout: SCI Overview, API Credentials & Authentication, Pairing, Transactions Overview, Purchase/Refund/Signature, Tipping & Surcharges, Cashout, MOTO, Settlement, Terminal receipt configuration, Transaction recovery (all under https://developer.mx51.io/docs).

### 1. Pairing — one record per terminal

`POST /v1/progress-pairing` returns and your POS must persist:

| Field | Notes |
| --- | --- |
| `pairing_id` (`pid_…`) | primary key, sent on every call |
| `key_id` (`kid_…`) | becomes `keyid` in the HTTP Message Signature |
| `signing_secret_part_b` | concatenated with Part A to form the full HMAC secret — stored in a secret store, **not** in the JSON contract |
| `sci_api_base_url` | per-pairing — different terminals can return different URLs |
| `tid`, `pairing_nickname`, `terminal_nickname` | display + reconciliation |
| status | active / unpaired (detected via 401 `no_active_pairings_found` from `GET /v1/pairing-info`) |

`POST /v1/unpair` returns 204 and must flip status. `pairing_not_initial` (422), `pairing_route_forbidden` (403), `test_api_key_forbidden_for_live_pairing` (403) and `pairing_not_found` (404) all need explicit error codes the POS UI can switch on.

### 2. Transaction request

Single endpoint `POST /v1/transactions` accepts exactly one of these blocks. All amounts are integers in **minor units (cents)**.

```
purchase_details              { purchase_amount, tip_amount?, surcharge_amount? }
refund_details                { refund_amount }
cashout_details               { cashout_amount, surcharge_amount? }
purchase_with_cashout_details { purchase_amount, cashout_amount, surcharge_amount? }
moto_details                  { moto_amount, surcharge_amount? }
```

Plus, on every request, four optional terminal/receipt flags:

```
print_merchant_receipt        bool
prompt_customer_receipt       bool
verify_signature_on_terminal  bool
pos_auto_print_signature_receipt bool
```

Critical correctness rule the contract must encode:

> Tip and surcharge MUST be sent as separate fields. Adding them into `purchase_amount` risks double-charging.
> Refund math MUST use `result_amounts.surcharge_amount` from the original transaction, not the request value.

This means the `Payment` row that stores the result has to carry the broken-down amounts, not just a single total — otherwise refunds will be wrong.

### 3. Transaction response — the polling state machine

```
status: PENDING | AWAITING_POS | FINALISED
version: int (monotonic, may skip — always pass back the response value as next min_version)
message: human-readable status to display
result_financial_status: APPROVED | DECLINED | CANCELLED | UNKNOWN     (when FINALISED)
result_amounts: { purchase_amount, tip_amount, surcharge_amount,
                  cashout_amount, refund_amount, moto_amount }         (in cents)
merchant_receipt, customer_receipt: string | null
pos_instructions: { auto_actions[], action_form{ layout, properties, details } }
```

`GET /v1/transactions/{id}?min_version=N` long-polls. Two distinct 404 sub-codes need to be handled differently:
- `transaction_not_found_within_timeout` → re-poll immediately (expected).
- `transaction_not_found` → genuine error.

`424 device_not_connected` is the terminal-offline case.

The Action Framework (`pos_instructions`) is dynamic UI: the API tells the POS what to render and which buttons to show. Buttons carry either `submit_url` (POS POSTs to the URL on click) or `action` (one of `PRINT_MERCHANT_RECEIPT`, `PRINT_CUSTOMER_RECEIPT`, `TRANSACTION_COMPLETE`, `RETRY_TRANSACTION`, `SETTLEMENT_COMPLETE`, `RETRY_SETTLEMENT`, `TEST_ACTION`). The doc explicitly warns not to key business logic on `details` keys.

### 4. Settlement

`POST /v1/settlements` with `{ type: "SETTLEMENT" | "ENQUIRY", enquiry_date?: { year, month, day } }`, then poll `GET /v1/settlements/{id}?min_version=N`. Same response shape as transactions; settlement-only Action Framework codes are `SETTLEMENT_COMPLETE` and `RETRY_SETTLEMENT`. The receipt rolls up totals coded `PUR | TIP | SUR | REF | CAS | TOT`.

### 5. Override flow — mandatory for certification

If the polling loop runs longer than a POS-defined timeout without progress, the POS MUST exit polling and present "Was the transaction successful? Yes/No". The chosen outcome is what the order/payment is closed against. This is mandatory for SCI certification, so the data model needs a place to record it.

### 6. Authentication (out-of-contract but worth flagging)

- Pairing API: `Authorization: ApiKey <SCI Pairing API Secret Key>`.
- SCI API: HTTP Message Signatures (RFC 9421) with HMAC-SHA-256 over `@method @authority @request-target [content-digest]`, `keyid` = `key_id`, secret = `Part A + Part B`. Belongs in the service that talks to MX51, not in the shared contract — but the contract must store everything that signing needs (`key_id`, `sci_api_base_url`).

---

## Current contract — what is in place

Reviewed:

- `pkg/contracts/sales/payment.go` — `Payment{ID, OrderID, Amount, Currency, Method, Status, Gateway, GatewayTransactionID, PaidAt, RefundedAt, RefundAmount, RefundReason, Metadata, CreatedAt}`
- `pkg/contracts/sales/sales_order.go` — `Order` with `Subtotal, DiscountAmount, ShippingAmount, TaxAmount, Total`
- `pkg/contracts/sales/cart.go`, `history.go`
- `pkg/enums/payment_method.go` — `card | cash | qr | bank_transfer | line_pay | ecpay | manual`
- `pkg/enums/payment_status.go` — `unpaid | pending | paid | partially_paid | refunded | partially_refunded`
- `pkg/enums/payment_record_status.go` — `pending | processing | completed | failed | cancelled | refunded`
- `pkg/enums/order_type.go` — includes `pos`
- `pkg/common/money.go` — minor-unit int + currency (correct shape for MX51 — all amounts are cents)
- `pkg/common/metadata.go`, `pkg/apiresponse/*`

What is good: `OrderType` already has a `pos` channel; `Money` is already minor-unit so MX51's "cents" amounts map directly; `Metadata` is available for storing the opaque MX51 blobs without contract churn; `Payment.GatewayTransactionID` can hold the MX51 transaction UUID.

What is missing is everything MX51-specific.

---

## Gap list (concrete, file-by-file)

### G1. No terminal / pairing model — **blocker**

Add new package `pkg/contracts/payments/terminal.go`:

```go
type Terminal struct {
    ID                string                 // your internal id
    TenantID          string
    StoreID           string
    Provider          enums.TerminalProvider // mx51_sci | mx51_spice | mx51_spi | …
    PairingID         string                 // pid_…
    KeyID             string                 // kid_…
    TID               string
    PairingNickname   string
    TerminalNickname  string
    SciAPIBaseURL     string                 // per pairing — must be persisted
    Status            enums.TerminalStatus   // pairing | active | unpaired | expired | error
    PairedAt          *time.Time
    UnpairedAt        *time.Time
    LastSeenAt        *time.Time
    Metadata          common.Metadata
    common.AuditFields
}
```

`signing_secret_part_b` and Part A are NOT stored in the JSON contract — they belong in your secret store. This file should explicitly say so in a comment so consumers don't put them in the struct later.

### G2. No terminal-transaction model — **blocker**

Add `pkg/contracts/payments/terminal_transaction.go`:

```go
type TerminalTransaction struct {
    ID                    string
    TerminalID            string
    OrderID               string
    PaymentID             string                       // links to sales.Payment
    ProviderTransactionID string                       // mx51 UUID
    Type                  enums.TerminalTxType         // purchase | refund | cashout | purchase_with_cashout | moto | settlement | enquiry
    Requested             RequestedAmounts             // amounts the POS asked for
    ReceiptOptions        ReceiptOptions               // the four boolean flags
    Status                enums.TerminalTxStatus       // pending | awaiting_pos | finalised | override_pending | override_resolved
    Version               int                          // for next min_version
    FinancialStatus       enums.TerminalTxFinancialStatus  // approved | declined | cancelled | unknown | ""
    Result                ResultAmounts                // the source-of-truth split
    MerchantReceipt       string
    CustomerReceipt       string
    Message               string
    POSInstructions       json.RawMessage              // opaque action-framework blob
    LastPolledAt          *time.Time
    FinalisedAt           *time.Time
    OverrideDecision      enums.RecoveryDecision       // pending | approved | declined
    common.AuditFields
}

type RequestedAmounts struct {
    Currency string
    Purchase int64; Tip int64; Surcharge int64; Cashout int64; Refund int64; MOTO int64
}

type ResultAmounts struct {
    Currency string
    Purchase int64; Tip int64; Surcharge int64; Cashout int64; Refund int64; MOTO int64
}

type ReceiptOptions struct {
    PrintMerchantReceipt          bool
    PromptCustomerReceipt         bool
    VerifySignatureOnTerminal     bool
    POSAutoPrintSignatureReceipt  bool
}
```

`POSInstructions` is intentionally `json.RawMessage` — MX51 explicitly says fields under `details` can change, and the `properties` dictionary is keyed dynamically. Decoding it strongly inside the shared contract creates a moving target.

### G3. No settlement model — **blocker for end-of-day**

Add `pkg/contracts/payments/settlement.go`:

```go
type Settlement struct {
    ID                   string
    TerminalID           string
    ProviderSettlementID string
    Type                 enums.SettlementType   // settlement | enquiry
    EnquiryDate          *common.Date           // YMD if enquiry
    Status               enums.TerminalTxStatus // reuse: pending | awaiting_pos | finalised | override_*
    FinancialStatus      enums.TerminalTxFinancialStatus
    Version              int
    Totals               SettlementTotals       // PUR / TIP / SUR / REF / CAS / TOT
    MerchantReceipt      string
    Message              string
    POSInstructions      json.RawMessage
    FinalisedAt          *time.Time
    common.AuditFields
}

type SettlementTotals struct {
    Currency  string
    Purchases int64; Tips int64; Surcharges int64; Refunds int64; Cashouts int64; Total int64
}
```

### G4. `enums.PaymentMethod` is too coarse for MX51

`card` collapses standard EFTPOS, MOTO and cashout into a single value. Add (or move to a sub-type):

```go
PaymentMethodEFTPOS  PaymentMethod = "eftpos"
PaymentMethodMOTO    PaymentMethod = "moto"
PaymentMethodCashout PaymentMethod = "cashout"
```

This matters for reconciliation against settlement reports (which split by transaction type) and for compliance — MX51 explicitly requires the POS UI to surface MOTO and cashout buttons separately from a standard purchase.

### G5. `enums.PaymentRecordStatus` is missing two MX51 outcomes — **blocker**

Add:

```go
PaymentRecordStatusAwaitingAction PaymentRecordStatus = "awaiting_action"  // ↔ MX51 AWAITING_POS (signature, override, MOTO entry)
PaymentRecordStatusUnknown        PaymentRecordStatus = "unknown"          // ↔ MX51 result_financial_status = UNKNOWN
```

Without `unknown`, the override flow can't be persisted truthfully; without `awaiting_action`, signature transactions have no state to live in while the merchant approves.

### G6. `sales.Payment` is missing the broken-down amounts

MX51's docs are explicit: `result_amounts` is the source of truth for tax invoices and refund calculations. Add:

```go
TipAmount        common.Money
SurchargeAmount  common.Money
CashoutAmount    common.Money
MOTOAmount       common.Money
TerminalID       string         // links to Terminal
MerchantReceipt  string
CustomerReceipt  string
RecoveryDecision enums.RecoveryDecision  // for UNKNOWN outcomes
```

Also: `Currency` duplicates `Money.Currency` and should be removed (or re-purposed and documented).

### G7. `sales.Order` does not surface tip / surcharge

For POS orders these change the customer-paid total and are line items on the receipt. Add to `Order`:

```go
TipAmount       common.Money
SurchargeAmount common.Money
```

(Cashout isn't part of the order total — it's a parallel cash withdrawal — so it belongs only on `Payment`/`TerminalTransaction`.)

### G8. New error codes in `apiresponse`

Add at minimum:

```go
CodeTerminalNotConnected   Code = "TERMINAL_NOT_CONNECTED"        // ↔ device_not_connected (424)
CodeTerminalUnpaired       Code = "TERMINAL_UNPAIRED"             // ↔ no_active_pairings_found (401)
CodeTerminalOutcomeUnknown Code = "TERMINAL_OUTCOME_UNKNOWN"      // ↔ result_financial_status = UNKNOWN
CodePairingNotFound        Code = "PAIRING_NOT_FOUND"             // 404
CodePairingExpired         Code = "PAIRING_EXPIRED"               // 422 pairing_not_initial
CodePairingForbidden       Code = "PAIRING_FORBIDDEN"             // 403 pairing_route_forbidden / test_api_key_forbidden
CodeSignatureRejected      Code = "TERMINAL_SIGNATURE_REJECTED"
```

Map them in `Code.HTTPStatus()`.

### G9. New enums

`pkg/enums/terminal_provider.go`, `terminal_status.go`, `terminal_tx_type.go`, `terminal_tx_status.go`, `terminal_tx_financial_status.go`, `recovery_decision.go`, `settlement_type.go`. All trivial string enums with `IsValid()` matching the existing house style.

### G10. `common.Date` for settlement enquiries

MX51 settlement enquiry takes `{year, month, day}` as three separate integers. The current `common.Date` is a `string` (`YYYY-MM-DD`), which works at storage time but has to be split out at the MX51 client edge. Either:

- keep `common.Date` as-is and split at the client (simpler, no contract change), or
- add a sibling `common.YMD struct{ Year, Month, Day int }` for fields that round-trip directly to MX51.

Recommendation: keep `Date` and split at the edge; the MX51 wire shape shouldn't leak into the shared contract.

### G11. Payment-status transitions

`enums.PaymentStatus` currently goes `unpaid → pending → paid`. MX51 introduces a real "money is in flight on the terminal" state. The cleanest split is to keep `PaymentStatus` as the order-level summary and let the new `PaymentRecordStatusAwaitingAction` carry the granular state on the `Payment` row. No change to `PaymentStatus` values is required, but the transition rules must allow `pending` to live for as long as the polling loop runs.

### G12. Idempotency

MX51 doesn't require an idempotency key on `POST /v1/transactions` (it returns its own UUID), but your own services need one to ensure the POS never double-creates a `TerminalTransaction` after a network blip. There's no shared "request envelope" with `Idempotency-Key` in the contract — recommend adding one (e.g. a `common.IdempotencyKey` header constant + a documented header).

---

## Suggested file plan

```
pkg/
  contracts/
    payments/                           ← new package
      terminal.go                       (G1)
      terminal_transaction.go           (G2)
      settlement.go                     (G3)
      action_framework.go               (constants for known action codes; opaque blob type alias)
  enums/
    payment_method.go                   (G4 — extend)
    payment_record_status.go            (G5 — extend)
    terminal_provider.go                (new)
    terminal_status.go                  (new)
    terminal_tx_type.go                 (new)
    terminal_tx_status.go               (new)
    terminal_tx_financial_status.go     (new)
    recovery_decision.go                (new)
    settlement_type.go                  (new)
  common/
    date.go                             (G10 — add Date if missing)
  apiresponse/
    app_error.go                        (G8 — extend)
  contracts/sales/
    payment.go                          (G6 — extend)
    sales_order.go                      (G7 — extend)
```

Existing consumers can ignore the new `payments` package until they wire MX51 in, so this is purely additive for them. The two extensions in `enums` (G4, G5) are additive too — the `IsValid()` switch just gets new cases.

---

## What is **already** OK and needs no change

- `common.Money` — minor-unit int + currency code is the exact shape MX51 uses for amounts.
- `common.Metadata` — fits nicely as the place to dump MX51 `details` if you don't want an opaque `RawMessage`.
- `OrderType.pos` already exists — no new channel needed.
- `apiresponse.APIResponse` envelope — works for both success and error sides.

---

## Conclusion

The shared contract is well-structured for an e-commerce/WMS world but has zero surface area for a card-terminal integration. To be ready to integrate MX51 SCI you need at minimum: the **Terminal**, **TerminalTransaction**, and **Settlement** entities (G1–G3), the two new `PaymentRecordStatus` values (G5, mandatory for the override flow that SCI certification requires), and the broken-down amounts on `Payment` (G6, otherwise refunds will be wrong). G4 (finer payment method), G7 (tip/surcharge on Order), G8 (error codes) and the new enums are needed for a clean, certifiable build.

Treat G1, G2, G3, G5, G6 as blockers. Everything else is "before you ship to production".
