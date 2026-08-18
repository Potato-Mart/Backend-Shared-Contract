// Package payments contains the shared contracts for EFTPOS payment
// terminals and the transactions they process.
package terminal

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/metadata"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/payments/terminal/terminal_enums"
)

// Terminal is an EFTPOS device registered for POS use at one depot.
//
// DepotCode is the trading site: depots are the only site identity in the
// platform, so a terminal is never keyed by a store code. MarketCode and
// CountryCode are the denormalized market and country of that depot, carried
// so a geographically scoped staff query is a plain indexed match.
type Terminal struct {
	ID          string                          `json:"id"`
	TenantID    string                          `json:"tenant_id"`
	DepotCode   string                          `json:"depot_code,omitempty"`
	MarketCode  string                          `json:"market_code,omitempty"`
	CountryCode geography.CountryCode           `json:"country_code,omitempty"`
	Provider    terminal_enums.TerminalProvider `json:"provider"`

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
