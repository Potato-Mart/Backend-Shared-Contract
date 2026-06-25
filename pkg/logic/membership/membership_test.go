package membershiplogic

import (
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

func TestPointReservationScenarios(t *testing.T) {
	now := time.Date(2026, 6, 23, 0, 0, 0, 0, time.UTC)
	owner := membership.MembershipOwnerRef{OwnerType: enums.MembershipOwnerTypeWholesaleOrganisation, OwnerID: "org_1"}
	balance := membership.PointBalanceBreakdown{
		MembershipAccountID: "mem_org_1",
		Owner:               owner,
		TotalPoints:         100,
		ReservedPoints:      30,
		AvailablePoints:     70,
		CalculatedAt:        now,
	}
	if !CanReserve(balance, 70) || CanReserve(balance, 71) {
		t.Fatalf("CanReserve did not respect available points: %+v", balance)
	}

	active := membership.PointReservation{
		ID:                   "res_1",
		MembershipAccountID:  "mem_org_1",
		Owner:                owner,
		OrganisationAccessID: "access_1",
		Points:               70,
		Status:               enums.PointReservationStatusReserved,
		RedemptionType:       enums.MembershipRedemptionTypeCheckoutDiscount,
		ExpiresAt:            now.Add(time.Minute),
	}
	expired := active
	expired.ExpiresAt = now.Add(-time.Minute)
	cancelled := active
	cancelled.Status = enums.PointReservationStatusCancelled

	if !CanCommit(active, now) {
		t.Fatal("active reservation should be committable")
	}
	if CanCommit(expired, now) || CanCommit(cancelled, now) {
		t.Fatal("expired or cancelled reservation should not be committable")
	}
	if HasRequiredSpendAccess(active, false) {
		t.Fatal("wholesale reservation should require spend permission")
	}
	if !HasRequiredSpendAccess(active, true) {
		t.Fatal("wholesale reservation with organisation access and permission should spend")
	}
	active.OrganisationAccessID = ""
	if HasRequiredSpendAccess(active, true) {
		t.Fatal("wholesale reservation without organisation access should not spend")
	}
}
