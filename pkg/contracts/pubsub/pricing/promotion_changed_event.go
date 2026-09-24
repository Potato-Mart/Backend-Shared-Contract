package pricing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/pubsub/pricing/promotion_enums"
)

// PromotionChangedEvent is the customer-safe storefront-events projection
// emitted after publish or unpublish. Consumers refetch authoritative pricing
// and linked campaign content; rule-engine internals never enter the event.
// For event v3, PublicationContext is required and valid. Event v2 producers
// must omit it; consumers retain existing v2 behavior when it is absent.
type PromotionChangedEvent struct {
	PromotionID        string                                      `json:"promotion_id"`
	SeriesCode         string                                      `json:"series_code,omitempty"`
	CampaignCode       string                                      `json:"campaign_code,omitempty"`
	ScopeMode          geography_enums.GeographicScopeMode         `json:"scope_mode"`
	ScopeRevision      int64                                       `json:"scope_revision"`
	Published          bool                                        `json:"published"`
	Revision           int64                                       `json:"revision"`
	PublicationContext promotion_enums.PromotionPublicationContext `json:"publication_context,omitempty"`
	RefetchRequired    bool                                        `json:"refetch_required"`
	ChangedAt          time.Time                                   `json:"changed_at"`
}
