package product_enums

// StorefrontStockState is the aggregate stock state safe for public product
// cards and detail pages. It deliberately contains no quantity, depot, lot, or
// location identity.
type StorefrontStockState string

const (
	StorefrontStockStateUnknown    StorefrontStockState = "unknown"
	StorefrontStockStateInStock    StorefrontStockState = "in_stock"
	StorefrontStockStateOutOfStock StorefrontStockState = "out_of_stock"
)

// IsValid reports whether s is a known StorefrontStockState value.
func (s StorefrontStockState) IsValid() bool {
	switch s {
	case StorefrontStockStateUnknown, StorefrontStockStateInStock, StorefrontStockStateOutOfStock:
		return true
	}
	return false
}

func (s StorefrontStockState) String() string { return string(s) }
