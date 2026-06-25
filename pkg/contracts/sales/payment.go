package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	paymentcontracts "github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/payments"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// Payment is the order-level record of money received against an order.
type Payment struct {
	ID      string       `json:"id"`
	OrderID string       `json:"order_id"`
	Amount  common.Money `json:"amount"`

	// Currency is retained for backward compatibility with consumers
	// that read it as a sibling of Amount. New code should read
	// Amount.Currency instead.
	//
	// Deprecated: use Amount.Currency.
	Currency string `json:"currency,omitempty"`

	Method            enums.PaymentMethod                `json:"method"`
	Status            enums.PaymentRecordStatus          `json:"status"`
	Provider          string                             `json:"provider,omitempty"`
	ProviderReference *paymentcontracts.PaymentReference `json:"provider_reference,omitempty"`

	// Terminal linkage, populated for eftpos / moto / cashout methods.
	TerminalID            string `json:"terminal_id,omitempty"`
	TerminalTransactionID string `json:"terminal_transaction_id,omitempty"`

	// Component breakdown in the same currency as Amount. Each is
	// optional and zero-valued when not applicable to the method.
	TipAmount       *common.Money `json:"tip_amount,omitempty"`
	SurchargeAmount *common.Money `json:"surcharge_amount,omitempty"`
	CashoutAmount   *common.Money `json:"cashout_amount,omitempty"`
	MOTOAmount      *common.Money `json:"moto_amount,omitempty"`

	// Receipt strings exactly as returned by the terminal. Optional;
	// when present, the POS is expected to print or store them.
	MerchantReceipt string `json:"merchant_receipt,omitempty"`
	CustomerReceipt string `json:"customer_receipt,omitempty"`

	// RecoveryDecision is the merchant's manual answer when the terminal
	// outcome stayed unknown after provider status checks.
	RecoveryDecision enums.RecoveryDecision `json:"recovery_decision,omitempty"`

	PaidAt       *time.Time            `json:"paid_at,omitempty"`
	RefundedAt   *time.Time            `json:"refunded_at,omitempty"`
	RefundAmount *common.Money         `json:"refund_amount,omitempty"`
	RefundReason string                `json:"refund_reason,omitempty"`
	Metadata     common.Metadata       `json:"metadata,omitempty"`
	History      []shared.HistoryEntry `json:"history,omitempty"`
	CreatedAt    time.Time             `json:"created_at"`

	common.DataProtectionFields `bson:",inline"`
}
