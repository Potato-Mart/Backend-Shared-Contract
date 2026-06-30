package enums

// WalletInstrumentType identifies the kind of value instrument linked in a
// customer wallet.
type WalletInstrumentType string

const (
	WalletInstrumentTypePoints   WalletInstrumentType = "points"
	WalletInstrumentTypeGiftCard WalletInstrumentType = "gift_card"
	WalletInstrumentTypeVoucher  WalletInstrumentType = "voucher"
	WalletInstrumentTypeCoupon   WalletInstrumentType = "coupon"
	WalletInstrumentTypeReward   WalletInstrumentType = "reward"
)

// IsValid reports whether t is a known WalletInstrumentType value.
func (t WalletInstrumentType) IsValid() bool {
	switch t {
	case WalletInstrumentTypePoints, WalletInstrumentTypeGiftCard,
		WalletInstrumentTypeVoucher, WalletInstrumentTypeCoupon,
		WalletInstrumentTypeReward:
		return true
	}
	return false
}

// String returns the wire value for t.
func (t WalletInstrumentType) String() string { return string(t) }

// GiftCardStatus is the lifecycle state of a stored-value gift card.
type GiftCardStatus string

const (
	GiftCardStatusActive            GiftCardStatus = "active"
	GiftCardStatusPartiallyRedeemed GiftCardStatus = "partially_redeemed"
	GiftCardStatusDepleted          GiftCardStatus = "depleted"
	GiftCardStatusExpired           GiftCardStatus = "expired"
	GiftCardStatusVoid              GiftCardStatus = "void"
)

// IsValid reports whether s is a known GiftCardStatus value.
func (s GiftCardStatus) IsValid() bool {
	switch s {
	case GiftCardStatusActive, GiftCardStatusPartiallyRedeemed,
		GiftCardStatusDepleted, GiftCardStatusExpired, GiftCardStatusVoid:
		return true
	}
	return false
}

// String returns the wire value for s.
func (s GiftCardStatus) String() string { return string(s) }

// GiftCardTransactionReason classifies a gift-card balance ledger entry.
type GiftCardTransactionReason string

const (
	GiftCardTransactionReasonIssue  GiftCardTransactionReason = "issue"
	GiftCardTransactionReasonRedeem GiftCardTransactionReason = "redeem"
	GiftCardTransactionReasonRefund GiftCardTransactionReason = "refund"
	GiftCardTransactionReasonTopUp  GiftCardTransactionReason = "top_up"
	GiftCardTransactionReasonExpire GiftCardTransactionReason = "expire"
	GiftCardTransactionReasonAdjust GiftCardTransactionReason = "adjust"
)

// IsValid reports whether r is a known GiftCardTransactionReason value.
func (r GiftCardTransactionReason) IsValid() bool {
	switch r {
	case GiftCardTransactionReasonIssue, GiftCardTransactionReasonRedeem,
		GiftCardTransactionReasonRefund, GiftCardTransactionReasonTopUp,
		GiftCardTransactionReasonExpire, GiftCardTransactionReasonAdjust:
		return true
	}
	return false
}

// String returns the wire value for r.
func (r GiftCardTransactionReason) String() string { return string(r) }
