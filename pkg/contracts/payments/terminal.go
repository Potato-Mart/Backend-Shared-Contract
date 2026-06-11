// Package payments contains the shared contracts for EFTPOS payment
// terminals and the transactions they process.
package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v4/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v4/pkg/enums"
)

// Terminal is an EFTPOS device registered for POS use.
type Terminal struct {
	ID       string                 `json:"id"`
	TenantID string                 `json:"tenant_id"`
	StoreID  string                 `json:"store_id,omitempty"`
	Provider enums.TerminalProvider `json:"provider"`

	ConnectionMode enums.TerminalConnectionMode `json:"connection_mode,omitempty"`

	ProviderMerchantID string `json:"provider_merchant_id,omitempty"`
	ProviderStoreID    string `json:"provider_store_id,omitempty"`
	ProviderTerminalID string `json:"provider_terminal_id,omitempty"`
	ProviderDeviceID   string `json:"provider_device_id,omitempty"`
	TerminalNickname   string `json:"terminal_nickname,omitempty"`
	ProviderBaseURL    string `json:"provider_base_url,omitempty"`

	Status enums.TerminalStatus `json:"status"`

	RegisteredAt   *time.Time `json:"registered_at,omitempty"`
	DeregisteredAt *time.Time `json:"deregistered_at,omitempty"`
	LastSeenAt     *time.Time `json:"last_seen_at,omitempty"`

	Metadata common.Metadata `json:"metadata,omitempty"`

	common.AuditFields
}

// RegisterTerminalRequest saves a provider terminal association for a
// tenant/store.
type RegisterTerminalRequest struct {
	TenantID           string                       `json:"tenant_id"`
	StoreID            string                       `json:"store_id,omitempty"`
	Provider           enums.TerminalProvider       `json:"provider"`
	ConnectionMode     enums.TerminalConnectionMode `json:"connection_mode,omitempty"`
	ProviderMerchantID string                       `json:"provider_merchant_id,omitempty"`
	ProviderStoreID    string                       `json:"provider_store_id,omitempty"`
	ProviderTerminalID string                       `json:"provider_terminal_id,omitempty"`
	ProviderDeviceID   string                       `json:"provider_device_id,omitempty"`
	TerminalNickname   string                       `json:"terminal_nickname,omitempty"`
	ProviderBaseURL    string                       `json:"provider_base_url,omitempty"`
	Metadata           common.Metadata              `json:"metadata,omitempty"`
}

// RegisterTerminalResponse is returned after a terminal association is
// stored and can be used by the POS.
type RegisterTerminalResponse struct {
	Terminal Terminal `json:"terminal"`
}

// TerminalConnectionInfo is a lightweight liveness check returned by a
// provider status call.
type TerminalConnectionInfo struct {
	TerminalID         string                 `json:"terminal_id"`
	Provider           enums.TerminalProvider `json:"provider"`
	ProviderMerchantID string                 `json:"provider_merchant_id,omitempty"`
	ProviderStoreID    string                 `json:"provider_store_id,omitempty"`
	ProviderTerminalID string                 `json:"provider_terminal_id,omitempty"`
	ProviderDeviceID   string                 `json:"provider_device_id,omitempty"`
	Status             enums.TerminalStatus   `json:"status"`
	Connected          bool                   `json:"connected"`
	ProviderStatus     string                 `json:"provider_status,omitempty"`
	CheckedAt          time.Time              `json:"checked_at"`
	Metadata           common.Metadata        `json:"metadata,omitempty"`
}
