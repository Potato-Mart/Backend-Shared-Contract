package security

// HistoryChange is one field-level before/after fragment in a timeline entry.
type HistoryChange struct {
	Field     string `json:"field"`
	FromValue string `json:"from_value,omitempty"`
	ToValue   string `json:"to_value,omitempty"`
}
