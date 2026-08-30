package points

import "time"

// PointsSummary is a customer-safe wallet headline projection. The points
// ledger remains the source of truth.
type PointsSummary struct {
	PointBalances
	NextExpiryAt *time.Time `json:"next_expiry_at,omitempty"`
	CalculatedAt time.Time  `json:"calculated_at"`
}
