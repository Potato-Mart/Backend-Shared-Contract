package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

// AdClickSnapshot is one paid-advertising click the client stored when the
// visitor arrived, handed to checkout so the resulting order can be credited
// to the ad that produced it.
//
// The UTM fields are the tracking parameters that were on the landing URL.
// UTMID, UTMTerm and UTMContent carry the platform's campaign, ad-set and ad
// numbers, which is what makes an order attributable without guessing from
// names. ClickIDType names which platform click identifier ClickID holds, for
// example a Google or Meta click id, and FBC and FBP are the browser values
// the Meta pixel sets; a server-side conversion needs them, so they are
// carried here and never returned by any ads read.
//
// LandingPath is the path of the first page, with the query string removed, so
// it cannot smuggle the original parameters back in. CapturedAt is when the
// client stored the click, not when checkout happened: a click captured after
// checkout started, or older than the longest attribution window, is evidence
// of nothing and the whole snapshot is dropped rather than failing checkout.
// SourcePlatform says which client captured it, and AppTrackingStatus is the
// iOS tracking answer; unless it is authorized, neither the click identifier
// nor any hashed contact digest may be sent to a platform.
//
// It holds no personal data: no name, email, phone number or address ever
// appears on an ad click, and a client that sends one has its snapshot
// rejected.
type AdClickSnapshot struct {
	UTMSource         string                        `json:"utm_source,omitempty"`
	UTMMedium         string                        `json:"utm_medium,omitempty"`
	UTMCampaign       string                        `json:"utm_campaign,omitempty"`
	UTMID             string                        `json:"utm_id,omitempty"`
	UTMTerm           string                        `json:"utm_term,omitempty"`
	UTMContent        string                        `json:"utm_content,omitempty"`
	ClickIDType       string                        `json:"click_id_type,omitempty"`
	ClickID           string                        `json:"click_id,omitempty"`
	FBC               string                        `json:"fbc,omitempty"`
	FBP               string                        `json:"fbp,omitempty"`
	LandingPath       string                        `json:"landing_path,omitempty"`
	CapturedAt        time.Time                     `json:"captured_at,omitempty"`
	SourcePlatform    ads_enums.AdSourcePlatform    `json:"source_platform,omitempty"`
	AppTrackingStatus ads_enums.AdAppTrackingStatus `json:"att_status,omitempty"`
}
