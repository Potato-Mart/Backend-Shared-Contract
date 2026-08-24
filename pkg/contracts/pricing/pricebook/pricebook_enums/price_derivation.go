package pricebook_enums

type PriceDerivation string

const (
	PriceDerivationManual                PriceDerivation = "manual"
	PriceDerivationSuggestedFromBaseCost PriceDerivation = "suggested_from_base_cost"
	PriceDerivationImported              PriceDerivation = "imported"
)

func (d PriceDerivation) IsValid() bool {
	switch d {
	case PriceDerivationManual, PriceDerivationSuggestedFromBaseCost, PriceDerivationImported:
		return true
	}
	return false
}
func (d PriceDerivation) String() string { return string(d) }
