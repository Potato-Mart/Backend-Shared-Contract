package inventory

// ProductStockQuantitySnapshot is a base-unit stock projection for one
// product and one optional stock dimension.
type ProductStockQuantitySnapshot struct {
	OnHandBaseUnits      int64 `json:"on_hand_base_units"`
	AvailableBaseUnits   int64 `json:"available_base_units"`
	SellableBaseUnits    int64 `json:"sellable_base_units"`
	ReducedBaseUnits     int64 `json:"reduced_base_units"`
	ReservedBaseUnits    int64 `json:"reserved_base_units"`
	StagedBaseUnits      int64 `json:"staged_base_units"`
	QualityHoldBaseUnits int64 `json:"quality_hold_base_units"`
}
