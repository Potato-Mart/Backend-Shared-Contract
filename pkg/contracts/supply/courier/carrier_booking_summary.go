package courier

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/courier/courier_enums"
)

// CarrierBookingSummary is the compact projection of an order's latest
// CarrierBookingAttempt, embedded on each shipment the attempt covers.
// OutcomeAt is the attempt's resolved_at once resolved, otherwise its
// attempted_at. The attempt remains the source of truth.
//
// DeliveryDate, IntakeDate, TimeWindow, BoxCount, DeliveryType and
// Instructions are the values actually sent to the carrier in that latest
// booking, after any retry edits were applied, written back onto the
// shipment. A nil field was not sent to the carrier. A present empty
// DeliveryType or Instructions was sent empty. TimeWindow uses wall-clock
// times in the delivery (recipient's) timezone.
type CarrierBookingSummary struct {
	AttemptID        string                              `json:"attempt_id"`
	Outcome          courier_enums.CarrierBookingOutcome `json:"outcome"`
	CarrierReference string                              `json:"carrier_reference,omitempty"`
	OutcomeAt        time.Time                           `json:"outcome_at"`
	DeliveryDate     *temporal.Date                      `json:"delivery_date,omitempty"`
	IntakeDate       *temporal.Date                      `json:"intake_date,omitempty"`
	TimeWindow       *CarrierBookingTimeWindow           `json:"time_window,omitempty"`
	BoxCount         *int                                `json:"box_count,omitempty"`
	DeliveryType     *string                             `json:"delivery_type,omitempty"`
	Instructions     *string                             `json:"instructions,omitempty"`
}
