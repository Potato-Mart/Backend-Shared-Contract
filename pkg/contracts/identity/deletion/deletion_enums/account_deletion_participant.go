package deletion_enums

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
