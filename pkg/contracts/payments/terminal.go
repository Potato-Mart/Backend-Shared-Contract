// Package payments contains the shared contracts for EFTPOS payment
// terminals and the transactions they process.
//
// The package is provider-agnostic but is shaped to fit MX51 SCI's
// surface area (pairing, polled transactions, settlement, override
// flow). Other providers - MX51 Spice, MX51 SPI, or future PaaS - reuse
// the same types by setting Terminal.Provider.
//
// Secrets policy: this package never holds API keys, signing secrets,
// or any credential material. The provider-side secrets (e.g. MX51
// Signing Secret Part A and Part B, or the SCI Pairing API Key) live
// exclusively in the secret store of the service that signs the
// outbound HTTP calls. Only public, non-secret identifiers - pairing
// id, key id, base url, tid - are carried on the wire.
package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// Terminal is a paired EFTPOS device. There is one Terminal per
// successful pairing; an unpair flips Status to Unpaired but the row
// is kept so historical TerminalTransactions stay joinable.
//
// For MX51 SCI specifically, the fields PairingID, KeyID, TID and
// SciAPIBaseURL all come from the POST /v1/progress-pairing response
// and are mandatory for any subsequent transaction call. SciAPIBaseURL
// in particular is per-pairing - different terminals can return
// different URLs - and must be stored alongside the pairing row.
type Terminal struct {
	ID       string                 `json:"id"`
	TenantID string                 `json:"tenant_id"`
	StoreID  string                 `json:"store_id,omitempty"`
	Provider enums.TerminalProvider `json:"provider"`

	// Provider-side identifiers. All are non-secret and safe to
	// transport in API responses.
	PairingID        string `json:"pairing_id"`
	KeyID            string `json:"key_id"`
	TID              string `json:"tid,omitempty"`
	PairingNickname  string `json:"pairing_nickname,omitempty"`
	TerminalNickname string `json:"terminal_nickname,omitempty"`

	// SciAPIBaseURL is the per-pairing base URL returned at pairing
	// time. All subsequent SCI API calls for this terminal must be
	// directed at this URL.
	SciAPIBaseURL string `json:"sci_api_base_url,omitempty"`

	Status enums.TerminalStatus `json:"status"`

	PairedAt   *time.Time `json:"paired_at,omitempty"`
	UnpairedAt *time.Time `json:"unpaired_at,omitempty"`
	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`

	Metadata common.Metadata `json:"metadata,omitempty"`

	common.AuditFields
}

// PairingStartRequest initiates a pairing flow on the merchant's POS.
// The pairing code is read off the EFTPOS terminal screen by the
// merchant and typed into the POS.
type PairingStartRequest struct {
	TenantID        string                 `json:"tenant_id"`
	StoreID         string                 `json:"store_id,omitempty"`
	Provider        enums.TerminalProvider `json:"provider"`
	PairingCode     string                 `json:"pairing_code"`
	PairingNickname string                 `json:"pairing_nickname,omitempty"`
}

// PairingStartResponse is returned to the POS once the pairing has
// progressed on the provider side. ConfirmationCode is shown on the
// terminal at the same time and the merchant must verify they match
// before confirming on the device. The pairing only becomes Active
// once the merchant confirms on the terminal itself.
type PairingStartResponse struct {
	Terminal         Terminal `json:"terminal"`
	ConfirmationCode string   `json:"confirmation_code"`
}

// PairingInfo is the lightweight liveness check returned by the
// provider-equivalent of MX51's GET /v1/pairing-info. Use it before
// navigating to the pairing screen and before initiating a transaction.
type PairingInfo struct {
	TerminalID string               `json:"terminal_id"`
	PairingID  string               `json:"pairing_id"`
	Status     enums.TerminalStatus `json:"status"`
	CheckedAt  time.Time            `json:"checked_at"`
}
