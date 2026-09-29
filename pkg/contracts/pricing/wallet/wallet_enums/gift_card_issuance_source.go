package wallet_enums

// GiftCardIssuanceSource classifies the committed business fact that created
// stored value. It does not assert a purchase, payment or denomination policy
// for membership rewards or refund replacements.
type GiftCardIssuanceSource string

const (
	GiftCardIssuanceSourcePurchase          GiftCardIssuanceSource = "purchase"
	GiftCardIssuanceSourceMembershipReward  GiftCardIssuanceSource = "membership_reward"
	GiftCardIssuanceSourceRefundReplacement GiftCardIssuanceSource = "refund_replacement"
)

func (s GiftCardIssuanceSource) IsValid() bool {
	switch s {
	case GiftCardIssuanceSourcePurchase, GiftCardIssuanceSourceMembershipReward, GiftCardIssuanceSourceRefundReplacement:
		return true
	default:
		return false
	}
}
func (s GiftCardIssuanceSource) String() string { return string(s) }
