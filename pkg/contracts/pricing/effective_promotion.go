package pricing

// This file defines the effective-promotion lookup (v5.2.0, ADR 0001
// family): resolving which single targeted promotion prices one product
// right now, with the fixed precedence product special_campaign >
// category-tag special_campaign > product normal_promotion > category-tag
// normal_promotion > none.
//
// Provider: Backend-Management (owner of promotion rules). Consumers:
// Backend-Operations (product detail/listing enrichment) and any other
// service that must display an effective price. The resolution maths
// itself lives in the shared contract (promotion.ResolveEffective) so
// every service agrees on the outcome.

// PathEffectivePromotion is the full request path of the internal
// effective-promotion endpoint. Requires scope
// serviceauth.ScopePricingQuote ("pricing:quote").
const PathEffectivePromotion = "/v1/internal/pricing/effective-promotion"
