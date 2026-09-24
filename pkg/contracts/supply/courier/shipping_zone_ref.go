package courier

// ShippingZoneRef identifies an Orders-owned shipping zone without copying its
// mutable name or coverage rules. Resolve the current Orders zone when making
// coverage decisions; this reference alone grants no coverage.
type ShippingZoneRef struct {
	ID string `json:"id"`
}
