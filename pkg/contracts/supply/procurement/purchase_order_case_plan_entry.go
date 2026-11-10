package procurement

// PurchaseOrderCasePlanEntry plans, before arrival, how many cartons of one
// purchase order line go to one case location.
type PurchaseOrderCasePlanEntry struct {
	// DepotCode is the depot of the case location. Blank means the order's
	// depot.
	DepotCode    string `json:"depot_code,omitempty"`
	LocationCode string `json:"location_code"`
	// Cartons is the number of whole cartons planned for the location, at
	// least 1.
	Cartons int `json:"cartons"`
}
