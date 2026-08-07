// Package payments contains the shared contracts for EFTPOS payment
// terminals and the transactions they process.
package terminal

import (
	security "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/security"
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
)

// Terminal is an EFTPOS device registered for POS use.
type Terminal struct {
	ID       string           `json:"id"`
	TenantID string           `json:"tenant_id"`
	StoreID  string           `json:"store_id,omitempty"`
	Provider TerminalProvider `json:"provider"`

	ConnectionMode TerminalConnectionMode `json:"connection_mode,omitempty"`

	ProviderDetails *TerminalProviderDetails `json:"provider_details,omitempty"`

	Status TerminalStatus `json:"status"`

	RegisteredAt   *time.Time `json:"registered_at,omitempty"`
	DeregisteredAt *time.Time `json:"deregistered_at,omitempty"`
	LastSeenAt     *time.Time `json:"last_seen_at,omitempty"`

	Metadata common.Metadata         `json:"metadata,omitempty"`
	History  []security.HistoryEntry `json:"history,omitempty"`

	common.AuditFields
}
