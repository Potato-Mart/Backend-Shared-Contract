package security

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/security/security_enums"
)

// DataProtectionFields are optional labels that let each service apply
// consistent classification, retention, deletion, and privacy controls.
type DataProtectionFields struct {
	Classification       security_enums.DataClassification  `json:"classification,omitempty"`
	ProtectionBasis      security_enums.DataProtectionBasis `json:"protection_basis,omitempty"`
	ContainsPII          bool                               `json:"contains_pii,omitempty"`
	ContainsSensitivePII bool                               `json:"contains_sensitive_pii,omitempty"`
	DataOwnerID          string                             `json:"data_owner_id,omitempty"`
	RetentionPolicyKey   string                             `json:"retention_policy_key,omitempty"`
	RetentionUntil       *time.Time                         `json:"retention_until,omitempty"`
	LegalHold            bool                               `json:"legal_hold,omitempty"`
	DeletionRequestedAt  *time.Time                         `json:"deletion_requested_at,omitempty"`
	DeletionCompletedAt  *time.Time                         `json:"deletion_completed_at,omitempty"`
}
