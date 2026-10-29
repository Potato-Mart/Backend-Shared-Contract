package coupon

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/coupon/coupon_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/promotion"
)

// Coupon is a code-based discount that customers enter at checkout.
// Unlike Promotion (auto-applied rule), a coupon is manually redeemed.
type Coupon struct {
	// ID is immutable and identifies entitlement history even if Code changes.
	ID      string        `json:"id"`
	Code    string        `json:"code"`
	Content CouponContent `json:"content"`
	// Policy defaults, validation, eligibility and periodic issuance are service-owned.
	Acquisition           coupon_enums.Acquisition     `json:"acquisition"`
	AutomaticPolicy       coupon_enums.AutomaticPolicy `json:"automatic_policy,omitempty"`
	Visibility            coupon_enums.Visibility      `json:"visibility"`
	ReceivingEligibility  ReceivingEligibility         `json:"receiving_eligibility"`
	RedemptionEligibility RedemptionEligibility        `json:"redemption_eligibility"`
	AcquisitionWindow     promotion.PromotionPeriod    `json:"acquisition_window"`
	Inheritance           Inheritance                  `json:"inheritance"`
	IssuanceLimit         *int64                       `json:"issuance_limit,omitempty"`
	RedemptionLimits      RedemptionLimits             `json:"redemption_limits"`

	Scope            promotion.PromotionScope    `json:"scope"`
	RedemptionWindow promotion.PromotionPeriod   `json:"redemption_window"`
	Benefit          Benefit                     `json:"benefit"`
	MinimumSpend     *money.Money                `json:"minimum_spend,omitempty"`
	Controls         promotion.PromotionControls `json:"controls"`
	History          []security.HistoryEntry     `json:"history,omitempty"`
	// MarketCode and CountryCode are the denormalized owning market and its
	// country, carried so a geographically scoped staff query is a plain
	// indexed match.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`

	audit.AuditFields
}
