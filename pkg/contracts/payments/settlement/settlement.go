package settlement

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/settlement/settlement_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/terminal/terminal_enums"
)

// Settlement is an end-of-day reconciliation/batch close
// (Type = SettlementTypeSettlement) or a read-only summary
// (Type = SettlementTypeEnquiry).
type Settlement struct {
	ID                   string                                   `json:"id"`
	TerminalID           string                                   `json:"terminal_id"`
	ProviderSettlementID string                                   `json:"provider_settlement_id,omitempty"`
	Type                 settlement_enums.SettlementType          `json:"type"`
	EnquiryDate          *temporal.Date                           `json:"enquiry_date,omitempty"`
	Status               terminal_enums.TerminalTxStatus          `json:"status"`
	FinancialStatus      terminal_enums.TerminalTxFinancialStatus `json:"financial_status,omitempty"`

	Totals          SettlementTotals `json:"totals"`
	Message         string           `json:"message,omitempty"`
	ProviderResult  string           `json:"provider_result,omitempty"`
	MerchantReceipt string           `json:"merchant_receipt,omitempty"`

	LastStatusCheckedAt *time.Time `json:"last_status_checked_at,omitempty"`
	FinalisedAt         *time.Time `json:"finalised_at,omitempty"`

	History []security.HistoryEntry `json:"history,omitempty"`

	audit.AuditFields
}
