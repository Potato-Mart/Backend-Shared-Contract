package enums

// TerminalTxStatus is the lifecycle state of a transaction created on a
// payment terminal.
//
// Pending covers a request accepted by the provider but not yet resolved.
// AwaitingAction covers terminal or merchant interaction such as signature,
// MOTO entry, AVS review, or manual recovery after a timeout.
type TerminalTxStatus string

const (
	TerminalTxStatusUnknown          TerminalTxStatus = "unknown"
	TerminalTxStatusPending          TerminalTxStatus = "pending"
	TerminalTxStatusAwaitingAction   TerminalTxStatus = "awaiting_action"
	TerminalTxStatusFinalised        TerminalTxStatus = "finalised"
	TerminalTxStatusOverridePending  TerminalTxStatus = "override_pending"
	TerminalTxStatusOverrideResolved TerminalTxStatus = "override_resolved"
)

// IsValid reports whether s is a known TerminalTxStatus.
func (s TerminalTxStatus) IsValid() bool {
	switch s {
	case TerminalTxStatusUnknown, TerminalTxStatusPending, TerminalTxStatusAwaitingAction, TerminalTxStatusFinalised,
		TerminalTxStatusOverridePending, TerminalTxStatusOverrideResolved:
		return true
	}
	return false
}

// IsTerminal reports whether transaction status checks should stop in
// this state.
func (s TerminalTxStatus) IsTerminal() bool {
	return s == TerminalTxStatusFinalised || s == TerminalTxStatusOverrideResolved
}

func (s TerminalTxStatus) String() string { return string(s) }
