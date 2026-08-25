// Package quote holds the immutable transaction pricing evidence Pricing
// returns and transaction owners freeze. A captured snapshot is the
// authoritative historical value: changing a price entry, tax rule,
// organisation, or merchant profile never changes an existing snapshot.
package quote

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/listing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/product/product_enums"
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
