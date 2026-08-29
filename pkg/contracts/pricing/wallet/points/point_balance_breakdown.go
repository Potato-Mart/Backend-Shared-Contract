package points

import "time"

// PointBalanceBreakdown is an operational points balance projection.
type PointBalanceBreakdown struct {
	CustomerNumber string `json:"customer_number"`
	PointBalances
	Buckets      []PointBucket `json:"buckets,omitempty"`
	CalculatedAt time.Time     `json:"calculated_at"`
}
