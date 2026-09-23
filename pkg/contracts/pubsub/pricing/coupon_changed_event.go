package pricing

import "time"

// CouponChangedEvent is a customer-safe storefront-events invalidation for a
// Pricing-owned coupon lifecycle change. Consumers refetch by opaque CouponID;
// this event carries no redeemable code, terms, eligibility rules, or customer
// data.
type CouponChangedEvent struct {
	CouponID        string    `json:"coupon_id"`
	Revision        int64     `json:"revision"`
	RefetchRequired bool      `json:"refetch_required"`
	ChangedAt       time.Time `json:"changed_at"`
}
