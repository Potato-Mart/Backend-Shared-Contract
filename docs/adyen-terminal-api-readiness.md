# Adyen Terminal API Readiness - Shared Contract Audit

Scope: `Backend-Shared-Contract` vs Adyen Terminal API, Cloud Device API,
Management API, and the official Go library:
https://github.com/Adyen/adyen-go-api-library.

## Decision

Use Adyen Terminal API for in-person EFTPOS payments. The shared contract
must stay SDK-free, but it must carry the identifiers and state that the
service using `github.com/adyen/adyen-go-api-library/v21` needs.

Do not add the Adyen SDK dependency to this shared contract module. Add it
to the payment service that actually builds and sends Adyen requests.

Because this removes exported provider enum values, setup DTOs, and
polling DTOs from the public contract, publish it as the next major
contract release when you tag it.

## What Changed

- Removed old provider enum values and setup DTOs.
- Added `TerminalProviderAdyenTerminalAPI`.
- Added `TerminalConnectionMode` for `cloud_sync`, `cloud_async`, and
  `local`.
- Replaced old terminal setup fields with Adyen setup fields:
  `merchant_account`, `poi_id`, `sale_id`, and optional
  `terminal_api_base_url`.
- Added Adyen/nexo request references on terminal transactions:
  `service_id`, `sale_id`, `poi_id`, and `sale_transaction_id`.
- Added Adyen reconciliation references:
  `psp_reference`, `tender_reference`, and provider transaction id.
- Added `TerminalRefundType` so services can distinguish referenced
  refunds (`ReversalRequest`) from unreferenced refunds
  (`PaymentRequest` with `PaymentType=Refund`).
- Added raw provider request/response/notification payload fields for
  diagnostics without forcing the shared contract to mirror every Adyen
  model.
- Replaced polling/min-version DTOs with status-check DTOs. Adyen async
  flows resolve by webhook or transaction status request, not provider
  version polling.

## Adyen Service Requirements

The payment service that depends on this module still needs to implement:

- Adyen SDK dependency:
  `go get github.com/adyen/adyen-go-api-library/v21@v21.2.0`
- API key storage outside the shared contract.
- Cloud Device API or Terminal API endpoint selection:
  test, live AU, live APSE, live EU, or live US.
- `ServiceID` generation: 1-10 alphanumeric characters, unique per
  `POIID` for at least 48 hours.
- `SaleTransactionID` generation: stable merchant reference for the order
  payment attempt.
- Minor-unit to Adyen decimal amount conversion at the service boundary.
- Synchronous and asynchronous result handling:
  - sync response populates normalized fields immediately;
  - async response marks the transaction pending until webhook/status check;
  - missing response after timeout triggers transaction status check;
  - unresolved outcome uses `unknown` plus manual recovery.
- Adyen webhook parsing and HMAC validation in the payment service.
- Receipt handling from Adyen `PaymentReceipt` and/or generated receipt
  strings.

## Field Mapping

| Shared contract | Adyen concept |
| --- | --- |
| `Terminal.MerchantAccount` | Cloud Device API `{merchantAccount}` and request merchant scope |
| `Terminal.POIID` | `MessageHeader.POIID` and Cloud Device API `{deviceId}` |
| `Terminal.SaleID` | `MessageHeader.SaleID` |
| `TerminalTransaction.ServiceID` | `MessageHeader.ServiceID` |
| `TerminalTransaction.SaleTransactionID` | `SaleData.SaleTransactionID.TransactionID` |
| `TerminalTransaction.ProviderTransactionID` | `POIData.POITransactionID.TransactionID` |
| `TerminalTransaction.TenderReference` | First part of `POITransactionID.TransactionID` |
| `TerminalTransaction.PSPReference` | Second part of `POITransactionID.TransactionID` |
| `Amounts.PurchaseMinor` | Original purchase amount in minor units |
| `Amounts.TipMinor` | `AmountsReq/AmountsResp.TipAmount` after conversion |
| `Amounts.CashoutMinor` | `AmountsReq/AmountsResp.CashBackAmount` after conversion |
| `Amounts.AuthorizedMinor` | Authorized total from Adyen response/additional response |

## Important Notes

- Adyen terminal IDs are POIIDs, usually `[terminal model]-[serial number]`.
- The Cloud Device API async call can return HTTP 200 before the payment
  result exists; the actual result arrives in an event notification or via
  transaction status request.
- Referenced refunds are generally asynchronous and should remain pending
  until Adyen refund webhooks confirm the outcome.
- MOTO through Terminal API is represented by tender option `MOTO`, not by a
  separate provider endpoint.
- Tip, surcharge, cashback/cashout, and authorized totals must be recorded
  from the applied result, not only from requested values.
