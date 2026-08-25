package security

import "time"

// CustodyEvent records one evidence custody action.
type CustodyEvent struct {
	OccurredAt time.Time `json:"occurred_at"`
	ActorID    string    `json:"actor_id,omitempty"`
	Action     string    `json:"action"` // e.g. "collected", "transferred", "reviewed", "sealed"
	Reason     string    `json:"reason,omitempty"`
}
