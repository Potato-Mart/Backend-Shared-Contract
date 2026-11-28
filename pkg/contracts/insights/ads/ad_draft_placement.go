package ads

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"

// AdDraftPlacement says where a draft ad will live on the platform.
//
// When Mode is existing, AdSetID names an ad set that is already there and the
// settings below are absent. When Mode is new, Campaign and AdSet carry the
// settings for the campaign and ad set to create, in the same normalised shape
// a delivery change uses, and AdSetID is empty. CampaignID may name an
// existing campaign that a new ad set is added to.
type AdDraftPlacement struct {
	Mode       ads_enums.AdPlacementMode `json:"mode,omitempty"`
	CampaignID string                    `json:"campaign_id,omitempty"`
	AdSetID    string                    `json:"ad_set_id,omitempty"`
	Campaign   *AdChangeSet              `json:"campaign,omitempty"`
	AdSet      *AdChangeSet              `json:"ad_set,omitempty"`
}
