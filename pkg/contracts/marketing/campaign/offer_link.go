package campaign

// OfferLink references commercial authority; it never supplies a price or benefit.
type OfferLink struct {
	Kind         string `json:"kind"`
	ID           string `json:"id"`
	DisplayOrder int    `json:"display_order"`
	Revision     string `json:"revision,omitempty"`
}
