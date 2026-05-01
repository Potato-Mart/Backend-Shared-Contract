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
)

func (p PaymentRecordStatus) IsValid() bool {
	switch p {
	case PaymentRecordStatusPending, PaymentRecordStatusProcessing, PaymentRecordStatusCompleted,
		PaymentRecordStatusFailed, PaymentRecordStatusCancelled, PaymentRecordStatusRefunded:
		return true
	}
	return false
}

func (p PaymentRecordStatus) String() string { return string(p) }
