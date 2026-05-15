package payments

import (
	"encoding/json"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// Settlement is an end-of-day reconciliation/batch close
// (Type = SettlementTypeSettlement) or a read-only summary
// (Type = SettlementTypeEnquiry). For Adyen Terminal API this maps to
// ReconciliationRequest and GetTotals-style flows.
type Settlement struct {
	ID                   string                          `json:"id"`
	TerminalID           string                          `json:"terminal_id"`
	ProviderSettlementID string                          `json:"provider_settlement_id,omitempty"`
	MerchantAccount      string                          `json:"merchant_account,omitempty"`
	POIID                string                          `json:"poi_id,omitempty"`
	ServiceID            string                          `json:"service_id,omitempty"`
	SaleID               string                          `json:"sale_id,omitempty"`
	Type                 enums.SettlementType            `json:"type"`
	EnquiryDate          *common.Date                    `json:"enquiry_date,omitempty"`
	Status               enums.TerminalTxStatus          `json:"status"`
	FinancialStatus      enums.TerminalTxFinancialStatus `json:"financial_status,omitempty"`

	Totals          SettlementTotals `json:"totals"`
	Message         string           `json:"message,omitempty"`
	ProviderResult  string           `json:"provider_result,omitempty"`
	MerchantReceipt string           `json:"merchant_receipt,omitempty"`

	ProviderRequest      json.RawMessage `json:"provider_request,omitempty"`
	ProviderResponse     json.RawMessage `json:"provider_response,omitempty"`
	ProviderNotification json.RawMessage `json:"provider_notification,omitempty"`

	IdempotencyKey string `json:"idempotency_key,omitempty"`

	LastStatusCheckedAt *time.Time `json:"last_status_checked_at,omitempty"`
	FinalisedAt         *time.Time `json:"finalised_at,omitempty"`

	Metadata common.Metadata `json:"metadata,omitempty"`

	common.AuditFields
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

// CreateSettlementRequest triggers a settlement or enquiry on a
// registered terminal. EnquiryDate is required for an enquiry on a date
// other than today; omit it for "today's" enquiry or for a real
// settlement.
type CreateSettlementRequest struct {
	TerminalID     string                       `json:"terminal_id"`
	Type           enums.SettlementType         `json:"type"`
	EnquiryDate    *common.Date                 `json:"enquiry_date,omitempty"`
	ConnectionMode enums.TerminalConnectionMode `json:"connection_mode,omitempty"`
	ServiceID      string                       `json:"service_id,omitempty"`
	IdempotencyKey string                       `json:"idempotency_key,omitempty"`
}

// CheckSettlementStatusRequest asks the backend to refresh settlement
// state using the provider's status mechanism.
type CheckSettlementStatusRequest struct {
	SettlementID string `json:"settlement_id"`
}
