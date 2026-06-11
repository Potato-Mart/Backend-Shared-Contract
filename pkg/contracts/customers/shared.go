package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v4/pkg/common"
)

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
