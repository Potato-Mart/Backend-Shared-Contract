package wallet

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/money"
)

// GiftCardDenominationPolicy is the versioned, server-authored set of purchase
// denominations for one market. Owning services remain responsible for
// validation, caching, cutover, and authorization behavior.
//
// The policy is stored per market, one active record per MarketCode, and
// CountryCode is the denormalized country of that market so a country-scoped
// editor can be authorized with a plain indexed match. Version is bumped on
// every accepted administrative write, and quotes pin the version they were
// computed from so a policy edit forces a re-quote instead of silently
// repricing an in-flight purchase.
type GiftCardDenominationPolicy struct {
	MarketCode          string                `json:"market_code"`
	CountryCode         geography.CountryCode `json:"country_code,omitempty"`
	Version             int                   `json:"version"`
	Currency            money.CurrencyCode    `json:"currency"`
	AllowedAmountsMinor []int64               `json:"allowed_amounts_minor"`

	// BonusAmountsMinor is an ordered slice of denomination/bonus pairs rather
	// than a map, because JSON stringifies integer map keys and reorders them;
	// a slice keeps the wire form stable for consumers that compare two policy
	// revisions byte for byte. A denomination absent from this slice carries no
	// bonus.
	BonusAmountsMinor []GiftCardDenominationBonus `json:"bonus_amounts_minor,omitempty"`

	audit.AuditFields
}

// GiftCardDenominationBonus records the bonus balance a buyer receives for one
// allowed denomination. The buyer is charged AmountMinor; the issued card
// carries AmountMinor + BonusMinor. Both are minor units of the policy's
// currency.
type GiftCardDenominationBonus struct {
	AmountMinor int64 `json:"amount_minor"`
	BonusMinor  int64 `json:"bonus_minor"`
}
