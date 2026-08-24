package pricebook_enums

// PriceVisibility controls whether a resolved price may be rendered to a
// customer. Hidden prices never appear in SellingProduct.
type PriceVisibility string

const (
	PriceVisibilityVisible PriceVisibility = "visible"
	PriceVisibilityHidden  PriceVisibility = "hidden"
)

func (v PriceVisibility) IsValid() bool {
	return v == PriceVisibilityVisible || v == PriceVisibilityHidden
}

func (v PriceVisibility) String() string { return string(v) }
