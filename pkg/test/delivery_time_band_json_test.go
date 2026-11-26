package pkg_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/shipping"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/shipping/shipping_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/fulfilment"
)

// V53.53 adds the delivery time band and the customer's acknowledgement of a
// late-night or early-morning window. Every new field is optional, so slots
// and selections written before the release read back unchanged.
func TestV5353DeliveryTimeBandFields(t *testing.T) {
	cases := []struct {
		name  string
		model any
		field string
		json  string
		typ   reflect.Type
	}{
		{"slot time band", shipping.DeliverySlot{}, "TimeBand", "time_band,omitempty", reflect.TypeOf(shipping_enums.DeliveryTimeBand(""))},
		{"slot acknowledgement", shipping.DeliverySlot{}, "RequiresTimeAcknowledgement", "requires_time_acknowledgement,omitempty", reflect.TypeOf(false)},
		{"selection time band", shipping.DeliverySelection{}, "TimeBand", "time_band,omitempty", reflect.TypeOf(shipping_enums.DeliveryTimeBand(""))},
		{"selection acknowledged at", shipping.DeliverySelection{}, "TimeAcknowledgedAt", "time_acknowledged_at,omitempty", reflect.TypeOf((*time.Time)(nil))},
	}
	for _, tc := range cases {
		field, ok := reflect.TypeOf(tc.model).FieldByName(tc.field)
		if !ok || field.Type != tc.typ || field.Tag.Get("json") != tc.json {
			t.Errorf("%s: field %s = %+v, want %s `json:%q`", tc.name, tc.field, field, tc.typ, tc.json)
		}
	}
}

func TestV5353DeliverySlotTimeBandRoundTrip(t *testing.T) {
	const wire = `{"id":"slot-night","start_at":"2026-11-27T10:00:00Z","end_at":"2026-11-27T12:00:00Z","availability":"available","time_band":"late_night","requires_time_acknowledgement":true}`
	var slot shipping.DeliverySlot
	if err := json.Unmarshal([]byte(wire), &slot); err != nil {
		t.Fatal(err)
	}
	if slot.TimeBand != shipping_enums.DeliveryTimeBandLateNight || !slot.RequiresTimeAcknowledgement {
		t.Fatalf("slot time band not decoded: %+v", slot)
	}
	assertDeliveryJSON(t, slot, wire)

	const standard = `{"id":"slot-day","start_at":"2026-11-27T00:00:00Z","end_at":"2026-11-27T03:00:00Z","availability":"available","time_band":"standard"}`
	var day shipping.DeliverySlot
	if err := json.Unmarshal([]byte(standard), &day); err != nil {
		t.Fatal(err)
	}
	assertDeliveryJSON(t, day, standard)
}

func TestV5353DeliverySelectionAcknowledgementSurvivesHandoff(t *testing.T) {
	const wire = `{"delivery_company":{"code":"melbourne-fleet","name":"Local delivery","revision":7},"schedule_code":"early","schedule_revision":12,"slot_id":"opaque-slot","date":"2026-11-28","start_at":"2026-11-27T19:00:00Z","end_at":"2026-11-27T21:00:00Z","timezone":"Australia/Melbourne","source":"configured_schedule","time_band":"early_morning","time_acknowledged_at":"2026-11-26T08:30:00Z"}`
	var selection shipping.DeliverySelection
	if err := json.Unmarshal([]byte(wire), &selection); err != nil {
		t.Fatal(err)
	}
	if selection.TimeBand != shipping_enums.DeliveryTimeBandEarlyMorning || selection.TimeAcknowledgedAt == nil ||
		!selection.TimeAcknowledgedAt.Equal(time.Date(2026, 11, 26, 8, 30, 0, 0, time.UTC)) {
		t.Fatalf("selection acknowledgement not decoded: %+v", selection)
	}
	assertDeliveryJSON(t, selection, wire)

	orderWire, err := json.Marshal(order.Order{OrderNumber: "test-order", DeliverySelection: &selection})
	if err != nil {
		t.Fatal(err)
	}
	var received order.Order
	if err := json.Unmarshal(orderWire, &received); err != nil {
		t.Fatal(err)
	}
	shipmentWire, err := json.Marshal(fulfilment.OutboundShipment{OrderNumber: received.OrderNumber, DeliverySelection: received.DeliverySelection})
	if err != nil {
		t.Fatal(err)
	}
	var persisted fulfilment.OutboundShipment
	if err := json.Unmarshal(shipmentWire, &persisted); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(persisted.DeliverySelection, &selection) {
		t.Fatalf("time band or acknowledgement lost in handoff: %+v", persisted.DeliverySelection)
	}
}

func TestV5353LegacyDeliveryRecordsOmitTimeBand(t *testing.T) {
	for name, model := range map[string]any{
		"slot":      shipping.DeliverySlot{ID: "old-slot", Availability: "available"},
		"selection": shipping.DeliverySelection{SlotID: "old-slot"},
	} {
		data, err := json.Marshal(model)
		if err != nil {
			t.Fatal(err)
		}
		var fields map[string]json.RawMessage
		if err := json.Unmarshal(data, &fields); err != nil {
			t.Fatal(err)
		}
		for _, key := range []string{"time_band", "requires_time_acknowledgement", "time_acknowledged_at"} {
			if _, present := fields[key]; present {
				t.Errorf("%s: legacy record unexpectedly emits %s: %s", name, key, data)
			}
		}
	}
}
