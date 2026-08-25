package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography/geography_enums"
)

// PromotionChangedEvent is the customer-safe storefront-events projection
// emitted after publish or unpublish. Consumers refetch authoritative pricing
// and linked campaign content; rule-engine internals never enter the event.
type PromotionChangedEvent struct {
	PromotionID     string                              `json:"promotion_id"`
	SeriesCode      string                              `json:"series_code,omitempty"`
	CampaignCode    string                              `json:"campaign_code,omitempty"`
	ScopeMode       geography_enums.GeographicScopeMode `json:"scope_mode"`
	ScopeRevision   int64                               `json:"scope_revision"`
	Published       bool                                `json:"published"`
	Revision        int64                               `json:"revision"`
	RefetchRequired bool                                `json:"refetch_required"`
	ChangedAt       time.Time                           `json:"changed_at"`
}
