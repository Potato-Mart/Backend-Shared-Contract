package inventory

// KnownLotPhysicalEvidence identifies a physical allocation or issue from a
// verified lot. Unit IDs are present only when the owner actually identified
// those units; the lot evidence does not by itself prove a verified cost.
type KnownLotPhysicalEvidence struct {
	LotID        string   `json:"lot_id"`
	StockUnitIDs []string `json:"stock_unit_ids,omitempty"`
}
