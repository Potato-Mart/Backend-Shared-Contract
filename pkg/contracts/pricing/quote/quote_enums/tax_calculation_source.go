package quote_enums

type TaxCalculationSource string

const (
	TaxCalculationSourceInclusiveExtraction TaxCalculationSource = "inclusive_extraction"
	TaxCalculationSourceExclusiveAddition   TaxCalculationSource = "exclusive_addition"
	TaxCalculationSourceZeroRated           TaxCalculationSource = "zero_rated"
	TaxCalculationSourceExempt              TaxCalculationSource = "exempt"
	TaxCalculationSourceOutOfScope          TaxCalculationSource = "out_of_scope"
)

func (s TaxCalculationSource) IsValid() bool {
	switch s {
	case TaxCalculationSourceInclusiveExtraction, TaxCalculationSourceExclusiveAddition, TaxCalculationSourceZeroRated, TaxCalculationSourceExempt, TaxCalculationSourceOutOfScope:
		return true
	}
	return false
}
func (s TaxCalculationSource) String() string { return string(s) }
