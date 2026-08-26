package event_test

import (
	"encoding/json"
	"testing"
	"time"

	event "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pubsub/event"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pubsub/event/event_enums"
)

func TestPriceChangedEventJSONAndPrivacy(t *testing.T) {
	changedAt := time.Date(2026, 8, 25, 1, 2, 3, 0, time.UTC)
	value := event.PriceChangedEvent{
		MarketCode:      "market_au",
		SKUCode:         "sku_a00001",
		Revision:        42,
		RefetchRequired: true,
		ChangedAt:       changedAt,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal price changed event: %v", err)
	}
	const want = `{"market_code":"market_au","sku_code":"sku_a00001","revision":42,"refetch_required":true,"changed_at":"2026-08-25T01:02:03Z"}`
	if got := string(payload); got != want {
		t.Fatalf("price changed event JSON = %s, want %s", got, want)
	}

	var shape map[string]json.RawMessage
	if err := json.Unmarshal(payload, &shape); err != nil {
		t.Fatalf("unmarshal price changed event shape: %v", err)
	}
	wantKeys := map[string]struct{}{
		"market_code": {}, "sku_code": {}, "revision": {}, "refetch_required": {}, "changed_at": {},
	}
	if len(shape) != len(wantKeys) {
		t.Fatalf("price changed event exposed unexpected fields: %s", payload)
	}
	for key := range wantKeys {
		if _, ok := shape[key]; !ok {
			t.Fatalf("price changed event missing %q: %s", key, payload)
		}
	}
	for _, forbidden := range []string{
		"amount", "amount_minor", "money", "currency", "price", "price_book", "pricebook",
		"rule", "rules", "pricing_rule", "actor", "actor_id", "changed_by", "provider",
		"provider_id", "device", "device_id", "customer", "customer_id", "customer_number",
	} {
		if _, ok := shape[forbidden]; ok {
			t.Fatalf("price changed event leaked %q: %s", forbidden, payload)
		}
	}
}

func TestPriceChangedEventTypeUsesStorefrontInvalidationChannel(t *testing.T) {
	if event_enums.EventTopicStorefrontEvents.String() != "storefront-events" {
		t.Fatal("price changed event must reuse storefront-events")
	}
	if event_enums.EventTypePriceChanged.String() != "price.changed" || !event_enums.EventTypePriceChanged.IsValid() {
		t.Fatal("price changed event type must be a valid price.changed value")
	}
	if event_enums.EventType("price.changed.invalid").IsValid() {
		t.Fatal("unknown price event type must be invalid")
	}
}
