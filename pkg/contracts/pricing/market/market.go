// Package market holds the Pricing-owned commercial market model. A Market is
// where and under which commercial configuration something is sold; it is not
// physical geography. MarketCode and CountryCode are separate concepts, and one
// country may carry several markets without any change to Product, SKU, or
// price models.
package market

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/market/market_enums"
)

// Market is a commercial selling region owned by Pricing. ID is the stable
// cross-service reference; Code is the human market key such as "AU".
//
// DefaultCurrency and CurrencyExponent describe the market's usual money
// shape, but a price book owns the currency its entries are stated in. A
// market never carries prices, tax rules, or inventory of its own.
type Market struct {
	ID               string                    `json:"id"`
	Code             string                    `json:"code"`
	Name             string                    `json:"name"`
	CountryCode      geography.CountryCode     `json:"country_code"`
	DefaultCurrency  money.CurrencyCode        `json:"default_currency"`
	CurrencyExponent money.CurrencyExponent    `json:"currency_exponent"`
	DefaultLocale    string                    `json:"default_locale"`
	DefaultTimezone  string                    `json:"default_timezone"`
	Status           market_enums.MarketStatus `json:"status"`

	// ExpiryLeadDays is the market default soon-expiry lead time: the number of
	// days before a lot's date mark at which the soon-expiry window opens. A
	// market listing may override it, and frozen sale-eligibility evidence
	// records the lead time that was actually in force.
	ExpiryLeadDays int32 `json:"expiry_lead_days"`

	Revision int64 `json:"revision"`

	audit.AuditFields
}
