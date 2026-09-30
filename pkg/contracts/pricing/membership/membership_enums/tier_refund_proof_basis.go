package membership_enums

// TierRefundProofBasis identifies the protected original evidence underlying a
// Pricing net tier-refund reservation. It is not inferred from absent fields.
type TierRefundProofBasis string

const (
	TierRefundProofBasisQuotedCredit         TierRefundProofBasis = "quoted_credit"
	TierRefundProofBasisExcludedGiftPurchase TierRefundProofBasis = "excluded_gift_purchase"
)

func (b TierRefundProofBasis) IsValid() bool {
	switch b {
	case TierRefundProofBasisQuotedCredit, TierRefundProofBasisExcludedGiftPurchase:
		return true
	}
	return false
}
func (b TierRefundProofBasis) String() string { return string(b) }
