package money

// CurrencyCode is an ISO 4217 alphabetic currency code such as AUD, TWD, JPY,
// or USD. It is an open typed string: the platform is global and the set of
// tradeable currencies is configuration owned by Pricing, not a closed
// contract enum. Symbols such as "$" are ambiguous and are never identities.
type CurrencyCode string

// CurrencyExponent is the immutable minor-unit evidence for one currency.
//
// Exponent is the number of decimal places between the currency's minor unit
// and its major unit, so a minor amount is major * 10^Exponent. AUD and USD
// use 2, JPY uses 0, and some currencies use 3. No service may assume every
// currency has two decimals; the exponent that was in force is frozen into
// every price, tax, and settlement snapshot alongside the amount.
type CurrencyExponent struct {
	Currency CurrencyCode `json:"currency"`
	Exponent int32        `json:"exponent"`
}
