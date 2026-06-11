package payments

type PaymentReference struct {
	Stripe *StripePaymentReference `json:"stripe,omitempty"`
	Mx51   *Mx51PaymentReference   `json:"mx51,omitempty"`
}

type StripePaymentReference struct {
	PaymentIntentID   string `json:"payment_intent_id,omitempty"`
	ChargeID          string `json:"charge_id,omitempty"`
	CheckoutSessionID string `json:"checkout_session_id,omitempty"`
	CustomerID        string `json:"customer_id,omitempty"`
	RefundID          string `json:"refund_id,omitempty"`
}

type Mx51PaymentReference struct {
	TransactionID            string `json:"transaction_id,omitempty"`
	ReferenceID              string `json:"reference_id,omitempty"`
	RetrievalReferenceNumber string `json:"retrieval_reference_number,omitempty"`
	SystemTraceAuditNumber   string `json:"system_trace_audit_number,omitempty"`
	AuthorizationCode        string `json:"authorization_code,omitempty"`
	MerchantID               string `json:"merchant_id,omitempty"`
	TerminalID               string `json:"terminal_id,omitempty"`
	SettlementID             string `json:"settlement_id,omitempty"`
}
