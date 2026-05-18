// Package payments contains the shared contracts for EFTPOS payment
// terminals and the transactions they process.
//
// The package is provider-aware but SDK-free. The current terminal
// contract is shaped for Adyen Terminal API / Cloud Device API while
// staying small enough for other terminal providers to reuse through
// Terminal.Provider and Metadata.
//
// Secrets policy: this package never holds API keys, HMAC keys, or any
// credential material. Provider-side secrets live exclusively in the
// secret store of the service that signs or sends outbound calls. Only
// public, non-secret identifiers such as merchant account, POIID, SaleID,
// ServiceID, PSP reference, and terminal base URL are carried here.
package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// Terminal is an EFTPOS device registered for POS use. For Adyen, POIID
// is the terminal identifier used in SaleToPOIRequest.MessageHeader and
// in Cloud Device API paths. MerchantAccount identifies the Adyen
// merchant account that owns the device, and SaleID identifies the POS
// system component sending Terminal API requests.
type Terminal struct {
	ID       string                 `json:"id"`
	TenantID string                 `json:"tenant_id"`
	StoreID  string                 `json:"store_id,omitempty"`
	Provider enums.TerminalProvider `json:"provider"`

	ConnectionMode enums.TerminalConnectionMode `json:"connection_mode,omitempty"`

	// Provider-side identifiers. All are non-secret and safe to
	// transport in API responses.
	MerchantAccount  string `json:"merchant_account,omitempty"`
	POIID            string `json:"poi_id,omitempty"`
	SaleID           string `json:"sale_id,omitempty"`
	TerminalNickname string `json:"terminal_nickname,omitempty"`

	// TerminalAPIBaseURL optionally pins the regional endpoint used by
	// the service layer. For Adyen AU live traffic this can be
	// https://terminal-api-live-au.adyen.com or the Cloud Device API
	// equivalent. Leave empty when the Adyen SDK/environment config owns
	// endpoint selection.
	TerminalAPIBaseURL string `json:"terminal_api_base_url,omitempty"`

	Status enums.TerminalStatus `json:"status"`

	RegisteredAt   *time.Time `json:"registered_at,omitempty"`
	DeregisteredAt *time.Time `json:"deregistered_at,omitempty"`
	LastSeenAt     *time.Time `json:"last_seen_at,omitempty"`

	Metadata common.Metadata `json:"metadata,omitempty"`

	common.AuditFields
}

// RegisterTerminalRequest saves a provider terminal association for a
// tenant/store. Adyen terminals are registered from merchant account,
// POIID, SaleID, and connection mode selected during setup.
type RegisterTerminalRequest struct {
	TenantID           string                       `json:"tenant_id"`
	StoreID            string                       `json:"store_id,omitempty"`
	Provider           enums.TerminalProvider       `json:"provider"`
	ConnectionMode     enums.TerminalConnectionMode `json:"connection_mode,omitempty"`
	MerchantAccount    string                       `json:"merchant_account,omitempty"`
	POIID              string                       `json:"poi_id,omitempty"`
	SaleID             string                       `json:"sale_id,omitempty"`
	TerminalNickname   string                       `json:"terminal_nickname,omitempty"`
	TerminalAPIBaseURL string                       `json:"terminal_api_base_url,omitempty"`
	Metadata           common.Metadata              `json:"metadata,omitempty"`
}

// RegisterTerminalResponse is returned after a terminal association is
// stored and can be used by the POS.
type RegisterTerminalResponse struct {
	Terminal Terminal `json:"terminal"`
}

// TerminalConnectionInfo is a lightweight liveness check returned by a
// provider status call, such as Adyen's connected terminals or device
// status endpoints.
type TerminalConnectionInfo struct {
	TerminalID      string                 `json:"terminal_id"`
	Provider        enums.TerminalProvider `json:"provider"`
	MerchantAccount string                 `json:"merchant_account,omitempty"`
	POIID           string                 `json:"poi_id,omitempty"`
	Status          enums.TerminalStatus   `json:"status"`
	Connected       bool                   `json:"connected"`
	ProviderStatus  string                 `json:"provider_status,omitempty"`
	CheckedAt       time.Time              `json:"checked_at"`
	Metadata        common.Metadata        `json:"metadata,omitempty"`
}
