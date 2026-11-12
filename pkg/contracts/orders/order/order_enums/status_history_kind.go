package order_enums

// StatusHistoryKind classifies a non-routine order status-history entry. An
// empty kind is an ordinary workflow transition. Override marks a staff
// override of the order's status, delivery, lines or fee; PaymentCorrection
// marks a staff correction of a recorded payment's amount or method.
type StatusHistoryKind string

const (
	StatusHistoryKindOverride          StatusHistoryKind = "override"
	StatusHistoryKindPaymentCorrection StatusHistoryKind = "payment_correction"
)

// IsValid reports whether k is a known StatusHistoryKind value.
func (k StatusHistoryKind) IsValid() bool {
	switch k {
	case StatusHistoryKindOverride, StatusHistoryKindPaymentCorrection:
		return true
	}
	return false
}

func (k StatusHistoryKind) String() string { return string(k) }
