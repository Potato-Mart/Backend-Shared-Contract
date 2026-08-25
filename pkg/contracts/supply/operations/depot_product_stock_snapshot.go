package operations

// DepotProductStockSnapshot qualifies product stock by depot.
type DepotProductStockSnapshot struct {
	DepotCode  string                       `json:"depot_code"`
	Quantities ProductStockQuantitySnapshot `json:"quantities"`
}
