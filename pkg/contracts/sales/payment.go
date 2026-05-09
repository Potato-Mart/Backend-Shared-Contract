package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// Payment is the order-level record of money received against an
// order. It is method-agnostic: cash, card, EFTPOS, MOTO, QR, bank
// transfer all share this shape.
//
// For terminal-backed payments (Method = eftpos / moto / cashout) the
// instrument-level state lives on payments.TerminalTransaction and is
// linked back via TerminalID + TerminalTransactionID. The split lets
// the polled, action-driven terminal lifecycle stay separate from the
// merchant's "did we get the money" view of the world.
//
// Amount on this struct is the total amount received from the
// customer for this payment line. The component breakdown (Tip,
// Surcharge, Cashout, MOTO) is captured separately because terminal
// providers (notably MX51 SCI) report them split out and the split is
// the source of truth for tax invoices and refund math.
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

	Method               enums.PaymentMethod       `json:"method"`
	Status               enums.PaymentRecordStatus `json:"status"`
	Gateway              string                    `json:"gateway,omitempty"`
	GatewayTransactionID string                    `json:"gateway_transaction_id,omitempty"`

	// Terminal linkage, populated for eftpos / moto / cashout methods.
	TerminalID            string `json:"terminal_id,omitempty"`
	TerminalTransactionID string `json:"terminal_transaction_id,omitempty"`

	// Component breakdown in the same currency as Amount. Each is
	// optional and zero-valued when not applicable to the method.
	// These mirror the result_amounts shape returned by MX51 SCI.
	TipAmount       *common.Money `json:"tip_amount,omitempty"`
	SurchargeAmount *common.Money `json:"surcharge_amount,omitempty"`
	CashoutAmount   *common.Money `json:"cashout_amount,omitempty"`
	MOTOAmount      *common.Money `json:"moto_amount,omitempty"`

	// Receipt strings exactly as returned by the terminal. Optional;
	// when present, the POS is expected to print them.
	MerchantReceipt string `json:"merchant_receipt,omitempty"`
	CustomerReceipt string `json:"customer_receipt,omitempty"`

	// RecoveryDecision is the merchant's manual answer to the override
	// dialog. Populated when Status is Unknown after the polling loop
	// timed out and the merchant resolved it manually.
	RecoveryDecision enums.RecoveryDecision `json:"recovery_decision,omitempty"`

	PaidAt       *time.Time      `json:"paid_at,omitempty"`
	RefundedAt   *time.Time      `json:"refunded_at,omitempty"`
	RefundAmount *common.Money   `json:"refund_amount,omitempty"`
	RefundReason string          `json:"refund_reason,omitempty"`
	Metadata     common.Metadata `json:"metadata,omitempty"`
	CreatedAt    time.Time       `json:"created_at"`
}
