package pricebook_enums

// SellingPromotionDiscountType identifies a customer-visible discount value.
type SellingPromotionDiscountType string

const (
	SellingPromotionDiscountTypePercentOff SellingPromotionDiscountType = "percent_off"
	SellingPromotionDiscountTypeAmountOff  SellingPromotionDiscountType = "amount_off"
	SellingPromotionDiscountTypeFixedPrice SellingPromotionDiscountType = "fixed_price"
)

func (v SellingPromotionDiscountType) IsValid() bool {
	return v == SellingPromotionDiscountTypePercentOff || v == SellingPromotionDiscountTypeAmountOff || v == SellingPromotionDiscountTypeFixedPrice
}

func (v SellingPromotionDiscountType) String() string { return string(v) }
