package pkg_test

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/marketing/campaign"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/notification/preference"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/orders/shipping"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/pricing/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/pricing/tax"
	events "github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/pubsub/orders"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/supply/warehouse/warehouse_enums"
	"os"
	"reflect"
	"testing"
)

func TestConfigurableContractFixtures(t *testing.T) {
	body, err := os.ReadFile("testdata/configurable_contract.json")
	if err != nil {
		t.Fatal(err)
	}
	var f map[string]json.RawMessage
	if err := json.Unmarshal(body, &f); err != nil {
		t.Fatal(err)
	}
	models := map[string]any{"preferences": &preference.NotificationPreferences{}, "arrival_rule": &shipping.ShippingArrivalRule{}, "deferred_payment": &order.DeferredPaymentAuthorization{}, "paid": &events.OrderPaidEvent{}, "progress": &membership.CustomerTierProgress{}, "category": &tax.TaxCategory{}, "rule": &tax.TaxRule{}, "campaign": &campaign.Campaign{}}
	for name, model := range models {
		t.Run(name, func(t *testing.T) {
			if err := json.Unmarshal(f[name], model); err != nil {
				t.Fatal(err)
			}
			encoded, err := json.Marshal(model)
			if err != nil {
				t.Fatal(err)
			}
			var expected, actual map[string]json.RawMessage
			_ = json.Unmarshal(f[name], &expected)
			_ = json.Unmarshal(encoded, &actual)
			for key, want := range expected {
				var a, b any
				_ = json.Unmarshal(want, &a)
				_ = json.Unmarshal(actual[key], &b)
				if !reflect.DeepEqual(a, b) {
					t.Errorf("%s lost: got %s want %s", key, actual[key], want)
				}
			}
		})
	}
	prefs := models["preferences"].(*preference.NotificationPreferences)
	if prefs.Revision != 7 || prefs.Channels[0].Topics[0].Enabled {
		t.Fatal("revision/mandatory opt-out fixture changed")
	}
	// A false choice is represented, but mandatory delivery is still service policy.
	wire, _ := json.Marshal(prefs)
	var keys map[string]json.RawMessage
	_ = json.Unmarshal(wire, &keys)
	if _, ok := keys["topics"]; ok {
		t.Fatal("retired root topics serialized")
	}
	for _, c := range prefs.Channels {
		wire, _ := json.Marshal(c)
		_ = json.Unmarshal(wire, &keys)
		if _, ok := keys["enabled"]; ok {
			t.Fatal("channel master toggle serialized")
		}
	}
	paid := models["paid"].(*events.OrderPaidEvent)
	if paid.QuoteKey != "quote-1" || paid.QuoteRevision != 4 || paid.DeferredPayment == nil {
		t.Fatal("paid evidence lost")
	}
	for _, event := range []any{events.OrderCreatedEvent{DeferredPayment: paid.DeferredPayment}, events.OrderEditedEvent{DeferredPayment: paid.DeferredPayment}, order.Order{DeferredPayment: paid.DeferredPayment}} {
		wire, _ := json.Marshal(event)
		var fields map[string]json.RawMessage
		_ = json.Unmarshal(wire, &fields)
		var got order.DeferredPaymentAuthorization
		if err := json.Unmarshal(fields["deferred_payment"], &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(&got, paid.DeferredPayment) {
			t.Fatal("authorization snapshot changed")
		}
	}

	for _, legacy := range []any{events.OrderPaidEvent{}, events.OrderCreatedEvent{}, events.OrderEditedEvent{}, order.Order{}, membership.CustomerTierProgress{}} {
		wire, _ := json.Marshal(legacy)
		var fields map[string]json.RawMessage
		_ = json.Unmarshal(wire, &fields)
		for _, key := range []string{"quote_key", "quote_revision", "quote_fingerprint", "deferred_payment", "tier_progress_basis_points"} {
			if _, ok := fields[key]; ok {
				t.Fatalf("legacy emits %s", key)
			}
		}
	}
	if !warehouse_enums.StockLocationPurposeExpiryHold.IsValid() {
		t.Fatal("expiry hold invalid")
	}
}
