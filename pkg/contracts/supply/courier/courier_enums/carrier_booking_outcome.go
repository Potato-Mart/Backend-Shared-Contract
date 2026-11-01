package courier_enums

// CarrierBookingOutcome is the state of one carrier booking attempt for an
// order. Pending means the request was sent and no final carrier answer has
// been recorded. Booked means the carrier accepted the job; AlreadyBooked
// means the carrier reported an existing job for the same order and returned
// its reference. Rejected means the carrier refused the request with a mapped
// error; Unreachable means no carrier answer was obtained (network failure or
// timeout), so the carrier-side state is unknown.
type CarrierBookingOutcome string

const (
	CarrierBookingOutcomePending       CarrierBookingOutcome = "pending"
	CarrierBookingOutcomeBooked        CarrierBookingOutcome = "booked"
	CarrierBookingOutcomeAlreadyBooked CarrierBookingOutcome = "already_booked"
	CarrierBookingOutcomeRejected      CarrierBookingOutcome = "rejected"
	CarrierBookingOutcomeUnreachable   CarrierBookingOutcome = "unreachable"
)

func (o CarrierBookingOutcome) IsValid() bool {
	switch o {
	case CarrierBookingOutcomePending, CarrierBookingOutcomeBooked, CarrierBookingOutcomeAlreadyBooked,
		CarrierBookingOutcomeRejected, CarrierBookingOutcomeUnreachable:
		return true
	default:
		return false
	}
}

func (o CarrierBookingOutcome) String() string { return string(o) }
