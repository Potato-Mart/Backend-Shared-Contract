package security

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security/security_enums"
)

// AuditLogEntry is one immutable record of an administrative action.
//
// Entries are written by middleware on every successful write request
// (HTTP method POST/PUT/PATCH/DELETE) plus explicit business events
// (login, role change, refund issued, etc.). Reads are never audited
// here – use access logs for that.
type AuditLogEntry struct {
	ID         string    `json:"id"`
	OccurredAt time.Time `json:"occurred_at"`
	ActorRef
	Action     string `json:"action"`             // dotted key e.g. "customer.update"
	Resource   string `json:"resource,omitempty"` // dotted key e.g. "customer:cust_123"
	ResourceID string `json:"resource_id,omitempty"`
	RecordOutcome
	RiskLevel     security_enums.SecurityRiskLevel `json:"risk_level,omitempty"`
	IntegrityHash string                           `json:"integrity_hash,omitempty"`
	Diff          metadata.Metadata                `json:"diff,omitempty"` // arbitrary before/after fragments

	DataProtectionFields
}
