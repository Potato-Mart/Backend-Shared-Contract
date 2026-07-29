// Package payments contains the shared contracts for EFTPOS payment
// terminals and the transactions they process.
package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/shared"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/payment"
)

// Terminal is an EFTPOS device registered for POS use.
type Terminal struct {
	ID       string                       `json:"id"`
	TenantID string                       `json:"tenant_id"`
	StoreID  string                       `json:"store_id,omitempty"`
	Provider paymentenum.TerminalProvider `json:"provider"`

	ConnectionMode paymentenum.TerminalConnectionMode `json:"connection_mode,omitempty"`

	ProviderDetails *TerminalProviderDetails `json:"provider_details,omitempty"`

	Status paymentenum.TerminalStatus `json:"status"`

	RegisteredAt   *time.Time `json:"registered_at,omitempty"`
	DeregisteredAt *time.Time `json:"deregistered_at,omitempty"`
	LastSeenAt     *time.Time `json:"last_seen_at,omitempty"`

	Metadata common.Metadata       `json:"metadata,omitempty"`
	History  []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields
}
