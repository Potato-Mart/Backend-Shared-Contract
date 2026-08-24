package deletion_enums

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
