package wallet_enums

// GiftCardIssuanceSource classifies the committed business fact that created
// stored value. It does not assert a purchase, payment or denomination policy
// for membership rewards, refund replacements or legacy migrations.
// LegacyMigration is stored value carried over from the retired point-of-sale
// system: the balance was bought there, so it is neither a purchase here nor
// value the business gave away.
type GiftCardIssuanceSource string

const (
	GiftCardIssuanceSourcePurchase          GiftCardIssuanceSource = "purchase"
	GiftCardIssuanceSourceMembershipReward  GiftCardIssuanceSource = "membership_reward"
	GiftCardIssuanceSourceRefundReplacement GiftCardIssuanceSource = "refund_replacement"
	GiftCardIssuanceSourceLegacyMigration   GiftCardIssuanceSource = "legacy_migration"
)

func (s GiftCardIssuanceSource) IsValid() bool {
	switch s {
	case GiftCardIssuanceSourcePurchase, GiftCardIssuanceSourceMembershipReward, GiftCardIssuanceSourceRefundReplacement, GiftCardIssuanceSourceLegacyMigration:
		return true
	default:
		return false
	}
}
func (s GiftCardIssuanceSource) String() string { return string(s) }
