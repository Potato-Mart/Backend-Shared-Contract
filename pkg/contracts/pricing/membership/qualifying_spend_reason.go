package membership

// QualifyingSpendReason identifies why qualifying spend changed.
type QualifyingSpendReason string

const (
	QualifyingSpendReasonOrderPaid QualifyingSpendReason = "order_paid"
	QualifyingSpendReasonRefund    QualifyingSpendReason = "refund"
)

func (r QualifyingSpendReason) IsValid() bool {
	switch r {
	case QualifyingSpendReasonOrderPaid, QualifyingSpendReasonRefund:
		return true
	default:
		return false
	}
}

func (r QualifyingSpendReason) String() string { return string(r) }
