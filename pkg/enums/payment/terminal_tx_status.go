package paymentenum

type TerminalTxStatus string

const (
	TerminalTxStatusUnknown          TerminalTxStatus = "unknown"
	TerminalTxStatusPending          TerminalTxStatus = "pending"
	TerminalTxStatusAwaitingAction   TerminalTxStatus = "awaiting_action"
	TerminalTxStatusFinalised        TerminalTxStatus = "finalised"
	TerminalTxStatusOverridePending  TerminalTxStatus = "override_pending"
	TerminalTxStatusOverrideResolved TerminalTxStatus = "override_resolved"
)

func (s TerminalTxStatus) IsValid() bool {
	switch s {
	case TerminalTxStatusUnknown, TerminalTxStatusPending, TerminalTxStatusAwaitingAction,
		TerminalTxStatusFinalised, TerminalTxStatusOverridePending, TerminalTxStatusOverrideResolved:
		return true
	default:
		return false
	}
}

func (s TerminalTxStatus) String() string { return string(s) }
