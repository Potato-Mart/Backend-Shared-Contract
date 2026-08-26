package wallet_enums

// RewardType describes the customer benefit a Membership reward can issue.
type RewardType string

const (
	RewardTypeOrderDiscount RewardType = "ORDER_DISCOUNT"
	RewardTypeProduct       RewardType = "PRODUCT"
	RewardTypeFreeShipping  RewardType = "FREE_SHIPPING"
	RewardTypeVoucher       RewardType = "VOUCHER"
	RewardTypeCoupon        RewardType = "COUPON"
	RewardTypeGiftCard      RewardType = "GIFT_CARD"
	// RewardTypeExternal is a redemption provisioned by a partner company, such as
	// another company's subscription or service.
	RewardTypeExternal RewardType = "EXTERNAL"
)

func (t RewardType) IsValid() bool {
	switch t {
	case RewardTypeOrderDiscount, RewardTypeProduct, RewardTypeFreeShipping, RewardTypeVoucher, RewardTypeCoupon, RewardTypeGiftCard, RewardTypeExternal:
		return true
	}
	return false
}
func (t RewardType) String() string { return string(t) }
