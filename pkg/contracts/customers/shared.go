package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/contracts/membership"
	customerenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/customer"
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
	RecencyDays   *int                   `json:"recency_days,omitempty"`
	R             *int                   `json:"r,omitempty"`
	F             *int                   `json:"f,omitempty"`
	M             *int                   `json:"m,omitempty"`
	Score         string                 `json:"score,omitempty"`
	Segment       string                 `json:"segment,omitempty"`
	ChurnRisk     customerenum.ChurnRisk `json:"churn_risk,omitempty"`
	AvgRepeatDays *float64               `json:"avg_repeat_days,omitempty"`
}

// RetailCustomerProfileCompletion is a computed read projection describing how
// complete a retail customer's profile is.
type RetailCustomerProfileCompletion struct {
	Percent         int      `json:"percent"`
	CompletedFields []string `json:"completed_fields,omitempty"`
	MissingFields   []string `json:"missing_fields,omitempty"`
}

// RetailCustomerReferralProfile groups the referral-programme state of a
// retail customer.
type RetailCustomerReferralProfile struct {
	Code                      string     `json:"code,omitempty"`
	ReferrerCustomerNumber    string     `json:"referrer_customer_number,omitempty"`
	UsedReferralCodeConfirmed bool       `json:"used_referral_code_confirmed,omitempty"`
	UsedByCount               int        `json:"used_by_count,omitempty"`
	RewardVouchersIssued      int        `json:"reward_vouchers_issued,omitempty"`
	RewardVoucherCodes        []string   `json:"reward_voucher_codes,omitempty"`
	LastRewardIssuedAt        *time.Time `json:"last_reward_issued_at,omitempty"`
}

// RetailCustomerManagementProfile groups CRM fields that are manually edited
// by staff and are never overwritten by sync jobs.
type RetailCustomerManagementProfile struct {
	Notes    string   `json:"notes,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	SalesRep string   `json:"sales_rep,omitempty"`
	CRMTier  string   `json:"crm_tier,omitempty"`
}
