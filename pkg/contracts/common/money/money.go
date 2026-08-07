package money

// Money is a currency amount represented in JSON minor units.
//
// AmountMinor is the amount in the currency's minor unit, for example cents
// for AUD/USD. Keeping this as an integer avoids float rounding drift across
// JSON and service boundaries.
type Money struct {
	AmountMinor int64  `json:"amount_minor"`
	Currency    string `json:"currency"`
}
