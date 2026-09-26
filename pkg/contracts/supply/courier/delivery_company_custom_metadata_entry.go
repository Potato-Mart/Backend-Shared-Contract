package courier

// DeliveryCompanyCustomMetadataEntry stores an extensible, non-secret value
// associated with a delivery company. ValueType is an open identifier and Value
// is typed JSON. Custom metadata is inert: it is not provider configuration and
// must never affect provider requests. Consumers preserve unknown types/values.
type DeliveryCompanyCustomMetadataEntry struct {
	Key       string `json:"key"`
	ValueType string `json:"value_type"`
	Value     any    `json:"value"`
}
