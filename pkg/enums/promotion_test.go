package enums_test

import (
	"testing"

	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/promotion"
)

func TestPromotionEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "promotionenum.CouponAppliesTo", valid: []stringEnum{promotionenum.CouponAppliesToAll, promotionenum.CouponAppliesToSpecificProducts, promotionenum.CouponAppliesToSpecificCategoryTags}, invalid: promotionenum.CouponAppliesTo("__invalid__")},
		{name: "promotionenum.CouponSource", valid: []stringEnum{promotionenum.CouponSourceManual, promotionenum.CouponSourceRFMComeback, promotionenum.CouponSourceBirthday, promotionenum.CouponSourceReferral, promotionenum.CouponSourceSignupBonus, promotionenum.CouponSourceCampaign}, invalid: promotionenum.CouponSource("__invalid__")},
		{name: "promotionenum.DiscountScope", valid: []stringEnum{promotionenum.DiscountScopeAll, promotionenum.DiscountScopeCategoryTag, promotionenum.DiscountScopeProduct}, invalid: promotionenum.DiscountScope("__invalid__")},
		{name: "promotionenum.DiscountType", valid: []stringEnum{promotionenum.DiscountTypePercentage, promotionenum.DiscountTypeFixedAmount, promotionenum.DiscountTypeFreeShipping, promotionenum.DiscountTypeFixedPrice}, invalid: promotionenum.DiscountType("__invalid__")},
		{name: "promotionenum.PromotionAddonTrigger", valid: []stringEnum{promotionenum.PromotionAddonTriggerAmount, promotionenum.PromotionAddonTriggerRequiredProducts}, invalid: promotionenum.PromotionAddonTrigger("__invalid__")},
		{name: "promotionenum.PromotionClass", valid: []stringEnum{promotionenum.PromotionClassNormal, promotionenum.PromotionClassSpecialCampaign}, invalid: promotionenum.PromotionClass("__invalid__")},
		{name: "promotionenum.PromotionDiscountTarget", valid: []stringEnum{promotionenum.PromotionDiscountTargetCart, promotionenum.PromotionDiscountTargetRequiredItems}, invalid: promotionenum.PromotionDiscountTarget("__invalid__")},
		{name: "promotionenum.PromotionQtyMode", valid: []stringEnum{promotionenum.PromotionQtyModePerProduct, promotionenum.PromotionQtyModeCombined}, invalid: promotionenum.PromotionQtyMode("__invalid__")},
		{name: "promotionenum.PromotionType", valid: []stringEnum{promotionenum.PromotionTypeAutoDiscount, promotionenum.PromotionTypeSpendGift, promotionenum.PromotionTypeAddonPurchase, promotionenum.PromotionTypeBOGO, promotionenum.PromotionTypeBundle, promotionenum.PromotionTypeTieredPricing}, invalid: promotionenum.PromotionType("__invalid__")},
		{name: "promotionenum.VolumeDiscountAppliesTo", valid: []stringEnum{promotionenum.VolumeDiscountAppliesToAll, promotionenum.VolumeDiscountAppliesToWholesale, promotionenum.VolumeDiscountAppliesToRetail}, invalid: promotionenum.VolumeDiscountAppliesTo("__invalid__")},
	})
}
