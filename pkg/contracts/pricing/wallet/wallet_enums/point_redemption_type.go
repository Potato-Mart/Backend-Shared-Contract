package wallet_enums

// PointRedemptionType identifies the operational use of a points reservation.
type PointRedemptionType string

const (
	PointRedemptionTypeCheckoutDiscount PointRedemptionType = "CHECKOUT_DISCOUNT"
	PointRedemptionTypeRewardCatalog    PointRedemptionType = "REWARD_CATALOG"
)

func (t PointRedemptionType) IsValid() bool {
	return t == PointRedemptionTypeCheckoutDiscount || t == PointRedemptionTypeRewardCatalog
}
func (t PointRedemptionType) String() string { return string(t) }
