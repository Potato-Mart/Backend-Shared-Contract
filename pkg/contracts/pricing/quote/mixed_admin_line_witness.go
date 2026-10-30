package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

// MixedAdminLineWitness binds one mixed-source line's tax-group accounting to
// MixedAdminDiscountWitness. GlobalWitnessHash is lowercase SHA-256 of the
// complete global witness encoded as structural canonical JSON: object keys
// sorted, exact JSON integers preserved, array order preserved, and no quote
// revision embedded in the witness. The transaction service validates the
// global and group proofs against authenticated predecessor and quote evidence.
// This records accounting reductions, not physical unit tax reclassification.
type MixedAdminLineWitness struct {
	Version           int32                     `json:"version"`
	Operation         string                    `json:"operation"`
	LineID            string                    `json:"line_id"`
	GlobalWitnessHash string                    `json:"global_witness_hash"`
	PreAdminGross     money.Money               `json:"pre_admin_gross"`
	PreAdminTax       money.Money               `json:"pre_admin_tax"`
	CarriedGross      money.Money               `json:"carried_gross"`
	CarriedTax        money.Money               `json:"carried_tax"`
	AppliedGross      money.Money               `json:"applied_gross"`
	AppliedTax        money.Money               `json:"applied_tax"`
	Buckets           []MixedAdminBucketWitness `json:"buckets"`
}

// MixedAdminBucketWitness records one predecessor tax treatment and original
// tie-break anchor. Carry allocates retained line base and tax reductions
// independently over predecessor bucket reductions by largest remainder using
// the canonical treatment/anchor key. A new line discount is allocated over
// pre-admin bucket gross by the same rule. Its nominal tax is half-up at the
// frozen effective gross tax rate (inclusive n/d, exclusive n/(d+n)), then
// clamped to [max(0, gross-pre-admin base), min(gross, pre-admin tax)]. The
// clamp keeps both accounting amounts nonnegative, including ghost groups.
// Exact rounding evidence, bounds and applied amounts are independently
// checked; empty optional proofs are permitted only where the operation does
// not require the corresponding allocation.
// The canonical tie key is `len(category):category|len(rule):rule|revision|`
// `inclusion|calculation_source|first_source_unit`, with UTF-8 byte lengths
// in decimal and revision and first_source_unit zero-padded to 20 digits.
type MixedAdminBucketWitness struct {
	TaxCategoryCode      string                            `json:"tax_category_code"`
	TaxRuleID            string                            `json:"tax_rule_id"`
	TaxRuleRevision      int64                             `json:"tax_rule_revision"`
	InclusionBasis       pricebook_enums.PriceTaxInclusion `json:"inclusion_basis"`
	RateNumerator        int64                             `json:"rate_numerator"`
	RateDenominator      int64                             `json:"rate_denominator"`
	CalculationSource    quote_enums.TaxCalculationSource  `json:"calculation_source"`
	FirstSourceUnit      int64                             `json:"first_source_unit"`
	PreAdminBase         money.Money                       `json:"pre_admin_base"`
	PreAdminTax          money.Money                       `json:"pre_admin_tax"`
	CarriedBase          money.Money                       `json:"carried_base"`
	CarriedTax           money.Money                       `json:"carried_tax"`
	CarryBaseRounding    *RoundingEvidence                 `json:"carry_base_rounding,omitempty"`
	CarryTaxRounding     *RoundingEvidence                 `json:"carry_tax_rounding,omitempty"`
	AppliedGross         money.Money                       `json:"applied_gross"`
	AppliedBase          money.Money                       `json:"applied_base"`
	AppliedTax           money.Money                       `json:"applied_tax"`
	AppliedGrossRounding *RoundingEvidence                 `json:"applied_gross_rounding,omitempty"`
	NominalTaxRounding   *RoundingEvidence                 `json:"nominal_tax_rounding,omitempty"`
	TaxLowerBound        money.Money                       `json:"tax_lower_bound"`
	TaxUpperBound        money.Money                       `json:"tax_upper_bound"`
}
