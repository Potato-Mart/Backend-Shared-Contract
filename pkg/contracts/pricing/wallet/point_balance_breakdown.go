package wallet

import "time"

// PointBalanceBreakdown is an operational points balance projection.
type PointBalanceBreakdown struct {
	CustomerNumber  string        `json:"customer_number"`
	TotalPoints     int           `json:"total_points"`
	ReservedPoints  int           `json:"reserved_points"`
	AvailablePoints int           `json:"available_points"`
	PointDebt       int           `json:"point_debt"`
	ExpiringPoints  int           `json:"expiring_points"`
	Buckets         []PointBucket `json:"buckets,omitempty"`
	CalculatedAt    time.Time     `json:"calculated_at"`
}
