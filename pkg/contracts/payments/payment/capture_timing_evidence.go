package payment

import "time"

// CaptureTimingEvidence preserves what is actually known about completed capture.
// CapturedAt is present only for an authoritative exact capture timestamp.
// Confirmation is a separate upper bound, not a substitute exact timestamp.
// Payments validates provenance, amount/payment correlation and consistency.
// At least one form of evidence is required by services using this model;
// absence means unknown timing and must never default to processing time.
type CaptureTimingEvidence struct {
	CapturedAt   *time.Time                   `json:"captured_at,omitempty"`
	Confirmation *CaptureConfirmationEvidence `json:"confirmation,omitempty"`
}
