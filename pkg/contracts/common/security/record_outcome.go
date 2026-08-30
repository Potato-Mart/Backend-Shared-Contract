package security

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security/security_enums"

// RecordOutcome is the result block shared by audit and access records.
type RecordOutcome struct {
	Outcome    security_enums.AuditOutcome `json:"outcome"`
	StatusCode int                         `json:"status_code,omitempty"`
	Reason     string                      `json:"reason,omitempty"` // failure reason / human note
}
