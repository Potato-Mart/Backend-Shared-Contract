package promotion

import (
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/geography"
	"time"
)

// EffectivePromotion is the result of resolving the single promotion that
// prices one product at one instant.
type EffectivePromotion struct {
	PromotionID   string         `json:"promotion_id"`
	SeriesKey     string         `json:"series_key"`
	PromotionName string         `json:"promotion_name,omitempty"`
	Class         PromotionClass `json:"class"`
	TargetScope   DiscountScope  `json:"target_scope"`

	OriginalPriceMinor   int64  `json:"original_price_minor"`
	DiscountedPriceMinor int64  `json:"discounted_price_minor"`
	Currency             string `json:"currency,omitempty"`

	StartsAt          *time.Time                  `json:"starts_at,omitempty"`
	ExpiresAt         *time.Time                  `json:"expires_at,omitempty"`
	ScheduleTimezone  string                      `json:"schedule_timezone"`
	GeographicContext geography.GeographicContext `json:"geographic_context"`

	// OverrideReason is set when a special_campaign displaced a
	// normal_promotion that also matched; OverriddenPromotionID names the
	// displaced promotion.
	OverrideReason        string `json:"override_reason,omitempty"`
	OverriddenPromotionID string `json:"overridden_promotion_id,omitempty"`
}
