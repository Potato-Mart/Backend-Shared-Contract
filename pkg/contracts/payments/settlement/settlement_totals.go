package settlement

import "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"

// SettlementTotals is the per-category roll-up reported on settlement
// or totals receipts. The fields are provider-neutral and stay in minor
// units for invoice/report consistency.
type SettlementTotals struct {
	Currency        money.CurrencyCode `json:"currency"`
	PurchasesMinor  int64              `json:"purchases_minor,omitempty"`
	TipsMinor       int64              `json:"tips_minor,omitempty"`
	SurchargesMinor int64              `json:"surcharges_minor,omitempty"`
	RefundsMinor    int64              `json:"refunds_minor,omitempty"`
	CashoutsMinor   int64              `json:"cashouts_minor,omitempty"`
	TotalMinor      int64              `json:"total_minor"`
}
