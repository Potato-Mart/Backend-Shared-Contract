package apiresponse

// APIResponse is the canonical envelope for every JSON response.
// Exactly one of Data or Error is populated.
type APIResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

// APIError is an alias of Error kept for backward compatibility; the two
// types previously duplicated each other field-for-field. The JSON shape
// is unchanged.
type APIError = Error
