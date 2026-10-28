package coupon

import "github.com/Potato-Mart/Backend-Shared-Contract/v47/pkg/contracts/common/money"

type Benefit struct {
	BasisPoints     *int64       `json:"basis_points,omitempty"`
	Amount          *money.Money `json:"amount,omitempty"`
	FreeShipping    bool         `json:"free_shipping,omitempty"`
	MaximumDiscount *money.Money `json:"maximum_discount,omitempty"`
}
