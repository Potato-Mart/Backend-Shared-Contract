package supply

// POSDeficitEvidence represents an unfilled physical issue, with a lot only
// when the owner has actual lot evidence. Missing lot is not an inferred lot.
type POSDeficitEvidence struct {
	LotID string `json:"lot_id,omitempty"`
}
