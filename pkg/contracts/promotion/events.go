package promotion

import (
	"time"

	geographyenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/geography"
)

// PromotionChangedEvent is the customer-safe storefront-events projection
// emitted after publish or unpublish. Consumers refetch authoritative pricing
// and linked campaign content; rule-engine internals never enter the event.
type PromotionChangedEvent struct {
	PromotionID     string                            `json:"promotion_id"`
	SeriesKey       string                            `json:"series_key,omitempty"`
	CampaignID      string                            `json:"campaign_id,omitempty"`
	CampaignKey     string                            `json:"campaign_key,omitempty"`
	ScopeMode       geographyenum.GeographicScopeMode `json:"scope_mode"`
	ScopeRevision   int64                             `json:"scope_revision"`
	Published       bool                              `json:"published"`
	Revision        int64                             `json:"revision"`
	RefetchRequired bool                              `json:"refetch_required"`
	ChangedAt       time.Time                         `json:"changed_at"`
}
