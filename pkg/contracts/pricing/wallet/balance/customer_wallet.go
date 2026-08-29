package balance

import "time"

// CustomerWallet is the retail read model of every value instrument held by a
// customer. Its per-instrument ledgers remain authoritative.
type CustomerWallet struct {
	CustomerNumber string                `json:"customer_number"`
	Instruments    []WalletInstrument    `json:"instruments,omitempty"`
	Summary        CustomerWalletSummary `json:"summary"`
	CalculatedAt   time.Time             `json:"calculated_at"`
}
