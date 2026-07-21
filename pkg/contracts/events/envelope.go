// Package events holds the transport-neutral event envelope shared by every
// service that publishes or consumes Potato Mart domain events. The envelope
// carries delivery metadata only; the typed payload models live with their
// owning domain package (for example contracts/sales or contracts/payments)
// and are serialized into Payload as JSON.
package events

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
