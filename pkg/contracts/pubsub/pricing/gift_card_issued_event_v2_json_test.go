package pricing_test

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/pricing/wallet/wallet_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/pubsub/pricing"
	"reflect"
	"testing"
	"time"
)

func TestGiftCardIssuedV2SourcesRoundTripWithoutPurchaseAssertions(t *testing.T) {
	for _, source := range []wallet_enums.GiftCardIssuanceSource{wallet_enums.GiftCardIssuanceSourcePurchase, wallet_enums.GiftCardIssuanceSourceMembershipReward, wallet_enums.GiftCardIssuanceSourceRefundReplacement} {
		t.Run(source.String(), func(t *testing.T) {
			fact := pricing.GiftCardIssuedEventV2{IssuanceID: "issuance-1", Source: source, IssuedValue: money.Money{AmountMinor: 5500, Currency: "AUD"}, MarketCode: "market-au", CountryCode: "AU", IssuedAt: time.Date(2026, 9, 29, 0, 0, 0, 0, time.UTC)}
			data, err := json.Marshal(fact)
			if err != nil {
				t.Fatal(err)
			}
			want := `{"issuance_id":"issuance-1","source":"` + source.String() + `","issued_value":{"amount_minor":5500,"currency":"AUD"},"market_code":"market-au","country_code":"AU","issued_at":"2026-09-29T00:00:00Z"}`
			if string(data) != want {
				t.Fatalf("unsafe or changed issuance fact: %s", data)
			}
			var decoded pricing.GiftCardIssuedEventV2
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(fact, decoded) {
				t.Fatal("immutable source/value changed on replay")
			}
		})
	}
}

func TestGiftCardIssuedV2SafeFieldSurface(t *testing.T) {
	model := reflect.TypeOf(pricing.GiftCardIssuedEventV2{})
	fields := map[string]string{"IssuanceID": "issuance_id", "Source": "source", "IssuedValue": "issued_value", "MarketCode": "market_code,omitempty", "CountryCode": "country_code,omitempty", "IssuedAt": "issued_at"}
	if model.NumField() != len(fields) {
		t.Fatal("issuance event must carry only reviewed safe fields")
	}
	for name, tag := range fields {
		field, ok := model.FieldByName(name)
		if !ok || field.Tag.Get("json") != tag {
			t.Fatalf("issuance field %s changed", name)
		}
	}
	data, err := json.Marshal(pricing.GiftCardIssuedEventV2{})
	if err != nil {
		t.Fatal(err)
	}
	var wire map[string]json.RawMessage
	if err := json.Unmarshal(data, &wire); err != nil {
		t.Fatal(err)
	}
	for _, required := range []string{"issuance_id", "source", "issued_value", "issued_at"} {
		if _, found := wire[required]; !found {
			t.Errorf("required %s omitted", required)
		}
	}
	for _, optional := range []string{"market_code", "country_code"} {
		if _, found := wire[optional]; found {
			t.Errorf("unknown %s must remain absent", optional)
		}
	}
}

func TestGiftCardIssuanceV1PurchaseEvidenceRemainsUnchanged(t *testing.T) {
	issuedAt := time.Date(2026, 9, 29, 0, 0, 0, 0, time.UTC)
	legacy := pricing.GiftCardIssuedEvent{IssuanceID: "issuance-1", DenominationPolicyVersion: 2, Amount: money.Money{AmountMinor: 5000, Currency: "AUD"}, BonusAmountMinor: 500, IssuedAt: issuedAt}
	data, err := json.Marshal(legacy)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"issuance_id":"issuance-1","denomination_policy_version":2,"amount":{"amount_minor":5000,"currency":"AUD"},"bonus_amount_minor":500,"issued_at":"2026-09-29T00:00:00Z"}`
	if string(data) != want {
		t.Fatalf("additive release changed v1 purchase evidence: %s", data)
	}
}
