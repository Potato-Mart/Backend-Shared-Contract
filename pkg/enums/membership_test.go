package enums_test

import (
	"testing"

	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/enums/membership"
)

func TestMembershipEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "membershipenum.MembershipAccountStatus", valid: []stringEnum{membershipenum.MembershipAccountStatusActive, membershipenum.MembershipAccountStatusSuspended, membershipenum.MembershipAccountStatusClosed}, invalid: membershipenum.MembershipAccountStatus("__invalid__")},
		{name: "membershipenum.MembershipOwnerType", valid: []stringEnum{membershipenum.MembershipOwnerTypeRetailCustomer, membershipenum.MembershipOwnerTypeWholesaleOrganisation}, invalid: membershipenum.MembershipOwnerType("__invalid__")},
		{name: "membershipenum.MembershipPointReason", valid: []stringEnum{membershipenum.MembershipPointReasonOrder, membershipenum.MembershipPointReasonBirthday, membershipenum.MembershipPointReasonRedeem, membershipenum.MembershipPointReasonRewardRedeem, membershipenum.MembershipPointReasonAdminAdjust, membershipenum.MembershipPointReasonExpired, membershipenum.MembershipPointReasonReferral, membershipenum.MembershipPointReasonSignupBonus, membershipenum.MembershipPointReasonTierUpgrade, membershipenum.MembershipPointReasonManual}, invalid: membershipenum.MembershipPointReason("__invalid__")},
		{name: "membershipenum.MembershipPromotionTarget", valid: []stringEnum{membershipenum.MembershipPromotionTargetAll, membershipenum.MembershipPromotionTargetWholesale, membershipenum.MembershipPromotionTargetRetail, membershipenum.MembershipPromotionTargetTierSpecific}, invalid: membershipenum.MembershipPromotionTarget("__invalid__")},
		{name: "membershipenum.MembershipRedemptionType", valid: []stringEnum{membershipenum.MembershipRedemptionTypeCheckoutDiscount, membershipenum.MembershipRedemptionTypeRewardCatalog}, invalid: membershipenum.MembershipRedemptionType("__invalid__")},
		{name: "membershipenum.MembershipRewardRedemptionStatus", valid: []stringEnum{membershipenum.MembershipRewardRedemptionStatusReserved, membershipenum.MembershipRewardRedemptionStatusRedeemed, membershipenum.MembershipRewardRedemptionStatusCancelled, membershipenum.MembershipRewardRedemptionStatusExpired}, invalid: membershipenum.MembershipRewardRedemptionStatus("__invalid__")},
		{name: "membershipenum.MembershipRewardType", valid: []stringEnum{membershipenum.MembershipRewardTypeOrderDiscount, membershipenum.MembershipRewardTypeProduct, membershipenum.MembershipRewardTypeFreeShipping, membershipenum.MembershipRewardTypeVoucher}, invalid: membershipenum.MembershipRewardType("__invalid__")},
		{name: "membershipenum.MembershipTierMetric", valid: []stringEnum{membershipenum.MembershipTierMetricAnnualSpend, membershipenum.MembershipTierMetricLifetimeSpend, membershipenum.MembershipTierMetricManual}, invalid: membershipenum.MembershipTierMetric("__invalid__")},
		{name: "membershipenum.MemberSubscriptionStatus", valid: []stringEnum{membershipenum.MemberSubscriptionStatusActive, membershipenum.MemberSubscriptionStatusPaused, membershipenum.MemberSubscriptionStatusCancelled}, invalid: membershipenum.MemberSubscriptionStatus("__invalid__")},
		{name: "membershipenum.PointReservationStatus", valid: []stringEnum{membershipenum.PointReservationStatusReserved, membershipenum.PointReservationStatusCommitted, membershipenum.PointReservationStatusCancelled, membershipenum.PointReservationStatusExpired}, invalid: membershipenum.PointReservationStatus("__invalid__")},
	})
}
