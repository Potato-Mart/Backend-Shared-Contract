package pkg_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/orders/shipping"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/supply/courier"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/supply/courier/courier_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/supply/fulfilment"
)

func TestDeliveryLegacyJSONRemainsUnchanged(t *testing.T) {
	for name, tc := range map[string]struct {
		wire  string
		model any
	}{
		"slot":      {`{"id":"old-slot","start_at":"2026-10-01T21:00:00Z","end_at":"2026-10-02T07:00:00Z","label":"Daytime","availability":"available"}`, &shipping.DeliverySlot{}},
		"schedule":  {`{"availability":"available","revision":4,"timezone":"Australia/Melbourne","carrier":"Be Cool Refrigerated Couriers","date_groups":[]}`, &shipping.DeliverySchedule{}},
		"preferred": {`{"date":"2026-10-02","slot_id":"old-slot","start_at":"2026-10-01T21:00:00Z","end_at":"2026-10-02T07:00:00Z","schedule_revision":4}`, &shipping.PreferredDeliverySlot{}},
	} {
		t.Run(name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(tc.wire), tc.model); err != nil {
				t.Fatal(err)
			}
			assertDeliveryJSON(t, tc.model, tc.wire)
		})
	}
	for name, model := range map[string]any{"order": order.Order{}, "shipment": fulfilment.OutboundShipment{}} {
		t.Run(name, func(t *testing.T) {
			data, err := json.Marshal(model)
			if err != nil {
				t.Fatal(err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(data, &fields); err != nil {
				t.Fatal(err)
			}
			if _, present := fields["delivery_selection"]; present {
				t.Fatal("legacy record unexpectedly emits delivery_selection")
			}
		})
	}
}

func TestDeliverySelectionSurvivesOrdersToSupplyHandoff(t *testing.T) {
	const legacyWire = `{"delivery_company":{"code":"melbourne-fleet","name":"Local delivery","integration":"api","adapter":"detrack","revision":7},"schedule_code":"daytime","schedule_revision":12,"slot_id":"opaque-slot","date":"2026-10-05","start_at":"2026-10-04T20:00:00Z","end_at":"2026-10-05T07:00:00Z","timezone":"Australia/Melbourne","source":"configured_schedule"}`
	const wire = `{"delivery_company":{"code":"melbourne-fleet","name":"Local delivery","revision":7},"schedule_code":"daytime","schedule_revision":12,"slot_id":"opaque-slot","date":"2026-10-05","start_at":"2026-10-04T20:00:00Z","end_at":"2026-10-05T07:00:00Z","timezone":"Australia/Melbourne","source":"configured_schedule"}`
	var selection shipping.DeliverySelection
	if err := json.Unmarshal([]byte(legacyWire), &selection); err != nil {
		t.Fatal(err)
	}
	assertDeliveryJSON(t, selection, wire)
	orderWire, err := json.Marshal(order.Order{OrderNumber: "test-order", OutsourcedCarrier: "Local delivery", DeliverySelection: &selection})
	if err != nil {
		t.Fatal(err)
	}
	var received order.Order
	if err := json.Unmarshal(orderWire, &received); err != nil {
		t.Fatal(err)
	}
	shipmentWire, err := json.Marshal(fulfilment.OutboundShipment{OrderNumber: received.OrderNumber, Carrier: received.DeliverySelection.DeliveryCompany.Code, DeliverySelection: received.DeliverySelection})
	if err != nil {
		t.Fatal(err)
	}
	var persisted fulfilment.OutboundShipment
	if err := json.Unmarshal(shipmentWire, &persisted); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(persisted.DeliverySelection, &selection) {
		t.Fatalf("selection lost in handoff: %+v", persisted.DeliverySelection)
	}
	if received.OutsourcedCarrier != "Local delivery" || persisted.Carrier != "melbourne-fleet" {
		t.Fatal("legacy carrier meanings changed")
	}
}

func TestCourierExplicitCoverageWireShape(t *testing.T) {
	// These postal strings exercise JSON shape only; they are not approved
	// provider coverage or rollout configuration.
	for _, wire := range []string{
		`{"code":"CONFIGURED-ZONE","country_code":"AU","market_code":"market_au","shipping_zone":{"id":"zone_test_1"},"postal_code_mode":"include_only","include_postal_codes":["0800","3000"],"exclude_postal_codes":["3000"],"enabled":true,"provider_coverage_required":false}`,
		`{"code":"PROVIDER-CHECKED-ZONE","country_code":"AU","market_code":"market_au","shipping_zone":{"id":"zone_test_2"},"postal_code_mode":"all_except","exclude_postal_codes":["3000"],"enabled":true,"provider_coverage_required":true}`,
		`{"code":"ZONE-BOUND","country_code":"AU","market_code":"market_au_second","shipping_zone":{"id":"zone_test_1"},"state_codes":["AU-VIC"],"postal_code_mode":"include_only","enabled":true,"provider_coverage_required":false}`,
	} {
		var area courier.DeliveryServiceArea
		if err := json.Unmarshal([]byte(wire), &area); err != nil {
			t.Fatal(err)
		}
		assertDeliveryJSON(t, area, wire)
	}
	const reference = `{"code":"new-fleet","name":"Future carrier","revision":2}`
	var ref courier.DeliveryCompanyRef
	if err := json.Unmarshal([]byte(reference), &ref); err != nil {
		t.Fatal(err)
	}
	assertDeliveryJSON(t, ref, reference)
}

func TestCourierLegacyAreaDoesNotInferMarketOrRetainPriority(t *testing.T) {
	const legacy = `{"code":"legacy-zone","country_code":"AU","shipping_zone":{"id":"zone_test_1"},"state_codes":["AU-VIC"],"postal_code_mode":"include_only","include_postal_codes":["0800"],"routing_priority":10,"enabled":true,"provider_coverage_required":false}`
	var area courier.DeliveryServiceArea
	if err := json.Unmarshal([]byte(legacy), &area); err != nil {
		t.Fatal(err)
	}
	assertDeliveryJSON(t, area, `{"code":"legacy-zone","country_code":"AU","market_code":"","shipping_zone":{"id":"zone_test_1"},"state_codes":["AU-VIC"],"postal_code_mode":"include_only","include_postal_codes":["0800"],"enabled":true,"provider_coverage_required":false}`)
	if _, found := reflect.TypeOf(area).FieldByName("RoutingPriority"); found {
		t.Fatal("removed RoutingPriority must not remain as a Go compatibility alias")
	}
}

func TestCourierCompanyCodeIgnoresLegacyClassifiers(t *testing.T) {
	for _, legacyMode := range []string{"api", "manual"} {
		t.Run(legacyMode, func(t *testing.T) {
			legacy := `{"code":"legacy-company","name":"Legacy","integration":"` + legacyMode + `","adapter":"legacy-provider","revision":1}`
			var ref courier.DeliveryCompanyRef
			if err := json.Unmarshal([]byte(legacy), &ref); err != nil {
				t.Fatal(err)
			}
			assertDeliveryJSON(t, ref, `{"code":"legacy-company","name":"Legacy","revision":1}`)
			var company courier.DeliveryCompany
			if err := json.Unmarshal([]byte(legacy), &company); err != nil {
				t.Fatal(err)
			}
			data, err := json.Marshal(company)
			if err != nil {
				t.Fatal(err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(data, &fields); err != nil {
				t.Fatal(err)
			}
			if _, ok := fields["adapter"]; ok {
				t.Fatal("removed provider discriminator was serialized")
			}
			if _, ok := fields["integration"]; ok {
				t.Fatal("removed integration classification was serialized")
			}
			if string(fields["code"]) != `"legacy-company"` {
				t.Fatalf("legacy company code changed: %s", fields["code"])
			}
		})
	}
	for _, code := range []string{"market-one", "market-two"} {
		company := courier.DeliveryCompany{Code: code}
		data, err := json.Marshal(company)
		if err != nil {
			t.Fatal(err)
		}
		var received courier.DeliveryCompany
		if err := json.Unmarshal(data, &received); err != nil {
			t.Fatal(err)
		}
		if received.Code != code {
			t.Fatalf("company code changed: %+v", received)
		}
	}
}

func TestDeliveryFeesAndOfferMetadata(t *testing.T) {
	expires := time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)
	for _, source := range []courier_enums.DeliverySlotSource{courier_enums.DeliverySlotSourceConfiguredSchedule, courier_enums.DeliverySlotSourceProvider} {
		slot := shipping.DeliverySlot{ID: "offer", Source: source, ScheduleCode: "day", ExpiresAt: &expires, Availability: "unavailable", UnavailableReason: "coverage_unknown", DeliveryCompany: &courier.DeliveryCompanyRef{Code: "fleet", Revision: 4}}
		data, err := json.Marshal(slot)
		if err != nil {
			t.Fatal(err)
		}
		var decoded shipping.DeliverySlot
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(slot, decoded) {
			t.Fatalf("offer metadata changed: %+v", decoded)
		}
	}
}

func assertDeliveryJSON(t *testing.T, model any, want string) {
	t.Helper()
	got, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}
	var actual, expected any
	if err := json.Unmarshal(got, &actual); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(want), &expected); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("JSON mismatch\ngot: %s\nwant: %s", got, want)
	}
}

func containsModelType(model, target reflect.Type, seen map[reflect.Type]bool) bool {
	for model.Kind() == reflect.Pointer || model.Kind() == reflect.Slice || model.Kind() == reflect.Array {
		model = model.Elem()
	}
	if model == target {
		return true
	}
	if seen[model] {
		return false
	}
	seen[model] = true
	switch model.Kind() {
	case reflect.Struct:
		for index := 0; index < model.NumField(); index++ {
			if containsModelType(model.Field(index).Type, target, seen) {
				return true
			}
		}
	case reflect.Map:
		return containsModelType(model.Key(), target, seen) || containsModelType(model.Elem(), target, seen)
	}
	return false
}
