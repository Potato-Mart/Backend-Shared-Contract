package coupon

type RedemptionLimits struct {
	Total          *int64 `json:"total,omitempty"`
	PerEntitlement *int64 `json:"per_entitlement,omitempty"`
}
