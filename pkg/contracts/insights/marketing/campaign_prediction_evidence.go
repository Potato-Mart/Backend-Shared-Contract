package marketing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/insights/marketing/marketing_enums"
)

// CampaignPredictionEvidence records an input window used by a forecast.
type CampaignPredictionEvidence struct {
	CampaignCode        string                                   `json:"campaign_code,omitempty"`
	Source              marketing_enums.CampaignPredictionSource `json:"source"`
	WindowStart         time.Time                                `json:"window_start"`
	WindowEnd           time.Time                                `json:"window_end"`
	RawNetBaseUnits     int64                                    `json:"raw_net_base_units"`
	NormalizedBaseUnits int64                                    `json:"normalized_base_units"`
	Weight              float64                                  `json:"weight"`
}
