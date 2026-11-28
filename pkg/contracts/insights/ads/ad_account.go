package ads

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// AdAccount identifies one advertising account a market's ads are bought
// through, in the masked form that is safe to show in admin.
//
// MaskedID is the account number with all but a trailing fragment hidden; it
// is for recognition only and is never used as a key or sent to a platform.
// Currency and TimeZone are the account's own reporting settings, which may
// differ from the market's and are then reported as a health check. MarketCode
// is the market this account buys for; the first market is a single state.
type AdAccount struct {
	MaskedID   string             `json:"masked_id,omitempty"`
	Name       string             `json:"name,omitempty"`
	Currency   money.CurrencyCode `json:"currency,omitempty"`
	TimeZone   string             `json:"time_zone,omitempty"`
	MarketCode string             `json:"market_code,omitempty"`
}
