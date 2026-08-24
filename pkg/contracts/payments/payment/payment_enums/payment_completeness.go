package payment_enums

// PaymentCompleteness states whether a customer payment summary could be
// fully assembled from the underlying records.
type PaymentCompleteness string

const (
	PaymentCompletenessComplete    PaymentCompleteness = "complete"
	PaymentCompletenessPartial     PaymentCompleteness = "partial"
	PaymentCompletenessUnavailable PaymentCompleteness = "unavailable"
)

func (c PaymentCompleteness) IsValid() bool {
	switch c {
	case PaymentCompletenessComplete, PaymentCompletenessPartial, PaymentCompletenessUnavailable:
		return true
	default:
		return false
	}
}

func (c PaymentCompleteness) String() string { return string(c) }
