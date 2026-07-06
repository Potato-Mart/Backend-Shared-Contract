package paymentenum

// TerminalRefundType distinguishes referenced and unreferenced in-person
// refund paths.
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
