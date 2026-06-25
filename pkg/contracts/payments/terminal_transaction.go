package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// TerminalTransaction is one card-terminal interaction.
type TerminalTransaction struct {
	ID         string `json:"id"`
	TerminalID string `json:"terminal_id"`
	OrderID    string `json:"order_id,omitempty"`
	PaymentID  string `json:"payment_id,omitempty"`

	ProviderReference *PaymentReference         `json:"provider_reference,omitempty"`
	ProviderDetails   *TerminalProviderDetails  `json:"provider_details,omitempty"`
	OperationContext  *ProviderOperationContext `json:"operation_context,omitempty"`

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

	Payloads *ProviderPayloads `json:"provider_payloads,omitempty"`

	// RecoveryDecision captures the merchant's manual answer when the
	// provider cannot determine the outcome after status checks.
	RecoveryDecision enums.RecoveryDecision `json:"recovery_decision,omitempty"`

	LastStatusCheckedAt *time.Time `json:"last_status_checked_at,omitempty"`
	FinalisedAt         *time.Time `json:"finalised_at,omitempty"`

	Metadata common.Metadata       `json:"metadata,omitempty"`
	History  []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields          `bson:",inline"`
	common.DataProtectionFields `bson:",inline"`
}
