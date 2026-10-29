package pricebook

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// SellingPromotionDiscountValue carries the authored customer-visible value.
// Percent discounts use BasisPoints; monetary discounts and fixed prices use
// Amount. The publishing service owns eligibility and validates the variant.
type SellingPromotionDiscountValue struct {
	BasisPoints *int64       `json:"basis_points,omitempty"`
	Amount      *money.Money `json:"amount,omitempty"`
}
