package shared

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/enums"
)

// SecurityEvent is a normalized event used for alerts, suspicious activity,
// access-control failures, data leakage indicators, and monitoring findings.
type SecurityEvent struct {
	ID                 string                      `json:"id"`
	DetectedAt         time.Time                   `json:"detected_at"`
	ReportedAt         *time.Time                  `json:"reported_at,omitempty"`
	Source             string                      `json:"source,omitempty"`
	Category           string                      `json:"category"` // e.g. "auth", "access", "data", "cloud", "payment"
	Title              string                      `json:"title"`
	Description        string                      `json:"description,omitempty"`
	Severity           enums.SecurityEventSeverity `json:"severity"`
	Status             enums.SecurityEventStatus   `json:"status"`
	RiskLevel          enums.SecurityRiskLevel     `json:"risk_level,omitempty"`
	ActorRef           `bson:",inline"`
	SubjectUserID      string `json:"subject_user_id,omitempty"`
	RequestContext     `bson:",inline"`
	Resource           string          `json:"resource,omitempty"`
	ResourceID         string          `json:"resource_id,omitempty"`
	RelatedAuditLogID  string          `json:"related_audit_log_id,omitempty"`
	RelatedAccessLogID string          `json:"related_access_log_id,omitempty"`
	EvidenceIDs        []string        `json:"evidence_ids,omitempty"`
	AssignedTo         string          `json:"assigned_to,omitempty"`
	ResolvedAt         *time.Time      `json:"resolved_at,omitempty"`
	Resolution         string          `json:"resolution,omitempty"`
	Metadata           common.Metadata `json:"metadata,omitempty"`
	History            []HistoryEntry  `json:"history,omitempty"`

	common.DataProtectionFields `bson:",inline"`
}

// SecurityIncident groups one or more security events into an incident
// response workflow with ownership, evidence, impact, and closure data.
type SecurityIncident struct {
	ID                 string                      `json:"id"`
	OpenedAt           time.Time                   `json:"opened_at"`
	ClosedAt           *time.Time                  `json:"closed_at,omitempty"`
	Title              string                      `json:"title"`
	Description        string                      `json:"description,omitempty"`
	Severity           enums.SecurityEventSeverity `json:"severity"`
	Status             enums.SecurityEventStatus   `json:"status"`
	OwnerID            string                      `json:"owner_id,omitempty"`
	CommanderID        string                      `json:"commander_id,omitempty"`
	EventIDs           []string                    `json:"event_ids,omitempty"`
	EvidenceIDs        []string                    `json:"evidence_ids,omitempty"`
	AffectedSystems    []string                    `json:"affected_systems,omitempty"`
	AffectedData       []string                    `json:"affected_data,omitempty"`
	CustomerImpact     string                      `json:"customer_impact,omitempty"`
	ContainmentSummary string                      `json:"containment_summary,omitempty"`
	RecoverySummary    string                      `json:"recovery_summary,omitempty"`
	RootCause          string                      `json:"root_cause,omitempty"`
	LessonsLearned     string                      `json:"lessons_learned,omitempty"`
	NotifiableBreach   bool                        `json:"notifiable_breach,omitempty"`
	NotificationDueAt  *time.Time                  `json:"notification_due_at,omitempty"`
	Metadata           common.Metadata             `json:"metadata,omitempty"`
	History            []HistoryEntry              `json:"history,omitempty"`

	common.DataProtectionFields `bson:",inline"`
}

// EvidenceRecord tracks security evidence and chain-of-custody metadata.
type EvidenceRecord struct {
	ID             string          `json:"id"`
	CollectedAt    time.Time       `json:"collected_at"`
	CollectedBy    string          `json:"collected_by,omitempty"`
	Source         string          `json:"source,omitempty"`
	EvidenceType   string          `json:"evidence_type"` // e.g. "log", "screenshot", "export", "provider_report"
	Description    string          `json:"description,omitempty"`
	StorageURI     string          `json:"storage_uri,omitempty"`
	HashAlgorithm  string          `json:"hash_algorithm,omitempty"`
	HashValue      string          `json:"hash_value,omitempty"`
	ChainOfCustody []CustodyEvent  `json:"chain_of_custody,omitempty"`
	Metadata       common.Metadata `json:"metadata,omitempty"`

	common.DataProtectionFields `bson:",inline"`
}

type CustodyEvent struct {
	OccurredAt time.Time `json:"occurred_at"`
	ActorID    string    `json:"actor_id,omitempty"`
	Action     string    `json:"action"` // e.g. "collected", "transferred", "reviewed", "sealed"
	Reason     string    `json:"reason,omitempty"`
}
