package identity

import "github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/common"

// IdentityRequestContext carries audit and correlation metadata for identity
// and access-management command DTOs.
type IdentityRequestContext struct {
	RequestedBy   string          `json:"requested_by,omitempty"`
	RequestID     string          `json:"request_id,omitempty"`
	CorrelationID string          `json:"correlation_id,omitempty"`
	TraceID       string          `json:"trace_id,omitempty"`
	IPAddress     string          `json:"ip_address,omitempty"`
	UserAgent     string          `json:"user_agent,omitempty"`
	Reason        string          `json:"reason,omitempty"`
	Metadata      common.Metadata `json:"metadata,omitempty"`
}
