package order

import "time"

type StatusHistory struct {
	FromValue  string    `json:"from_value,omitempty"`
	ToValue    string    `json:"to_value"`
	Note       string    `json:"note,omitempty"`
	ActorID    string    `json:"actor_id,omitempty"`
	ActorEmail string    `json:"actor_email,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}
