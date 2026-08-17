package order

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/money"
	paymentcontracts "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/payments/payment"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/payments/payment/payment_enums"
)

// Payment is the order-level record of money received against an order.
type Payment struct {
	ID          string      `json:"id"`
	OrderNumber string      `json:"order_number"`
	Amount      money.Money `json:"amount"`

	Method            payment_enums.PaymentMethod        `json:"method"`
	Status            payment_enums.PaymentRecordStatus  `json:"status"`
	Provider          string                             `json:"provider,omitempty"`
	ProviderReference *paymentcontracts.PaymentReference `json:"provider_reference,omitempty"`

	// ProcessorFee and NetAmount are the payment processor's transaction
	// fee and the settled amount net of that fee (e.g. from the Stripe
	// balance transaction). Present only when the processor reports them.
	ProcessorFee *money.Money `json:"processor_fee,omitempty"`
	NetAmount    *money.Money `json:"net_amount,omitempty"`

	// Terminal linkage, populated for eftpos / moto / cashout methods.
	TerminalID            string `json:"terminal_id,omitempty"`
	TerminalTransactionID string `json:"terminal_transaction_id,omitempty"`

	// Component breakdown in the same currency as Amount. Each is
	// optional and zero-valued when not applicable to the method.
	TipAmount       *money.Money `json:"tip_amount,omitempty"`
	SurchargeAmount *money.Money `json:"surcharge_amount,omitempty"`
	CashoutAmount   *money.Money `json:"cashout_amount,omitempty"`
	MOTOAmount      *money.Money `json:"moto_amount,omitempty"`

	// Receipt strings exactly as returned by the terminal. Optional;
	// when present, the POS is expected to print or store them.
	MerchantReceipt string `json:"merchant_receipt,omitempty"`
	CustomerReceipt string `json:"customer_receipt,omitempty"`

	// RecoveryDecision is the merchant's manual answer when the terminal
	// outcome stayed unknown after provider status checks.
	RecoveryDecision payment_enums.RecoveryDecision `json:"recovery_decision,omitempty"`

	PaidAt       *time.Time              `json:"paid_at,omitempty"`
	RefundedAt   *time.Time              `json:"refunded_at,omitempty"`
	RefundAmount *money.Money            `json:"refund_amount,omitempty"`
	RefundReason string                  `json:"refund_reason,omitempty"`
	Metadata     metadata.Metadata       `json:"metadata,omitempty"`
	History      []security.HistoryEntry `json:"history,omitempty"`
	CreatedAt    time.Time               `json:"created_at"`

	security.DataProtectionFields
}
