package courier

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier/courier_enums"
)

// CarrierBookingSummary is the compact projection of an order's latest
// CarrierBookingAttempt, embedded on each shipment the attempt covers.
// OutcomeAt is the attempt's resolved_at once resolved, otherwise its
// attempted_at. The attempt remains the source of truth.
type CarrierBookingSummary struct {
	AttemptID        string                              `json:"attempt_id"`
	Outcome          courier_enums.CarrierBookingOutcome `json:"outcome"`
	CarrierReference string                              `json:"carrier_reference,omitempty"`
	OutcomeAt        time.Time                           `json:"outcome_at"`
}
