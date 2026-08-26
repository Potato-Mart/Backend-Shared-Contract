package membership

import "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"

// RewardBenefit carries the type-specific configuration for one Reward. The
// populated arm must match Reward.Type: a discount amount or basis points for
// ORDER_DISCOUNT, SKUCode for PRODUCT, the voucher template for VOUCHER,
// CouponCode for COUPON, GiftCardValue for GIFT_CARD, and External for
// EXTERNAL. FREE_SHIPPING carries no configuration. The owning service
// enforces that invariant, not the model.
type RewardBenefit struct {
	DiscountAmount      *money.Money           `json:"discount_amount,omitempty"`
	DiscountBasisPoints *int64                 `json:"discount_basis_points,omitempty"`
	SKUCode             string                 `json:"sku_code,omitempty"`
	VoucherCodePrefix   string                 `json:"voucher_code_prefix,omitempty"`
	VoucherValue        *money.Money           `json:"voucher_value,omitempty"`
	CouponCode          string                 `json:"coupon_code,omitempty"`
	GiftCardValue       *money.Money           `json:"gift_card_value,omitempty"`
	External            *ExternalRewardBenefit `json:"external,omitempty"`
}
