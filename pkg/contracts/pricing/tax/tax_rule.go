package tax

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pricing/pricebook/pricebook_enums"
	"time"
)

// TaxRule records a revisioned market/category rate. Pricing owns validation,
// effective-window overlap and calculation; rates are exact rational numbers.
type TaxRule struct {
	ID              string                            `json:"id"`
	TaxCategoryCode string                            `json:"tax_category_code"`
	MarketCode      string                            `json:"market_code"`
	RateNumerator   int64                             `json:"rate_numerator"`
	RateDenominator int64                             `json:"rate_denominator"`
	InclusionBasis  pricebook_enums.PriceTaxInclusion `json:"inclusion_basis"`
	EffectiveFrom   time.Time                         `json:"effective_from"`
	EffectiveTo     *time.Time                        `json:"effective_to,omitempty"`
	Revision        int64                             `json:"revision"`
	audit.AuditFields
}
