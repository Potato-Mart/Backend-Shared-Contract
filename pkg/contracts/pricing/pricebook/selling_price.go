package pricebook

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/product/product_enums"
)

// SellingPrice is the customer-safe market-scoped price snapshot that may be
// published with a SellingProduct. It deliberately omits price-book identity,
// revision, approval, derivation, source cost, and promotion calculations.
type SellingPrice struct {
	UnitPrice        money.Money                       `json:"unit_price"`
	CurrencyExponent money.CurrencyExponent            `json:"currency_exponent"`
	MarketCode       string                            `json:"market_code"`
	Channel          commerce_enums.OrderType          `json:"channel"`
	Audience         product_enums.PriceAudience       `json:"audience"`
	PriceVisibility  pricebook_enums.PriceVisibility   `json:"price_visibility"`
	TaxInclusion     pricebook_enums.PriceTaxInclusion `json:"tax_inclusion"`
	ValidFrom        time.Time                         `json:"valid_from"`
	ValidUntil       *time.Time                        `json:"valid_until,omitempty"`
	AsOf             time.Time                         `json:"as_of"`
}
