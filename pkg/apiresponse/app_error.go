// Package apiresponse defines the canonical error codes that every
// service in the MeatAndMe backend returns. Holding these codes in a
// shared module means API consumers (web, mobile, future
// microservices) can switch on stable string codes instead of matching
// message strings.
package apiresponse

import apiresponseenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/apiresponse"

type Error struct {
	Code    apiresponseenum.Code `json:"code"`
	Message string               `json:"message"`
	Details map[string]string    `json:"details,omitempty"`
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return string(e.Code) + ": " + e.Message
}

// New builds an *Error with a code and message.
func New(code apiresponseenum.Code, msg string) *Error {
	return &Error{Code: code, Message: msg}
}

// WithDetails adds a structured details map (used for field-level
// validation errors).
func (e *Error) WithDetails(details map[string]string) *Error {
	e.Details = details
	return e
}
