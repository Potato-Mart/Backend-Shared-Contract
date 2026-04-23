package enums

// PaymentStatus tracks the money side of an order separately from the
// fulfilment side (OrderStatus). Separating them makes it possible to
// model refunds, partial refunds and offline payments cleanly.
type PaymentStatus string

const (
	PaymentStatusUnpaid          PaymentStatus = "UNPAID"
	PaymentStatusPending         PaymentStatus = "PENDING"
	PaymentStatusPaid            PaymentStatus = "PAID"
	PaymentStatusPartiallyPaid   PaymentStatus = "PARTIALLY_PAID"
	PaymentStatusRefunded        PaymentStatus = "REFUNDED"
	PaymentStatusPartialRefunded PaymentStatus = "PARTIALLY_REFUNDED"
)

// IsValid reports whether p is a known PaymentStatus.
func (p PaymentStatus) IsValid() bool {
	switch p {
	case PaymentStatusUnpaid, PaymentStatusPending, PaymentStatusPaid,
		PaymentStatusPartiallyPaid, PaymentStatusRefunded, PaymentStatusPartialRefunded:
		return true
	}
	return false
}

func (p PaymentStatus) String() string { return string(p) }
