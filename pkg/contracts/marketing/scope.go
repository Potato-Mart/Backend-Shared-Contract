package marketing

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/marketing/marketing_enums"
)

// ScopeTarget identifies one named, public target inside a coupon or
// promotion scope. Its code is authoritative; Name is a display snapshot.
type ScopeTarget struct {
	Code string                       `json:"code"`
	Name []localization.LocalizedName `json:"name"`
}

// DiscountValue uses a typed arm for percentage and monetary discount forms.
// For a percentage, BasisPoints is populated. For fixed amount or fixed
// package price, Amount is populated. Free shipping carries neither arm.
type DiscountValue struct {
	BasisPoints *int64       `json:"basis_points,omitempty"`
	Amount      *money.Money `json:"amount,omitempty"`
}

// ScopeDetail holds the discount and threshold values attached to one scope.
// Pricing services own cross-field validation for the selected discount type.
type ScopeDetail struct {
	DiscountType          marketing_enums.DiscountType `json:"discount_type"`
	DiscountValue         DiscountValue                `json:"discount_value"`
	MinimumOrderAmount    *money.Money                 `json:"minimum_order_amount,omitempty"`
	MinimumUnits          *int64                       `json:"minimum_units,omitempty"`
	MaximumDiscountAmount *money.Money                 `json:"maximum_discount_amount,omitempty"`
}
