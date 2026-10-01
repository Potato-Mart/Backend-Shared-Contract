package coupon

import "github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/pricing/coupon/coupon_enums"

// TierRestriction is the redemption-time membership restriction, independent
// of receiving eligibility. The service validates active tier references.
type TierRestriction struct {
	Active  bool                   `json:"active"`
	TierKey string                 `json:"tier_key,omitempty"`
	Match   coupon_enums.TierMatch `json:"match,omitempty"`
}
