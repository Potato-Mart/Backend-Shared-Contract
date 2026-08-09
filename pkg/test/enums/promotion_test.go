package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/promotion/promotion_enums"
)

func TestPromotionEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "promotionenum.DiscountScope", valid: []stringEnum{promotion_enums.DiscountScopeAll, promotion_enums.DiscountScopeCategoryTag, promotion_enums.DiscountScopeProduct}, invalid: promotion_enums.DiscountScope("__invalid__")},
		{name: "promotionenum.DiscountType", valid: []stringEnum{promotion_enums.DiscountTypePercentage, promotion_enums.DiscountTypeFixedAmount, promotion_enums.DiscountTypeFreeShipping, promotion_enums.DiscountTypeFixedPrice}, invalid: promotion_enums.DiscountType("__invalid__")},
		{name: "promotionenum.PromotionAddonTrigger", valid: []stringEnum{promotion_enums.PromotionAddonTriggerAmount, promotion_enums.PromotionAddonTriggerRequiredProducts}, invalid: promotion_enums.PromotionAddonTrigger("__invalid__")},
		{name: "promotionenum.PromotionClass", valid: []stringEnum{promotion_enums.PromotionClassNormal, promotion_enums.PromotionClassSpecialCampaign}, invalid: promotion_enums.PromotionClass("__invalid__")},
		{name: "promotionenum.PromotionDiscountTarget", valid: []stringEnum{promotion_enums.PromotionDiscountTargetCart, promotion_enums.PromotionDiscountTargetRequiredItems}, invalid: promotion_enums.PromotionDiscountTarget("__invalid__")},
		{name: "promotionenum.PromotionQtyMode", valid: []stringEnum{promotion_enums.PromotionQtyModePerProduct, promotion_enums.PromotionQtyModeCombined}, invalid: promotion_enums.PromotionQtyMode("__invalid__")},
		{name: "promotionenum.PromotionType", valid: []stringEnum{promotion_enums.PromotionTypeAutoDiscount, promotion_enums.PromotionTypeSpendGift, promotion_enums.PromotionTypeAddonPurchase, promotion_enums.PromotionTypeBOGO, promotion_enums.PromotionTypeBundle, promotion_enums.PromotionTypeTieredPricing}, invalid: promotion_enums.PromotionType("__invalid__")},
		{name: "promotionenum.VolumeDiscountAppliesTo", valid: []stringEnum{promotion_enums.VolumeDiscountAppliesToAll, promotion_enums.VolumeDiscountAppliesToWholesale, promotion_enums.VolumeDiscountAppliesToRetail}, invalid: promotion_enums.VolumeDiscountAppliesTo("__invalid__")},
	})
}
