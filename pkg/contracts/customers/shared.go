package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// RetailCustomerMembershipProfile links a retail customer profile to the
// global membership programme. Wallet values are projections; membership
// ledger contracts remain the source of truth.
type RetailCustomerMembershipProfile struct {
	MembershipAccountID string                               `json:"membership_account_id,omitempty"`
	Summary             *membership.MembershipAccountSummary `json:"summary,omitempty"`
}

// RetailCustomerCommerceProfile groups aggregated commerce statistics. Values
// are computed by sync jobs and must never be manually edited.
type RetailCustomerCommerceProfile struct {
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

// RetailCustomerMarketingProfile groups per-channel marketing opt-ins together
// with the provenance of the consent decision.
type RetailCustomerMarketingProfile struct {
	EmailOptIn bool       `json:"email_opt_in"`
	SMSOptIn   bool       `json:"sms_opt_in"`
	LineOptIn  bool       `json:"line_opt_in"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	Source     string     `json:"source,omitempty"`
}

// RetailCustomerAnalyticsProfile groups the recency/frequency/monetary
// analytics computed by the stats sync job.
type RetailCustomerAnalyticsProfile struct {
	RecencyDays   *int            `json:"recency_days,omitempty"`
	R             *int            `json:"r,omitempty"`
	F             *int            `json:"f,omitempty"`
	M             *int            `json:"m,omitempty"`
	Score         string          `json:"score,omitempty"`
	Segment       string          `json:"segment,omitempty"`
	ChurnRisk     enums.ChurnRisk `json:"churn_risk,omitempty"`
	AvgRepeatDays *float64        `json:"avg_repeat_days,omitempty"`
}

// RetailCustomerReferralProfile groups the referral-programme state of a
// retail customer.
type RetailCustomerReferralProfile struct {
	Code       string `json:"code,omitempty"`
	ReferrerID string `json:"referrer_id,omitempty"`
	Credited   bool   `json:"credited,omitempty"`
}

// RetailCustomerManagementProfile groups CRM fields that are manually edited
// by staff and are never overwritten by sync jobs.
type RetailCustomerManagementProfile struct {
	Notes    string   `json:"notes,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	SalesRep string   `json:"sales_rep,omitempty"`
	CRMTier  string   `json:"crm_tier,omitempty"`
}
