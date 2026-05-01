package common

// Money is a database-neutral currency amount.
//
// AmountMinor is the amount in the currency's minor unit, for example cents
// for AUD/USD. Keeping this as an integer avoids float rounding drift across
// JSON, Postgres NUMERIC, MongoDB Decimal128, and service boundaries.
type Money struct {
	AmountMinor int64  `json:"amount_minor"`
	Currency    string `json:"currency"`
}
