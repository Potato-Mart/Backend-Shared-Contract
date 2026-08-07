package security

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/security/security_enums"
)

// AccessLogEntry records read/list/search/export access to protected data.
// Administrative writes should continue to use AuditLogEntry.
type AccessLogEntry struct {
	ID         string    `json:"id"`
	OccurredAt time.Time `json:"occurred_at"`
	ActorRef
	Action      string `json:"action"`             // e.g. "customer.read" or "order.export"
	Resource    string `json:"resource,omitempty"` // e.g. "customer:cust_123"
	ResourceID  string `json:"resource_id,omitempty"`
	RecordCount int    `json:"record_count,omitempty"`
	RequestContext
	RecordOutcome
	RiskLevel     security_enums.SecurityRiskLevel `json:"risk_level,omitempty"`
	IntegrityHash string                           `json:"integrity_hash,omitempty"`
	Metadata      metadata.Metadata                `json:"metadata,omitempty"`

	DataProtectionFields
}
