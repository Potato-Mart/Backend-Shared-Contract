// Package envelope holds the transport-neutral event envelope shared by every
// service that publishes or consumes Potato Mart domain events. The envelope
// carries delivery metadata only; the routed typed payload models live in
// producer-owned pubsub packages and are serialized into Payload as JSON.
package envelope

import (
	"encoding/json"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pubsub/routing"
)

// EventVersion is an open publisher-controlled schema code, such as "v4".
// Consumers must preserve unknown values for forward compatibility.
type EventVersion string

// EventEnvelope wraps every message published to a Potato Mart event topic.
// EventID is globally unique and is the consumer-side dedupe key; AggregateID
// doubles as the Pub/Sub ordering key (for example order_number or sku).
type EventEnvelope struct {
	EventID          string            `json:"event_id"`
	EventType        routing.EventType `json:"event_type"`
	EventVersion     EventVersion      `json:"event_version"`
	OccurredAt       time.Time         `json:"occurred_at"`
	AggregateID      string            `json:"aggregate_id"`
	CorrelationID    string            `json:"correlation_id,omitempty"`
	CausationEventID string            `json:"causation_event_id,omitempty"`
	RequestID        string            `json:"request_id,omitempty"`
	Payload          json.RawMessage   `json:"payload"`
}
