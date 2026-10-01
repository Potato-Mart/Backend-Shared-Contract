package coupon

import "github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/pricing/coupon/coupon_enums"

// ReceivingTierPolicy applies when issuing an entitlement. TierKey and Match
// apply to SELECTED_TIER; defaults, tier existence and ordering are service-owned.
type ReceivingTierPolicy struct {
	Specification coupon_enums.TierSpecification `json:"specification"`
	TierKey       string                         `json:"tier_key,omitempty"`
	Match         coupon_enums.TierMatch         `json:"match,omitempty"`
}
