package audit

import "time"

// LifecycleAction records who performed one lifecycle action, when it
// happened, and the optional human/business reason for it.
type LifecycleAction struct {
	By     string    `json:"by"`
	At     time.Time `json:"at"`
	Reason string    `json:"reason,omitempty"`
}
