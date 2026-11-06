package orders

import "time"

// PickingFactPicker is one PDA picking assignment half: who picked it, from
// the first pick to the half being done, and what it picked.
type PickingFactPicker struct {
	Role          string                  `json:"role"`
	PickerSubject string                  `json:"picker_subject"`
	StartedAt     *time.Time              `json:"started_at,omitempty"`
	CompletedAt   *time.Time              `json:"completed_at,omitempty"`
	Lines         []PickingFactPickerLine `json:"lines,omitempty"`
}
