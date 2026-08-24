package pricebook_enums

type PriceTaxInclusion string

const (
	PriceTaxInclusionInclusive PriceTaxInclusion = "tax_inclusive"
	PriceTaxInclusionExclusive PriceTaxInclusion = "tax_exclusive"
)

func (i PriceTaxInclusion) IsValid() bool {
	return i == PriceTaxInclusionInclusive || i == PriceTaxInclusionExclusive
}
func (i PriceTaxInclusion) String() string { return string(i) }
