package payments

import (
	"encoding/json"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/enums"
)

// TerminalTransaction is one card-terminal interaction.
type TerminalTransaction struct {
	ID         string `json:"id"`
	TerminalID string `json:"terminal_id"`
	OrderID    string `json:"order_id,omitempty"`
	PaymentID  string `json:"payment_id,omitempty"`

	ProviderReference  *PaymentReference `json:"provider_reference,omitempty"`
	ProviderRequestID  string            `json:"provider_request_id,omitempty"`
	ProviderTerminalID string            `json:"provider_terminal_id,omitempty"`
	MerchantReference  string            `json:"merchant_reference,omitempty"`

	Type           enums.TerminalTxType `json:"type"`
	Requested      Amounts              `json:"requested"`
	ReceiptOptions ReceiptOptions       `json:"receipt_options,omitempty"`

	Status          enums.TerminalTxStatus          `json:"status"`
	FinancialStatus enums.TerminalTxFinancialStatus `json:"financial_status,omitempty"`
	Result          Amounts                         `json:"result"`

	Message                string          `json:"message,omitempty"`
	ProviderResult         string          `json:"provider_result,omitempty"`
	ProviderErrorCondition string          `json:"provider_error_condition,omitempty"`
	ProviderData           common.Metadata `json:"provider_data,omitempty"`
	MerchantReceipt        string          `json:"merchant_receipt,omitempty"`
	CustomerReceipt        string          `json:"customer_receipt,omitempty"`

	// Raw provider payloads are stored for diagnostics and future
	// compatibility. They should not be used as the source of truth for
	// business decisions once the normalized fields above are populated.
	ProviderRequest      json.RawMessage `json:"provider_request,omitempty"`
	ProviderResponse     json.RawMessage `json:"provider_response,omitempty"`
	ProviderNotification json.RawMessage `json:"provider_notification,omitempty"`
	DisplayNotification  json.RawMessage `json:"display_notification,omitempty"`

	// RecoveryDecision captures the merchant's manual answer when the
	// provider cannot determine the outcome after status checks.
	RecoveryDecision enums.RecoveryDecision `json:"recovery_decision,omitempty"`

	// IdempotencyKey lets the POS safely retry a Create call after a
	// network blip without double-creating the transaction. It is the
	// caller's responsibility to generate and persist; the contract
	// only guarantees the field exists.
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	LastStatusCheckedAt *time.Time `json:"last_status_checked_at,omitempty"`
	FinalisedAt         *time.Time `json:"finalised_at,omitempty"`

	Metadata common.Metadata       `json:"metadata,omitempty"`
	History  []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields          `bson:",inline"`
	common.DataProtectionFields `bson:",inline"`
}

// CreateTerminalTransactionRequest is the input the POS submits to
// start a new terminal transaction.
type CreateTerminalTransactionRequest struct {
	TerminalID        string                       `json:"terminal_id"`
	OrderID           string                       `json:"order_id,omitempty"`
	PaymentID         string                       `json:"payment_id,omitempty"`
	Currency          string                       `json:"currency"`
	ConnectionMode    enums.TerminalConnectionMode `json:"connection_mode,omitempty"`
	IdempotencyKey    string                       `json:"idempotency_key,omitempty"`
	ProviderRequestID string                       `json:"provider_request_id,omitempty"`
	MerchantReference string                       `json:"merchant_reference,omitempty"`

	ReceiptOptions  ReceiptOptions  `json:"receipt_options,omitempty"`
	ProviderOptions common.Metadata `json:"provider_options,omitempty"`

	Purchase            *PurchaseDetails            `json:"purchase_details,omitempty"`
	Refund              *RefundDetails              `json:"refund_details,omitempty"`
	Cashout             *CashoutDetails             `json:"cashout_details,omitempty"`
	PurchaseWithCashout *PurchaseWithCashoutDetails `json:"purchase_with_cashout_details,omitempty"`
	MOTO                *MOTODetails                `json:"moto_details,omitempty"`
}

// PurchaseDetails models a standard purchase.
type PurchaseDetails struct {
	PurchaseAmountMinor int64  `json:"purchase_amount_minor"`
	TipMinor            *int64 `json:"tip_minor,omitempty"`
	SurchargeMinor      *int64 `json:"surcharge_minor,omitempty"`
}

// RefundDetails models both referenced and unreferenced refunds.
type RefundDetails struct {
	RefundType                enums.TerminalRefundType `json:"refund_type,omitempty"`
	RefundAmountMinor         int64                    `json:"refund_amount_minor"`
	OriginalProviderReference *PaymentReference        `json:"original_provider_reference,omitempty"`
}

// CashoutDetails models a cashout-only transaction.
type CashoutDetails struct {
	CashoutAmountMinor int64  `json:"cashout_amount_minor"`
	SurchargeMinor     *int64 `json:"surcharge_minor,omitempty"`
}

// PurchaseWithCashoutDetails combines a purchase and a cashout in one
// transaction.
type PurchaseWithCashoutDetails struct {
	PurchaseAmountMinor int64  `json:"purchase_amount_minor"`
	CashoutAmountMinor  *int64 `json:"cashout_amount_minor,omitempty"`
	SurchargeMinor      *int64 `json:"surcharge_minor,omitempty"`
}

// MOTODetails models a Mail Order / Telephone Order card-not-present
// transaction initiated through the terminal.
type MOTODetails struct {
	MOTOAmountMinor int64  `json:"moto_amount_minor"`
	SurchargeMinor  *int64 `json:"surcharge_minor,omitempty"`
}

// CheckTerminalTransactionStatusRequest asks the backend to refresh a
// terminal transaction outcome using the provider's status mechanism.
type CheckTerminalTransactionStatusRequest struct {
	TerminalTransactionID string `json:"terminal_transaction_id"`
}

// ResolveOverrideRequest closes a transaction stuck in OverridePending
// status with the merchant's manual decision.
type ResolveOverrideRequest struct {
	TerminalTransactionID string                 `json:"terminal_transaction_id"`
	Decision              enums.RecoveryDecision `json:"decision"`
	Note                  string                 `json:"note,omitempty"`
}
