package settlement

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/audit"
	terminal "github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/payments/terminal"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/payments/settlement/settlement_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/payments/terminal/terminal_enums"
)

// Settlement is an end-of-day reconciliation/batch close
// (Type = SettlementTypeSettlement) or a read-only summary
// (Type = SettlementTypeEnquiry).
type Settlement struct {
	ID                   string                                   `json:"id"`
	TerminalID           string                                   `json:"terminal_id"`
	ProviderSettlementID string                                   `json:"provider_settlement_id,omitempty"`
	ProviderDetails      *terminal.TerminalProviderDetails        `json:"provider_details,omitempty"`
	OperationContext     *terminal.ProviderOperationContext       `json:"operation_context,omitempty"`
	Type                 settlement_enums.SettlementType          `json:"type"`
	EnquiryDate          *temporal.Date                           `json:"enquiry_date,omitempty"`
	Status               terminal_enums.TerminalTxStatus          `json:"status"`
	FinancialStatus      terminal_enums.TerminalTxFinancialStatus `json:"financial_status,omitempty"`

	Totals          SettlementTotals `json:"totals"`
	Message         string           `json:"message,omitempty"`
	ProviderResult  string           `json:"provider_result,omitempty"`
	MerchantReceipt string           `json:"merchant_receipt,omitempty"`

	Payloads *terminal.ProviderPayloads `json:"provider_payloads,omitempty"`

	LastStatusCheckedAt *time.Time `json:"last_status_checked_at,omitempty"`
	FinalisedAt         *time.Time `json:"finalised_at,omitempty"`

	Metadata metadata.Metadata       `json:"metadata,omitempty"`
	History  []security.HistoryEntry `json:"history,omitempty"`

	audit.AuditFields
}

// SettlementTotals is the per-category roll-up reported on settlement
// or totals receipts. The fields are provider-neutral and stay in minor
// units for invoice/report consistency.
type SettlementTotals struct {
	Currency        string `json:"currency"`
	PurchasesMinor  int64  `json:"purchases_minor,omitempty"`
	TipsMinor       int64  `json:"tips_minor,omitempty"`
	SurchargesMinor int64  `json:"surcharges_minor,omitempty"`
	RefundsMinor    int64  `json:"refunds_minor,omitempty"`
	CashoutsMinor   int64  `json:"cashouts_minor,omitempty"`
	TotalMinor      int64  `json:"total_minor"`
}
