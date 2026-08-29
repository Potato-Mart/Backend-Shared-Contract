// Package pricebook holds the Pricing-owned authoritative price container:
// the price book, its per-SKU entries, and the assignments that bind a book to
// buyers in a market. Authoritative prices live here and never inside Product
// or SKU, and one country's price is never derived from another country's
// price through currency conversion.
package pricebook

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/product/product_enums"
)

// PriceBook is one pricing context inside exactly one market. It owns the
// currency and exponent its entries are stated in, the channel and audience it
// serves, whether its amounts include tax, and the price-ending policy applied
// when drafts are generated for it.
type PriceBook struct {
	ID               string                            `json:"id"`
	Code             string                            `json:"code"`
	Name             string                            `json:"name"`
	MarketCode       string                            `json:"market_code"`
	Currency         money.CurrencyCode                `json:"currency"`
	CurrencyExponent money.CurrencyExponent            `json:"currency_exponent"`
	Channel          commerce_enums.OrderType          `json:"channel"`
	Audience         product_enums.PriceAudience       `json:"audience"`
	TaxInclusion     pricebook_enums.PriceTaxInclusion `json:"tax_inclusion"`
	PriceEnding      pricebook_enums.PriceEndingPolicy `json:"price_ending"`
	Status           pricebook_enums.PriceBookStatus   `json:"status"`
	ValidFrom        time.Time                         `json:"valid_from"`
	ValidUntil       *time.Time                        `json:"valid_until,omitempty"`
	Revision         int64                             `json:"revision"`

	audit.AuditFields
}
