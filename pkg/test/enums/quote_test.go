package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/quote/quote_enums"
)

func TestQuoteEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "quoteenum.TaxCalculationSource", valid: []stringEnum{quote_enums.TaxCalculationSourceInclusiveExtraction, quote_enums.TaxCalculationSourceExclusiveAddition, quote_enums.TaxCalculationSourceZeroRated, quote_enums.TaxCalculationSourceExempt, quote_enums.TaxCalculationSourceOutOfScope}, invalid: quote_enums.TaxCalculationSource("__invalid__")},
		{name: "quoteenum.TaxRoundingMethod", valid: []stringEnum{quote_enums.TaxRoundingMethodSumExactThenRound, quote_enums.TaxRoundingMethodPerLineRound}, invalid: quote_enums.TaxRoundingMethod("__invalid__")},
		{name: "quoteenum.RoundingMode", valid: []stringEnum{quote_enums.RoundingModeHalfUp, quote_enums.RoundingModeHalfCentUp, quote_enums.RoundingModeCharmNineUp, quote_enums.RoundingModeCashIncrement, quote_enums.RoundingModeLargestRemainder}, invalid: quote_enums.RoundingMode("__invalid__")},
		{name: "quoteenum.CostComparison", valid: []stringEnum{quote_enums.CostComparisonBelowCost, quote_enums.CostComparisonAtOrAboveCost, quote_enums.CostComparisonUnavailable}, invalid: quote_enums.CostComparison("__invalid__")},
	})
}
