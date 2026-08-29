// Package envelop holds the transport-neutral event envelope shared by every
// service that publishes or consumes Potato Mart domain events. The envelope
// carries delivery metadata only; the routed typed payload models live in
// pkg/contracts/pubsub/event and are serialized into Payload as JSON.
package envelope

import (
	"encoding/json"
	"time"
)

// EventEnvelope wraps every message published to a Potato Mart event topic.
// EventID is globally unique and is the consumer-side dedupe key; AggregateID
// doubles as the Pub/Sub ordering key (for example order_number or sku).
type EventEnvelope struct {
	EventID      string          `json:"event_id"`
	EventType    string          `json:"event_type"`
	EventVersion string          `json:"event_version"`
	OccurredAt   time.Time       `json:"occurred_at"`
	AggregateID  string          `json:"aggregate_id"`
	Payload      json.RawMessage `json:"payload"`
}
