package provider

type StripePaymentReference struct {
	PaymentIntentID   string `json:"payment_intent_id,omitempty"`
	ChargeID          string `json:"charge_id,omitempty"`
	CheckoutSessionID string `json:"checkout_session_id,omitempty"`
	CustomerID        string `json:"customer_id,omitempty"`
	RefundID          string `json:"refund_id,omitempty"`
	// BalanceTransactionID is the Stripe balance transaction that carries
	// the processor fee and net amount for the charge.
	BalanceTransactionID string `json:"balance_transaction_id,omitempty"`
}
