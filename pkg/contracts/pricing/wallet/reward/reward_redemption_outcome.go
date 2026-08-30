package reward

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"

// RewardRedemptionOutcome records what one reward redemption actually issued.
// Only the arm matching the redeemed reward's type is set: a checkout discount
// amount, the issued voucher, coupon, or gift-card business code, or the
// external partner fulfilment evidence. The owning service enforces that
// invariant, not the model.
type RewardRedemptionOutcome struct {
	DiscountAmount *money.Money              `json:"discount_amount,omitempty"`
	VoucherCode    string                    `json:"voucher_code,omitempty"`
	CouponCode     string                    `json:"coupon_code,omitempty"`
	GiftCardCode   string                    `json:"gift_card_code,omitempty"`
	External       *ExternalRewardFulfilment `json:"external,omitempty"`
}
