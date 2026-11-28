package ads

// AdTargetingSummary is a reader-friendly digest of who one platform object
// targets. It is a summary, not the platform's full targeting spec: it carries
// enough to show and to restate a targeting change, while the platform keeps
// the authoritative definition.
//
// Locations, Interests, Languages and the audience lists hold platform
// identifiers or codes, never personal data about any individual.
// AgeMin and AgeMax are absent when the object does not narrow by age.
// AutomaticAudience is the platform's own broad-audience mode: when it is on,
// the explicit interest and audience fields are advisory signals only.
type AdTargetingSummary struct {
	Locations                 []string `json:"locations,omitempty"`
	Languages                 []string `json:"languages,omitempty"`
	Interests                 []string `json:"interests,omitempty"`
	Genders                   []string `json:"genders,omitempty"`
	AgeMin                    int32    `json:"age_min,omitempty"`
	AgeMax                    int32    `json:"age_max,omitempty"`
	IncludedCustomAudienceIDs []string `json:"included_custom_audience_ids,omitempty"`
	ExcludedCustomAudienceIDs []string `json:"excluded_custom_audience_ids,omitempty"`
	AutomaticAudience         bool     `json:"automatic_audience,omitempty"`
}
