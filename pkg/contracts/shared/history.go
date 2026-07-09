package shared

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/common"
	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/security"
)

// HistoryChange is one field-level before/after fragment in a timeline entry.
type HistoryChange struct {
	Field     string `json:"field"`
	FromValue string `json:"from_value,omitempty"`
	ToValue   string `json:"to_value,omitempty"`
}

// HistoryEntry is a concise, UI-safe timeline record for process visibility.
//
// It is not a replacement for immutable audit/access logs. Use it on domain
// records to show what changed, who or what caused it, why it changed, and
// whether the change introduces operational, financial, customer, or security
// risk.
type HistoryEntry struct {
	ID                string                         `json:"id,omitempty"`
	Sequence          int64                          `json:"sequence,omitempty"`
	OccurredAt        time.Time                      `json:"occurred_at"`
	Type              string                         `json:"type"`
	Summary           string                         `json:"summary,omitempty"`
	Changes           []HistoryChange                `json:"changes,omitempty"`
	Source            string                         `json:"source,omitempty"`
	ReasonCode        string                         `json:"reason_code,omitempty"`
	Note              string                         `json:"note,omitempty"`
	RiskLevel         securityenum.SecurityRiskLevel `json:"risk_level,omitempty"`
	RiskFlags         []string                       `json:"risk_flags,omitempty"`
	RelatedResource   string                         `json:"related_resource,omitempty"`
	RelatedResourceID string                         `json:"related_resource_id,omitempty"`
	Metadata          common.Metadata                `json:"metadata,omitempty"`
	ActorRef
	RequestContext
}
