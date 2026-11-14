package order

import "time"

// TenderPlanSnapshot is the plan for paying one POS sale with several
// tenders (split tender), stored on the Order. Orders owns the plan; Payments
// records one Payment per collected part and names the part on it.
//
// PlanID is stable for the life of the plan. Revision increases each time the
// plan's parts change before it is frozen. FrozenAt is when the plan was
// frozen (no part may be added, removed or re-priced after it); it is absent
// while the plan can still be revised.
//
// Invariants, which Orders and Payments validate:
//   - every money value in every part shares one currency, the order's
//     currency;
//   - the parts' Sequence values are exactly 1..n, each used once, and give
//     the order in which the parts are collected;
//   - the sum of the parts' Amount equals Order.Total less the order's
//     gift-card redemptions (GiftCardRedemptions[].AppliedAmount) and points
//     redemption (PointRedemption.DiscountAmount). Gift cards and points are
//     never parts of the plan.
//
// Order.PaymentMethod stays the primary external tender of the sale; the
// plan never introduces a "split" or "mixed" payment method.
type TenderPlanSnapshot struct {
	PlanID   string           `json:"plan_id"`
	Revision int64            `json:"revision"`
	Parts    []TenderPlanPart `json:"parts"`
	FrozenAt *time.Time       `json:"frozen_at,omitempty"`
}
