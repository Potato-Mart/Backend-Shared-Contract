package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
)

// AmendmentAllocationAdjustment binds an accounting adjustment on a retained
// amendment line to the immutable predecessor and native unit selection.
// Version 1 preserves the transaction service's component-proportional package
// allocation: the predecessor component amount is divided by PackageCount and
// any minor remainder is assigned to earlier package ordinals. Version 1 covers
// whole packages only, so the within-package range is 1 through UnitsPerPackage.
// Each Delta is signed in
// minor units and equals the retained line amount minus the sum of its actual
// selected source unit evidence. TaxAdjustments partition TaxableBaseDelta and
// TaxAmountDelta by predecessor tax treatment; they are accounting evidence,
// not a claim that an unselected source was physically retained.
//
// The source ranges, adjustment, predecessor binding and selected native
// units must be covered by the accepted quote fingerprint and retained for
// replay. Transaction services resolve the immutable predecessor, recompute
// the proportional target from it and the selected ordinals, verify every delta
// and reject overlapping selections or negative adjusted totals. Arithmetic
// equality to a caller-supplied target alone is not authority to adjust money.
// PhysicalRoundingProof, when set to "inherited_predecessor_v1", means
// retained source and unit amounts preserve the predecessor's physical
// rounding rather than rerounding the new line. Unit range ordinals are
// rebased; exact inputs, amounts, awards and original tie identities remain
// bound to the selected predecessor ordinals. When clipping a unit run, its
// inherited first rank and original-ordinal tie key advance by the clipped
// offset. Source consideration, source tax and target-promotion aggregate
// proofs sum the inherited unit exact values and rounded amounts; their rank
// is zero, tie key empty, and remainder applied is rounded sum minus floor of
// the exact sum (which may be negative or greater than one). They make no
// fresh largest-remainder claim. Consumers must authenticate and recursively
// validate the immutable predecessor and every selected unit before accepting
// this mode; the signed accounting deltas remain separate. An omitted mode
// preserves the legacy rounding semantics.
type AmendmentAllocationAdjustment struct {
	Version                           int32                      `json:"version"`
	SourceVersionID                   string                     `json:"source_version_id"`
	SourceItemID                      string                     `json:"source_item_id"`
	SourceComponentID                 string                     `json:"source_component_id"`
	SourceQuoteID                     string                     `json:"source_quote_id"`
	SourceQuoteRevision               string                     `json:"source_quote_revision"`
	SourceQuoteFingerprint            string                     `json:"source_quote_fingerprint"`
	SourceLineID                      string                     `json:"source_line_id"`
	PhysicalRoundingProof             string                     `json:"physical_rounding_proof,omitempty"`
	AdminDiscountWitness              *MixedAdminLineWitness     `json:"admin_discount_witness,omitempty"`
	PackageOption                     packaging.PackageOptionRef `json:"package_option"`
	PackageCount                      int64                      `json:"package_count"`
	UnitsPerPackage                   int64                      `json:"units_per_package"`
	FirstPackageOrdinal               int64                      `json:"first_package_ordinal"`
	LastPackageOrdinal                int64                      `json:"last_package_ordinal"`
	FirstWithinPackageBaseUnitOrdinal int64                      `json:"first_within_package_base_unit_ordinal"`
	LastWithinPackageBaseUnitOrdinal  int64                      `json:"last_within_package_base_unit_ordinal"`
	ListAmountDelta                   money.Money                `json:"list_amount_delta"`
	DiscountAmountDelta               money.Money                `json:"discount_amount_delta"`
	ConsiderationDelta                money.Money                `json:"consideration_delta"`
	TaxableBaseDelta                  money.Money                `json:"taxable_base_delta"`
	TaxAmountDelta                    money.Money                `json:"tax_amount_delta"`
	LineTotalDelta                    money.Money                `json:"line_total_delta"`
	TaxAdjustments                    []AmendmentTaxAdjustment   `json:"tax_adjustments,omitempty"`
}
