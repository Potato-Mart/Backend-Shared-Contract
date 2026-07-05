// Package payments contains the shared contracts for EFTPOS payment
// terminals and the transactions they process.
package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/enums"
)

// Terminal is an EFTPOS device registered for POS use.
type Terminal struct {
	ID       string                 `json:"id"`
	TenantID string                 `json:"tenant_id"`
	StoreID  string                 `json:"store_id,omitempty"`
	Provider enums.TerminalProvider `json:"provider"`

	ConnectionMode enums.TerminalConnectionMode `json:"connection_mode,omitempty"`

	ProviderDetails *TerminalProviderDetails `json:"provider_details,omitempty"`

	Status enums.TerminalStatus `json:"status"`

	RegisteredAt   *time.Time `json:"registered_at,omitempty"`
	DeregisteredAt *time.Time `json:"deregistered_at,omitempty"`
	LastSeenAt     *time.Time `json:"last_seen_at,omitempty"`

	Metadata common.Metadata       `json:"metadata,omitempty"`
	History  []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields
}

// TerminalConnectionInfo is a lightweight liveness check returned by a
// provider status call.
type TerminalConnectionInfo struct {
	TerminalID      string                   `json:"terminal_id"`
	Provider        enums.TerminalProvider   `json:"provider"`
	ProviderDetails *TerminalProviderDetails `json:"provider_details,omitempty"`
	Status          enums.TerminalStatus     `json:"status"`
	Connected       bool                     `json:"connected"`
	ProviderStatus  string                   `json:"provider_status,omitempty"`
	CheckedAt       time.Time                `json:"checked_at"`
	Metadata        common.Metadata          `json:"metadata,omitempty"`
}
