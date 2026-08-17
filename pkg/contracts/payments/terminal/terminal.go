// Package payments contains the shared contracts for EFTPOS payment
// terminals and the transactions they process.
package terminal

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/metadata"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/payments/terminal/terminal_enums"
)

// Terminal is an EFTPOS device registered for POS use.
type Terminal struct {
	ID       string                          `json:"id"`
	TenantID string                          `json:"tenant_id"`
	StoreID  string                          `json:"store_id,omitempty"`
	Provider terminal_enums.TerminalProvider `json:"provider"`

	ConnectionMode terminal_enums.TerminalConnectionMode `json:"connection_mode,omitempty"`

	ProviderDetails *TerminalProviderDetails `json:"provider_details,omitempty"`

	Status terminal_enums.TerminalStatus `json:"status"`

	RegisteredAt   *time.Time `json:"registered_at,omitempty"`
	DeregisteredAt *time.Time `json:"deregistered_at,omitempty"`
	LastSeenAt     *time.Time `json:"last_seen_at,omitempty"`

	Metadata metadata.Metadata       `json:"metadata,omitempty"`
	History  []security.HistoryEntry `json:"history,omitempty"`

	audit.AuditFields
}
