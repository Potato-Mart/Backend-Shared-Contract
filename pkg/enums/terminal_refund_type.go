package enums

// TerminalRefundType distinguishes Adyen's two in-person refund paths:
// a referenced refund using a ReversalRequest, or an unreferenced refund
// using a PaymentRequest with PaymentType=Refund.
type TerminalRefundType string

const (
	TerminalRefundTypeReferenced   TerminalRefundType = "referenced"
	TerminalRefundTypeUnreferenced TerminalRefundType = "unreferenced"
)

// IsValid reports whether t is a known TerminalRefundType.
func (t TerminalRefundType) IsValid() bool {
	switch t {
	case TerminalRefundTypeReferenced, TerminalRefundTypeUnreferenced:
		return true
	}
	return false
}

func (t TerminalRefundType) String() string { return string(t) }
