package wallet

import "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"

// GiftCardDenominationPolicy is the versioned, server-authored set of purchase
// denominations for one currency. Owning services remain responsible for
// validation, caching, cutover, and authorization behavior.
type GiftCardDenominationPolicy struct {
	Version             int                `json:"version"`
	Currency            money.CurrencyCode `json:"currency"`
	AllowedAmountsMinor []int64            `json:"allowed_amounts_minor"`
}
