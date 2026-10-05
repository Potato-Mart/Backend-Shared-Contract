package membership_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pricing/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pricing/membership/membership_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pricing/wallet/points"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/supply/catalogue/classification"
)

func TestMembershipSelectedMarketFixturesRoundTrip(t *testing.T) {
	now := time.Date(2026, 10, 6, 0, 0, 0, 0, time.UTC)
	auditFields := audit.AuditFields{CreatedAt: now, CreatedBy: "STAFF-1", UpdatedAt: now, UpdatedBy: "STAFF-1"}
	threshold := money.Money{AmountMinor: 250000, Currency: "AUD"}
	tier := membership.MembershipTier{
		TierKey:             "au-gold",
		Label:               []localization.LocalizedText{{Language: "en", Text: "Gold"}, {Language: "zh-TW", Text: "黃金"}, {Language: "zh-CN", Text: "黄金"}},
		QualificationMetric: membership_enums.MembershipTierMetricAnnualSpend,
		MinQualifyingSpend:  threshold,
		PointMultiplier:     1.25,
		Metadata:            []membership.MembershipTierMetadataEntry{{Key: "description", Value: ""}},
		TierCard:            &classification.ObjectMediaRef{Code: "tier-au-gold"},
		Benefits: []membership.TierBenefit{
			{BenefitKey: "qualifying-spend", Kind: membership_enums.TierBenefitKindQualifyingSpend, Value: membership.TierBenefitValue{Money: &threshold}},
			{BenefitKey: "points-multiplier", Kind: membership_enums.TierBenefitKindPointsMultiplier, Value: membership.TierBenefitValue{Decimal: "1.25"}},
		},
		IsSystem:    true,
		MarketCodes: []string{"AU-VIC", "AU-NSW"},
		CountryCode: "AU",
		AuditFields: auditFields,
	}
	account := membership.MembershipAccount{
		ID:          "RC-1",
		TierKey:     tier.TierKey,
		Status:      membership_enums.MembershipAccountStatusActive,
		Wallet:      points.PointsSummary{PointBalances: points.PointBalances{TotalPoints: 120, AvailablePoints: 120}, CalculatedAt: now},
		EnrolledAt:  now,
		Metadata:    metadata.Metadata{"source": "membership-fixture"},
		MarketCodes: []string{"AU-VIC", "AU-NSW"},
		CountryCode: "AU",
		AuditFields: auditFields,
		DataProtectionFields: security.DataProtectionFields{
			ContainsPII: true, DataOwnerID: "RC-1", RetentionPolicyKey: "membership",
		},
	}

	for _, tc := range []struct {
		name    string
		value   any
		decoded any
		fields  map[string]string
	}{
		{
			name: "tier", value: tier, decoded: &membership.MembershipTier{},
			fields: map[string]string{
				"min_qualifying_spend": `{"amount_minor":250000,"currency":"AUD"}`,
				"metadata":             `[{"key":"description","value":""}]`,
				"tier_card":            `{"code":"tier-au-gold"}`,
				"is_system":            `true`,
			},
		},
		{
			name: "account", value: account, decoded: &membership.MembershipAccount{},
			fields: map[string]string{
				"id":                   `"RC-1"`,
				"tier_key":             `"au-gold"`,
				"status":               `"active"`,
				"contains_pii":         `true`,
				"data_owner_id":        `"RC-1"`,
				"retention_policy_key": `"membership"`,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			payload, fields := membershipJSONFields(t, tc.value)
			for key, want := range map[string]string{
				"market_codes": `["AU-VIC","AU-NSW"]`,
				"country_code": `"AU"`,
				"created_at":   `"2026-10-06T00:00:00Z"`,
				"updated_at":   `"2026-10-06T00:00:00Z"`,
			} {
				if string(fields[key]) != want {
					t.Fatalf("%s = %s, want %s", key, fields[key], want)
				}
			}
			for key, want := range tc.fields {
				if string(fields[key]) != want {
					t.Fatalf("%s = %s, want %s", key, fields[key], want)
				}
			}
			if err := json.Unmarshal(payload, tc.decoded); err != nil {
				t.Fatalf("unmarshal fixture: %v", err)
			}
			if decoded := reflect.ValueOf(tc.decoded).Elem().Interface(); !reflect.DeepEqual(tc.value, decoded) {
				t.Fatalf("round trip changed fixture: %#v", decoded)
			}
		})
	}
}

func TestMembershipUnconfiguredMarketsRemainExplicit(t *testing.T) {
	for _, tc := range []struct {
		name    string
		markets []string
		want    string
	}{
		{name: "nil", markets: nil, want: `null`},
		{name: "empty", markets: []string{}, want: `[]`},
	} {
		for _, model := range []struct {
			name    string
			value   any
			decoded any
		}{
			{name: "tier", value: membership.MembershipTier{MarketCodes: tc.markets}, decoded: &membership.MembershipTier{}},
			{name: "account", value: membership.MembershipAccount{MarketCodes: tc.markets}, decoded: &membership.MembershipAccount{}},
		} {
			t.Run(model.name+"/"+tc.name, func(t *testing.T) {
				payload, fields := membershipJSONFields(t, model.value)
				if string(fields["market_codes"]) != tc.want {
					t.Fatalf("market_codes = %s, want %s", fields["market_codes"], tc.want)
				}
				for _, optional := range []string{"benefits", "tier_card", "metadata"} {
					if _, exists := fields[optional]; exists {
						t.Fatalf("empty optional field %s emitted", optional)
					}
				}
				if err := json.Unmarshal(payload, model.decoded); err != nil {
					t.Fatalf("unmarshal unconfigured markets: %v", err)
				}
				if decoded := reflect.ValueOf(model.decoded).Elem().Interface(); !reflect.DeepEqual(model.value, decoded) {
					t.Fatalf("round trip changed unconfigured markets: %#v", decoded)
				}
			})
		}
	}
}

func TestMembershipLegacyMarketDoesNotPopulateSelectedMarkets(t *testing.T) {
	for _, model := range []struct {
		name  string
		value any
	}{
		{name: "tier", value: &membership.MembershipTier{}},
		{name: "account", value: &membership.MembershipAccount{}},
	} {
		t.Run(model.name, func(t *testing.T) {
			modelType := reflect.TypeOf(model.value).Elem()
			if _, exists := modelType.FieldByName("MarketCode"); exists {
				t.Fatal("legacy MarketCode Go field remains")
			}
			for _, input := range []string{
				`{"country_code":"AU"}`,
				`{"country_code":"AU","market_code":"AU-VIC"}`,
			} {
				if err := json.Unmarshal([]byte(input), model.value); err != nil {
					t.Fatalf("unmarshal input: %v", err)
				}
				_, fields := membershipJSONFields(t, model.value)
				if string(fields["market_codes"]) != `null` || string(fields["country_code"]) != `"AU"` {
					t.Fatalf("missing or legacy market was converted: %v", fields)
				}
			}
		})
	}
}

func membershipJSONFields(t *testing.T, value any) ([]byte, map[string]json.RawMessage) {
	t.Helper()
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal membership: %v", err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(payload, &fields); err != nil {
		t.Fatalf("decode membership fields: %v", err)
	}
	if _, exists := fields["market_code"]; exists {
		t.Fatalf("legacy market_code emitted: %s", payload)
	}
	return payload, fields
}
