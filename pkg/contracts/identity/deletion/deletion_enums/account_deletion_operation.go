package deletion_enums

// AccountDeletionOperation identifies one idempotent service command in the
// account-deletion workflow.
type AccountDeletionOperation string

const (
	// AccountDeletionOperationRestrict blocks commerce and sensitive profile
	// changes while leaving only the explicitly permitted deletion actions.
	AccountDeletionOperationRestrict AccountDeletionOperation = "restrict"
	// AccountDeletionOperationEligibility rechecks whether the receiving
	// service has a durable obligation that blocks execution.
	AccountDeletionOperationEligibility AccountDeletionOperation = "eligibility"
	// AccountDeletionOperationEraseOrDeidentify removes subject-bound data or
	// irreversibly deidentifies evidence that a lawful retention policy keeps.
	AccountDeletionOperationEraseOrDeidentify AccountDeletionOperation = "erase_or_deidentify"
	// AccountDeletionOperationUnrestrict restores access after cancellation
	// before the coordinator has atomically entered execution.
	AccountDeletionOperationUnrestrict AccountDeletionOperation = "unrestrict"
)

// IsValid reports whether o is a supported account-deletion operation.
func (o AccountDeletionOperation) IsValid() bool {
	switch o {
	case AccountDeletionOperationRestrict, AccountDeletionOperationEligibility,
		AccountDeletionOperationEraseOrDeidentify, AccountDeletionOperationUnrestrict:
		return true
	default:
		return false
	}
}

// String returns the wire value for o.
func (o AccountDeletionOperation) String() string { return string(o) }
