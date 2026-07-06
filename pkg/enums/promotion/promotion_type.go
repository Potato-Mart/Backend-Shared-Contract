package promotionenum

// PromotionType classifies which engine rule applies to a promotion.
type PromotionType string

const (
	PromotionTypeAutoDiscount  PromotionType = "auto_discount"
	PromotionTypeSpendGift     PromotionType = "spend_gift"
	PromotionTypeAddonPurchase PromotionType = "addon_purchase"
	PromotionTypeBOGO          PromotionType = "bogo"
	PromotionTypeBundle        PromotionType = "bundle"
	PromotionTypeTieredPricing PromotionType = "tiered_pricing"
)

func (p PromotionType) IsValid() bool {
	switch p {
	case PromotionTypeAutoDiscount, PromotionTypeSpendGift, PromotionTypeAddonPurchase,
		PromotionTypeBOGO, PromotionTypeBundle, PromotionTypeTieredPricing:
		return true
	}
	return false
}

func (p PromotionType) String() string { return string(p) }

// PromotionDiscountTarget controls which items the discount is applied to.
type PromotionDiscountTarget string

const (
	PromotionDiscountTargetCart          PromotionDiscountTarget = "cart"
	PromotionDiscountTargetRequiredItems PromotionDiscountTarget = "required_items"
)

func (p PromotionDiscountTarget) IsValid() bool {
	switch p {
	case PromotionDiscountTargetCart, PromotionDiscountTargetRequiredItems:
		return true
	}
	return false
}

func (p PromotionDiscountTarget) String() string { return string(p) }

// PromotionQtyMode specifies whether required_qty_each applies per product
// or as a combined total across all required products.
type PromotionQtyMode string

const (
	PromotionQtyModePerProduct PromotionQtyMode = "per_product"
	PromotionQtyModeCombined   PromotionQtyMode = "combined"
)

func (p PromotionQtyMode) IsValid() bool {
	switch p {
	case PromotionQtyModePerProduct, PromotionQtyModeCombined:
		return true
	}
	return false
}

func (p PromotionQtyMode) String() string { return string(p) }

// PromotionAddonTrigger determines how an addon_purchase promotion is unlocked.
type PromotionAddonTrigger string

const (
	PromotionAddonTriggerAmount           PromotionAddonTrigger = "amount"
	PromotionAddonTriggerRequiredProducts PromotionAddonTrigger = "required_products"
)

func (p PromotionAddonTrigger) IsValid() bool {
	switch p {
	case PromotionAddonTriggerAmount, PromotionAddonTriggerRequiredProducts:
		return true
	}
	return false
}

func (p PromotionAddonTrigger) String() string { return string(p) }
