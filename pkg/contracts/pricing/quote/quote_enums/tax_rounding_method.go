package quote_enums

type TaxRoundingMethod string

const (
	TaxRoundingMethodSumExactThenRound TaxRoundingMethod = "sum_exact_then_round"
	TaxRoundingMethodPerLineRound      TaxRoundingMethod = "per_line_round"
)

func (m TaxRoundingMethod) IsValid() bool {
	return m == TaxRoundingMethodSumExactThenRound || m == TaxRoundingMethodPerLineRound
}
func (m TaxRoundingMethod) String() string { return string(m) }
