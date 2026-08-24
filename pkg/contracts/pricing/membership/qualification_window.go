package membership

import "time"

// QualificationWindow is the period over which qualifying spend is measured.
type QualificationWindow struct {
	StartsAt *time.Time `json:"starts_at,omitempty"`
	EndsAt   *time.Time `json:"ends_at,omitempty"`
}
