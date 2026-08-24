package wallet

// GiftCardDenominationBonus records bonus value for one purchase denomination.
type GiftCardDenominationBonus struct {
	AmountMinor int64 `json:"amount_minor"`
	BonusMinor  int64 `json:"bonus_minor"`
}
