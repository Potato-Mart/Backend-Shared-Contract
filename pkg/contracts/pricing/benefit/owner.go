// Package benefit contains ownership models for coupons, vouchers, gift cards,
// and checkout benefits. Membership ownership remains retail-only.
package benefit

// OwnerRef is the stable business owner of a non-membership benefit.
type OwnerRef struct {
	OwnerType OwnerType `json:"owner_type"`
	OwnerID   string    `json:"owner_id"`
}
