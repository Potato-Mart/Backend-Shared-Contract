package ads

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"

// AdChangeSet is the normalised set of delivery settings a change touches. The
// same shape records what the object looked like before and after, so a change
// history row can be read as a field-by-field diff.
//
// Every field is optional and absence means "not part of this change", never a
// cleared value. AutomaticPlacements is a pointer for that reason: nil leaves
// placement selection alone, while false asks for manual placements and names
// them in Placements.
type AdChangeSet struct {
	Status              ads_enums.AdEntityStatus `json:"status,omitempty"`
	Budget              *AdBudget                `json:"budget,omitempty"`
	Bidding             *AdBidding               `json:"bid_strategy,omitempty"`
	Targeting           *AdTargetingSummary      `json:"targeting,omitempty"`
	Schedule            *AdSchedule              `json:"schedule,omitempty"`
	AutomaticPlacements *bool                    `json:"automatic_placements,omitempty"`
	Placements          []string                 `json:"placements,omitempty"`
}
