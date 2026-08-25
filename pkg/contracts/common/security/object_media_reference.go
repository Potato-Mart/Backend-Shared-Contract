package security

// ObjectMediaReference records an owning aggregate attachment without
// embedding the aggregate itself in the asset record.
type ObjectMediaReference struct {
	EntityType string `json:"entity_type"`
	EntityID   string `json:"entity_id"`
	Field      string `json:"field,omitempty"`
}
