package pricing

// This file defines the effective-promotion lookup (v5.2.0, ADR 0001
// family): resolving which single targeted promotion prices one product
// right now, with the fixed precedence product special_campaign >
// category special_campaign > product normal_promotion > category
// normal_promotion > none.
//
// Provider: Backend-Management (owner of promotion rules). Consumers:
// Backend-Operations (product detail/listing enrichment) and any other
// service that must display an effective price. The resolution maths
// itself lives in the shared contract (promotion.ResolveEffective) so
// every service agrees on the outcome.

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/contracts/promotion"
)

// PathEffectivePromotion is the full request path of the internal
// effective-promotion endpoint. Requires scope
// serviceauth.ScopePricingQuote ("pricing:quote").
const PathEffectivePromotion = "/v1/internal/pricing/effective-promotion"

// EffectivePromotionItem identifies one product to resolve.
type EffectivePromotionItem struct {
	ProductID string `json:"product_id" binding:"required"`
	// CategoryPath is the canonical category key chain (root→leaf,
	// leaf last) from the product master. Optional: without it only
	// product-targeted promotions can match.
	CategoryPath   []string `json:"category_path,omitempty"`
	UnitPriceMinor int64    `json:"unit_price_minor" binding:"gte=0"`
}

// EffectivePromotionRequest resolves the effective promotion for one or
// more products in a single round trip (listing pages batch their page
// of products into one call).
type EffectivePromotionRequest struct {
	Items    []EffectivePromotionItem `json:"items" binding:"required"`
	Currency string                   `json:"currency" binding:"required"`
}

// EffectivePromotionResponse maps product id → resolved promotion.
// Products with no applicable promotion are absent from the map.
type EffectivePromotionResponse struct {
	Effective map[string]*promotion.EffectivePromotion `json:"effective,omitempty"`
}
