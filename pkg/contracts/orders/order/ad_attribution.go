package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

// AdAttribution is the server-derived, read-only record of which paid ad an
// order came from. It is what the order keeps; the client-supplied AdClick is
// only the input it is derived from.
//
// Platform is decided from the tracking parameters by the receiving service
// and is never taken from the client. CampaignID, AdSetID and AdID are the
// platform's own numbers, which is what lets reporting roll an order up to a
// campaign, an ad set and an ad without matching on names. Click is the
// snapshot that was actually resolved under the attribution rule in force, so
// a reader does not have to re-apply that rule, and ClickedAt is when that
// click happened. SourcePlatform says which client the order was placed from.
//
// It holds no personal data: no name, email, phone number or address appears
// here, and the fields that identify a person on a platform stay on the click
// snapshot and in the paid-order event.
type AdAttribution struct {
	Platform       ads_enums.AdPlatform       `json:"platform,omitempty"`
	CampaignID     string                     `json:"campaign_id,omitempty"`
	AdSetID        string                     `json:"ad_set_id,omitempty"`
	AdID           string                     `json:"ad_id,omitempty"`
	Click          *AdClickSnapshot           `json:"click,omitempty"`
	ClickedAt      *time.Time                 `json:"clicked_at,omitempty"`
	SourcePlatform ads_enums.AdSourcePlatform `json:"source_platform,omitempty"`
}
