package supply_test

import (
	"encoding/json"
	"strings"
	"testing"

	event "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pubsub/supply"
)

func TestFulfilmentPayloadsRelyOnEnvelopeOccurredAt(t *testing.T) {
	payloads := []any{
		event.FulfilmentShippedEvent{OrderNumber: "SO-1"},
		event.FulfilmentDeliveredEvent{OrderNumber: "SO-1"},
		event.FulfilmentCompletedEvent{OrderNumber: "SO-1"},
		event.FulfilmentTrackingEvent{OrderNumber: "SO-1", ShipmentID: "shipment_1", Carrier: "auspost", TrackingNumber: "tracking_1"},
	}
	for _, value := range payloads {
		encoded, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal %T: %v", value, err)
		}
		var shape map[string]any
		if err := json.Unmarshal(encoded, &shape); err != nil {
			t.Fatalf("unmarshal %T: %v", value, err)
		}
		if _, ok := shape["occurred_at"]; ok {
			t.Fatalf("%T payload must rely on envelope occurred_at: %+v", value, shape)
		}
	}
}

func TestFulfilmentTrackingEventCarriesShipmentAndCarrier(t *testing.T) {
	value := event.FulfilmentTrackingEvent{
		OrderNumber: "SO-1", ShipmentID: "shipment_1", Carrier: "auspost", TrackingNumber: "tracking_1",
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal fulfilment tracking event: %v", err)
	}
	for _, want := range []string{`"shipment_id":"shipment_1"`, `"carrier":"auspost"`} {
		if !strings.Contains(string(encoded), want) {
			t.Fatalf("fulfilment tracking JSON missing %s: %s", want, encoded)
		}
	}
}
