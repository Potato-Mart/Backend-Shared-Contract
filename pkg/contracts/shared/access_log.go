package shared

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// AccessLogEntry records read/list/search/export access to protected data.
// Administrative writes should continue to use AuditLogEntry.
type AccessLogEntry struct {
	ID            string                  `json:"id"`
	OccurredAt    time.Time               `json:"occurred_at"`
	ActorID       string                  `json:"actor_id,omitempty"`
	ActorEmail    string                  `json:"actor_email,omitempty"`
	ActorRole     enums.UserRole          `json:"actor_role,omitempty"`
	Action        string                  `json:"action"`             // e.g. "customer.read" or "order.export"
	Resource      string                  `json:"resource,omitempty"` // e.g. "customer:cust_123"
	ResourceID    string                  `json:"resource_id,omitempty"`
	RecordCount   int                     `json:"record_count,omitempty"`
	DeviceID      string                  `json:"device_id,omitempty"`
	SessionID     string                  `json:"session_id,omitempty"`
	IPAddress     string                  `json:"ip_address,omitempty"`
	UserAgent     string                  `json:"user_agent,omitempty"`
	RequestID     string                  `json:"request_id,omitempty"`
	CorrelationID string                  `json:"correlation_id,omitempty"`
	TraceID       string                  `json:"trace_id,omitempty"`
	Outcome       AuditOutcome            `json:"outcome"`
	StatusCode    int                     `json:"status_code,omitempty"`
	Reason        string                  `json:"reason,omitempty"`
	RiskLevel     enums.SecurityRiskLevel `json:"risk_level,omitempty"`
	IntegrityHash string                  `json:"integrity_hash,omitempty"`
	Metadata      common.Metadata         `json:"metadata,omitempty"`

	common.DataProtectionFields
}
