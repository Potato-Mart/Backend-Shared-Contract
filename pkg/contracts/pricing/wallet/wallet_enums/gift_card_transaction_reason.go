package wallet_enums

// GiftCardTransactionReason classifies a stored-value ledger entry.
type GiftCardTransactionReason string

const (
	GiftCardTransactionReasonIssue  GiftCardTransactionReason = "issue"
	GiftCardTransactionReasonRedeem GiftCardTransactionReason = "redeem"
	GiftCardTransactionReasonRefund GiftCardTransactionReason = "refund"
	GiftCardTransactionReasonTopUp  GiftCardTransactionReason = "top_up"
	GiftCardTransactionReasonExpire GiftCardTransactionReason = "expire"
	GiftCardTransactionReasonAdjust GiftCardTransactionReason = "adjust"
)

func (r GiftCardTransactionReason) IsValid() bool {
	switch r {
	case GiftCardTransactionReasonIssue, GiftCardTransactionReasonRedeem, GiftCardTransactionReasonRefund, GiftCardTransactionReasonTopUp, GiftCardTransactionReasonExpire, GiftCardTransactionReasonAdjust:
		return true
	}
	return false
}
func (r GiftCardTransactionReason) String() string { return string(r) }
