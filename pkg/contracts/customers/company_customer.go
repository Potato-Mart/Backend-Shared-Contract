package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// CompanyCustomer is the unified CRM entity for wholesale and retail
// customers. It is separate from Customer (which is auth-linked B2C) and
// serves as the hub for loyalty, RFM analytics, marketing, subscriptions,
// and referrals.
// Uniqueness: (customer_name, phone, segment).
type CompanyCustomer struct {
	ID            string                      `json:"id"`
	CustomerName  string                      `json:"customer_name"`
	NameEN        string                      `json:"name_en,omitempty"`
	Phone         string                      `json:"phone"`
	Email         string                      `json:"email,omitempty"`
	ContactPerson string                      `json:"contact_person,omitempty"`
	Address       string                      `json:"address,omitempty"`
	DateOfBirth   *time.Time                  `json:"date_of_birth,omitempty"` // contact's birth date, used for birthday bonus points
	Segment       enums.CustomerSegment       `json:"segment"`
	Status        enums.CustomerProfileStatus `json:"status"`
	Hidden        bool                        `json:"hidden,omitempty"`

	// ── CRM fields (manually edited, never overwritten by sync) ───────
	Notes        string   `json:"notes,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	SalesRep     string   `json:"sales_rep,omitempty"`
	CRMTier      string   `json:"crm_tier,omitempty"`      // CRM cooperation level: VIP / A / B / C (distinct from Loyalty.TierKey)
	PaymentTerms string   `json:"payment_terms,omitempty"` // e.g. "NET30"
	TaxID        string   `json:"tax_id,omitempty"`        // ABN / tax number

	// ── Grouped state shared with Customer ────────────────────────────
	Loyalty          LoyaltyStatus    `json:"loyalty"`
	OrderStats       OrderStats       `json:"order_stats"` // synced, not manually edited
	MarketingConsent MarketingConsent `json:"marketing_consent"`

	// ── B2B-specific ──────────────────────────────────────────────────
	Wholesale *WholesaleTerms `json:"wholesale,omitempty"`

	// ── RFM analytics (computed by sync job) ──────────────────────────
	RFM *RFMMetrics `json:"rfm,omitempty"`

	// ── Referral ──────────────────────────────────────────────────────
	Referral *Referral `json:"referral,omitempty"`

	common.AuditFields
	common.DataProtectionFields
}

// WholesaleTerms groups the B2B price-tier configuration and freight terms
// for wholesale customers.
type WholesaleTerms struct {
	TierKey         string          `json:"tier_key,omitempty"` // references wholesale_tier_presets.tier_key
	PriceTier       int             `json:"price_tier,omitempty"`
	PriceTierSea    *int            `json:"price_tier_sea,omitempty"`
	PriceTierAir    *int            `json:"price_tier_air,omitempty"`
	PriceTierRebate *float64        `json:"price_tier_rebate,omitempty"`
	ShippingFee     *common.Money   `json:"shipping_fee,omitempty"`
	FreightRule     common.Metadata `json:"freight_rule,omitempty"`
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
