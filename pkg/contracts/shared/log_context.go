package shared

import "github.com/Potato-Mart/Backend-Shared-Contract/v4/pkg/enums"

// ActorRef identifies the authenticated principal that performed an
// action. It is shared by audit, access, and security records.
type ActorRef struct {
	ActorID    string         `json:"actor_id,omitempty"`
	ActorEmail string         `json:"actor_email,omitempty"`
	ActorRole  enums.UserRole `json:"actor_role,omitempty"`
}

// RequestContext carries the transport-level correlation fields captured
// alongside audit, access, and security records.
type RequestContext struct {
	DeviceID      string `json:"device_id,omitempty"`
	SessionID     string `json:"session_id,omitempty"`
	IPAddress     string `json:"ip_address,omitempty"`
	UserAgent     string `json:"user_agent,omitempty"`
	RequestID     string `json:"request_id,omitempty"`
	CorrelationID string `json:"correlation_id,omitempty"`
	TraceID       string `json:"trace_id,omitempty"`
}

// RecordOutcome is the result block shared by audit and access records.
type RecordOutcome struct {
	Outcome    AuditOutcome `json:"outcome"`
	StatusCode int          `json:"status_code,omitempty"`
	Reason     string       `json:"reason,omitempty"` // failure reason / human note
}
