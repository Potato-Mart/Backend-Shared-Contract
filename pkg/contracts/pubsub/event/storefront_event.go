package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/marketing/campaign/campaign_enums"
)

// CampaignChangedEvent is the customer-safe storefront-events projection used
// only to invalidate/refetch authoritative campaign content. It contains no
// authored copy, targeting rules, provider destination, or customer data.
type CampaignChangedEvent struct {
	CampaignCode       string                              `json:"campaign_code"`
	SeriesCode         string                              `json:"series_code,omitempty"`
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
	SeriesCode      string                              `json:"series_code,omitempty"`
	CampaignCode    string                              `json:"campaign_code,omitempty"`
	ScopeMode       geography_enums.GeographicScopeMode `json:"scope_mode"`
	ScopeRevision   int64                               `json:"scope_revision"`
	Published       bool                                `json:"published"`
	Revision        int64                               `json:"revision"`
	RefetchRequired bool                                `json:"refetch_required"`
	ChangedAt       time.Time                           `json:"changed_at"`
}
