// Package pricing defines stable internal pricing endpoint paths.
package pricing

// PathQuote is the full request path of the internal pricing quote endpoint.
// Provider: Backend-Management; consumer: Backend-Commerce (checkout).
// Requires scope serviceauth.ScopePricingQuote ("pricing:quote").
const PathQuote = "/v1/internal/pricing/quote"
