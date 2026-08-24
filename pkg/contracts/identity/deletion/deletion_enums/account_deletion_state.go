package deletion_enums

// AccountDeletionState is the durable coordinator lifecycle for one request.
// It deliberately separates action_required from approval and retains
// execution_failed as a restricted, retryable state after execution begins.
type AccountDeletionState string

const (
	AccountDeletionStateRestricting     AccountDeletionState = "restricting"
	AccountDeletionStateUnderReview     AccountDeletionState = "under_review"
	AccountDeletionStateActionRequired  AccountDeletionState = "action_required"
	AccountDeletionStateApproved        AccountDeletionState = "approved"
	AccountDeletionStateExecuting       AccountDeletionState = "executing"
	AccountDeletionStateCompleted       AccountDeletionState = "completed"
	AccountDeletionStateCancelling      AccountDeletionState = "cancelling"
	AccountDeletionStateCancelled       AccountDeletionState = "cancelled"
	AccountDeletionStateExecutionFailed AccountDeletionState = "execution_failed"
)

// IsValid reports whether s is a supported account-deletion workflow state.
func (s AccountDeletionState) IsValid() bool {
	switch s {
	case AccountDeletionStateRestricting, AccountDeletionStateUnderReview,
		AccountDeletionStateActionRequired, AccountDeletionStateApproved,
		AccountDeletionStateExecuting, AccountDeletionStateCompleted,
		AccountDeletionStateCancelling, AccountDeletionStateCancelled,
		AccountDeletionStateExecutionFailed:
		return true
	default:
		return false
	}
}

// String returns the wire value for s.
func (s AccountDeletionState) String() string { return string(s) }
