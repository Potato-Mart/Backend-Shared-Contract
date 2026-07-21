package paymentenum

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

// PaymentCompleteness states whether a customer payment summary could be
// fully assembled from the underlying records.
type PaymentCompleteness string

const (
	PaymentCompletenessComplete    PaymentCompleteness = "complete"
	PaymentCompletenessPartial     PaymentCompleteness = "partial"
	PaymentCompletenessUnavailable PaymentCompleteness = "unavailable"
)

func (c PaymentCompleteness) IsValid() bool {
	switch c {
	case PaymentCompletenessComplete, PaymentCompletenessPartial, PaymentCompletenessUnavailable:
		return true
	default:
		return false
	}
}

func (c PaymentCompleteness) String() string { return string(c) }

// PaymentSummaryComponent names a summary component that could not be
// assembled when completeness is partial.
type PaymentSummaryComponent string

const (
	PaymentSummaryComponentAllocationHistory    PaymentSummaryComponent = "allocation_history"
	PaymentSummaryComponentPaymentTotals        PaymentSummaryComponent = "payment_totals"
	PaymentSummaryComponentRedemptionTimestamps PaymentSummaryComponent = "redemption_timestamps"
	PaymentSummaryComponentPointsEarned         PaymentSummaryComponent = "points_earned"
)

func (c PaymentSummaryComponent) IsValid() bool {
	switch c {
	case PaymentSummaryComponentAllocationHistory, PaymentSummaryComponentPaymentTotals,
		PaymentSummaryComponentRedemptionTimestamps, PaymentSummaryComponentPointsEarned:
		return true
	default:
		return false
	}
}

func (c PaymentSummaryComponent) String() string { return string(c) }

// InvoiceResendStatus tracks a customer-requested invoice email redelivery.
type InvoiceResendStatus string

const (
	InvoiceResendStatusQueued InvoiceResendStatus = "queued"
	InvoiceResendStatusSent   InvoiceResendStatus = "sent"
	InvoiceResendStatusFailed InvoiceResendStatus = "failed"
)

func (s InvoiceResendStatus) IsValid() bool {
	switch s {
	case InvoiceResendStatusQueued, InvoiceResendStatusSent, InvoiceResendStatusFailed:
		return true
	default:
		return false
	}
}

func (s InvoiceResendStatus) String() string { return string(s) }
