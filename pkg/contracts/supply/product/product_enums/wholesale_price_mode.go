package product_enums

// WholesalePriceMode distinguishes products with an authoritative wholesale
// price from products that approved buyers may browse but must enquire about
// before they can be added to a priced cart.
type WholesalePriceMode string

const (
	WholesalePriceModeFixed     WholesalePriceMode = "fixed"
	WholesalePriceModeOnRequest WholesalePriceMode = "on_request"
)

// IsValid reports whether m is a known WholesalePriceMode value.
func (m WholesalePriceMode) IsValid() bool {
	switch m {
	case WholesalePriceModeFixed, WholesalePriceModeOnRequest:
		return true
	}
	return false
}

// String returns the wire value for m.
func (m WholesalePriceMode) String() string { return string(m) }
