package points

import "time"

// PointsSummary is a customer-safe wallet headline projection. The points
// ledger remains the source of truth.
type PointsSummary struct {
	TotalPoints     int        `json:"total_points"`
	ReservedPoints  int        `json:"reserved_points"`
	AvailablePoints int        `json:"available_points"`
	PointDebt       int        `json:"point_debt"`
	ExpiringPoints  int        `json:"expiring_points"`
	NextExpiryAt    *time.Time `json:"next_expiry_at,omitempty"`
	CalculatedAt    time.Time  `json:"calculated_at"`
}
