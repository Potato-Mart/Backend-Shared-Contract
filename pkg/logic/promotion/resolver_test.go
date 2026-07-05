package promotionlogic

import (
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/enums"
)

var resolverNow = time.Date(2026, 6, 1, 12, 0, 0, 0, time.UTC)

func targetedPromo(id string, class enums.PromotionClass, scope enums.DiscountScope, ref string) promotion.Promotion {
	p := promotion.Promotion{
		ID:    id,
		Name:  id,
		Type:  enums.PromotionTypeAutoDiscount,
		Class: class,
		DiscountSpec: promotion.DiscountSpec{
			DiscountType:  enums.DiscountTypePercentage,
			DiscountValue: "10",
		},
		TargetScope: scope,
	}
	p.IsActive = true
	switch scope {
	case enums.DiscountScopeProduct:
		p.TargetProductSKUCode = ref
	case enums.DiscountScopeCategoryTag:
		p.TargetCategoryTagID = ref
		p.TargetCategoryTagName = localizedName(ref + " name")
	}
	return p
}

func localizedName(name string) []common.LocalizedName {
	return []common.LocalizedName{{Language: "en", Name: name}}
}

func porkBelly() ResolveTarget {
	return ResolveTarget{
		ProductSKUCode: "PORK-BELLY-001",
		CategoryTags: []product.CategoryTag{
			{ID: "tag_pork", Name: localizedName("Pork"), CollectionID: "col_meat", CollectionName: localizedName("Meat")},
			{ID: "tag_hotpot", Name: localizedName("Hotpot"), CollectionID: "col_featured", CollectionName: localizedName("Featured")},
		},
		UnitPriceMinor: 1000,
		Currency:       "AUD",
		Now:            resolverNow,
	}
}

// 1. NORMAL_PROMOTION applies when no SPECIAL_CAMPAIGN exists.
func TestResolveNormalAppliesWithoutSpecial(t *testing.T) {
	active := []promotion.Promotion{targetedPromo("prm_normal", enums.PromotionClassNormal, enums.DiscountScopeProduct, "PORK-BELLY-001")}
	eff := ResolveEffective(active, porkBelly())
	if eff == nil || eff.PromotionID != "prm_normal" {
		t.Fatalf("effective = %+v, want prm_normal", eff)
	}
	if eff.Class != enums.PromotionClassNormal {
		t.Errorf("class = %s, want normal_promotion", eff.Class)
	}
	if eff.DiscountedPriceMinor != 900 {
		t.Errorf("discounted = %d, want 900", eff.DiscountedPriceMinor)
	}
	if eff.OverrideReason != "" {
		t.Errorf("override reason = %q, want empty", eff.OverrideReason)
	}
}

// 2. SPECIAL_CAMPAIGN overrides NORMAL_PROMOTION for the same product.
func TestResolveSpecialOverridesNormalSameProduct(t *testing.T) {
	active := []promotion.Promotion{
		targetedPromo("prm_normal", enums.PromotionClassNormal, enums.DiscountScopeProduct, "PORK-BELLY-001"),
		targetedPromo("prm_special", enums.PromotionClassSpecialCampaign, enums.DiscountScopeProduct, "PORK-BELLY-001"),
	}
	eff := ResolveEffective(active, porkBelly())
	if eff == nil || eff.PromotionID != "prm_special" {
		t.Fatalf("effective = %+v, want prm_special", eff)
	}
	if eff.OverrideReason != OverrideReasonSpecialCampaign {
		t.Errorf("override reason = %q, want %q", eff.OverrideReason, OverrideReasonSpecialCampaign)
	}
	if eff.OverriddenPromotionID != "prm_normal" {
		t.Errorf("overridden id = %q, want prm_normal", eff.OverriddenPromotionID)
	}
}

// 3. SPECIAL_CAMPAIGN overrides NORMAL_PROMOTION for the same category tag.
func TestResolveSpecialOverridesNormalSameCategoryTag(t *testing.T) {
	active := []promotion.Promotion{
		targetedPromo("prm_normal_tag", enums.PromotionClassNormal, enums.DiscountScopeCategoryTag, "tag_pork"),
		targetedPromo("prm_special_tag", enums.PromotionClassSpecialCampaign, enums.DiscountScopeCategoryTag, "tag_pork"),
	}
	eff := ResolveEffective(active, porkBelly())
	if eff == nil || eff.PromotionID != "prm_special_tag" {
		t.Fatalf("effective = %+v, want prm_special_tag", eff)
	}
	if eff.OverriddenPromotionID != "prm_normal_tag" {
		t.Errorf("overridden id = %q, want prm_normal_tag", eff.OverriddenPromotionID)
	}
}

