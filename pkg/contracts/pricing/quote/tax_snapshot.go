package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/quote/quote_enums"
)

// TaxSnapshot is the frozen tax evidence for one line. The rate is kept as an
// exact numerator over denominator so an inclusive extraction and an exclusive
// addition are both reproducible without floats.
type TaxSnapshot struct {
	TaxCategoryCode string `json:"tax_category_code"`
	TaxRuleID       string `json:"tax_rule_id"`
	TaxRuleRevision int64  `json:"tax_rule_revision"`

	InclusionBasis pricebook_enums.PriceTaxInclusion `json:"inclusion_basis"`
	// RateNumerator over RateDenominator is the exact rate that was
	// applied, for example 1/11 to extract tax from a tax-inclusive
	// consideration or 1/10 to add tax to a tax-exclusive one.
	RateNumerator   int64 `json:"rate_numerator"`
	RateDenominator int64 `json:"rate_denominator"`

	TaxableBase  money.Money `json:"taxable_base"`
	AllocatedTax money.Money `json:"allocated_tax"`

	CalculationSource quote_enums.TaxCalculationSource `json:"calculation_source"`
	RoundingMethod    quote_enums.TaxRoundingMethod    `json:"rounding_method"`
}
