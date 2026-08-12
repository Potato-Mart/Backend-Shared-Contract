package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/customers/campaign/campaign_enums"
)

// CampaignChangedEvent is the customer-safe storefront-events projection used
// only to invalidate/refetch authoritative campaign content. It contains no
// authored copy, targeting rules, provider destination, or customer data.
type CampaignChangedEvent struct {
	CampaignID         string                              `json:"campaign_id"`
	CampaignKey        string                              `json:"campaign_key"`
	SeriesKey          string                              `json:"series_key,omitempty"`
	PromotionID        string                              `json:"promotion_id,omitempty"`
	Status             campaign_enums.CampaignStatus       `json:"status"`
	IsActive           bool                                `json:"is_active"`
	ScopeMode          geography_enums.GeographicScopeMode `json:"scope_mode"`
	ScopeRevision      int64                               `json:"scope_revision"`
	ActivationRevision int64                               `json:"activation_revision"`
	ContentRevision    int64                               `json:"content_revision"`
	RefetchRequired    bool                                `json:"refetch_required"`
	ChangedAt          time.Time                           `json:"changed_at"`
}

// PromotionChangedEvent is the customer-safe storefront-events projection
// emitted after publish or unpublish. Consumers refetch authoritative pricing
// and linked campaign content; rule-engine internals never enter the event.
type PromotionChangedEvent struct {
	PromotionID     string                              `json:"promotion_id"`
	SeriesKey       string                              `json:"series_key,omitempty"`
	CampaignID      string                              `json:"campaign_id,omitempty"`
	CampaignKey     string                              `json:"campaign_key,omitempty"`
	ScopeMode       geography_enums.GeographicScopeMode `json:"scope_mode"`
	ScopeRevision   int64                               `json:"scope_revision"`
	Published       bool                                `json:"published"`
	Revision        int64                               `json:"revision"`
	RefetchRequired bool                                `json:"refetch_required"`
	ChangedAt       time.Time                           `json:"changed_at"`
}
