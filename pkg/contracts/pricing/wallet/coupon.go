package wallet

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/promotion"
)

// Coupon is a code-based discount that customers enter at checkout.
// Unlike Promotion (auto-applied rule), a coupon is manually redeemed.
type Coupon struct {
	ID      string        `json:"id"`
	Code    string        `json:"code"`
	Content CouponContent `json:"content"`

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
