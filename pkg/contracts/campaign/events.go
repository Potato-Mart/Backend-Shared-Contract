package campaign

import (
	"time"

	campaignenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/campaign"
)

// CampaignChangedEvent is the customer-safe storefront-events projection used
// only to invalidate/refetch authoritative campaign content. It contains no
// authored copy, targeting rules, provider destination, or customer data.
type CampaignChangedEvent struct {
	CampaignID         string                      `json:"campaign_id"`
	CampaignKey        string                      `json:"campaign_key"`
	PromotionID        string                      `json:"promotion_id,omitempty"`
	Status             campaignenum.CampaignStatus `json:"status"`
	IsActive           bool                        `json:"is_active"`
	ActivationRevision int64                       `json:"activation_revision"`
	ContentRevision    int64                       `json:"content_revision"`
	ChangedAt          time.Time                   `json:"changed_at"`
}
