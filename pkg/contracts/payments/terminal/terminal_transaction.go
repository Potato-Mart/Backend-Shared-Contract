package terminal

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/audit"
	payment "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/payments/payment"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/payments/payment/payment_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/payments/terminal/terminal_enums"
)

// TerminalTransaction is one card-terminal interaction.
type TerminalTransaction struct {
	ID          string `json:"id"`
	TerminalID  string `json:"terminal_id"`
	OrderNumber string `json:"order_number,omitempty"`
	PaymentID   string `json:"payment_id,omitempty"`

	ProviderReference *payment.PaymentReference `json:"provider_reference,omitempty"`
	ProviderDetails   *TerminalProviderDetails  `json:"provider_details,omitempty"`
	OperationContext  *ProviderOperationContext `json:"operation_context,omitempty"`

	Type           terminal_enums.TerminalTxType `json:"type"`
	Requested      Amounts                       `json:"requested"`
	ReceiptOptions ReceiptOptions                `json:"receipt_options,omitempty"`

	Status          terminal_enums.TerminalTxStatus          `json:"status"`
	FinancialStatus terminal_enums.TerminalTxFinancialStatus `json:"financial_status,omitempty"`
	Result          Amounts                                  `json:"result"`

	Message                string            `json:"message,omitempty"`
	ProviderResult         string            `json:"provider_result,omitempty"`
	ProviderErrorCondition string            `json:"provider_error_condition,omitempty"`
	ProviderData           metadata.Metadata `json:"provider_data,omitempty"`
	MerchantReceipt        string            `json:"merchant_receipt,omitempty"`
	CustomerReceipt        string            `json:"customer_receipt,omitempty"`

	Payloads *ProviderPayloads `json:"provider_payloads,omitempty"`

	// RecoveryDecision captures the merchant's manual answer when the
	// provider cannot determine the outcome after status checks.
	RecoveryDecision payment_enums.RecoveryDecision `json:"recovery_decision,omitempty"`

	LastStatusCheckedAt *time.Time `json:"last_status_checked_at,omitempty"`
	FinalisedAt         *time.Time `json:"finalised_at,omitempty"`

	Metadata metadata.Metadata       `json:"metadata,omitempty"`
	History  []security.HistoryEntry `json:"history,omitempty"`

	audit.AuditFields
	security.DataProtectionFields
}
