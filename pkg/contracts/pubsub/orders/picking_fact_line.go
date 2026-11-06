package orders

// PickingFactLine is one order line's picking outcome in base units.
type PickingFactLine struct {
	OrderItemID       string `json:"order_item_id"`
	SKUCode           string `json:"sku_code"`
	OrderedBaseUnits  int64  `json:"ordered_base_units"`
	PickedBaseUnits   int64  `json:"picked_base_units"`
	ShortageBaseUnits int64  `json:"shortage_base_units"`
}
