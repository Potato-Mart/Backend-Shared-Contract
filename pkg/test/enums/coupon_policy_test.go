package enums_test

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/pricing/coupon/coupon_enums"
	"testing"
)

func TestCouponEnums(t *testing.T) {
	for _, tc := range []struct {
		name    string
		valid   []stringEnum
		invalid stringEnum
	}{
		{"coupon.Acquisition", []stringEnum{coupon_enums.AcquisitionSelfClaim, coupon_enums.AcquisitionAutomatic, coupon_enums.AcquisitionManualAssignment, coupon_enums.AcquisitionReferral}, coupon_enums.Acquisition("SELF_CLAIM_BY_CODE")},
		{"coupon.AutomaticPolicy", []stringEnum{coupon_enums.AutomaticPolicyOnceWhenLive, coupon_enums.AutomaticPolicyWhileActive}, coupon_enums.AutomaticPolicy("automatic")},
		{"coupon.InheritanceSource", []stringEnum{coupon_enums.InheritanceIndependent, coupon_enums.InheritanceCampaign}, coupon_enums.InheritanceSource("legacy")},
		{"coupon.Visibility", []stringEnum{coupon_enums.VisibilityPublic, coupon_enums.VisibilityCampaignOnly, coupon_enums.VisibilityUnlisted}, coupon_enums.Visibility("PUBLIC")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, value := range tc.valid {
				if !value.IsValid() || value.String() == "" {
					t.Fatalf("invalid canonical value %v", value)
				}
			}
			if tc.invalid.IsValid() {
				t.Fatalf("retired value accepted %v", tc.invalid)
			}
		})
	}
}
