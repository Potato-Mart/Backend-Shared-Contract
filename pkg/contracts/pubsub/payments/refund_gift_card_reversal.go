package payments

// RefundGiftCardReversal is one sold gift card a completed refund reversed.
// LineID is the order's gift-card line (OrderGiftCardIssue.LineID) and
// IssuanceID is Pricing's issuance reference for the card. State is the
// Payments-owned outcome of the reversal for that card, as an open code (for
// example reversed); consumers must not reject a state they do not know.
type RefundGiftCardReversal struct {
	LineID     string `json:"line_id"`
	IssuanceID string `json:"issuance_id"`
	State      string `json:"state"`
}
