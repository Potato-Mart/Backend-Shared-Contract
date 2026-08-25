package terminal

// TerminalProviderDetails groups the external identifiers needed to associate
// a platform terminal with its payment-terminal provider. Provider endpoints,
// credentials, request correlation, and retry state stay owned by the
// provider integration rather than this shared domain contract.
type TerminalProviderDetails struct {
	MerchantID string `json:"merchant_id,omitempty"`
	StoreID    string `json:"store_id,omitempty"`
	TerminalID string `json:"terminal_id,omitempty"`
	DeviceID   string `json:"device_id,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
}
