package payments

import (
	"encoding/json"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// Settlement is an end-of-day batch close (Type = SettlementTypeSettlement)
// or a read-only summary (Type = SettlementTypeEnquiry). It uses the
// same polling envelope as TerminalTransaction - same Status, Version,
// FinancialStatus, MerchantReceipt and POSInstructions semantics.
type Settlement struct {
	ID                   string                          `json:"id"`
	TerminalID           string                          `json:"terminal_id"`
	ProviderSettlementID string                          `json:"provider_settlement_id"`
	Type                 enums.SettlementType            `json:"type"`
	EnquiryDate          *common.Date                    `json:"enquiry_date,omitempty"`
	Status               enums.TerminalTxStatus          `json:"status"`
	Version              int                             `json:"version"`
	FinancialStatus      enums.TerminalTxFinancialStatus `json:"financial_status,omitempty"`

	Totals          SettlementTotals `json:"totals"`
	Message         string           `json:"message,omitempty"`
	MerchantReceipt string           `json:"merchant_receipt,omitempty"`

	POSInstructions json.RawMessage `json:"pos_instructions,omitempty"`

	IdempotencyKey string `json:"idempotency_key,omitempty"`

	LastPolledAt *time.Time `json:"last_polled_at,omitempty"`
	FinalisedAt  *time.Time `json:"finalised_at,omitempty"`

	Metadata common.Metadata `json:"metadata,omitempty"`

	common.AuditFields
}

// SettlementTotals is the per-category roll-up reported on the
// settlement receipt. Field names mirror the MX51 receipt codes
// (PUR, TIP, SUR, REF, CAS, TOT) so reconciliation against the printed
// receipt is direct.
type SettlementTotals struct {
	Currency        string `json:"currency"`
	PurchasesMinor  int64  `json:"purchases_minor,omitempty"`
	TipsMinor       int64  `json:"tips_minor,omitempty"`
	SurchargesMinor int64  `json:"surcharges_minor,omitempty"`
	RefundsMinor    int64  `json:"refunds_minor,omitempty"`
	CashoutsMinor   int64  `json:"cashouts_minor,omitempty"`
	TotalMinor      int64  `json:"total_minor"`
}

// CreateSettlementRequest triggers a settlement or enquiry on a
// paired terminal. EnquiryDate is required for an enquiry on a date
// other than today; omit it for "today's" enquiry or for a real
// settlement.
type CreateSettlementRequest struct {
	TerminalID     string               `json:"terminal_id"`
	Type           enums.SettlementType `json:"type"`
	EnquiryDate    *common.Date         `json:"enquiry_date,omitempty"`
	IdempotencyKey string               `json:"idempotency_key,omitempty"`
}

// PollSettlementRequest is the input to the long-poll GET on a
// settlement.
type PollSettlementRequest struct {
	SettlementID string `json:"settlement_id"`
	MinVersion   int    `json:"min_version"`
}
