package apiresponse

// APIResponse is the canonical envelope for every JSON response.
// Exactly one of Data or Error is populated.
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

// APIError is the JSON shape of an error response. The Code field is
// a stable machine-readable identifier; Message is human-readable.
type APIError struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}
