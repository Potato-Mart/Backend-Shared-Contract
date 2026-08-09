package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/membership/membership_enums"
)

func TestMembershipEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "membershipenum.MembershipAccountStatus", valid: []stringEnum{membership_enums.MembershipAccountStatusActive, membership_enums.MembershipAccountStatusSuspended, membership_enums.MembershipAccountStatusClosed}, invalid: membership_enums.MembershipAccountStatus("__invalid__")},
		{name: "membershipenum.MembershipPointReason", valid: []stringEnum{membership_enums.MembershipPointReasonOrder, membership_enums.MembershipPointReasonBirthday, membership_enums.MembershipPointReasonRedeem, membership_enums.MembershipPointReasonRefund, membership_enums.MembershipPointReasonRewardRedeem, membership_enums.MembershipPointReasonAdminAdjust, membership_enums.MembershipPointReasonExpired, membership_enums.MembershipPointReasonReferral, membership_enums.MembershipPointReasonSignupBonus, membership_enums.MembershipPointReasonTierUpgrade, membership_enums.MembershipPointReasonDebtIncurred, membership_enums.MembershipPointReasonDebtRepaid, membership_enums.MembershipPointReasonManual}, invalid: membership_enums.MembershipPointReason("__invalid__")},
		{name: "membershipenum.PointAwardStatus", valid: []stringEnum{membership_enums.PointAwardStatusIneligible, membership_enums.PointAwardStatusDisabled, membership_enums.PointAwardStatusPending, membership_enums.PointAwardStatusAwarded, membership_enums.PointAwardStatusFailed}, invalid: membership_enums.PointAwardStatus("__invalid__")},
		{name: "membershipenum.MembershipRedemptionType", valid: []stringEnum{membership_enums.MembershipRedemptionTypeCheckoutDiscount, membership_enums.MembershipRedemptionTypeRewardCatalog}, invalid: membership_enums.MembershipRedemptionType("__invalid__")},
		{name: "membershipenum.MembershipRewardRedemptionStatus", valid: []stringEnum{membership_enums.MembershipRewardRedemptionStatusReserved, membership_enums.MembershipRewardRedemptionStatusRedeemed, membership_enums.MembershipRewardRedemptionStatusCancelled, membership_enums.MembershipRewardRedemptionStatusExpired}, invalid: membership_enums.MembershipRewardRedemptionStatus("__invalid__")},
		{name: "membershipenum.MembershipRewardType", valid: []stringEnum{membership_enums.MembershipRewardTypeOrderDiscount, membership_enums.MembershipRewardTypeProduct, membership_enums.MembershipRewardTypeFreeShipping, membership_enums.MembershipRewardTypeVoucher}, invalid: membership_enums.MembershipRewardType("__invalid__")},
		{name: "membershipenum.MembershipTierMetric", valid: []stringEnum{membership_enums.MembershipTierMetricAnnualSpend, membership_enums.MembershipTierMetricLifetimeSpend, membership_enums.MembershipTierMetricManual}, invalid: membership_enums.MembershipTierMetric("__invalid__")},
		{name: "membershipenum.MemberSubscriptionStatus", valid: []stringEnum{membership_enums.MemberSubscriptionStatusActive, membership_enums.MemberSubscriptionStatusPaused, membership_enums.MemberSubscriptionStatusCancelled}, invalid: membership_enums.MemberSubscriptionStatus("__invalid__")},
		{name: "membershipenum.PointReservationStatus", valid: []stringEnum{membership_enums.PointReservationStatusReserved, membership_enums.PointReservationStatusCommitted, membership_enums.PointReservationStatusCancelled, membership_enums.PointReservationStatusExpired}, invalid: membership_enums.PointReservationStatus("__invalid__")},
	})
}
