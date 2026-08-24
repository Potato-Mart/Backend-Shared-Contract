package deletion_enums

// CommandReceiptStatus describes whether the participant accepted, is still
// processing, completed, blocked, or failed the idempotent command.
type CommandReceiptStatus string

const (
	CommandReceiptStatusAccepted   CommandReceiptStatus = "accepted"
	CommandReceiptStatusInProgress CommandReceiptStatus = "in_progress"
	CommandReceiptStatusCompleted  CommandReceiptStatus = "completed"
	CommandReceiptStatusBlocked    CommandReceiptStatus = "blocked"
	CommandReceiptStatusFailed     CommandReceiptStatus = "failed"
)

// IsValid reports whether s is a supported command-receipt status.
func (s CommandReceiptStatus) IsValid() bool {
	switch s {
	case CommandReceiptStatusAccepted, CommandReceiptStatusInProgress,
		CommandReceiptStatusCompleted, CommandReceiptStatusBlocked,
		CommandReceiptStatusFailed:
		return true
	default:
		return false
	}
}

// String returns the wire value for s.
func (s CommandReceiptStatus) String() string { return string(s) }
