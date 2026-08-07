// Package benefit contains ownership models for coupons, vouchers, gift cards,
// and checkout benefits. Membership ownership remains retail-only.
package benefit

import "github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/pricing/benefit/benefit_enums"

// OwnerRef is the stable business owner of a non-membership benefit.
type OwnerRef struct {
	OwnerType benefit_enums.OwnerType `json:"owner_type"`
	OwnerID   string                  `json:"owner_id"`
}
