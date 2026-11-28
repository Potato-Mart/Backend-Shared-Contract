package ads

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"

// AdCopy is the wording of a draft ad in one language.
//
// Locale is a BCP 47 language tag, and a draft carries one entry per language
// it runs in; the landing URL is built per locale from the draft's landing
// destination. Headline, LongHeadline, Description and BusinessName map onto
// the platform's own slots, so a platform that has no slot for a field leaves
// it empty rather than merging it elsewhere. Length limits differ per platform
// and per slot and are enforced by the ads module; a reader may show a live
// character count but must not treat its own limit as authoritative.
// Translations are reviewed by a person before a draft can be submitted.
type AdCopy struct {
	Locale       string                   `json:"locale,omitempty"`
	PrimaryText  string                   `json:"primary_text,omitempty"`
	Headline     string                   `json:"headline,omitempty"`
	LongHeadline string                   `json:"long_headline,omitempty"`
	Description  string                   `json:"description,omitempty"`
	BusinessName string                   `json:"business_name,omitempty"`
	CallToAction ads_enums.AdCallToAction `json:"call_to_action,omitempty"`
	Reviewed     bool                     `json:"reviewed,omitempty"`
}
