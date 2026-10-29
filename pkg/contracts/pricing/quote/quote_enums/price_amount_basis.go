package quote_enums

// PriceAmountBasis identifies the quantity for which a price amount is stated.
type PriceAmountBasis string

const (
	PriceAmountBasisBaseUnit PriceAmountBasis = "base_unit"
	PriceAmountBasisPackage  PriceAmountBasis = "package"
)

func (b PriceAmountBasis) IsValid() bool {
	return b == PriceAmountBasisBaseUnit || b == PriceAmountBasisPackage
}

func (b PriceAmountBasis) String() string { return string(b) }
