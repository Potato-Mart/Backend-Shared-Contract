package enums

// TerminalTxStatus is the polling-loop state of a transaction created
// on a payment terminal.
//
// MX51 SCI surfaces three native states - PENDING, AWAITING_POS,
// FINALISED - mapped here to Pending, AwaitingPOS, and Finalised.
// OverridePending and OverrideResolved capture the mandatory recovery
// flow when the API stops responding within the POS-defined timeout
// and the merchant must manually confirm the outcome.
type TerminalTxStatus string

const (
	TerminalTxStatusPending          TerminalTxStatus = "pending"
	TerminalTxStatusAwaitingPOS      TerminalTxStatus = "awaiting_pos"
	TerminalTxStatusFinalised        TerminalTxStatus = "finalised"
	TerminalTxStatusOverridePending  TerminalTxStatus = "override_pending"
	TerminalTxStatusOverrideResolved TerminalTxStatus = "override_resolved"
)

// IsValid reports whether s is a known TerminalTxStatus.
func (s TerminalTxStatus) IsValid() bool {
	switch s {
	case TerminalTxStatusPending, TerminalTxStatusAwaitingPOS, TerminalTxStatusFinalised,
		TerminalTxStatusOverridePending, TerminalTxStatusOverrideResolved:
		return true
	}
	return false
}

// IsTerminal reports whether polling should stop in this state.
func (s TerminalTxStatus) IsTerminal() bool {
	return s == TerminalTxStatusFinalised || s == TerminalTxStatusOverrideResolved
}

func (s TerminalTxStatus) String() string { return string(s) }
