package campaign

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v47/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v47/pkg/contracts/marketing/campaign/campaign_enums"
)

// Placement has independent ordering and artwork for one presentation position.
type Placement struct {
	Position  campaign_enums.CampaignPlacement  `json:"position"`
	Platforms []campaign_enums.CampaignPlatform `json:"platforms"`
	Priority  int                               `json:"priority"`
	Media     *security.ObjectMedia             `json:"media,omitempty"`
}
