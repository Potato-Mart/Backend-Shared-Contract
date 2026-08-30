package classification

// ObjectMediaRef is the code-only media relationship persisted by catalog
// records. Render URLs are resolved from the managed media asset.
type ObjectMediaRef struct {
	Code string `json:"code"`
}
