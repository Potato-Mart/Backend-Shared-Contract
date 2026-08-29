package forecasting

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/marketing/audience"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/marketing/campaign/campaign_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/forecasting/marketing_enums"
)

// CampaignComparableEvent is historical evidence used to forecast a campaign.
type CampaignComparableEvent struct {
	CampaignCode     string                                   `json:"campaign_code"`
	SeriesCode       string                                   `json:"series_code,omitempty"`
	MatchSource      marketing_enums.CampaignPredictionSource `json:"match_source"`
	ResolvedSKUCodes []string                                 `json:"resolved_sku_codes"`
	StartsAt         time.Time                                `json:"starts_at"`
	EndsAt           time.Time                                `json:"ends_at"`
	ScheduleTimezone string                                   `json:"schedule_timezone"`
	Audience         *audience.Audience                       `json:"audience,omitempty"`
	GeographicScope  geography.GeographicScope                `json:"geographic_scope"`
	Placement        campaign_enums.CampaignPlacement         `json:"placement"`
}
