package courier_enums

// CarrierCancellationOutcome is the result of asking a carrier to cancel a
// booking. Cancelled means the carrier accepted the cancellation;
// AlreadyCancelled means the carrier reported the job as already cancelled.
// NotConfigured means the company's configuration defines no cancel request.
// Failed means the carrier refused or answered with an error; Unreachable
// means no carrier answer was obtained, so the carrier-side state is unknown.
// Every outcome other than cancelled and already_cancelled needs a manual
// action.
type CarrierCancellationOutcome string

const (
	CarrierCancellationOutcomeCancelled        CarrierCancellationOutcome = "cancelled"
	CarrierCancellationOutcomeAlreadyCancelled CarrierCancellationOutcome = "already_cancelled"
	CarrierCancellationOutcomeNotConfigured    CarrierCancellationOutcome = "not_configured"
	CarrierCancellationOutcomeFailed           CarrierCancellationOutcome = "failed"
	CarrierCancellationOutcomeUnreachable      CarrierCancellationOutcome = "unreachable"
)

func (o CarrierCancellationOutcome) IsValid() bool {
	switch o {
	case CarrierCancellationOutcomeCancelled, CarrierCancellationOutcomeAlreadyCancelled,
		CarrierCancellationOutcomeNotConfigured, CarrierCancellationOutcomeFailed, CarrierCancellationOutcomeUnreachable:
		return true
	default:
		return false
	}
}

func (o CarrierCancellationOutcome) String() string { return string(o) }
