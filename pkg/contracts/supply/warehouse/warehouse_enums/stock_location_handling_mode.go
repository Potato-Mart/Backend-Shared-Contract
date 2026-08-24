package warehouse_enums

// StockLocationHandlingMode describes the package forms accepted by a stock
// location.
type StockLocationHandlingMode string

const (
	StockLocationHandlingEach  StockLocationHandlingMode = "EACH"
	StockLocationHandlingCase  StockLocationHandlingMode = "CASE"
	StockLocationHandlingMixed StockLocationHandlingMode = "MIXED"
)

func (m StockLocationHandlingMode) IsValid() bool {
	switch m {
	case StockLocationHandlingEach, StockLocationHandlingCase, StockLocationHandlingMixed:
		return true
	}
	return false
}

func (m StockLocationHandlingMode) String() string { return string(m) }
