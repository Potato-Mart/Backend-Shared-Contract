// Package membershiplogic holds derived predicates for membership contract
// structs. It lives outside pkg/contracts so the contract files stay pure data
// shapes.
package membershiplogic

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/contracts/membership"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/enums/membership"
)

// PermissionMembershipPointsSpend is the permission a wholesale organisation
// access grant needs before it may spend organisation-owned points.
const PermissionMembershipPointsSpend = "membership.points.spend"

// RequiresOrganisationAccess reports whether spending from this owner should be
// backed by a valid wholesale organisation access grant.
func RequiresOrganisationAccess(o membership.MembershipOwnerRef) bool {
	return o.OwnerType == membershipenum.MembershipOwnerTypeWholesaleOrganisation
}

// CanReserve reports whether the projected balance can reserve points without
// going negative.
func CanReserve(b membership.PointBalanceBreakdown, points int) bool {
	return points > 0 && b.AvailablePoints >= points
}

// CanCommit reports whether a reservation is still eligible to become a
// committed ledger spend at the supplied time.
func CanCommit(r membership.PointReservation, now time.Time) bool {
	return r.Status == membershipenum.PointReservationStatusReserved && !now.After(r.ExpiresAt)
}

// HasRequiredSpendAccess reports whether this reservation has the required
// access context for the owner. Retail owners do not need organisation access;
// wholesale organisation owners need both an access ID and spend permission.
func HasRequiredSpendAccess(r membership.PointReservation, hasSpendPermission bool) bool {
	if !RequiresOrganisationAccess(r.Owner) {
		return true
	}
	return r.OrganisationAccessID != "" && hasSpendPermission
}
