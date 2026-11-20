package security

// ObjectMedia is the safe render projection of an object-storage asset.
// It intentionally contains only the stable identity and URL needed by
// consumer-facing contracts, plus optional display copy carried from the
// owning record's media reference.
//
// AltText and Caption are the owning record's copy for this use of the image
// (for example product cover, gallery and detail images): AltText for screen
// readers and failed loads, Caption as visible text shown with the image.
// Both are omitted when empty.
type ObjectMedia struct {
	Code    string `json:"code"`
	URL     string `json:"url,omitempty"`
	AltText string `json:"alt_text,omitempty"`
	Caption string `json:"caption,omitempty"`
}
