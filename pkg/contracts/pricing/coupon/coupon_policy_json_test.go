package coupon_test

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/pricing/coupon"
	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/pricing/coupon/coupon_enums"
	"reflect"
	"testing"
)

func TestReceivingAndRedemptionPoliciesRemainIndependent(t *testing.T) {
	record := coupon.Coupon{
		Distribution:                 coupon_enums.DistributionAutoMemberClaim,
		Visibility:                   coupon_enums.VisibilityUnlisted,
		ReceivingTier:                coupon.ReceivingTierPolicy{Specification: coupon_enums.TierSpecificationSelectedTier, TierKey: "silver", Match: coupon_enums.TierMatchAtOrAbove},
		TierRestriction:              coupon.TierRestriction{Active: true, TierKey: "gold", Match: coupon_enums.TierMatchExact},
		ProfileCompletionRestriction: coupon.ProfileCompletionRestriction{Active: true},
	}
	payload, err := json.Marshal(record)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(payload, &fields); err != nil {
		t.Fatal(err)
	}
	for key, want := range map[string]string{
		"distribution": `"AUTO_MEMBER_CLAIM"`, "visibility": `"UNLISTED"`,
		"receiving_tier":                 `{"specification":"SELECTED_TIER","tier_key":"silver","match":"AT_OR_ABOVE"}`,
		"tier_restriction":               `{"active":true,"tier_key":"gold","match":"EXACT"}`,
		"profile_completion_restriction": `{"active":true}`,
	} {
		if string(fields[key]) != want {
			t.Errorf("%s = %s, want %s", key, fields[key], want)
		}
	}
	var decoded coupon.Coupon
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded, record) {
		t.Fatalf("round trip changed policy: %#v", decoded)
	}
	all, err := json.Marshal(coupon.ReceivingTierPolicy{Specification: coupon_enums.TierSpecificationAllTiers})
	if err != nil {
		t.Fatal(err)
	}
	if string(all) != `{"specification":"ALL_TIERS"}` {
		t.Fatalf("unexpected all-tier policy: %s", all)
	}
	disabled, err := json.Marshal(coupon.TierRestriction{})
	if err != nil {
		t.Fatal(err)
	}
	if string(disabled) != `{"active":false}` {
		t.Fatalf("unexpected disabled restriction: %s", disabled)
	}
}

func TestCouponPolicyEnumsRejectUnknownValues(t *testing.T) {
	for _, v := range []coupon_enums.Distribution{coupon_enums.DistributionSelfClaimByCode, coupon_enums.DistributionAssigned, coupon_enums.DistributionAutoMemberClaim} {
		if !v.IsValid() || v.String() != string(v) {
			t.Fatal(v)
		}
	}
	for _, v := range []coupon_enums.Visibility{coupon_enums.VisibilityPublic, coupon_enums.VisibilityUnlisted} {
		if !v.IsValid() {
			t.Fatal(v)
		}
	}
	for _, v := range []coupon_enums.TierMatch{coupon_enums.TierMatchExact, coupon_enums.TierMatchAtOrAbove} {
		if !v.IsValid() {
			t.Fatal(v)
		}
	}
	for _, v := range []coupon_enums.TierSpecification{coupon_enums.TierSpecificationAllTiers, coupon_enums.TierSpecificationSelectedTier} {
		if !v.IsValid() {
			t.Fatal(v)
		}
	}
	for _, bad := range []string{"", "unknown", "all_tiers"} {
		if coupon_enums.Distribution(bad).IsValid() || coupon_enums.Visibility(bad).IsValid() || coupon_enums.TierMatch(bad).IsValid() || coupon_enums.TierSpecification(bad).IsValid() {
			t.Fatal(bad)
		}
	}
}
