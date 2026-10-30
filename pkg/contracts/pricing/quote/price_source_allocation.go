package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// PriceSourceAllocation records the immutable price and tax facts for a
// contiguous range of units on a mixed-source quote line. FirstUnit is
// one-based. Together, a line's allocations cover each of its BaseUnits once.
// Package evidence is present when the source uses a package price entry.
// The snapshot's fingerprint binds these allocations to the accepted quote.
//
// Source ranges record the original price decision and do not constitute
// authorization to refund individual source prices. Refund allocation remains
// owned by the transaction and refund services.
//
// Consideration and Tax.AllocatedTax are allocated amounts, not independently
// rounded unit prices. First sum exact tax for the complete document and round
// once; allocate the document tax across original lines using the established
// largest-remainder rule and line-ID tie-break. Then allocate each line's tax
// across its source ranges by largest remainder, breaking equal remainders by
// FirstUnit ascending. TaxRounding records the exact source tax, allocated tax,
// rank and stable tie-break identity. Rounding records the corresponding
// consideration evidence. A range's tax is computed from its own exact
// consideration, inclusion basis and tax rule. The range LineTotal is its
// allocated consideration plus allocated tax only for an exclusive source.
type PriceSourceAllocation struct {
	FirstUnit                 int64                      `json:"first_unit"`
	BaseUnits                 int64                      `json:"base_units"`
	PriceBookCode             string                     `json:"price_book_code"`
	PriceBookRevision         int64                      `json:"price_book_revision"`
	PriceEntryID              string                     `json:"price_entry_id"`
	PriceEntryRevision        int64                      `json:"price_entry_revision"`
	PackagePriceEntryID       string                     `json:"package_price_entry_id,omitempty"`
	PackagePriceEntryRevision int64                      `json:"package_price_entry_revision,omitempty"`
	PackageListAmount         *money.Money               `json:"package_list_amount,omitempty"`
	ListUnitPrice             money.Money                `json:"list_unit_price"`
	ListAmount                money.Money                `json:"list_amount"`
	DiscountAmount            money.Money                `json:"discount_amount"`
	Consideration             money.Money                `json:"consideration"`
	Rounding                  RoundingEvidence           `json:"rounding"`
	Tax                       TaxSnapshot                `json:"tax"`
	TaxRounding               RoundingEvidence           `json:"tax_rounding"`
	LineTotal                 money.Money                `json:"line_total"`
	AppliedRules              []AppliedPriceRule         `json:"applied_rules,omitempty"`
	PromotionAllocations      []PricePromotionAllocation `json:"promotion_allocations,omitempty"`
}
