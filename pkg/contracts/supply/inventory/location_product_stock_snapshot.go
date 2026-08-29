package inventory

// LocationProductStockSnapshot qualifies product stock by depot and location.
type LocationProductStockSnapshot struct {
	DepotCode    string                       `json:"depot_code"`
	LocationCode string                       `json:"location_code"`
	Quantities   ProductStockQuantitySnapshot `json:"quantities"`
}
