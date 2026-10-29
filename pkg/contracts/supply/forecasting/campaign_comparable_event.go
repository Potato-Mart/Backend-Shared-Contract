package forecasting

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/marketing/campaign"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/marketing/campaign/campaign_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/forecasting/marketing_enums"
)

// CampaignComparableEvent is historical evidence used to forecast a campaign.
type CampaignComparableEvent struct {
	CampaignCode     string                                   `json:"campaign_code"`
	MatchSource      marketing_enums.CampaignPredictionSource `json:"match_source"`
	ResolvedSKUCodes []string                                 `json:"resolved_sku_codes"`
	StartsAt         time.Time                                `json:"starts_at"`
	EndsAt           time.Time                                `json:"ends_at"`
	ScheduleTimezone string                                   `json:"schedule_timezone"`
	Audience         []campaign_enums.CampaignCustomerType    `json:"audience"`
	GeographicScope  geography.GeographicScope                `json:"geographic_scope"`
	Placements       []campaign.Placement                     `json:"placements"`
}
