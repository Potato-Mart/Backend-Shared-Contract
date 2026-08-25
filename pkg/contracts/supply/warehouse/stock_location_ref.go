package warehouse

// StockLocationRef qualifies a location code by its depot.
type StockLocationRef struct {
	DepotCode    string `json:"depot_code"`
	LocationCode string `json:"location_code"`
}
