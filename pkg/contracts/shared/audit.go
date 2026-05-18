package shared

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// AuditLogEntry is one immutable record of an administrative action.
//
// Entries are written by middleware on every successful write request
// (HTTP method POST/PUT/PATCH/DELETE) plus explicit business events
// (login, role change, refund issued, etc.). Reads are never audited
// here – use access logs for that.
type AuditLogEntry struct {
	ID         string          `json:"id"`
	OccurredAt time.Time       `json:"occurred_at"`
	ActorID    string          `json:"actor_id,omitempty"`
	ActorEmail string          `json:"actor_email,omitempty"`
	ActorRole  enums.UserRole  `json:"actor_role,omitempty"`
	Action     string          `json:"action"`              // dotted key e.g. "customer.update"
	Resource   string          `json:"resource,omitempty"`  // dotted key e.g. "customer:cust_123"
	ResourceID string          `json:"resource_id,omitempty"`
	IPAddress  string          `json:"ip_address,omitempty"`
	UserAgent  string          `json:"user_agent,omitempty"`
	RequestID  string          `json:"request_id,omitempty"`
	Outcome    AuditOutcome    `json:"outcome"`
	StatusCode int             `json:"status_code,omitempty"`
	Reason     string          `json:"reason,omitempty"` // failure reason / human note
	Diff       common.Metadata `json:"diff,omitempty"`   // arbitrary before/after fragments
}

// AuditOutcome is a coarse success/failure flag for audit reporting.
type AuditOutcome string

const (
	AuditOutcomeSuccess AuditOutcome = "success"
	AuditOutcomeFailure AuditOutcome = "failure"
	AuditOutcomeDenied  AuditOutcome = "denied"
)
