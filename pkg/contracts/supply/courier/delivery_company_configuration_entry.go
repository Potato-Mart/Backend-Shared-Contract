package courier

// DeliveryCompanyConfigurationEntry is one user-authored configuration value,
// including API endpoints or credentials. Key is an arbitrary company-scoped
// identifier, not a preset catalogue. Value retains its JSON type. Supply
// validates duplicate keys and the configured normalized protocol.
// Sensitive values must be encrypted at rest and replaced with a masked value
// in read projections; the masking and update DTO semantics are Supply-owned.
// This data-only type does not mask itself during JSON serialization. Only
// authorized operations may receive unmasked values.
type DeliveryCompanyConfigurationEntry struct {
	Key       string `json:"key"`
	Value     any    `json:"value"`
	Sensitive bool   `json:"sensitive"`
}
