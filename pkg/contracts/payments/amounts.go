package payments

// Amounts is the per-component breakdown of a terminal transaction in
// minor units (cents for AUD). The same shape is used both for what the
// POS requested and for what the terminal/provider actually applied.
//
// The shared contract stays in minor units so services do not lose
// precision at boundaries.
//
// Result amounts are the source of truth for tax invoices, receipts, and
// refund calculations. In particular, when refunding a transaction that
// included a tip, surcharge, or cashout, use the applied result fields
// instead of the originally requested values.
type Amounts struct {
	Currency string `json:"currency"`

	PurchaseMinor   int64 `json:"purchase_minor,omitempty"`
	TipMinor        int64 `json:"tip_minor,omitempty"`
	SurchargeMinor  int64 `json:"surcharge_minor,omitempty"`
	CashoutMinor    int64 `json:"cashout_minor,omitempty"`
	RefundMinor     int64 `json:"refund_minor,omitempty"`
	MOTOMinor       int64 `json:"moto_minor,omitempty"`
	AuthorizedMinor int64 `json:"authorized_minor,omitempty"`
}

// ReceiptOptions are the per-request flags that control whether the
// terminal or the POS owns receipt printing and signature verification.
// Provider services map these flags to the relevant Terminal API fields
// or ignore unsupported options.
type ReceiptOptions struct {
	PrintMerchantReceipt         bool `json:"print_merchant_receipt,omitempty"`
	PromptCustomerReceipt        bool `json:"prompt_customer_receipt,omitempty"`
	VerifySignatureOnTerminal    bool `json:"verify_signature_on_terminal,omitempty"`
	POSAutoPrintSignatureReceipt bool `json:"pos_auto_print_signature_receipt,omitempty"`
}
