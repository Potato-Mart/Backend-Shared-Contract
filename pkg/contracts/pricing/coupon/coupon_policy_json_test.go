package coupon_test

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v45/pkg/contracts/pricing/coupon"
	"github.com/Potato-Mart/Backend-Shared-Contract/v45/pkg/contracts/pricing/coupon/coupon_enums"
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
	for _, v := range []coupon_enums.Distribution{coupon_enums.DistributionSelfClaimByCode, coupon_enums.DistributionAssigned, coupon_enums.DistributionReferralAssigned, coupon_enums.DistributionAutoMemberClaim} {
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

func TestCouponDistributionWireValuesRemainDistinct(t *testing.T) {
	for _, tc := range []struct {
		value coupon_enums.Distribution
		wire  string
	}{
		{coupon_enums.DistributionSelfClaimByCode, "SELF_CLAIM_BY_CODE"},
		{coupon_enums.DistributionAssigned, "ASSIGNED"},
		{coupon_enums.DistributionReferralAssigned, "REFERRAL_ASSIGNED"},
		{coupon_enums.DistributionAutoMemberClaim, "AUTO_MEMBER_CLAIM"},
	} {
		t.Run(tc.wire, func(t *testing.T) {
			record := coupon.Coupon{Distribution: tc.value}
			payload, err := json.Marshal(record)
			if err != nil {
				t.Fatal(err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(payload, &fields); err != nil {
				t.Fatal(err)
			}
			if string(fields["distribution"]) != `"`+tc.wire+`"` || tc.value.String() != tc.wire || !tc.value.IsValid() {
				t.Fatalf("distribution %q has unexpected wire value or validity: %s", tc.value, fields["distribution"])
			}
			var decoded coupon.Coupon
			if err := json.Unmarshal(payload, &decoded); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(record, decoded) {
				t.Fatalf("distribution round trip changed coupon: %#v", decoded)
			}
		})
	}
	// Ordinary JSON decoding retains unknown strings; service-owned validation
	// decides whether a value is accepted and must not map it to cohort issuance.
	for _, value := range []string{"referral_assigned", "REFERRAL_ASSIGNED_UNKNOWN"} {
		var decoded coupon.Coupon
		if err := json.Unmarshal([]byte(`{"distribution":"`+value+`"}`), &decoded); err != nil {
			t.Fatal(err)
		}
		if decoded.Distribution.String() != value || decoded.Distribution.IsValid() {
			t.Fatalf("unknown distribution was accepted or normalized: %q", decoded.Distribution)
		}
	}
}
