package promotion

// PromotionSource records an optional upstream source with open identifiers.
type PromotionSource struct {
	Kind string `json:"kind,omitempty"`
	Ref  string `json:"ref,omitempty"`
}
