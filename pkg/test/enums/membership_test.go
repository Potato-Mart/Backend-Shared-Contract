package enums_test

import (
	"testing"

	subscription_enums "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/subscription/subscription_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/membership/membership_enums"
)

func TestMembershipEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "membershipenum.MembershipAccountStatus", valid: []stringEnum{membership_enums.MembershipAccountStatusActive, membership_enums.MembershipAccountStatusSuspended, membership_enums.MembershipAccountStatusClosed}, invalid: membership_enums.MembershipAccountStatus("__invalid__")},
		{name: "membershipenum.MembershipTierMetric", valid: []stringEnum{membership_enums.MembershipTierMetricAnnualSpend, membership_enums.MembershipTierMetricLifetimeSpend, membership_enums.MembershipTierMetricManual}, invalid: membership_enums.MembershipTierMetric("__invalid__")},
		{name: "membershipenum.MemberSubscriptionStatus", valid: []stringEnum{subscription_enums.MemberSubscriptionStatusActive, subscription_enums.MemberSubscriptionStatusPaused, subscription_enums.MemberSubscriptionStatusCancelled}, invalid: subscription_enums.MemberSubscriptionStatus("__invalid__")},
	})
}
