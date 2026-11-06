package orders

// PickingFactPickerLine is one picker half's share of an order line.
type PickingFactPickerLine struct {
	OrderItemID       string `json:"order_item_id"`
	PickedBaseUnits   int64  `json:"picked_base_units"`
	ShortageBaseUnits int64  `json:"shortage_base_units"`
}
