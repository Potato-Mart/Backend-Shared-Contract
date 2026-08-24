package membership

// PointsPolicy is server-authored membership redemption policy metadata.
// Customer-specific available balances are Wallet-owned.
type PointsPolicy struct {
	PointsPerMinorUnit      int `json:"points_per_minor_unit"`
	MinimumEligibleBalance  int `json:"minimum_eligible_balance"`
	RedemptionStepPoints    int `json:"redemption_step_points"`
	MaximumRedemptionPoints int `json:"maximum_redemption_points"`
}
