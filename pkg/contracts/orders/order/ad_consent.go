package order

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"

// AdConsent is the advertising consent the visitor gave, recorded alongside
// the ad click it governs.
//
// All three default to denied, and an absent answer is denied, never assumed
// consent. AdStorage permits storing the click at all. AdUserData is the one
// that gates contact data: only when it is granted may a hashed email, phone
// or external-id digest be sent to a platform with the purchase, and even then
// only digests, never the raw values. AdPersonalization permits using the
// purchase for personalised advertising. Consent is recorded as it was at the
// time of the click and is not back-dated by a later change of mind.
type AdConsent struct {
	AdStorage         ads_enums.AdConsentState `json:"ad_storage,omitempty"`
	AdUserData        ads_enums.AdConsentState `json:"ad_user_data,omitempty"`
	AdPersonalization ads_enums.AdConsentState `json:"ad_personalization,omitempty"`
}
