package order_enums

// GiftCardIssueStatus is the state of a gift card sold on an order: pending
// until Pricing issues it, issued once it has a code, failed when issuing it
// failed for good, reversal_pending while a refund is reversing it, and
// reversed once that refund has completed.
type GiftCardIssueStatus string

const (
	GiftCardIssueStatusPending         GiftCardIssueStatus = "pending"
	GiftCardIssueStatusIssued          GiftCardIssueStatus = "issued"
	GiftCardIssueStatusFailed          GiftCardIssueStatus = "failed"
	GiftCardIssueStatusReversalPending GiftCardIssueStatus = "reversal_pending"
	GiftCardIssueStatusReversed        GiftCardIssueStatus = "reversed"
)

// IsValid reports whether s is a known GiftCardIssueStatus.
func (s GiftCardIssueStatus) IsValid() bool {
	switch s {
	case GiftCardIssueStatusPending, GiftCardIssueStatusIssued, GiftCardIssueStatusFailed,
		GiftCardIssueStatusReversalPending, GiftCardIssueStatusReversed:
		return true
	}
	return false
}

func (s GiftCardIssueStatus) String() string { return string(s) }
