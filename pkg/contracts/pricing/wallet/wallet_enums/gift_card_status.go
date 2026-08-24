package wallet_enums

// GiftCardStatus is the lifecycle state of stored value.
type GiftCardStatus string

const (
	GiftCardStatusActive            GiftCardStatus = "active"
	GiftCardStatusPartiallyRedeemed GiftCardStatus = "partially_redeemed"
	GiftCardStatusDepleted          GiftCardStatus = "depleted"
	GiftCardStatusExpired           GiftCardStatus = "expired"
	GiftCardStatusVoid              GiftCardStatus = "void"
)

func (s GiftCardStatus) IsValid() bool {
	switch s {
	case GiftCardStatusActive, GiftCardStatusPartiallyRedeemed, GiftCardStatusDepleted, GiftCardStatusExpired, GiftCardStatusVoid:
		return true
	}
	return false
}
func (s GiftCardStatus) String() string { return string(s) }
