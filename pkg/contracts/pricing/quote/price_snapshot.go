// Package quote holds the immutable transaction pricing evidence Pricing
// returns and transaction owners freeze. A captured snapshot is the
// authoritative historical value: changing a price entry, tax rule,
// organisation, or merchant profile never changes an existing snapshot.
package quote

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/quote/quote_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/listing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/product/product_enums"
)

// TaxSnapshot is the frozen tax evidence for one line. The rate is kept as an
// exact numerator over denominator so an inclusive extraction and an exclusive
// addition are both reproducible without floats.
type TaxSnapshot struct {
	TaxCategoryID   string `json:"tax_category_id"`
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

// RoundingEvidence records the exact value a rounded minor amount came from
// and the rule that reduced it, so every rounded cent is reproducible.
type RoundingEvidence struct {
	Mode        quote_enums.RoundingMode          `json:"mode"`
	PriceEnding pricebook_enums.PriceEndingPolicy `json:"price_ending"`
	Exponent    int32                             `json:"exponent"`
	// ExactNumerator over ExactDenominator is the unrounded value in minor
	// units before the mode was applied.
	ExactNumerator   int64       `json:"exact_numerator"`
	ExactDenominator int64       `json:"exact_denominator"`
	RoundedAmount    money.Money `json:"rounded_amount"`
	// RemainderRank and RemainderMinorApplied record this line's position
	// and share when document minor units were allocated by largest
	// remainder. TieBreakKey is the stable line identity used to break
	// equal remainders.
	RemainderRank         int32  `json:"remainder_rank,omitempty"`
	RemainderMinorApplied int64  `json:"remainder_minor_applied,omitempty"`
	TieBreakKey           string `json:"tie_break_key,omitempty"`
}

// AppliedPriceRule is one frozen rule outcome. Kind is an open string so new
// commercial mechanics do not require a contract major bump; current values
// include group_15, group_20, expiry_20, expiry_bogo, damage_tier_30,
// damage_tier_50, damage_tier_80, and custom_pos. Exclusive rules never stack
// with each other or with ordinary promotions.
type AppliedPriceRule struct {
	Kind         string `json:"kind"`
	RuleID       string `json:"rule_id,omitempty"`
	RuleRevision int64  `json:"rule_revision,omitempty"`
	Exclusive    bool   `json:"exclusive"`
	// FactorNumerator over FactorDenominator is the exact multiplier the
	// rule applied, for example 85/100 for group_15 or 8/10 for
	// damage_tier_30.
	FactorNumerator   int64       `json:"factor_numerator,omitempty"`
	FactorDenominator int64       `json:"factor_denominator,omitempty"`
	AmountBefore      money.Money `json:"amount_before"`
	AmountAfter       money.Money `json:"amount_after"`
	// ChargeableBaseUnits records how many base units the rule charged for,
	// which differs from the reserved quantity under a buy-one-get-one
	// expiry mechanic.
	ChargeableBaseUnits int64     `json:"chargeable_base_units,omitempty"`
	Reason              string    `json:"reason,omitempty"`
	AppliedAt           time.Time `json:"applied_at"`
}

// UnitPriceEvidence is the comparison-unit price shown beside an ordinary
// grocery listing. Exempt records a listing that is not required to display a
// comparison price, with the reason a markdown created that exemption.
type UnitPriceEvidence struct {
	NetContent       measurement.NetContent `json:"net_content"`
	ComparisonAmount money.Money            `json:"comparison_amount"`
	Exempt           bool                   `json:"exempt"`
	ExemptionReason  string                 `json:"exemption_reason,omitempty"`
}

// CustomPriceOverrideEvidence freezes a cashier-entered tax-inclusive unit
// amount together with the actor, the mandatory reason, the approved price it
// replaced, and the cost comparison recorded at the time.
type CustomPriceOverrideEvidence struct {
	ActorUserID string `json:"actor_user_id"`
	Reason      string `json:"reason"`
	// SourceApprovedPrice is the approved price the override replaced.
	SourceApprovedPrice money.Money `json:"source_approved_price"`
	// OverrideGrossAmount is the tax-inclusive unit amount the cashier
	// entered. Tax is extracted from it rather than added to it.
	OverrideGrossAmount money.Money                `json:"override_gross_amount"`
	CostComparison      quote_enums.CostComparison `json:"cost_comparison"`
	ComparedCost        *money.Money               `json:"compared_cost,omitempty"`
	BelowCostWarning    bool                       `json:"below_cost_warning"`
	OverriddenAt        time.Time                  `json:"overridden_at"`
}

// PriceSnapshot is the immutable commercial evidence for one priced line. It
// is the authoritative historical value for the transaction that captured it.
type PriceSnapshot struct {
	QuoteID     string `json:"quote_id"`
	LineID      string `json:"line_id"`
	SKUID       string `json:"sku_id"`
	MarketID    string `json:"market_id"`
	PriceBookID string `json:"price_book_id"`
	// PriceBookRevision and PriceEntryRevision pin the exact price
	// definitions that resolved. A newer revision requires a re-quote.
	PriceBookRevision  int64                       `json:"price_book_revision"`
	PriceEntryID       string                      `json:"price_entry_id"`
	PriceEntryRevision int64                       `json:"price_entry_revision"`
	Channel            commerce_enums.OrderType    `json:"channel"`
	Audience           product_enums.PriceAudience `json:"audience"`

	BaseUnits          int64                                `json:"base_units"`
	PackageComposition packaging.PackageCompositionSnapshot `json:"package_composition"`

	Currency         money.CurrencyCode     `json:"currency"`
	CurrencyExponent money.CurrencyExponent `json:"currency_exponent"`
	ListUnitPrice    money.Money            `json:"list_unit_price"`
	DiscountAmount   money.Money            `json:"discount_amount"`
	TaxableBase      money.Money            `json:"taxable_base"`
	TaxAmount        money.Money            `json:"tax_amount"`
	LineTotal        money.Money            `json:"line_total"`

	Tax            TaxSnapshot                  `json:"tax"`
	AppliedRules   []AppliedPriceRule           `json:"applied_rules,omitempty"`
	Rounding       RoundingEvidence             `json:"rounding"`
	UnitPrice      *UnitPriceEvidence           `json:"unit_price,omitempty"`
	CustomOverride *CustomPriceOverrideEvidence `json:"custom_override,omitempty"`

	// Eligibility is the Supply-owned listing and inventory evidence the
	// price was resolved against.
	Eligibility listing.SaleEligibilitySnapshot `json:"eligibility"`

	// Fingerprint is the Pricing-owned immutable identity of the resolved
	// quote inputs and outputs.
	Fingerprint string    `json:"fingerprint"`
	CapturedAt  time.Time `json:"captured_at"`
}
