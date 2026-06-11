package apiresponse

// APIResponse is the canonical envelope for every JSON response.
// Exactly one of Data or Error is populated.
//
// This includes internal service-to-service endpoints (the token endpoint
// in serviceauth, the stockops endpoints and the pricing quote endpoint):
// their responses are wrapped in this same {success,data,error} envelope,
// so service clients unwrap Data exactly like user-facing clients do.
type APIResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

// APIError is an alias of Error kept for backward compatibility; the two
// types previously duplicated each other field-for-field. The JSON shape
// is unchanged.
type APIError = Error
