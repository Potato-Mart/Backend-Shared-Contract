package pricing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v49/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v49/pkg/contracts/marketing/campaign/campaign_enums"
)

// CampaignChangedEvent is version 3.
// CampaignChangedEvent is the customer-safe storefront-events projection used
// only to invalidate/refetch authoritative campaign content. It contains no
// authored copy, targeting rules, provider destination, or customer data.
type CampaignChangedEvent struct {
	CampaignCode       string                              `json:"campaign_code"`
	Status             campaign_enums.CampaignStatus       `json:"publication_status"`
	Published          bool                                `json:"published"`
	ScopeMode          geography_enums.GeographicScopeMode `json:"scope_mode"`
	ScopeRevision      int64                               `json:"scope_revision"`
	ActivationRevision int64                               `json:"publication_revision"`
	ContentRevision    int64                               `json:"content_revision"`
	RefetchRequired    bool                                `json:"refetch_required"`
	ChangedAt          time.Time                           `json:"changed_at"`
}
