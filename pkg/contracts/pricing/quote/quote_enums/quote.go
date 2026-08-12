package quote_enums

// TaxCalculationSource records how a line's tax amount was produced from the
// consideration that was in force.
type TaxCalculationSource string

const (
	// TaxCalculationSourceInclusiveExtraction extracts tax out of a
	// tax-inclusive consideration using the exact inclusive fraction.
	TaxCalculationSourceInclusiveExtraction TaxCalculationSource = "inclusive_extraction"
	// TaxCalculationSourceExclusiveAddition adds tax onto a tax-exclusive
	// consideration using the exact rate.
	TaxCalculationSourceExclusiveAddition TaxCalculationSource = "exclusive_addition"
	// TaxCalculationSourceZeroRated is a supply taxed at a zero rate.
	TaxCalculationSourceZeroRated TaxCalculationSource = "zero_rated"
	// TaxCalculationSourceExempt is a supply that is not taxable.
	TaxCalculationSourceExempt TaxCalculationSource = "exempt"
	// TaxCalculationSourceOutOfScope is a supply outside the market's tax
	// system.
	TaxCalculationSourceOutOfScope TaxCalculationSource = "out_of_scope"
)

// IsValid reports whether s is a known TaxCalculationSource.
func (s TaxCalculationSource) IsValid() bool {
	switch s {
	case TaxCalculationSourceInclusiveExtraction, TaxCalculationSourceExclusiveAddition,
		TaxCalculationSourceZeroRated, TaxCalculationSourceExempt, TaxCalculationSourceOutOfScope:
		return true
	}
	return false
}

func (s TaxCalculationSource) String() string { return string(s) }

// TaxRoundingMethod records how exact line tax was reduced to whole minor
// units across a document.
type TaxRoundingMethod string

const (
	// TaxRoundingMethodSumExactThenRound sums the exact line tax of every
	// taxable line first and rounds the document total once, then allocates
	// the rounded minor units back to the lines.
	TaxRoundingMethodSumExactThenRound TaxRoundingMethod = "sum_exact_then_round"
	// TaxRoundingMethodPerLineRound rounds each line's tax independently.
	TaxRoundingMethodPerLineRound TaxRoundingMethod = "per_line_round"
)

// IsValid reports whether m is a known TaxRoundingMethod.
func (m TaxRoundingMethod) IsValid() bool {
	switch m {
	case TaxRoundingMethodSumExactThenRound, TaxRoundingMethodPerLineRound:
		return true
	}
	return false
}

func (m TaxRoundingMethod) String() string { return string(m) }

// RoundingMode names the exact rounding rule that produced a stored minor
// amount from an exact rational value.
type RoundingMode string

const (
	// RoundingModeHalfUp rounds to the currency exponent with halves going
	// up.
	RoundingModeHalfUp RoundingMode = "half_up"
	// RoundingModeHalfCentUp rounds a tax total to the nearest minor unit
	// with an exact half minor unit going up.
	RoundingModeHalfCentUp RoundingMode = "half_cent_up"
	// RoundingModeCharmNineUp keeps an exact ten-minor-unit value and
	// otherwise rounds upward to one minor unit below the next ten.
	RoundingModeCharmNineUp RoundingMode = "charm_nine_up"
	// RoundingModeCashIncrement rounds a final all-cash tender to the
	// market's cash increment.
	RoundingModeCashIncrement RoundingMode = "cash_increment"
	// RoundingModeLargestRemainder allocates already-rounded document minor
	// units back to lines by largest remainder.
	RoundingModeLargestRemainder RoundingMode = "largest_remainder"
)

// IsValid reports whether m is a known RoundingMode.
func (m RoundingMode) IsValid() bool {
	switch m {
	case RoundingModeHalfUp, RoundingModeHalfCentUp, RoundingModeCharmNineUp,
		RoundingModeCashIncrement, RoundingModeLargestRemainder:
		return true
	}
	return false
}

func (m RoundingMode) String() string { return string(m) }

// CostComparison records whether a cashier-entered custom price was
// below the current cost when the override was frozen.
type CostComparison string

const (
	// CostComparisonBelowCost means the override was below the
	// comparable current cost. The sale is allowed and warned about.
	CostComparisonBelowCost CostComparison = "below_cost"
	// CostComparisonAtOrAboveCost means the override was at or above
	// the comparable current cost.
	CostComparisonAtOrAboveCost CostComparison = "at_or_above_cost"
	// CostComparisonUnavailable means no comparable cost existed, for
	// example because cost was absent or held in another currency.
	CostComparisonUnavailable CostComparison = "unavailable"
)

// IsValid reports whether r is a known CostComparison.
func (r CostComparison) IsValid() bool {
	switch r {
	case CostComparisonBelowCost, CostComparisonAtOrAboveCost, CostComparisonUnavailable:
		return true
	}
	return false
}

func (r CostComparison) String() string { return string(r) }
