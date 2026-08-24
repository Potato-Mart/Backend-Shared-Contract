package wallet_enums

// WalletInstrumentType identifies one customer-wallet holding family.
type WalletInstrumentType string

const (
	WalletInstrumentTypePoints   WalletInstrumentType = "points"
	WalletInstrumentTypeGiftCard WalletInstrumentType = "gift_card"
	WalletInstrumentTypeVoucher  WalletInstrumentType = "voucher"
	WalletInstrumentTypeCoupon   WalletInstrumentType = "coupon"
	WalletInstrumentTypeReward   WalletInstrumentType = "reward"
)

func (t WalletInstrumentType) IsValid() bool {
	switch t {
	case WalletInstrumentTypePoints, WalletInstrumentTypeGiftCard, WalletInstrumentTypeVoucher, WalletInstrumentTypeCoupon, WalletInstrumentTypeReward:
		return true
	}
	return false
}
func (t WalletInstrumentType) String() string { return string(t) }
