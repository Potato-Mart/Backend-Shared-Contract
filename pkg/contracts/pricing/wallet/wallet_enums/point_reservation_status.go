package wallet_enums

// PointReservationStatus is the lifecycle of a wallet points reservation.
type PointReservationStatus string

const (
	PointReservationStatusReserved  PointReservationStatus = "RESERVED"
	PointReservationStatusCommitted PointReservationStatus = "COMMITTED"
	PointReservationStatusCancelled PointReservationStatus = "CANCELLED"
	PointReservationStatusExpired   PointReservationStatus = "EXPIRED"
)

func (s PointReservationStatus) IsValid() bool {
	switch s {
	case PointReservationStatusReserved, PointReservationStatusCommitted, PointReservationStatusCancelled, PointReservationStatusExpired:
		return true
	}
	return false
}
func (s PointReservationStatus) String() string { return string(s) }
