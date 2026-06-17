package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/enums"
)

// CompanyCustomer is the unified CRM entity for wholesale and retail
// customers. It is separate from Customer (which is auth-linked B2C) and
// serves as the hub for loyalty, RFM analytics, marketing, subscriptions,
// and referrals.
// Uniqueness: (name, phone, segment).
type CompanyCustomer struct {
	common.PartyRef `bson:",inline"`      // id / name / phone / email
	AuthUserID      string                `json:"auth_user_id,omitempty"`
	NameEN          string                `json:"name_en,omitempty"`
	CustomerProfile CustomerProfile       `json:"customer_profile"`
	ContactPerson   string                `json:"contact_person,omitempty"`
	DateOfBirth     *time.Time            `json:"date_of_birth,omitempty"` // contact's birth date, used for birthday bonus points
	Source          enums.OrderType       `json:"source,omitempty"`        // acquisition channel: "online" | "pos" | "import"
	CompanyDetail   *common.CompanyDetail `json:"company_detail,omitempty"`

	// ── Grouped state shared with Customer ────────────────────────────
	CRM              CRM                     `json:"crm"`
	Loyalty          LoyaltyStatus           `json:"loyalty"`
	OrderStats       OrderStats              `json:"order_stats"` // synced, not manually edited
	MarketingConsent MarketingConsent        `json:"marketing_consent"`
	DefaultShipping  *common.ContactAddress  `json:"default_shipping,omitempty"`
	ShippingList     []common.ContactAddress `json:"shipping_list,omitempty"`

	// ── B2B-specific ──────────────────────────────────────────────────
	Wholesale *WholesaleTerms `json:"wholesale,omitempty"`

	// ── RFM analytics (computed by sync job) ──────────────────────────
	RFM *RFMMetrics `json:"rfm,omitempty"`

	// ── Referral ──────────────────────────────────────────────────────
	Referral *Referral             `json:"referral,omitempty"`
	History  []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields          `bson:",inline"`
	common.DataProtectionFields `bson:",inline"`
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
