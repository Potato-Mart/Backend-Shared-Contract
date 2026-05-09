package payments

import (
	"encoding/json"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// TerminalTransaction is one card-terminal interaction. It is the long-
// lived state behind a polled transaction: created when the POS submits
// the request, mutated on every poll, and finalised when the terminal
// reaches a terminal status (Finalised or OverrideResolved).
//
// Relationship to sales.Payment: a TerminalTransaction is the
// instrument-level record (one row per terminal interaction). A
// sales.Payment is the order-level record (one row per "money received
// against the order"). For terminal-backed payments they are 1:1; for
// other methods (cash, bank transfer) only Payment exists.
//
// Polling contract: callers must always pass back the response
// Version as the basis for the next min_version. The provider may
// skip versions if multiple updates happened between polls.
type TerminalTransaction struct {
	ID         string `json:"id"`
	TerminalID string `json:"terminal_id"`
	OrderID    string `json:"order_id,omitempty"`
	PaymentID  string `json:"payment_id,omitempty"`

	// ProviderTransactionID is the UUID assigned by the terminal
	// provider (e.g. MX51's transaction id). Use this when calling
	// GET /transactions/{id}.
	ProviderTransactionID string `json:"provider_transaction_id"`

	Type           enums.TerminalTxType `json:"type"`
	Requested      Amounts              `json:"requested"`
	ReceiptOptions ReceiptOptions       `json:"receipt_options,omitempty"`

	// Polling state.
	Status          enums.TerminalTxStatus          `json:"status"`
	Version         int                             `json:"version"`
	FinancialStatus enums.TerminalTxFinancialStatus `json:"financial_status,omitempty"`
	Result          Amounts                         `json:"result"`

	Message         string `json:"message,omitempty"`
	MerchantReceipt string `json:"merchant_receipt,omitempty"`
	CustomerReceipt string `json:"customer_receipt,omitempty"`

	// POSInstructions holds the provider's Action Framework payload
	// verbatim. It is intentionally an opaque blob - MX51 explicitly
	// warns that the keys inside `details` and the `properties`
	// dictionary are dynamic and must not be relied on for business
	// logic. Decode it at the rendering layer, not here.
	POSInstructions json.RawMessage `json:"pos_instructions,omitempty"`

	// OverrideDecision captures the merchant's manual answer to the
	// recovery dialog when polling times out or FinancialStatus is
	// Unknown. Required for SCI certification.
	OverrideDecision enums.RecoveryDecision `json:"override_decision,omitempty"`

	// IdempotencyKey lets the POS safely retry a Create call after a
	// network blip without double-creating the transaction. It is the
	// caller's responsibility to generate and persist; the contract
	// only guarantees the field exists.
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	LastPolledAt *time.Time `json:"last_polled_at,omitempty"`
	FinalisedAt  *time.Time `json:"finalised_at,omitempty"`

	Metadata common.Metadata `json:"metadata,omitempty"`

	common.AuditFields
}

// CreateTerminalTransactionRequest is the input the POS submits to
// start a new terminal transaction. Exactly one of the *Details
// pointers must be non-nil; the type is implied by which one is set
// (and recorded explicitly on the resulting TerminalTransaction.Type).
type CreateTerminalTransactionRequest struct {
	TerminalID     string         `json:"terminal_id"`
	OrderID        string         `json:"order_id,omitempty"`
	PaymentID      string         `json:"payment_id,omitempty"`
	IdempotencyKey string         `json:"idempotency_key,omitempty"`
	ReceiptOptions ReceiptOptions `json:"receipt_options,omitempty"`

	Purchase            *PurchaseDetails            `json:"purchase_details,omitempty"`
	Refund              *RefundDetails              `json:"refund_details,omitempty"`
	Cashout             *CashoutDetails             `json:"cashout_details,omitempty"`
	PurchaseWithCashout *PurchaseWithCashoutDetails `json:"purchase_with_cashout_details,omitempty"`
	MOTO                *MOTODetails                `json:"moto_details,omitempty"`
}

// PurchaseDetails models a standard purchase. Tip and surcharge are
// optional and MUST be passed as separate fields, never folded into
// PurchaseAmountMinor (otherwise the customer can be double-charged
// when terminal-based tipping or surcharging is also active).
//
// A nil pointer for TipMinor / SurchargeMinor means "let the terminal
// decide" (it may prompt the customer); a *int64 set to 0 means
// "explicitly suppress this component". This distinction matters for
// MX51 and is preserved by using pointer fields.
type PurchaseDetails struct {
	PurchaseAmountMinor int64  `json:"purchase_amount_minor"`
	TipMinor            *int64 `json:"tip_minor,omitempty"`
	SurchargeMinor      *int64 `json:"surcharge_minor,omitempty"`
}

// RefundDetails models a refund. When refunding a previous purchase
// that included a surcharge, callers must compute RefundAmountMinor
// from the original transaction's Result.SurchargeMinor, not from the
// originally requested value.
type RefundDetails struct {
	RefundAmountMinor int64 `json:"refund_amount_minor"`
}

// CashoutDetails models a cashout-only transaction (customer withdraws
// cash without making a purchase). Cashout cannot be combined with
// tipping and cannot be used on a credit card.
type CashoutDetails struct {
	CashoutAmountMinor int64  `json:"cashout_amount_minor"`
	SurchargeMinor     *int64 `json:"surcharge_minor,omitempty"`
}

// PurchaseWithCashoutDetails combines a purchase and a cashout in one
// transaction. CashoutAmountMinor may be nil to let the terminal
// prompt the customer for the cashout amount.
type PurchaseWithCashoutDetails struct {
	PurchaseAmountMinor int64  `json:"purchase_amount_minor"`
	CashoutAmountMinor  *int64 `json:"cashout_amount_minor,omitempty"`
	SurchargeMinor      *int64 `json:"surcharge_minor,omitempty"`
}

// MOTODetails models a Mail Order / Telephone Order (card-not-present)
// transaction. Tipping and cashout cannot be used with MOTO.
type MOTODetails struct {
	MOTOAmountMinor int64  `json:"moto_amount_minor"`
	SurchargeMinor  *int64 `json:"surcharge_minor,omitempty"`
}

// PollTerminalTransactionRequest is the input to the long-poll GET.
// MinVersion is the value the caller wants to wait beyond.
type PollTerminalTransactionRequest struct {
	TerminalTransactionID string `json:"terminal_transaction_id"`
	MinVersion            int    `json:"min_version"`
}

// SubmitActionRequest is the POS-to-server action sent when the
// merchant clicks an Action-Framework button whose submit_url is
// non-null (e.g. APPROVE_SIGNATURE).
type SubmitActionRequest struct {
	TerminalTransactionID string          `json:"terminal_transaction_id"`
	Action                string          `json:"action"`
	Inputs                common.Metadata `json:"inputs,omitempty"`
}

// ResolveOverrideRequest closes a transaction stuck in
// OverridePending status with the merchant's manual decision.
type ResolveOverrideRequest struct {
	TerminalTransactionID string                 `json:"terminal_transaction_id"`
	Decision              enums.RecoveryDecision `json:"decision"`
	Note                  string                 `json:"note,omitempty"`
}
