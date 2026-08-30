package coupon

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"

// CouponContent contains approved, locale-aware customer-facing coupon copy.
// Pricing owns this presentation because a coupon is a wallet instrument.
type CouponContent struct {
	Names        []localization.LocalizedName        `json:"names"`
	Descriptions []localization.LocalizedDescription `json:"descriptions,omitempty"`
}
