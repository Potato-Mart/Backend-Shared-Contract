package warehouse_enums

// StockAvailabilityDirection identifies a customer-visible zero crossing.
type StockAvailabilityDirection string

const (
	StockAvailabilityOutOfStock StockAvailabilityDirection = "OUT_OF_STOCK"
	StockAvailabilityRestocked  StockAvailabilityDirection = "RESTOCKED"
)

func (d StockAvailabilityDirection) IsValid() bool {
	switch d {
	case StockAvailabilityOutOfStock, StockAvailabilityRestocked:
		return true
	}
	return false
}

func (d StockAvailabilityDirection) String() string { return string(d) }
