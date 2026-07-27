package payments

type PaymentReference struct {
	Stripe *StripePaymentReference `json:"stripe,omitempty"`
	Mx51   *Mx51PaymentReference   `json:"mx51,omitempty"`
	Wallet *WalletPaymentReference `json:"wallet,omitempty"`
}

// WalletPaymentReference links a completed gift-card order payment to the
// authoritative wallet ledger transaction returned by Pricing.
type WalletPaymentReference struct {
	GiftCardCode        string `json:"gift_card_code"`
	WalletTransactionID string `json:"wallet_transaction_id"`
}

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
