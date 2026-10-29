package supply

// POSKnownLotPhysicalEvidence identifies a physically issued lot. The lot must
// be proved by the owner; this variant alone does not prove a verified cost.
type POSKnownLotPhysicalEvidence struct {
	LotID string `json:"lot_id"`
}
