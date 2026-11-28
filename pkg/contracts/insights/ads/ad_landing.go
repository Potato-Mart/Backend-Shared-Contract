package ads

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"

// AdLanding is where a draft ad sends the person who clicks it.
//
// Kind says what Ref refers to: a campaign code, a product, a selection, or a
// path on our own storefront domain. The ads module builds the final URL per
// copy locale and appends the agreed tracking parameters, so Ref itself never
// carries tracking parameters, never points at another domain, and is never a
// shortened link: a redirect breaks domain checks and can lose the platform
// click id.
type AdLanding struct {
	Kind ads_enums.AdLandingKind `json:"kind,omitempty"`
	Ref  string                  `json:"ref,omitempty"`
}
