package supply

// POSUnknownPhysicalReturnSource binds sold units to one recorded return.
type POSUnknownPhysicalReturnSource struct {
	ReceivingID string `json:"receiving_id"`
	BaseUnits   int64  `json:"base_units"`
}
