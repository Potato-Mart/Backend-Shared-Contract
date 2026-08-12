package money

// Money is a currency amount represented in JSON minor units.
//
// AmountMinor is the amount in the currency's minor unit, for example cents
// for AUD/USD. Keeping this as an integer avoids float rounding drift across
// JSON and service boundaries. Currency is the typed ISO 4217 code; the number
// of minor units per major unit is carried separately by CurrencyExponent
// because not every currency has two decimals.
type Money struct {
	AmountMinor int64        `json:"amount_minor"`
	Currency    CurrencyCode `json:"currency"`
}
