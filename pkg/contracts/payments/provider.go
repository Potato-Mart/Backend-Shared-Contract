package payments

import "encoding/json"

// TerminalProviderDetails groups provider-side terminal identifiers and
// connection hints without hiding the platform TerminalID from parent DTOs.
type TerminalProviderDetails struct {
	MerchantID string `json:"merchant_id,omitempty"`
	StoreID    string `json:"store_id,omitempty"`
	TerminalID string `json:"terminal_id,omitempty"`
	DeviceID   string `json:"device_id,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
	BaseURL    string `json:"base_url,omitempty"`
}

// ProviderOperationContext carries provider request correlation and retry
// metadata shared by terminal transactions and settlements.
type ProviderOperationContext struct {
	RequestID         string `json:"request_id,omitempty"`
	MerchantReference string `json:"merchant_reference,omitempty"`
	IdempotencyKey    string `json:"idempotency_key,omitempty"`
}

// ProviderPayloads keeps raw provider messages together for diagnostics and
// future compatibility. Normalized contract fields remain the source of truth.
type ProviderPayloads struct {
	Request             json.RawMessage `json:"request,omitempty"`
	Response            json.RawMessage `json:"response,omitempty"`
	Notification        json.RawMessage `json:"notification,omitempty"`
	DisplayNotification json.RawMessage `json:"display_notification,omitempty"`
}
