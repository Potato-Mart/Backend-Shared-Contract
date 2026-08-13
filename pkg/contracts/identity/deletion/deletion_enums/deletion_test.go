package deletion_enums

import "testing"

func TestAccountDeletionEnumsAcceptOnlyKnownWireValues(t *testing.T) {
	tests := []struct {
		name  string
		valid bool
	}{
		{"operation restrict", AccountDeletionOperationRestrict.IsValid()},
		{"operation eligibility", AccountDeletionOperationEligibility.IsValid()},
		{"operation erase or deidentify", AccountDeletionOperationEraseOrDeidentify.IsValid()},
		{"operation unrestrict", AccountDeletionOperationUnrestrict.IsValid()},
		{"participant identity", AccountDeletionParticipantIdentity.IsValid()},
		{"participant customers", AccountDeletionParticipantCustomers.IsValid()},
		{"participant orders", AccountDeletionParticipantOrders.IsValid()},
		{"participant payments", AccountDeletionParticipantPayments.IsValid()},
		{"participant pricing", AccountDeletionParticipantPricing.IsValid()},
		{"participant supply", AccountDeletionParticipantSupply.IsValid()},
		{"participant insights", AccountDeletionParticipantInsights.IsValid()},
		{"state restricting", AccountDeletionStateRestricting.IsValid()},
		{"state under review", AccountDeletionStateUnderReview.IsValid()},
		{"state action required", AccountDeletionStateActionRequired.IsValid()},
		{"state approved", AccountDeletionStateApproved.IsValid()},
		{"state executing", AccountDeletionStateExecuting.IsValid()},
		{"state completed", AccountDeletionStateCompleted.IsValid()},
		{"state cancelling", AccountDeletionStateCancelling.IsValid()},
		{"state cancelled", AccountDeletionStateCancelled.IsValid()},
		{"state execution failed", AccountDeletionStateExecutionFailed.IsValid()},
		{"receipt status accepted", CommandReceiptStatusAccepted.IsValid()},
		{"receipt status in progress", CommandReceiptStatusInProgress.IsValid()},
		{"receipt status completed", CommandReceiptStatusCompleted.IsValid()},
		{"receipt status blocked", CommandReceiptStatusBlocked.IsValid()},
		{"receipt status failed", CommandReceiptStatusFailed.IsValid()},
		{"blocker pending order", AccountDeletionBlockerPendingOrder.IsValid()},
		{"blocker pending fulfilment", AccountDeletionBlockerPendingFulfillment.IsValid()},
		{"blocker pending capture", AccountDeletionBlockerPendingCapture.IsValid()},
		{"blocker payto", AccountDeletionBlockerPayToProcessing.IsValid()},
		{"blocker pending refund", AccountDeletionBlockerPendingRefund.IsValid()},
		{"blocker dispute", AccountDeletionBlockerOpenDispute.IsValid()},
		{"blocker chargeback", AccountDeletionBlockerOpenChargeback.IsValid()},
		{"blocker invoice", AccountDeletionBlockerOpenInvoice.IsValid()},
		{"blocker debt", AccountDeletionBlockerOutstandingDebt.IsValid()},
		{"blocker subscription", AccountDeletionBlockerActiveSubscription.IsValid()},
		{"blocker group order", AccountDeletionBlockerActiveGroupOrder.IsValid()},
		{"blocker reservation", AccountDeletionBlockerActiveReservation.IsValid()},
		{"blocker persona", AccountDeletionBlockerRecoverablePersona.IsValid()},
		{"blocker legal hold", AccountDeletionBlockerLegalHold.IsValid()},
		{"blocker stored value", AccountDeletionBlockerProtectedStoredValue.IsValid()},
		{"blocker acknowledgement", AccountDeletionBlockerStoredValueAcknowledgement.IsValid()},
		{"blocker provider revocation", AccountDeletionBlockerProviderRevocationFailure.IsValid()},
		{"disposition erased", ErasureDispositionErased.IsValid()},
		{"disposition deidentified", ErasureDispositionDeidentified.IsValid()},
		{"disposition mixed", ErasureDispositionMixed.IsValid()},
		{"disposition no subject data", ErasureDispositionNoSubjectData.IsValid()},
		{"disposition retained", ErasureDispositionRetainedByPolicy.IsValid()},
		{"invalid operation", !AccountDeletionOperation("delete_now").IsValid()},
		{"invalid participant", !AccountDeletionParticipant("warehouse").IsValid()},
		{"invalid state", !AccountDeletionState("deleted").IsValid()},
		{"invalid receipt status", !CommandReceiptStatus("retry").IsValid()},
		{"invalid blocker", !AccountDeletionBlocker("provider_reference").IsValid()},
		{"invalid disposition", !ErasureDisposition("deleted").IsValid()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !test.valid {
				t.Fatal("enum validity result was false")
			}
		})
	}
}
