package security

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security/security_enums"
)

// SecurityIncident groups one or more security events into an incident
// response workflow with ownership, evidence, impact, and closure data.
type SecurityIncident struct {
	ID                 string                               `json:"id"`
	OpenedAt           time.Time                            `json:"opened_at"`
	ClosedAt           *time.Time                           `json:"closed_at,omitempty"`
	Title              string                               `json:"title"`
	Description        string                               `json:"description,omitempty"`
	Severity           security_enums.SecurityEventSeverity `json:"severity"`
	Status             security_enums.SecurityEventStatus   `json:"status"`
	OwnerID            string                               `json:"owner_id,omitempty"`
	CommanderID        string                               `json:"commander_id,omitempty"`
	EventIDs           []string                             `json:"event_ids,omitempty"`
	EvidenceIDs        []string                             `json:"evidence_ids,omitempty"`
	AffectedSystems    []string                             `json:"affected_systems,omitempty"`
	AffectedData       []string                             `json:"affected_data,omitempty"`
	CustomerImpact     string                               `json:"customer_impact,omitempty"`
	ContainmentSummary string                               `json:"containment_summary,omitempty"`
	RecoverySummary    string                               `json:"recovery_summary,omitempty"`
	RootCause          string                               `json:"root_cause,omitempty"`
	LessonsLearned     string                               `json:"lessons_learned,omitempty"`
	NotifiableBreach   bool                                 `json:"notifiable_breach,omitempty"`
	NotificationDueAt  *time.Time                           `json:"notification_due_at,omitempty"`
	Metadata           metadata.Metadata                    `json:"metadata,omitempty"`
	History            []HistoryEntry                       `json:"history,omitempty"`

	DataProtectionFields
}
