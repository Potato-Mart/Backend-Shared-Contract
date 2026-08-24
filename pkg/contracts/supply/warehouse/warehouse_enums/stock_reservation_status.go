package warehouse_enums

// StockReservationStatus describes the logical lifecycle of a reservation.
type StockReservationStatus string

const (
	StockReservationStatusReserved        StockReservationStatus = "RESERVED"
	StockReservationStatusPartiallyStaged StockReservationStatus = "PARTIALLY_STAGED"
	StockReservationStatusStaged          StockReservationStatus = "STAGED"
	StockReservationStatusCommitted       StockReservationStatus = "COMMITTED"
	StockReservationStatusReleased        StockReservationStatus = "RELEASED"
	StockReservationStatusExpired         StockReservationStatus = "EXPIRED"
	StockReservationStatusCancelled       StockReservationStatus = "CANCELLED"
)

func (s StockReservationStatus) IsValid() bool {
	switch s {
	case StockReservationStatusReserved, StockReservationStatusPartiallyStaged,
		StockReservationStatusStaged, StockReservationStatusCommitted,
		StockReservationStatusReleased, StockReservationStatusExpired,
		StockReservationStatusCancelled:
		return true
	}
	return false
}

func (s StockReservationStatus) String() string { return string(s) }
