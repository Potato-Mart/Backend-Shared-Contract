package coupon

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/geography"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/coupon/coupon_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/promotion"
)

// Coupon is a code-based discount that customers enter at checkout.
// Unlike Promotion (auto-applied rule), a coupon is manually redeemed.
type Coupon struct {
	// ID is immutable and identifies entitlement history even if Code changes.
	ID      string        `json:"id"`
	Code    string        `json:"code"`
	Content CouponContent `json:"content"`
	// Policy defaults, validation, eligibility and periodic issuance are service-owned.
	Distribution                 coupon_enums.Distribution    `json:"distribution"`
	Visibility                   coupon_enums.Visibility      `json:"visibility"`
	ReceivingTier                ReceivingTierPolicy          `json:"receiving_tier"`
	TierRestriction              TierRestriction              `json:"tier_restriction"`
	ProfileCompletionRestriction ProfileCompletionRestriction `json:"profile_completion_restriction"`

	Scope    promotion.PromotionScope    `json:"scope"`
	Period   promotion.PromotionPeriod   `json:"period"`
	Terms    []promotion.PromotionTerm   `json:"terms,omitempty"`
	Controls promotion.PromotionControls `json:"controls"`
	History  []security.HistoryEntry     `json:"history,omitempty"`
	// MarketCode and CountryCode are the denormalized owning market and its
	// country, carried so a geographically scoped staff query is a plain
	// indexed match.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`

	audit.AuditFields
}
