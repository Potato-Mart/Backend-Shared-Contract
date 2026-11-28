package ads

import "time"

// AdSchedule is the delivery window of one platform object.
//
// Both ends are optional: an absent StartAt means the object is already
// running, and an absent EndAt means it runs until it is paused. Instants are
// absolute, so a reader converts them into the ad account time zone before
// comparing them with a platform screen. A start time that has already passed
// cannot be changed, which the ads module enforces when it validates a change.
type AdSchedule struct {
	StartAt *time.Time `json:"start_at,omitempty"`
	EndAt   *time.Time `json:"end_at,omitempty"`
}
