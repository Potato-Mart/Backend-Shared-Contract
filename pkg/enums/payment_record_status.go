package enums

// PaymentRecordStatus is the processing state of an individual payment
// transaction record. Distinct from PaymentStatus which tracks the
// aggregate payment state of an order.
type PaymentRecordStatus string

const (
	PaymentRecordStatusPending    PaymentRecordStatus = "pending"
	PaymentRecordStatusProcessing PaymentRecordStatus = "processing"
	PaymentRecordStatusCompleted  PaymentRecordStatus = "completed"
	PaymentRecordStatusFailed     PaymentRecordStatus = "failed"
	PaymentRecordStatusCancelled  PaymentRecordStatus = "cancelled"
	PaymentRecordStatusRefunded   PaymentRecordStatus = "refunded"

	// AwaitingAction is the state where the terminal has paused and is
	// waiting for the POS or merchant to take an action - signature
	// approval, MOTO card entry, or the override decision after a
	// timeout. Maps to MX51 SCI's AWAITING_POS.
	PaymentRecordStatusAwaitingAction PaymentRecordStatus = "awaiting_action"

	// Unknown is used when the terminal could not confirm the outcome
	// (network drop, terminal offline) and the recovery / override flow
	// has not yet been resolved by the merchant. Maps to MX51 SCI's
	// result_financial_status = UNKNOWN. Required for SCI certification.
	PaymentRecordStatusUnknown PaymentRecordStatus = "unknown"
)

func (p PaymentRecordStatus) IsValid() bool {
	switch p {
	case PaymentRecordStatusPending, PaymentRecordStatusProcessing, PaymentRecordStatusCompleted,
		PaymentRecordStatusFailed, PaymentRecordStatusCancelled, PaymentRecordStatusRefunded,
		PaymentRecordStatusAwaitingAction, PaymentRecordStatusUnknown:
		return true
	}
	return false
}

func (p PaymentRecordStatus) String() string { return string(p) }
