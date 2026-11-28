package ads_enums

// AdLevel names the level of the platform object hierarchy that a performance
// row, a delivery change or a budget belongs to. The three levels are the same
// on both platforms: a Google ad group is an ad set, and a Google ad is an ad.
type AdLevel string

const (
	// AdLevelCampaign is the campaign, the top level of the hierarchy.
	AdLevelCampaign AdLevel = "campaign"
	// AdLevelAdSet is the Meta ad set or the Google ad group.
	AdLevelAdSet AdLevel = "ad_set"
	// AdLevelAd is the individual ad that carries the creative.
	AdLevelAd AdLevel = "ad"
)

// IsValid reports whether l is a known AdLevel.
func (l AdLevel) IsValid() bool {
	switch l {
	case AdLevelCampaign, AdLevelAdSet, AdLevelAd:
		return true
	}
	return false
}

func (l AdLevel) String() string { return string(l) }
