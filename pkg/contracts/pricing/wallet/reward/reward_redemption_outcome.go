package reward

import "github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/common/money"

// RewardRedemptionOutcome records what one reward redemption actually issued.
// Only the arm matching the redeemed reward's type is set: a checkout discount
// amount, the issued coupon or gift-card business code, or the
// external partner fulfilment evidence. The owning service enforces that
// invariant, not the model.
type RewardRedemptionOutcome struct {
	DiscountAmount *money.Money              `json:"discount_amount,omitempty"`
	CouponCode     string                    `json:"coupon_code,omitempty"`
	GiftCardCode   string                    `json:"gift_card_code,omitempty"`
	External       *ExternalRewardFulfilment `json:"external,omitempty"`
}
