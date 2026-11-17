package fulfilment

// PackageSubstitutionPart is one source of the loose units in an assembled
// case (a PackageSubstitutionSnapshot with reason CASE_ASSEMBLED_FROM_LOOSE):
// BaseUnits taken under one stock allocation from one bucket and lot.
// AllocationID, BucketID and LotID are required, BaseUnits is positive, and
// StockUnitIDs lists the scanned stock units when they are tracked.
type PackageSubstitutionPart struct {
	AllocationID string   `json:"allocation_id"`
	BucketID     string   `json:"bucket_id"`
	LotID        string   `json:"lot_id"`
	BaseUnits    int64    `json:"base_units"`
	StockUnitIDs []string `json:"stock_unit_ids,omitempty"`
}
