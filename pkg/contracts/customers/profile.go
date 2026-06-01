package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// CustomerProfile is the unified CRM entity for wholesale and retail customers.
// It is separate from Customer (which is auth-linked B2C) and serves as the
// hub for loyalty, RFM analytics, marketing, subscriptions, and referrals.
// Uniqueness: (customer_name, phone, segment).
type CustomerProfile struct {
	ID           string                `json:"id"`
	CustomerName string                `json:"customer_name"`
	Phone        string                `json:"phone"`
	Segment      enums.CustomerSegment `json:"segment"`

	// ── Aggregated order statistics (synced, not manually edited) ─────
	OrderCount     int          `json:"order_count"`
	TotalUnits     int          `json:"total_units"`
	TotalSpend     common.Money `json:"total_spend"`
	AvgOrderValue  common.Money `json:"avg_order_value"`
	FirstOrderDate *time.Time   `json:"first_order_date,omitempty"`
	LastOrderDate  *time.Time   `json:"last_order_date,omitempty"`
	LastOrderTS    *time.Time   `json:"last_order_ts,omitempty"`
	Provinces      []string     `json:"provinces,omitempty"`
	Suburbs        []string     `json:"suburbs,omitempty"`
	StatsSyncedAt  *time.Time   `json:"stats_synced_at,omitempty"`

	// ── CRM fields (manually edited, never overwritten by sync) ───────
	Notes         string                      `json:"notes,omitempty"`
	Tags          []string                    `json:"tags,omitempty"`
	SalesRep      string                      `json:"sales_rep,omitempty"`
	Tier          string                      `json:"tier,omitempty"`          // CRM cooperation level: VIP / A / B / C
	PaymentTerms  string                      `json:"payment_terms,omitempty"` // e.g. "NET30"
	TaxID         string                      `json:"tax_id,omitempty"`        // ABN / tax number
	ContactPerson string                      `json:"contact_person,omitempty"`
	Email         string                      `json:"email,omitempty"`
	Status        enums.CustomerProfileStatus `json:"status"`
	Hidden        bool                        `json:"hidden,omitempty"`

	// ── B2B-specific fields ───────────────────────────────────────────
	NameEN          string          `json:"name_en,omitempty"`
	Address         string          `json:"address,omitempty"`
	ABN             string          `json:"abn,omitempty"`
	PriceTier       int             `json:"price_tier,omitempty"`
	PriceTierSea    *int            `json:"price_tier_sea,omitempty"`
	PriceTierAir    *int            `json:"price_tier_air,omitempty"`
	PriceTierRebate *float64        `json:"price_tier_rebate,omitempty"`
	ShippingFee     *common.Money   `json:"shipping_fee,omitempty"`
	WholesaleTier   string          `json:"wholesale_tier,omitempty"` // references wholesale_tier_presets.tier_key
	FreightRule     common.Metadata `json:"freight_rule,omitempty"`

	// ── RFM analytics (computed by sync job) ─────────────────────────
	RecencyDays   *int            `json:"recency_days,omitempty"`
	RFMR          *int            `json:"rfm_r,omitempty"`
	RFMF          *int            `json:"rfm_f,omitempty"`
	RFMM          *int            `json:"rfm_m,omitempty"`
	RFMScore      string          `json:"rfm_score,omitempty"`   // e.g. "545"
	RFMSegment    string          `json:"rfm_segment,omitempty"` // e.g. "VIP" / "沉睡"
	ChurnRisk     enums.ChurnRisk `json:"churn_risk,omitempty"`
	AvgRepeatDays *float64        `json:"avg_repeat_days,omitempty"`

	// ── Marketing consent ─────────────────────────────────────────────
	AcceptsEmailMarketing bool       `json:"accepts_email_marketing"`
	AcceptsSMSMarketing   bool       `json:"accepts_sms_marketing"`
	AcceptsLineMarketing  bool       `json:"accepts_line_marketing"`
	ConsentUpdatedAt      *time.Time `json:"consent_updated_at,omitempty"`
	ConsentSource         string     `json:"consent_source,omitempty"` // "website"|"pos"|"import"|"manual"

	// ── Loyalty ───────────────────────────────────────────────────────
	LoyaltyTierKey       string       `json:"loyalty_tier_key,omitempty"` // references loyalty_tiers.tier_key
	LoyaltyPoints        int          `json:"loyalty_points"`
	LifetimeLoyaltySpend common.Money `json:"lifetime_loyalty_spend"`
	Birthday             *time.Time   `json:"birthday,omitempty"`
	TierEvaluatedAt      *time.Time   `json:"tier_evaluated_at,omitempty"`

	// ── Referral ──────────────────────────────────────────────────────
	ReferralCode     string `json:"referral_code,omitempty"`
	ReferrerID       string `json:"referrer_id,omitempty"`
	ReferralCredited bool   `json:"referral_credited,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	common.DataProtectionFields
}
