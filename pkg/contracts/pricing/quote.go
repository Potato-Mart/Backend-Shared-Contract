// Package pricing defines the wire DTOs for the internal pricing quote that
// Backend-Management exposes to other services (currently Commerce, at
// checkout) over the service-to-service auth flow. It pins the previously
// convention-only handoff in which the storefront prices a cart against
// Management and feeds the resulting discount/shipping/tax minor amounts and
// applied promotions into checkout. See ADR 0001.
//
// These are intentionally minimal and transport-only: the rich promotion and
// coupon domain types stay private to Management.
package pricing

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/contracts/sales"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/enums"
)

// PathQuote is the full request path of the internal pricing quote endpoint.
// Provider: Backend-Management; consumer: Backend-Commerce (checkout).
// Requires scope serviceauth.ScopePricingQuote ("pricing:quote").
const PathQuote = "/v1/internal/pricing/quote"

// QuoteLine is one cart line to price, mirroring what checkout has from the
// authoritative cart: the product snapshot's id and SKU, the captured unit
// price in minor units, and the quantity.
type QuoteLine struct {
	ProductID      string `json:"product_id" binding:"required"`
	SKU            string `json:"sku,omitempty"`
	UnitPriceMinor int64  `json:"unit_price_minor" binding:"gte=0"`
	Qty            int    `json:"qty" binding:"required,gt=0"`
	// CategoryPath is the product's canonical category key chain
	// (root→leaf, leaf last), as denormalised on the product master.
	// Optional and additive (v5.2.0): when present, category-targeted
	// promotions are evaluated against the line; when absent, only
	// product-targeted promotions can match it.
	CategoryPath []string `json:"category_path,omitempty"`
}

// QuoteRequest asks Management to price a cart: which coupon and promotions
// apply, and the resulting discount, shipping and tax amounts. CustomerID is
// optional (guest carts have none); CouponCode is optional. ShippingMethod
// and Currency mirror checkout's StartInput / cart fields.
type QuoteRequest struct {
	Lines          []QuoteLine            `json:"lines" binding:"required"`
	CustomerID     string                 `json:"customer_id,omitempty"`
	CouponCode     string                 `json:"coupon_code,omitempty"`
	ShippingMethod enums.ShippingRateName `json:"shipping_method,omitempty"`
	Currency       string                 `json:"currency" binding:"required"`
}

// QuoteResponse carries the priced result back to checkout in the same shape
// checkout already consumes: minor-unit money components plus the applied
// promotions to snapshot on the order.
type QuoteResponse struct {
	DiscountMinor     int64                    `json:"discount_minor"`
	ShippingMinor     int64                    `json:"shipping_minor"`
	TaxMinor          int64                    `json:"tax_minor"`
	AppliedPromotions []sales.AppliedPromotion `json:"applied_promotions,omitempty"`
}
