package pricing_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/membership/membership_enums"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/pricing"
	event_enums "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pubsub/routing"
)

func TestMembershipTierChangedEventJSONShape(t *testing.T) {
	gold, silver := "gold", "silver"
	value := event.MembershipTierChangedEvent{
		CustomerNumber:  "RC-1",
		CountryCode:     "AU",
		TierKey:         &gold,
		PreviousTierKey: &silver,
		Revision:        3,
		Reason:          membership_enums.MembershipTierChangeReasonSpend,
		ChangedAt:       time.Date(2026, 11, 7, 1, 2, 3, 0, time.UTC),
	}
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal membership.tier_changed payload: %v", err)
	}
	const want = `{"customer_number":"RC-1","country_code":"AU","tier_key":"gold","previous_tier_key":"silver","revision":3,"reason":"spend","changed_at":"2026-11-07T01:02:03Z"}`
	if got := string(payload); got != want {
		t.Fatalf("membership.tier_changed JSON = %s, want %s", got, want)
	}

	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	var decoded event.MembershipTierChangedEvent
	if err := decoder.Decode(&decoded); err != nil {
		t.Fatalf("strict decode membership.tier_changed payload: %v", err)
	}
	if !reflect.DeepEqual(decoded, value) {
		t.Fatalf("membership.tier_changed round trip = %+v, want %+v", decoded, value)
	}
	if !decoded.Reason.IsValid() {
		t.Fatalf("membership.tier_changed reason %q must be valid", decoded.Reason)
	}
}

func TestMembershipTierChangedEventCarriesExplicitNullTiers(t *testing.T) {
	previous := "gold"
	value := event.MembershipTierChangedEvent{
		CustomerNumber:  "RC-1",
		CountryCode:     "AU",
		PreviousTierKey: &previous,
		Revision:        4,
		Reason:          membership_enums.MembershipTierChangeReasonRefund,
		ChangedAt:       time.Date(2026, 11, 7, 1, 2, 3, 0, time.UTC),
	}
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal membership.tier_changed payload: %v", err)
	}
	const want = `{"customer_number":"RC-1","country_code":"AU","tier_key":null,"previous_tier_key":"gold","revision":4,"reason":"refund","changed_at":"2026-11-07T01:02:03Z"}`
	if got := string(payload); got != want {
		t.Fatalf("membership.tier_changed JSON = %s, want %s", got, want)
	}

	enrolment := []byte(`{"customer_number":"RC-2","country_code":"TW","tier_key":"basic","previous_tier_key":null,"revision":1,"reason":"enrolment","changed_at":"2026-11-07T01:02:03Z"}`)
	decoder := json.NewDecoder(bytes.NewReader(enrolment))
	decoder.DisallowUnknownFields()
	var decoded event.MembershipTierChangedEvent
	if err := decoder.Decode(&decoded); err != nil {
		t.Fatalf("strict decode membership.tier_changed payload: %v", err)
	}
	if decoded.PreviousTierKey != nil || decoded.TierKey == nil || *decoded.TierKey != "basic" || decoded.Revision != 1 || decoded.Reason != membership_enums.MembershipTierChangeReasonEnrolment {
		t.Fatalf("membership.tier_changed enrolment decode = %+v", decoded)
	}
}

func TestMembershipTierChangedEventRejectsUnknownFieldsAndHasReviewedSurface(t *testing.T) {
	unknown := []byte(`{"customer_number":"RC-1","country_code":"AU","tier_key":null,"previous_tier_key":null,"revision":2,"reason":"reconcile","changed_at":"2026-11-07T01:02:03Z","qualifying_spend":100}`)
	decoder := json.NewDecoder(bytes.NewReader(unknown))
	decoder.DisallowUnknownFields()
	var decoded event.MembershipTierChangedEvent
	if err := decoder.Decode(&decoded); err == nil {
		t.Fatal("strict decode must reject fields outside the membership.tier_changed contract")
	}

	model := reflect.TypeOf(event.MembershipTierChangedEvent{})
	fields := map[string]string{
		"CustomerNumber":  "customer_number",
		"CountryCode":     "country_code",
		"TierKey":         "tier_key",
		"PreviousTierKey": "previous_tier_key",
		"Revision":        "revision",
		"Reason":          "reason",
		"ChangedAt":       "changed_at",
	}
	if model.NumField() != len(fields) {
		t.Fatalf("membership.tier_changed has %d fields, want only the %d reviewed fields", model.NumField(), len(fields))
	}
	for name, tag := range fields {
		field, ok := model.FieldByName(name)
		if !ok || field.Tag.Get("json") != tag {
			t.Fatalf("membership.tier_changed field %s changed", name)
		}
	}
}

func TestMembershipTierChangedEventTypeUsesStorefrontEvents(t *testing.T) {
	if got := event_enums.EventTypeMembershipTierChanged; got.String() != "membership.tier_changed" || !got.IsValid() {
		t.Fatalf("membership tier changed event type = %q valid=%v", got, got.IsValid())
	}
	if event_enums.EventTopicStorefrontEvents.String() != "storefront-events" {
		t.Fatal("membership.tier_changed must reuse storefront-events")
	}
}
