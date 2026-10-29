package inventory

// UnknownReturnSource binds a physical quantity to one recorded return.
// Its owner verifies that source and prevents cumulative over-attribution.
type UnknownReturnSource struct {
	ReceivingID string `json:"receiving_id"`
	BaseUnits   int64  `json:"base_units"`
}
