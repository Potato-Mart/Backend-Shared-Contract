package promotion

import "time"

// PromotionChangedEvent is the customer-safe storefront-events projection
// emitted after publish or unpublish. Consumers refetch authoritative pricing
// and linked campaign content; rule-engine internals never enter the event.
type PromotionChangedEvent struct {
	PromotionID string    `json:"promotion_id"`
	CampaignID  string    `json:"campaign_id,omitempty"`
	CampaignKey string    `json:"campaign_key,omitempty"`
	Published   bool      `json:"published"`
	Revision    int64     `json:"revision"`
	ChangedAt   time.Time `json:"changed_at"`
}
