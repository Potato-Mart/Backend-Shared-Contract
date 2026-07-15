package walletenum

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

// VoucherStatus is the lifecycle state of a single-use voucher. Reserved is a
// temporary checkout hold and returns to issued when cancelled or expired.
type VoucherStatus string

const (
	VoucherStatusIssued   VoucherStatus = "issued"
	VoucherStatusReserved VoucherStatus = "reserved"
	VoucherStatusRedeemed VoucherStatus = "redeemed"
	VoucherStatusExpired  VoucherStatus = "expired"
	VoucherStatusVoid     VoucherStatus = "void"
)

// IsValid reports whether s is a known VoucherStatus value.
func (s VoucherStatus) IsValid() bool {
	switch s {
	case VoucherStatusIssued, VoucherStatusReserved, VoucherStatusRedeemed,
		VoucherStatusExpired, VoucherStatusVoid:
		return true
	}
	return false
}

// String returns the wire value for s.
func (s VoucherStatus) String() string { return string(s) }

// CheckoutBenefitReservationStatus is the durable lifecycle shared by a
// checkout's coupon, voucher, and gift-card allocations.
type CheckoutBenefitReservationStatus string

const (
	CheckoutBenefitReservationStatusReserved          CheckoutBenefitReservationStatus = "reserved"
	CheckoutBenefitReservationStatusCommitted         CheckoutBenefitReservationStatus = "committed"
	CheckoutBenefitReservationStatusCancelled         CheckoutBenefitReservationStatus = "cancelled"
	CheckoutBenefitReservationStatusExpired           CheckoutBenefitReservationStatus = "expired"
	CheckoutBenefitReservationStatusPartiallyRefunded CheckoutBenefitReservationStatus = "partially_refunded"
	CheckoutBenefitReservationStatusRefunded          CheckoutBenefitReservationStatus = "refunded"
)

// IsValid reports whether s is a known CheckoutBenefitReservationStatus value.
func (s CheckoutBenefitReservationStatus) IsValid() bool {
	switch s {
	case CheckoutBenefitReservationStatusReserved,
		CheckoutBenefitReservationStatusCommitted,
		CheckoutBenefitReservationStatusCancelled,
		CheckoutBenefitReservationStatusExpired,
		CheckoutBenefitReservationStatusPartiallyRefunded,
		CheckoutBenefitReservationStatusRefunded:
		return true
	}
	return false
}

// String returns the wire value for s.
func (s CheckoutBenefitReservationStatus) String() string { return string(s) }
