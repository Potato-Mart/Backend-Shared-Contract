package security

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security/security_enums"
)

// SecurityEvent is a normalized event used for alerts, suspicious activity,
// access-control failures, data leakage indicators, and monitoring findings.
type SecurityEvent struct {
	ID          string                               `json:"id"`
	DetectedAt  time.Time                            `json:"detected_at"`
	ReportedAt  *time.Time                           `json:"reported_at,omitempty"`
	Source      string                               `json:"source,omitempty"`
	Category    string                               `json:"category"` // e.g. "auth", "access", "data", "cloud", "payment"
	Title       string                               `json:"title"`
	Description string                               `json:"description,omitempty"`
	Severity    security_enums.SecurityEventSeverity `json:"severity"`
	Status      security_enums.SecurityEventStatus   `json:"status"`
	RiskLevel   security_enums.SecurityRiskLevel     `json:"risk_level,omitempty"`
	ActorRef
	SubjectUserID      string                        `json:"subject_user_id,omitempty"`
	SubjectAccountID   string                        `json:"subject_account_id,omitempty"`
	AuthIdentityID     string                        `json:"auth_identity_id,omitempty"`
	IdentityDomain     security_enums.IdentityDomain `json:"identity_domain,omitempty"`
	Resource           string                        `json:"resource,omitempty"`
	ResourceID         string                        `json:"resource_id,omitempty"`
	RelatedAuditLogID  string                        `json:"related_audit_log_id,omitempty"`
	RelatedAccessLogID string                        `json:"related_access_log_id,omitempty"`
	EvidenceIDs        []string                      `json:"evidence_ids,omitempty"`
	AssignedTo         string                        `json:"assigned_to,omitempty"`
	ResolvedAt         *time.Time                    `json:"resolved_at,omitempty"`
	Resolution         string                        `json:"resolution,omitempty"`
	Metadata           metadata.Metadata             `json:"metadata,omitempty"`
	History            []HistoryEntry                `json:"history,omitempty"`

	DataProtectionFields
}
