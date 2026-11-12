package courier_enums

// CarrierManualActionKind is the work a carrier manual action asks staff to
// do with the carrier. CancelBooking means a superseded carrier booking must
// be cancelled with the carrier by hand.
type CarrierManualActionKind string

const (
	CarrierManualActionKindCancelBooking CarrierManualActionKind = "cancel_booking"
)

func (k CarrierManualActionKind) IsValid() bool {
	switch k {
	case CarrierManualActionKindCancelBooking:
		return true
	default:
		return false
	}
}

func (k CarrierManualActionKind) String() string { return string(k) }
