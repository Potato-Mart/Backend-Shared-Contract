package membership_test

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/pricing/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/pricing/membership/membership_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/supply/catalogue/classification"
	"reflect"
	"testing"
)

func TestTierLocalizedMetadataAndMediaReference(t *testing.T) {
	tier := membership.MembershipTier{TierKey: "silver", Label: []localization.LocalizedText{{Language: "en", Text: "Silver"}, {Language: "zh-TW", Text: "銀卡"}, {Language: "zh-CN", Text: "银卡"}, {Language: "ja", Text: "シルバー"}}, Metadata: []membership.MembershipTierMetadataEntry{{Key: "description", Value: ""}}, TierCard: &classification.ObjectMediaRef{Code: "tier-silver"}}
	payload, err := json.Marshal(tier)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(payload, &fields); err != nil {
		t.Fatal(err)
	}
	if string(fields["metadata"]) != `[{"key":"description","value":""}]` || string(fields["tier_card"]) != `{"code":"tier-silver"}` {
		t.Fatalf("unexpected fields: %s", payload)
	}
	for _, retired := range []string{"is_active", "discount_percent", "free_shipping_threshold", "birthday_bonus_points", "rank"} {
		if _, ok := fields[retired]; ok {
			t.Fatalf("retired field %s", retired)
		}
	}
	var decoded membership.MembershipTier
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(tier, decoded) {
		t.Fatalf("round trip changed tier: %#v", decoded)
	}
	progress, err := json.Marshal(membership.TierProgressTier{TierKey: "silver", TierCard: tier.TierCard})
	if err != nil {
		t.Fatal(err)
	}
	if string(progress) != `{"tier_key":"silver","tier_card":{"code":"tier-silver"}}` {
		t.Fatalf("unexpected progress: %s", progress)
	}
	empty, err := json.Marshal(membership.TierProgressTier{TierKey: "silver"})
	if err != nil {
		t.Fatal(err)
	}
	if string(empty) != `{"tier_key":"silver"}` {
		t.Fatalf("optional card emitted: %s", empty)
	}
	for _, retired := range []string{"discount_percent", "free_shipping_threshold", "birthday_bonus_points"} {
		if membership_enums.TierBenefitKind(retired).IsValid() {
			t.Fatalf("retired benefit %s accepted", retired)
		}
	}
}
