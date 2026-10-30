package quote

// PromotionGroupedNetEventPeer records the actual original rounded line net,
// chosen units and floor result. FloorNumerator is the exact nonnegative
// product as a canonical decimal integer string, because two int64 inputs may
// exceed int64; FloorDenominator is original base units. For bundle allocation
// every peer in the actual allocator population is present, including zero
// shares, in original request order. Rank/award/tie are present only when the
// allocator actually ran. BOGO floor-only events do not claim such an event.
type PromotionGroupedNetEventPeer struct {
	RootLineID            string      `json:"root_line_id"`
	RequestOrder          int32       `json:"request_order"`
	OriginalLineNet       money.Money `json:"original_line_net"`
	OriginalBaseUnits     int64       `json:"original_base_units"`
	SelectedBaseUnits     int64       `json:"selected_base_units"`
	FloorNumerator        string      `json:"floor_numerator"`
	FloorDenominator      int64       `json:"floor_denominator"`
	FlooredSelectedNet    money.Money `json:"floored_selected_net"`
	RecordedNetReduction  money.Money `json:"recorded_net_reduction"`
	AllocationRank        *int64      `json:"allocation_rank,omitempty"`
	AllocationAwarded     *bool       `json:"allocation_awarded,omitempty"`
	AllocationTieBreakKey string      `json:"allocation_tie_break_key,omitempty"`
}
