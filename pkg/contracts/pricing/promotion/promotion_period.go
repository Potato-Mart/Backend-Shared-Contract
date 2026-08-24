package promotion

import "time"

// PromotionPeriod is the scheduled period for a promotion.
type PromotionPeriod struct {
	StartsAt *time.Time `json:"starts_at,omitempty"`
	EndsAt   *time.Time `json:"ends_at,omitempty"`
	Timezone string     `json:"timezone"`
}
