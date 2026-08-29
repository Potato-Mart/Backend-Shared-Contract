package security

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security/security_enums"
)

// HistoryEntry is a concise, UI-safe timeline record for process visibility.
//
// It is not a replacement for immutable audit/access logs. Use it on domain
// records to show what changed, who or what caused it, why it changed, and
// whether the change introduces operational, financial, customer, or security
// risk.
type HistoryEntry struct {
	ID                string                           `json:"id,omitempty"`
	Sequence          int64                            `json:"sequence,omitempty"`
	OccurredAt        time.Time                        `json:"occurred_at"`
	Type              string                           `json:"type"`
	Summary           string                           `json:"summary,omitempty"`
	Changes           []HistoryChange                  `json:"changes,omitempty"`
	Source            string                           `json:"source,omitempty"`
	ReasonCode        string                           `json:"reason_code,omitempty"`
	Note              string                           `json:"note,omitempty"`
	RiskLevel         security_enums.SecurityRiskLevel `json:"risk_level,omitempty"`
	RiskFlags         []string                         `json:"risk_flags,omitempty"`
	RelatedResource   string                           `json:"related_resource,omitempty"`
	RelatedResourceID string                           `json:"related_resource_id,omitempty"`
	Metadata          metadata.Metadata                `json:"metadata,omitempty"`
	ActorRef
}
