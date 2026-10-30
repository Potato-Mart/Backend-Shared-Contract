package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

// PromotionCandidatePeer identifies original eligible facts and every selected
// ordinal interval for this (requirement, root line, role). Spans partition the
// full manifest group range without gaps or overlapping units. The exact unit
// basis is the original approved book basis, before any promotion. Source and
// tax identities are authenticated against frozen root evidence, not inferred
// from an adjusted scalar line. The stage-zero proof additionally requires an
// independently authenticated initial source and absence of an earlier
// monetary stage on these units. SourceEvidenceID is a Pricing-owned opaque
// identity for the complete original source, authenticated by the owner against
// the immutable root decision. It excludes the finalized quote revision and
// must match the corresponding source-before evidence when that is present.
// TaxRate describes the recorded original tax calculation; the separate
// ExclusiveRate operands are the actual frozen rule values used by Pricing's
// inclusive from-net conversion. Consumers must not infer one rate from the
// other. Denominators and arithmetic are verified against the owner source.
type PromotionCandidatePeer struct {
	RootLineID                string                            `json:"root_line_id"`
	SourceEvidenceID          string                            `json:"source_evidence_id"`
	SKUCode                   string                            `json:"sku_code"`
	PackageOption             packaging.PackageOptionRef        `json:"package_option"`
	UnitsPerPackage           int64                             `json:"units_per_package"`
	UnitsPerGroup             int64                             `json:"units_per_group"`
	OriginalLineBaseUnits     int64                             `json:"original_line_base_units"`
	OriginalLinePackageCount  int64                             `json:"original_line_package_count"`
	ExactOriginalLineBasis    PromotionExactMinor               `json:"exact_original_line_basis"`
	RoundedOriginalLineBasis  money.Money                       `json:"rounded_original_line_basis"`
	SourceBefore              *PromotionSourceBeforeEvidence    `json:"source_before,omitempty"`
	QualifierTransitions      []PromotionQualifierIdentityProof `json:"qualifier_transitions,omitempty"`
	Currency                  money.CurrencyCode                `json:"currency"`
	CurrencyExponent          money.CurrencyExponent            `json:"currency_exponent"`
	BookTaxInclusion          pricebook_enums.PriceTaxInclusion `json:"book_tax_inclusion"`
	TaxCategoryCode           string                            `json:"tax_category_code"`
	TaxRuleID                 string                            `json:"tax_rule_id"`
	TaxRuleRevision           int64                             `json:"tax_rule_revision"`
	TaxCalculationSource      quote_enums.TaxCalculationSource  `json:"tax_calculation_source"`
	TaxRateNumerator          int64                             `json:"tax_rate_numerator"`
	TaxRateDenominator        int64                             `json:"tax_rate_denominator"`
	ExclusiveRateNumerator    int64                             `json:"exclusive_rate_numerator"`
	ExclusiveRateDenominator  int64                             `json:"exclusive_rate_denominator"`
	EligibleUnitBasis         PromotionExactMinor               `json:"eligible_unit_basis"`
	PriceBookCode             string                            `json:"price_book_code"`
	PriceBookRevision         int64                             `json:"price_book_revision"`
	PriceEntryID              string                            `json:"price_entry_id"`
	PriceEntryRevision        int64                             `json:"price_entry_revision"`
	PackagePriceEntryID       string                            `json:"package_price_entry_id,omitempty"`
	PackagePriceEntryRevision int64                             `json:"package_price_entry_revision,omitempty"`
	Spans                     []PromotionCandidatePeerSpan      `json:"spans"`
}
