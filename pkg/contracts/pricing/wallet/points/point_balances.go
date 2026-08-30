package points

// PointBalances is the shared headline balance set used by customer-safe and
// operational points projections. Ledger entries remain the source of truth.
type PointBalances struct {
	TotalPoints     int `json:"total_points"`
	ReservedPoints  int `json:"reserved_points"`
	AvailablePoints int `json:"available_points"`
	PointDebt       int `json:"point_debt"`
	ExpiringPoints  int `json:"expiring_points"`
}
