package promotion

import (
	"math"
	"strconv"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/enums"
)

// OverrideReasonSpecialCampaign is the canonical reason string written
// onto an EffectivePromotion when a special_campaign displaced an
// otherwise-applicable normal_promotion.
const OverrideReasonSpecialCampaign = "special_campaign overrides normal_promotion"

// EffectivePromotion is the result of resolving the single promotion
// that prices one product at one instant. Targeted promotions never
// stack: exactly one wins (or none).
type EffectivePromotion struct {
	PromotionID   string               `json:"promotion_id"`
	PromotionName string               `json:"promotion_name,omitempty"`
	Class         enums.PromotionClass `json:"class"`
	TargetScope   enums.DiscountScope  `json:"target_scope"`

	OriginalPriceMinor   int64  `json:"original_price_minor"`
	DiscountedPriceMinor int64  `json:"discounted_price_minor"`
	Currency             string `json:"currency,omitempty"`

	StartsAt  *time.Time `json:"starts_at,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`

	// OverrideReason is set when a special_campaign displaced a
	// normal_promotion that also matched; OverriddenPromotionID names
	// the displaced promotion.
	OverrideReason        string `json:"override_reason,omitempty"`
	OverriddenPromotionID string `json:"overridden_promotion_id,omitempty"`
}

// ResolveTarget describes the product being priced: its id, its
// canonical category path (root→leaf category keys, leaf last), the
// undiscounted unit price in minor units, and the pricing instant.
type ResolveTarget struct {
	ProductID      string
	CategoryPath   []string
	UnitPriceMinor int64
	Currency       string
	Now            time.Time
}

// precedence tiers, lower wins. Product beats category within a class;
// special_campaign beats normal_promotion across classes:
//
//	0: product-level special_campaign
//	1: category-level special_campaign
//	2: product-level normal_promotion
//	3: category-level normal_promotion
func tierOf(p Promotion) int {
	t := 0
	if p.EffectiveClass() != enums.PromotionClassSpecialCampaign {
		t += 2
	}
	if p.TargetScope == enums.DiscountScopeCategory {
		t++
	}
	return t
}

// Matches reports whether a targeted promotion applies to the given
// product at the given instant. Untargeted (cart-wide) promotions never
// match here — they stay on the cart-quote path.
func Matches(p Promotion, t ResolveTarget) bool {
	if !p.IsTargeted() {
		return false
	}
	if !p.IsActive {
		return false
	}
	if p.StartsAt != nil && t.Now.Before(*p.StartsAt) {
		return false
	}
	if p.ExpiresAt != nil && t.Now.After(*p.ExpiresAt) {
		return false
	}
	switch p.TargetScope {
	case enums.DiscountScopeProduct:
		return p.TargetProductID != "" && p.TargetProductID == t.ProductID
	case enums.DiscountScopeCategory:
		if p.TargetCategoryKey == "" || len(t.CategoryPath) == 0 {
			return false
		}
		if p.IncludesDescendants() {
			for _, key := range t.CategoryPath {
				if key == p.TargetCategoryKey {
					return true
				}
			}
			return false
		}
		return t.CategoryPath[len(t.CategoryPath)-1] == p.TargetCategoryKey
	}
	return false
}

// DiscountedUnitPriceMinor applies the promotion's DiscountSpec to a
// unit price. The bool result is false when the spec cannot price a
// single unit (free_shipping, malformed value) — such promotions are
// skipped by the resolver.
func DiscountedUnitPriceMinor(p Promotion, unitPriceMinor int64) (int64, bool) {
	switch p.DiscountType {
	case enums.DiscountTypePercentage:
		pct, err := strconv.ParseFloat(p.DiscountValue, 64)
		if err != nil || pct < 0 || pct > 100 {
			return 0, false
		}
		d := int64(math.Round(float64(unitPriceMinor) * pct / 100))
		if p.MaxDiscount != nil && d > p.MaxDiscount.AmountMinor {
			d = p.MaxDiscount.AmountMinor
		}
		return clampPrice(unitPriceMinor-d, unitPriceMinor), true
	case enums.DiscountTypeFixedAmount:
		major, err := strconv.ParseFloat(p.DiscountValue, 64)
		if err != nil || major < 0 {
			return 0, false
		}
		return clampPrice(unitPriceMinor-int64(math.Round(major*100)), unitPriceMinor), true
	case enums.DiscountTypeFixedPrice:
		major, err := strconv.ParseFloat(p.DiscountValue, 64)
		if err != nil || major < 0 {
			return 0, false
		}
		return clampPrice(int64(math.Round(major*100)), unitPriceMinor), true
	}
	return 0, false
}

func clampPrice(price, original int64) int64 {
	if price < 0 {
		return 0
	}
	if price > original {
		return original
	}
	return price
}

// ResolveEffective picks the single promotion that prices the product,
// applying the fixed precedence: product special_campaign > category
// special_campaign > product normal_promotion > category
// normal_promotion > none. Ties inside a tier break on Priority desc,
// then CreatedAt desc, then ID asc — the same ordering the promotion
// repository already uses.
//
// Targeted promotions never stack with each other; whether the winner
// additionally combines with cart-level promotions is the caller's
// concern (the cart quote path keeps its existing stacking rules).
func ResolveEffective(active []Promotion, t ResolveTarget) *EffectivePromotion {
	var winner, bestNormal *Promotion
	var winnerPrice, bestNormalPrice int64

	for i := range active {
		p := active[i]
		if !Matches(p, t) {
			continue
		}
		price, ok := DiscountedUnitPriceMinor(p, t.UnitPriceMinor)
		if !ok {
			continue
		}
		if p.EffectiveClass() == enums.PromotionClassNormal {
			if bestNormal == nil || beats(p, *bestNormal) {
				cp := p
				bestNormal, bestNormalPrice = &cp, price
			}
		}
		if winner == nil || beats(p, *winner) {
			cp := p
			winner, winnerPrice = &cp, price
		}
	}
	if winner == nil {
		return nil
	}

	eff := &EffectivePromotion{
		PromotionID:          winner.ID,
		PromotionName:        winner.Name,
		Class:                winner.EffectiveClass(),
		TargetScope:          winner.TargetScope,
		OriginalPriceMinor:   t.UnitPriceMinor,
		DiscountedPriceMinor: winnerPrice,
		Currency:             t.Currency,
		StartsAt:             winner.StartsAt,
		ExpiresAt:            winner.ExpiresAt,
	}
	if winner.EffectiveClass() == enums.PromotionClassSpecialCampaign && bestNormal != nil {
		eff.OverrideReason = OverrideReasonSpecialCampaign
		eff.OverriddenPromotionID = bestNormal.ID
		_ = bestNormalPrice // displaced price intentionally not exposed
	}
	return eff
}

// beats reports whether a should be preferred over b.
func beats(a, b Promotion) bool {
	ta, tb := tierOf(a), tierOf(b)
	if ta != tb {
		return ta < tb
	}
	if a.Priority != b.Priority {
		return a.Priority > b.Priority
	}
	if !a.CreatedAt.Equal(b.CreatedAt) {
		return a.CreatedAt.After(b.CreatedAt)
	}
	return a.ID < b.ID
}
