package payment_enums

// PaymentStatus tracks the money side of an order separately from the
// fulfilment side (OrderStatus). Separating them makes it possible to
// model refunds, partial refunds and offline payments cleanly.
type PaymentStatus string

const (
	PaymentStatusUnknown         PaymentStatus = "unknown"
	PaymentStatusUnpaid          PaymentStatus = "unpaid"
	PaymentStatusPending         PaymentStatus = "pending"
	PaymentStatusPaid            PaymentStatus = "paid"
	PaymentStatusPartiallyPaid   PaymentStatus = "partially_paid"
	PaymentStatusRefunded        PaymentStatus = "refunded"
	PaymentStatusPartialRefunded PaymentStatus = "partially_refunded"
)

// IsValid reports whether p is a known PaymentStatus.
func (p PaymentStatus) IsValid() bool {
	switch p {
	case PaymentStatusUnknown, PaymentStatusUnpaid, PaymentStatusPending, PaymentStatusPaid,
		PaymentStatusPartiallyPaid, PaymentStatusRefunded, PaymentStatusPartialRefunded:
		return true
	}
	return false
}

func (p PaymentStatus) String() string { return string(p) }
