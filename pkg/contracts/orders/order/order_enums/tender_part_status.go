package order_enums

// TenderPartStatus is the state of one part of a POS split-tender plan:
// pending until the till starts collecting it, in_progress while it is being
// collected (for example on the card terminal), captured once its payment is
// recorded, declined when the tender was refused, cancelled when the part was
// abandoned before capture, and refunded once a refund has returned the money
// it collected.
type TenderPartStatus string

const (
	TenderPartStatusPending    TenderPartStatus = "pending"
	TenderPartStatusInProgress TenderPartStatus = "in_progress"
	TenderPartStatusCaptured   TenderPartStatus = "captured"
	TenderPartStatusDeclined   TenderPartStatus = "declined"
	TenderPartStatusCancelled  TenderPartStatus = "cancelled"
	TenderPartStatusRefunded   TenderPartStatus = "refunded"
)

// IsValid reports whether s is a known TenderPartStatus.
func (s TenderPartStatus) IsValid() bool {
	switch s {
	case TenderPartStatusPending, TenderPartStatusInProgress, TenderPartStatusCaptured,
		TenderPartStatusDeclined, TenderPartStatusCancelled, TenderPartStatusRefunded:
		return true
	}
	return false
}

func (s TenderPartStatus) String() string { return string(s) }
