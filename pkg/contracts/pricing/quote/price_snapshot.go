// Package quote holds the immutable transaction pricing evidence Pricing
// returns and transaction owners freeze. A captured snapshot is the
// authoritative historical value: changing a price entry, tax rule,
// organisation, or merchant profile never changes an existing snapshot.
package quote

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/market/market_enums"
)

// PriceSnapshot is the immutable commercial evidence for one priced line. It
// is the authoritative historical value for the transaction that captured it.
type PriceSnapshot struct {
	QuoteID       string `json:"quote_id"`
	LineID        string `json:"line_id"`
	SKUCode       string `json:"sku_code"`
	MarketCode    string `json:"market_code"`
	PriceBookCode string `json:"price_book_code"`
	// PriceBookRevision and PriceEntryRevision pin the exact price
	// definitions that resolved. A newer revision requires a re-quote.
	PriceBookRevision                         int64                      `json:"price_book_revision"`
	PriceEntryID                              string                     `json:"price_entry_id"`
	PriceEntryRevision                        int64                      `json:"price_entry_revision"`
	PackagePriceEntryID                       string                     `json:"package_price_entry_id,omitempty"`
	PackagePriceEntryRevision                 int64                      `json:"package_price_entry_revision,omitempty"`
	PackageListAmount                         *money.Money               `json:"package_list_amount,omitempty"`
	MembershipTierPriceBookAssignmentID       string                     `json:"membership_tier_price_book_assignment_id,omitempty"`
	MembershipTierPriceBookAssignmentRevision int64                      `json:"membership_tier_price_book_assignment_revision,omitempty"`
	MembershipTierKey                         string                     `json:"membership_tier_key,omitempty"`
	Channel                                   commerce_enums.OrderType   `json:"channel"`
	Audience                                  market_enums.PriceAudience `json:"audience"`

	BaseUnits          int64                                `json:"base_units"`
	PackageComposition packaging.PackageCompositionSnapshot `json:"package_composition"`

	Currency         money.CurrencyCode     `json:"currency"`
	CurrencyExponent money.CurrencyExponent `json:"currency_exponent"`
	ListUnitPrice    money.Money            `json:"list_unit_price"`
	DiscountAmount   money.Money            `json:"discount_amount"`
	TaxableBase      money.Money            `json:"taxable_base"`
	TaxAmount        money.Money            `json:"tax_amount"`
	LineTotal        money.Money            `json:"line_total"`

	Tax          TaxSnapshot        `json:"tax"`
	AppliedRules []AppliedPriceRule `json:"applied_rules,omitempty"`
	Rounding     RoundingEvidence   `json:"rounding"`
	UnitPrice    *UnitPriceEvidence `json:"unit_price,omitempty"`
	// CustomOverride records a staff-set price separately from applied rule
	// evidence. Services own the rule that an offline manual override replaces
	// automatic promotions; the model does not enforce or infer that outcome.
	CustomOverride *CustomPriceOverrideEvidence `json:"custom_override,omitempty"`

	// SourceAllocationVersion and SourceAllocations are present together for a
	// mixed-source line, after the consumer explicitly opts into that version.
	// Version 1 uses contiguous source ranges ordered by FirstUnit. The ranges
	// cover all BaseUnits exactly once. Each range is the authoritative source
	// for its price and tax; without an adjustment, the line money fields are
	// their reconciled aggregates. Version 2 additionally carries exact unit
	// allocations and may carry a separately evidenced proportional adjustment
	// on a retained amendment, in which case source sums plus the adjustment
	// reconcile to the line money fields.
	// The legacy scalar source fields identify the first range only and must
	// not be interpreted as the source of every unit of a mixed-source line.
	// Omitting both fields retains single-source and historical semantics.
	SourceAllocationVersion       int32                          `json:"source_allocation_version,omitempty"`
	SourceAllocations             []PriceSourceAllocation        `json:"source_allocations,omitempty"`
	AmendmentAllocationAdjustment *AmendmentAllocationAdjustment `json:"amendment_allocation_adjustment,omitempty"`

	// SaleEvidence identifies the single Supply evidence basis used for a
	// Pricing quote line. Transaction owners reject missing or multiple variants
	// except for a verified Orders-owned digital gift face-value instrument,
	// which is not a Pricing quote or a stock sale and has no Supply evidence.
	SaleEvidence SaleEvidence `json:"sale_evidence"`

	// Fingerprint is the Pricing-owned immutable identity of the resolved
	// quote inputs and outputs.
	Fingerprint string    `json:"fingerprint"`
	CapturedAt  time.Time `json:"captured_at"`
}
