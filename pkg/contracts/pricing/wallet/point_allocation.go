package wallet

import "time"

// PointAllocation records one earned-points row consumed by an operational
// redemption or expiry entry.
type PointAllocation struct {
	LedgerEntryID string     `json:"ledger_entry_id"`
	Points        int        `json:"points"`
	ExpiresAt     *time.Time `json:"expires_at,omitempty"`
}
