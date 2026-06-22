package shared

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/enums"
)

// AccessLogEntry records read/list/search/export access to protected data.
// Administrative writes should continue to use AuditLogEntry.
type AccessLogEntry struct {
	ID             string    `json:"id"`
	OccurredAt     time.Time `json:"occurred_at"`
	ActorRef       `bson:",inline"`
	Action         string `json:"action"`             // e.g. "customer.read" or "order.export"
	Resource       string `json:"resource,omitempty"` // e.g. "customer:cust_123"
	ResourceID     string `json:"resource_id,omitempty"`
	RecordCount    int    `json:"record_count,omitempty"`
	RequestContext `bson:",inline"`
	RecordOutcome  `bson:",inline"`
	RiskLevel      enums.SecurityRiskLevel `json:"risk_level,omitempty"`
	IntegrityHash  string                  `json:"integrity_hash,omitempty"`
	Metadata       common.Metadata         `json:"metadata,omitempty"`

	common.DataProtectionFields `bson:",inline"`
}
