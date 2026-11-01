package courier

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/temporal"

// CarrierBookingEdits are the typed values staff changed before retrying a
// carrier booking. A nil field was not edited and the attempt used the value
// derived from the order and its shipments. For the text fields Instructions
// and DeliveryType, a non-nil empty string means staff cleared the value.
// Edits apply only to the attempt that carries them; they do not rewrite the
// order's accepted delivery selection.
//
// IntakeDate is the depot intake (production) date a carrier may require.
// TimeWindow holds wall-clock times in the delivery (recipient's) timezone;
// either side may be omitted to keep the value originally sent.
// DeliveryType is the carrier's free-text delivery-type label (for example
// Home delivery). A carrier ignores fields its booking does not carry.
// Instructions may contain personal data.
type CarrierBookingEdits struct {
	IntakeDate   *temporal.Date            `json:"intake_date,omitempty"`
	DeliveryDate *temporal.Date            `json:"delivery_date,omitempty"`
	TimeWindow   *CarrierBookingTimeWindow `json:"time_window,omitempty"`
	BoxCount     *int                      `json:"box_count,omitempty"`
	Instructions *string                   `json:"instructions,omitempty"`
	DeliveryType *string                   `json:"delivery_type,omitempty"`
}
