// Package benefit contains ownership models for coupons, vouchers, gift cards,
// and checkout benefits. Membership ownership remains retail-only.
package benefit

import benefitenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/benefit"

// OwnerRef is the stable business owner of a non-membership benefit.
type OwnerRef struct {
	OwnerType benefitenum.OwnerType `json:"owner_type"`
	OwnerID   string                `json:"owner_id"`
}
