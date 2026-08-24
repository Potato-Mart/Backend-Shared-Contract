package payment_enums

// CustomerPaymentAllocationKind classifies one row of how an order was paid,
// customer-safe (no processor internals).
type CustomerPaymentAllocationKind string

const (
	CustomerPaymentAllocationKindExternalPayment CustomerPaymentAllocationKind = "external_payment"
	CustomerPaymentAllocationKindWalletBalance   CustomerPaymentAllocationKind = "wallet_balance"
	CustomerPaymentAllocationKindWalletPoints    CustomerPaymentAllocationKind = "wallet_points"
	CustomerPaymentAllocationKindGiftCard        CustomerPaymentAllocationKind = "gift_card"
	CustomerPaymentAllocationKindVoucher         CustomerPaymentAllocationKind = "voucher"
	CustomerPaymentAllocationKindOther           CustomerPaymentAllocationKind = "other"
)

func (k CustomerPaymentAllocationKind) IsValid() bool {
	switch k {
	case CustomerPaymentAllocationKindExternalPayment, CustomerPaymentAllocationKindWalletBalance,
		CustomerPaymentAllocationKindWalletPoints, CustomerPaymentAllocationKindGiftCard,
		CustomerPaymentAllocationKindVoucher, CustomerPaymentAllocationKindOther:
		return true
	default:
		return false
	}
}

func (k CustomerPaymentAllocationKind) String() string { return string(k) }
