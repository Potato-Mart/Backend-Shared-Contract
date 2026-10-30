package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

// AmendmentTaxAdjustment records the signed accounting adjustment attributed
// to one frozen predecessor tax treatment. It does not assert that a unit
// with that treatment was physically retained. The adjustment entries sum to
// the parent adjustment's taxable-base and tax deltas. TargetTaxRounding and
// TargetBaseRounding evidence the target line tax and taxable base allocated
// across predecessor treatment groups independently by largest remainder.
// Exact weights are the corresponding frozen predecessor group amounts;
// equal remainders use the canonical group key (tax category, rule identity
// and revision, inclusion basis, calculation source, then first original unit
// ordinal). Zero predecessor totals cannot receive a positive target. Each
// signed delta is that target share less the selected unit evidence for the
// same predecessor group, and the adjusted share must remain nonnegative.
// FirstSourceUnit is the one-based original-line ordinal establishing the
// group's canonical tie-break. It remains unchanged in later amendments,
// including when the group has no physically retained source unit. It is
// optional for legacy records; consumers of repeated mixed-source amendments
// must verify it against the authenticated predecessor evidence.
type AmendmentTaxAdjustment struct {
	TaxCategoryCode    string                            `json:"tax_category_code"`
	TaxRuleID          string                            `json:"tax_rule_id"`
	TaxRuleRevision    int64                             `json:"tax_rule_revision"`
	InclusionBasis     pricebook_enums.PriceTaxInclusion `json:"inclusion_basis"`
	RateNumerator      int64                             `json:"rate_numerator"`
	RateDenominator    int64                             `json:"rate_denominator"`
	CalculationSource  quote_enums.TaxCalculationSource  `json:"calculation_source"`
	FirstSourceUnit    int64                             `json:"first_source_unit,omitempty"`
	TaxableBaseDelta   money.Money                       `json:"taxable_base_delta"`
	TaxAmountDelta     money.Money                       `json:"tax_amount_delta"`
	TargetBaseRounding RoundingEvidence                  `json:"target_base_rounding"`
	TargetTaxRounding  RoundingEvidence                  `json:"target_tax_rounding"`
}
