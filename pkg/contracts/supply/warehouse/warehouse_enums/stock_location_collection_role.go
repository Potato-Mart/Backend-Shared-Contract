package warehouse_enums

// StockLocationCollectionRole identifies the primary or overflow role of a
// collection's eligible location.
type StockLocationCollectionRole string

const (
	StockLocationCollectionPrimary  StockLocationCollectionRole = "PRIMARY"
	StockLocationCollectionOverflow StockLocationCollectionRole = "OVERFLOW"
)

func (r StockLocationCollectionRole) IsValid() bool {
	switch r {
	case StockLocationCollectionPrimary, StockLocationCollectionOverflow:
		return true
	}
	return false
}

func (r StockLocationCollectionRole) String() string { return string(r) }
