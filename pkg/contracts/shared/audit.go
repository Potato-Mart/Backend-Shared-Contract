package shared

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/security"
)

// AuditLogEntry is one immutable record of an administrative action.
//
// Entries are written by middleware on every successful write request
// (HTTP method POST/PUT/PATCH/DELETE) plus explicit business events
// (login, role change, refund issued, etc.). Reads are never audited
// here â€“ use access logs for that.
type AuditLogEntry struct {
	ID         string    `json:"id"`
	OccurredAt time.Time `json:"occurred_at"`
	ActorRef
	Action     string `json:"action"`             // dotted key e.g. "customer.update"
	Resource   string `json:"resource,omitempty"` // dotted key e.g. "customer:cust_123"
	ResourceID string `json:"resource_id,omitempty"`
	RequestContext
	RecordOutcome
	RiskLevel     securityenum.SecurityRiskLevel `json:"risk_level,omitempty"`
	IntegrityHash string                         `json:"integrity_hash,omitempty"`
	Diff          common.Metadata                `json:"diff,omitempty"` // arbitrary before/after fragments

	common.DataProtectionFields
}
