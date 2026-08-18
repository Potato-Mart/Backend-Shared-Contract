package marketing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/marketing/marketing_enums"
)

// CampaignChangedEvent is the minimal PII-free invalidation fact for a
// campaign aggregate. Consumers refetch public content from its authority.
type CampaignChangedEvent struct {
	CampaignCode string                         `json:"campaign_code"`
	Status       marketing_enums.CampaignStatus `json:"status"`
	ChangedAt    time.Time                      `json:"changed_at"`
}

// CouponChangedEvent is the minimal PII-free invalidation fact for a coupon
// aggregate.
type CouponChangedEvent struct {
	CouponCode string                       `json:"coupon_code"`
	Status     marketing_enums.CouponStatus `json:"status"`
	ChangedAt  time.Time                    `json:"changed_at"`
}

// PromotionChangedEvent is the minimal PII-free invalidation fact for a
// promotion aggregate.
type PromotionChangedEvent struct {
	PromotionCode string                          `json:"promotion_code"`
	Status        marketing_enums.PromotionStatus `json:"status"`
	ChangedAt     time.Time                       `json:"changed_at"`
}

// MarketingMessageChangedEvent is the minimal PII-free invalidation fact for
// a marketing message aggregate.
type MarketingMessageChangedEvent struct {
	Code         string                                 `json:"code"`
	CampaignCode string                                 `json:"campaign_code"`
	Status       marketing_enums.MarketingMessageStatus `json:"status"`
	ChangedAt    time.Time                              `json:"changed_at"`
}
