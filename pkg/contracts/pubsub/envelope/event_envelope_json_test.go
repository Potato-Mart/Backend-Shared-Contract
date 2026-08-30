package envelope_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pubsub/envelope"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pubsub/routing"
)

func TestEventEnvelopeCarriesTypedRoutingAndDeliveryEvidence(t *testing.T) {
	value := envelope.EventEnvelope{
		EventID:          "event_1",
		EventType:        routing.EventTypeOrderPaid,
		EventVersion:     envelope.EventVersion("v3"),
		OccurredAt:       time.Date(2026, 8, 30, 1, 2, 3, 0, time.UTC),
		AggregateID:      "SO-1",
		CorrelationID:    "corr_1",
		CausationEventID: "event_0",
		RequestID:        "req_1",
		Payload:          json.RawMessage(`{"order_number":"SO-1"}`),
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal event envelope: %v", err)
	}
	var shape map[string]any
	if err := json.Unmarshal(encoded, &shape); err != nil {
		t.Fatalf("unmarshal event envelope: %v", err)
	}
	for _, key := range []string{"event_type", "event_version", "correlation_id", "causation_event_id", "request_id"} {
		if _, ok := shape[key]; !ok {
			t.Fatalf("event envelope missing %q: %+v", key, shape)
		}
	}
	if shape["event_type"] != "order.paid" || shape["event_version"] != "v3" {
		t.Fatalf("event envelope routing did not marshal: %+v", shape)
	}
}
