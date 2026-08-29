package classification

// CategoryTagRef identifies a category tag by its immutable code without
// embedding mutable classification or audit details.
type CategoryTagRef struct {
	Code string `json:"code"`
}
