package courier

// DeliveryAPIRequestField is one manually authored request field. Exactly one
// of Value, ValueIsNull, or SecretRef is supplied. Value is non-secret,
// non-null JSON data; ValueIsNull represents an explicit JSON null. A
// SecretRef names a key in the protected provider-extension credential values
// and never carries the secret itself. Supply validates the one-of rule and
// requires header/query values to be JSON strings.
type DeliveryAPIRequestField struct {
	Key         string `json:"key"`
	Value       any    `json:"value,omitempty"`
	ValueIsNull bool   `json:"value_is_null,omitempty"`
	SecretRef   string `json:"secret_ref,omitempty"`
}
