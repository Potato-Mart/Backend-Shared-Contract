package pricing

import "time"

// PriceChangedEvent is the customer-safe storefront-events invalidation fact
// emitted after an authoritative price revision. Consumers refetch the
// current price when RefetchRequired is true and use Revision to discard stale
// deliveries. It intentionally contains no monetary value, currency,
// price-book, pricing-rule, actor, provider, device, or customer data.
type PriceChangedEvent struct {
	MarketCode      string    `json:"market_code"`
	SKUCode         string    `json:"sku_code"`
	Revision        int64     `json:"revision"`
	RefetchRequired bool      `json:"refetch_required"`
	ChangedAt       time.Time `json:"changed_at"`
}
