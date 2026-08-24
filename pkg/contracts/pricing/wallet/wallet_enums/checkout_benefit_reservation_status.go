package wallet_enums

// CheckoutBenefitReservationStatus is the lifecycle of a checkout holding.
type CheckoutBenefitReservationStatus string

const (
	CheckoutBenefitReservationStatusReserved          CheckoutBenefitReservationStatus = "reserved"
	CheckoutBenefitReservationStatusCommitted         CheckoutBenefitReservationStatus = "committed"
	CheckoutBenefitReservationStatusCancelled         CheckoutBenefitReservationStatus = "cancelled"
	CheckoutBenefitReservationStatusExpired           CheckoutBenefitReservationStatus = "expired"
	CheckoutBenefitReservationStatusPartiallyRefunded CheckoutBenefitReservationStatus = "partially_refunded"
	CheckoutBenefitReservationStatusRefunded          CheckoutBenefitReservationStatus = "refunded"
)

func (s CheckoutBenefitReservationStatus) IsValid() bool {
	switch s {
	case CheckoutBenefitReservationStatusReserved, CheckoutBenefitReservationStatusCommitted, CheckoutBenefitReservationStatusCancelled, CheckoutBenefitReservationStatusExpired, CheckoutBenefitReservationStatusPartiallyRefunded, CheckoutBenefitReservationStatusRefunded:
		return true
	}
	return false
}
func (s CheckoutBenefitReservationStatus) String() string { return string(s) }