// 4. Product-level promotion beats category-tag-level within the same class.
func TestResolveProductBeatsCategoryTagWithinClass(t *testing.T) {
	active := []promotion.Promotion{
		targetedPromo("prm_tag", enums.PromotionClassNormal, enums.DiscountScopeCategoryTag, "tag_pork"),
		targetedPromo("prm_prod", enums.PromotionClassNormal, enums.DiscountScopeProduct, "PORK-BELLY-001"),
	}
	eff := ResolveEffective(active, porkBelly())
	if eff == nil || eff.PromotionID != "prm_prod" {
		t.Fatalf("effective = %+v, want prm_prod", eff)
	}
	// And a category-tag special_campaign still beats a product normal_promotion.
	active = append(active, targetedPromo("prm_tag_special", enums.PromotionClassSpecialCampaign, enums.DiscountScopeCategoryTag, "tag_pork"))
	eff = ResolveEffective(active, porkBelly())
	if eff == nil || eff.PromotionID != "prm_tag_special" {
		t.Fatalf("effective = %+v, want prm_tag_special", eff)
	}
}

// 5. Category-tag promotion applies to products with the target tag ID.
func TestResolveCategoryTagPromotionAppliesToMember(t *testing.T) {
	active := []promotion.Promotion{targetedPromo("prm_tag", enums.PromotionClassNormal, enums.DiscountScopeCategoryTag, "tag_pork")}
	if eff := ResolveEffective(active, porkBelly()); eff == nil || eff.PromotionID != "prm_tag" {
		t.Fatalf("effective = %+v, want prm_tag", eff)
	}
	other := porkBelly()
	other.CategoryTags = []product.CategoryTag{{ID: "tag_veg", Name: localizedName("Vegetables"), CollectionID: "col_produce", CollectionName: localizedName("Produce")}}
	if eff := ResolveEffective(active, other); eff != nil {
		t.Fatalf("effective = %+v for non-member, want nil", eff)
	}
}

// 6. Untargeted (cart-wide / SKU-less) promotions never match the
// product price resolver — there is no SKU-level targeting.
func TestResolveIgnoresUntargetedPromotions(t *testing.T) {
	p := targetedPromo("prm_cartwide", enums.PromotionClassNormal, "", "")
	p.TargetScope = enums.DiscountScopeAll
	if eff := ResolveEffective([]promotion.Promotion{p}, porkBelly()); eff != nil {
		t.Fatalf("effective = %+v for untargeted promotion, want nil", eff)
	}
}

// Window edges: inactive, not-yet-started, and expired promotions never match.
func TestResolveRespectsActiveWindow(t *testing.T) {
	p := targetedPromo("prm_window", enums.PromotionClassSpecialCampaign, enums.DiscountScopeProduct, "PORK-BELLY-001")
	later := resolverNow.Add(time.Hour)
	p.StartsAt = &later
	if eff := ResolveEffective([]promotion.Promotion{p}, porkBelly()); eff != nil {
		t.Fatalf("effective = %+v before starts_at, want nil", eff)
	}
	earlier := resolverNow.Add(-2 * time.Hour)
	ended := resolverNow.Add(-time.Hour)
	p.StartsAt, p.ExpiresAt = &earlier, &ended
	if eff := ResolveEffective([]promotion.Promotion{p}, porkBelly()); eff != nil {
		t.Fatalf("effective = %+v after expires_at, want nil", eff)
	}
	p.StartsAt, p.ExpiresAt = &earlier, nil
	p.IsActive = false
	if eff := ResolveEffective([]promotion.Promotion{p}, porkBelly()); eff != nil {
		t.Fatalf("effective = %+v for unpublished promotion, want nil", eff)
	}
}

// Tie-break inside one tier: higher priority wins.
func TestResolvePriorityTieBreak(t *testing.T) {
	a := targetedPromo("prm_a", enums.PromotionClassNormal, enums.DiscountScopeProduct, "PORK-BELLY-001")
	b := targetedPromo("prm_b", enums.PromotionClassNormal, enums.DiscountScopeProduct, "PORK-BELLY-001")
	b.Priority = 5
	if eff := ResolveEffective([]promotion.Promotion{a, b}, porkBelly()); eff == nil || eff.PromotionID != "prm_b" {
		t.Fatalf("effective = %+v, want prm_b (priority 5)", eff)
	}
}

// Discount maths: fixed_amount, fixed_price, and floor at zero.
func TestDiscountedUnitPriceMinor(t *testing.T) {
	cases := []struct {
		name  string
		dtype enums.DiscountType
		value string
		want  int64
		ok    bool
	}{
		{"percentage", enums.DiscountTypePercentage, "25", 750, true},
		{"fixed_amount", enums.DiscountTypeFixedAmount, "2.50", 750, true},
		{"fixed_amount_floor", enums.DiscountTypeFixedAmount, "99", 0, true},
		{"fixed_price", enums.DiscountTypeFixedPrice, "7.50", 750, true},
		{"fixed_price_above_original_clamped", enums.DiscountTypeFixedPrice, "99", 1000, true},
		{"free_shipping_not_priceable", enums.DiscountTypeFreeShipping, "", 0, false},
		{"malformed", enums.DiscountTypePercentage, "ten", 0, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := promotion.Promotion{DiscountSpec: promotion.DiscountSpec{DiscountType: tc.dtype, DiscountValue: tc.value}}
			got, ok := DiscountedUnitPriceMinor(p, 1000)
			if ok != tc.ok || (ok && got != tc.want) {
				t.Errorf("got (%d,%v), want (%d,%v)", got, ok, tc.want, tc.ok)
			}
		})
	}
}
