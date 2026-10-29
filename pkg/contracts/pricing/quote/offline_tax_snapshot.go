package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/pricing/quote/quote_enums"
)

// OfflineTaxSnapshot records the signed book's tax class and rate, and the
// basis and rounding used for the charged amount. It does not assert an online
// tax-rule ID or revision. TaxableBase is the line's amount excluding tax;
// AllocatedTax is the tax allocated to the line, both in minor currency units.
type OfflineTaxSnapshot struct {
	TaxClass            string                            `json:"tax_class"`
	RateBasisPoints     int64                             `json:"rate_basis_points"`
	SourceTaxInclusion  pricebook_enums.PriceTaxInclusion `json:"source_tax_inclusion"`
	ChargedTaxInclusion pricebook_enums.PriceTaxInclusion `json:"charged_tax_inclusion"`
	RoundingMethod      quote_enums.TaxRoundingMethod     `json:"rounding_method"`
	TaxableBase         money.Money                       `json:"taxable_base"`
	AllocatedTax        money.Money                       `json:"allocated_tax"`
}
