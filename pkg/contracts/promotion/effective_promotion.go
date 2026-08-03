package promotion

import (
	"time"

	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/promotion"
)

// EffectivePromotion is the result of resolving the single promotion that
// prices one product at one instant. Targeted promotions never stack: exactly
// one wins (or none). It is produced by promotionlogic.ResolveEffective.
type EffectivePromotion struct {
	PromotionID   string                       `json:"promotion_id"`
	PromotionName string                       `json:"promotion_name,omitempty"`
	Class         promotionenum.PromotionClass `json:"class"`
	TargetScope   promotionenum.DiscountScope  `json:"target_scope"`

	OriginalPriceMinor   int64  `json:"original_price_minor"`
	DiscountedPriceMinor int64  `json:"discounted_price_minor"`
	Currency             string `json:"currency,omitempty"`

	StartsAt  *time.Time `json:"starts_at,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`

	// OverrideReason is set when a special_campaign displaced a
	// normal_promotion that also matched; OverriddenPromotionID names the
	// displaced promotion.
	OverrideReason        string `json:"override_reason,omitempty"`
	OverriddenPromotionID string `json:"overridden_promotion_id,omitempty"`
}
