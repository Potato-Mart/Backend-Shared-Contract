package money

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
