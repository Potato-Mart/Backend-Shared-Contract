package ads

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"

// AdAttributionPolicy is the rule that decides which paid orders are credited
// to an ad.
//
// Model picks the click, and ClickWindowDays bounds how long before the order
// that click may be: the agreed default is the last ad click within seven days.
// Only paid orders count, and an order that is cancelled or fully refunded
// leaves the attributed set again. IncludeWholesale is false by default, so
// wholesale orders are not credited to ads unless the rule says so. Changing
// the policy changes past figures, so the ads module recomputes in the
// background rather than letting two readers disagree.
type AdAttributionPolicy struct {
	Model            ads_enums.AdAttributionModel `json:"model,omitempty"`
	ClickWindowDays  int                          `json:"click_window_days,omitempty"`
	IncludeWholesale bool                         `json:"include_wholesale,omitempty"`
}
