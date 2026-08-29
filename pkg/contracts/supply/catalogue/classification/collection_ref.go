package classification

// CollectionRef identifies a collection by its immutable business code.
// Display names and presentation slugs are resolved from the root master.
type CollectionRef struct {
	Code string `json:"code"`
}
