// Package deletion_enums defines the finite wire values used by trusted
// account-deletion coordination messages.
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

// AccountDeletionParticipant identifies the seven services that acknowledge
// restriction, eligibility, erasure/deidentification, and unrestriction.
type AccountDeletionParticipant string

const (
	AccountDeletionParticipantIdentity  AccountDeletionParticipant = "identity"
	AccountDeletionParticipantCustomers AccountDeletionParticipant = "customers"
	AccountDeletionParticipantOrders    AccountDeletionParticipant = "orders"
	AccountDeletionParticipantPayments  AccountDeletionParticipant = "payments"
	AccountDeletionParticipantPricing   AccountDeletionParticipant = "pricing"
	AccountDeletionParticipantSupply    AccountDeletionParticipant = "supply"
	AccountDeletionParticipantInsights  AccountDeletionParticipant = "insights"
)

// IsValid reports whether p is a configured deletion-workflow participant.
func (p AccountDeletionParticipant) IsValid() bool {
	switch p {
	case AccountDeletionParticipantIdentity, AccountDeletionParticipantCustomers,
		AccountDeletionParticipantOrders, AccountDeletionParticipantPayments,
		AccountDeletionParticipantPricing, AccountDeletionParticipantSupply,
		AccountDeletionParticipantInsights:
		return true
	default:
		return false
	}
}

// String returns the wire value for p.
func (p AccountDeletionParticipant) String() string { return string(p) }

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

// AccountDeletionBlocker is a code-only reason that execution cannot proceed.
// A receipt must report no record identifier, amount, contact detail, provider
// reference, or free-form explanation alongside these values.
type AccountDeletionBlocker string

const (
	AccountDeletionBlockerPendingOrder               AccountDeletionBlocker = "pending_order"
	AccountDeletionBlockerPendingFulfillment         AccountDeletionBlocker = "pending_fulfillment"
	AccountDeletionBlockerPendingCapture             AccountDeletionBlocker = "pending_capture"
	AccountDeletionBlockerPayToProcessing            AccountDeletionBlocker = "payto_processing"
	AccountDeletionBlockerPendingRefund              AccountDeletionBlocker = "pending_refund"
	AccountDeletionBlockerOpenDispute                AccountDeletionBlocker = "open_dispute"
	AccountDeletionBlockerOpenChargeback             AccountDeletionBlocker = "open_chargeback"
	AccountDeletionBlockerOpenInvoice                AccountDeletionBlocker = "open_invoice"
	AccountDeletionBlockerOutstandingDebt            AccountDeletionBlocker = "outstanding_debt"
	AccountDeletionBlockerActiveSubscription         AccountDeletionBlocker = "active_subscription"
	AccountDeletionBlockerActiveGroupOrder           AccountDeletionBlocker = "active_group_order"
	AccountDeletionBlockerActiveReservation          AccountDeletionBlocker = "active_reservation"
	AccountDeletionBlockerRecoverablePersona         AccountDeletionBlocker = "recoverable_persona"
	AccountDeletionBlockerLegalHold                  AccountDeletionBlocker = "legal_hold"
	AccountDeletionBlockerProtectedStoredValue       AccountDeletionBlocker = "protected_stored_value"
	AccountDeletionBlockerStoredValueAcknowledgement AccountDeletionBlocker = "stored_value_acknowledgement_required"
	AccountDeletionBlockerProviderRevocationFailure  AccountDeletionBlocker = "provider_revocation_failure"
)

// IsValid reports whether b is a supported account-deletion blocker.
func (b AccountDeletionBlocker) IsValid() bool {
	switch b {
	case AccountDeletionBlockerPendingOrder,
		AccountDeletionBlockerPendingFulfillment,
		AccountDeletionBlockerPendingCapture,
		AccountDeletionBlockerPayToProcessing,
		AccountDeletionBlockerPendingRefund,
		AccountDeletionBlockerOpenDispute,
		AccountDeletionBlockerOpenChargeback,
		AccountDeletionBlockerOpenInvoice,
		AccountDeletionBlockerOutstandingDebt,
		AccountDeletionBlockerActiveSubscription,
		AccountDeletionBlockerActiveGroupOrder,
		AccountDeletionBlockerActiveReservation,
		AccountDeletionBlockerRecoverablePersona,
		AccountDeletionBlockerLegalHold,
		AccountDeletionBlockerProtectedStoredValue,
		AccountDeletionBlockerStoredValueAcknowledgement,
		AccountDeletionBlockerProviderRevocationFailure:
		return true
	default:
		return false
	}
}

// String returns the wire value for b.
func (b AccountDeletionBlocker) String() string { return string(b) }

// ErasureDisposition is the record-free summary of what an
// erase_or_deidentify command achieved. Mixed means the service both erased
// data and deidentified policy-retained evidence; it never indicates which
// collection or source record was involved.
type ErasureDisposition string

const (
	ErasureDispositionErased           ErasureDisposition = "erased"
	ErasureDispositionDeidentified     ErasureDisposition = "deidentified"
	ErasureDispositionMixed            ErasureDisposition = "mixed"
	ErasureDispositionNoSubjectData    ErasureDisposition = "no_subject_data"
	ErasureDispositionRetainedByPolicy ErasureDisposition = "retained_by_policy"
)

// IsValid reports whether d is a supported erasure disposition.
func (d ErasureDisposition) IsValid() bool {
	switch d {
	case ErasureDispositionErased, ErasureDispositionDeidentified,
		ErasureDispositionMixed, ErasureDispositionNoSubjectData,
		ErasureDispositionRetainedByPolicy:
		return true
	default:
		return false
	}
}

// String returns the wire value for d.
func (d ErasureDisposition) String() string { return string(d) }
