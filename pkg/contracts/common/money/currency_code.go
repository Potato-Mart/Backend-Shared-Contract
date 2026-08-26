package money

// CurrencyCode is an ISO 4217 alphabetic currency code such as AUD, TWD, JPY,
// or USD. It is an open typed string: the platform is global and the set of
// tradeable currencies is configuration owned by Pricing, not a closed
// contract enum. Symbols such as "$" are ambiguous and are never identities.
type CurrencyCode string
