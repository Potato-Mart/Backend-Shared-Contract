package paymentenum

// TerminalTxFinancialStatus is the final outcome of a terminal
// transaction. Populated only when TerminalTxStatus reaches Finalised
// or OverrideResolved.
//
// Unknown is a real outcome: it means the API could not determine the
// result after retries/status checks (network drop, terminal offline,
// missing webhook) and the merchant must manually confirm via the
// recovery flow.
type TerminalTxFinancialStatus string

const (
	TerminalTxFinancialStatusApproved  TerminalTxFinancialStatus = "approved"
	TerminalTxFinancialStatusDeclined  TerminalTxFinancialStatus = "declined"
	TerminalTxFinancialStatusCancelled TerminalTxFinancialStatus = "cancelled"
	TerminalTxFinancialStatusUnknown   TerminalTxFinancialStatus = "unknown"
)

// IsValid reports whether f is a known TerminalTxFinancialStatus.
func (f TerminalTxFinancialStatus) IsValid() bool {
	switch f {
	case TerminalTxFinancialStatusApproved, TerminalTxFinancialStatusDeclined,
		TerminalTxFinancialStatusCancelled, TerminalTxFinancialStatusUnknown:
		return true
	}
	return false
}

func (f TerminalTxFinancialStatus) String() string { return string(f) }
