// Package campaignlogic holds derived helpers for campaign contract structs.
// It lives outside pkg/contracts so the contract files stay pure data shapes.
package campaignlogic

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/campaign"
)

// LiveAt reports whether the campaign should be shown at instant t: active and
// within its scheduling window.
func LiveAt(c campaign.Campaign, t time.Time) bool {
	if !c.IsActive {
		return false
	}
	if c.StartsAt != nil && t.Before(*c.StartsAt) {
		return false
	}
	if c.EndsAt != nil && t.After(*c.EndsAt) {
		return false
	}
	return true
}
