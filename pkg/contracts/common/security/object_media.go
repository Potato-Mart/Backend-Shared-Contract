package security

// ObjectMedia is the safe render projection of an object-storage asset.
// It intentionally contains only the stable identity and URL needed by
// consumer-facing contracts.
type ObjectMedia struct {
	Code string `json:"code"`
	URL  string `json:"url,omitempty"`
}
