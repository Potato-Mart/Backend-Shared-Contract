package security

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security/security_enums"
)

// SecurityLogEntryFields contains the immutable evidence shared by protected
// data-access and administrative-write log entries. The embedding roots retain
// their distinct access and audit details.
type SecurityLogEntryFields struct {
	ID         string    `json:"id"`
	OccurredAt time.Time `json:"occurred_at"`
	ActorRef
	Action        string                           `json:"action"`
	Resource      string                           `json:"resource,omitempty"`
	ResourceID    string                           `json:"resource_id,omitempty"`
	RecordOutcome
	RiskLevel     security_enums.SecurityRiskLevel `json:"risk_level,omitempty"`
	IntegrityHash string                           `json:"integrity_hash,omitempty"`

	DataProtectionFields
}
