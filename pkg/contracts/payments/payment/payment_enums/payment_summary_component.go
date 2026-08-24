package payment_enums

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
