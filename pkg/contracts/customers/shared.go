package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/enums"
)

type CustomerProfile struct {
	Segment enums.CustomerSegment       `json:"segment"`
	Status  enums.CustomerProfileStatus `json:"status"`
}

// LoyaltyStatus groups the loyalty-programme state of a customer: points
// balance, membership tier, and the spend counters used for tier evaluation.
// Shared by Customer and CompanyCustomer.
type LoyaltyStatus struct {
	Points          int           `json:"points"`
	TierKey         string        `json:"tier_key,omitempty"`       // references loyalty_tiers.tier_key
	TierSpend       *common.Money `json:"tier_spend,omitempty"`     // spend counted toward the current tier window
	LifetimeSpend   *common.Money `json:"lifetime_spend,omitempty"` // lifetime spend recognised by the loyalty programme
	TierEvaluatedAt *time.Time    `json:"tier_evaluated_at,omitempty"`
}

// OrderStats groups aggregated order statistics for a customer. All values
// are computed by sync jobs and must never be manually edited.
type OrderStats struct {
	TotalOrders       int          `json:"total_orders"`
	TotalUnits        int          `json:"total_units,omitempty"`
	TotalSpend        common.Money `json:"total_spend"`
	AverageOrderValue common.Money `json:"average_order_value"`
	FirstOrderAt      *time.Time   `json:"first_order_at,omitempty"`
	LastOrderAt       *time.Time   `json:"last_order_at,omitempty"`
	Provinces         []string     `json:"provinces,omitempty"`
	Suburbs           []string     `json:"suburbs,omitempty"`
	SyncedAt          *time.Time   `json:"synced_at,omitempty"`
}

// MarketingConsent groups per-channel marketing opt-ins together with the
// provenance of the consent decision.
type MarketingConsent struct {
	Email     bool       `json:"email"`
	SMS       bool       `json:"sms"`
	Line      bool       `json:"line"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Source    string     `json:"source,omitempty"` // "website" | "pos" | "import" | "manual"
}

// RFMMetrics groups the recency/frequency/monetary analytics computed by
// the stats sync job.
type RFMMetrics struct {
	RecencyDays   *int            `json:"recency_days,omitempty"`
	R             *int            `json:"r,omitempty"`
	F             *int            `json:"f,omitempty"`
	M             *int            `json:"m,omitempty"`
	Score         string          `json:"score,omitempty"`   // e.g. "545"
	Segment       string          `json:"segment,omitempty"` // e.g. "VIP" / "沉睡"
	ChurnRisk     enums.ChurnRisk `json:"churn_risk,omitempty"`
	AvgRepeatDays *float64        `json:"avg_repeat_days,omitempty"`
}

// Referral groups the referral-programme state of a customer.
type Referral struct {
	Code       string `json:"code,omitempty"`
	ReferrerID string `json:"referrer_id,omitempty"`
	Credited   bool   `json:"credited,omitempty"`
}

// ── CRM fields (manually edited, never overwritten by sync) ───────
type CRM struct {
	Notes        string   `json:"notes,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	SalesRep     string   `json:"sales_rep,omitempty"`
	CRMTier      string   `json:"crm_tier,omitempty"`      // CRM cooperation level: VIP / A / B / C (distinct from Loyalty.TierKey)
	PaymentTerms string   `json:"payment_terms,omitempty"` // e.g. "NET30"
	TaxID        string   `json:"tax_id,omitempty"`        // ABN / tax number

}
