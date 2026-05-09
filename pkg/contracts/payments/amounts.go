package payments

// Amounts is the per-component breakdown of a terminal transaction in
// minor units (cents). The same shape is used both for what the POS
// requested and for what the terminal actually applied.
//
// MX51 SCI is explicit on this point: result_amounts is the source of
// truth for tax invoices and refund calculations. In particular, when
// refunding a transaction that included a surcharge, you MUST refund
// against the applied SurchargeMinor - not the value originally sent
// in the request - because terminal-based surcharging may have changed
// it.
//
// Currency is carried once at this level rather than on each minor-unit
// field so the values stay simple integers.
type Amounts struct {
	Currency string `json:"currency"`

	PurchaseMinor  int64 `json:"purchase_minor,omitempty"`
	TipMinor       int64 `json:"tip_minor,omitempty"`
	SurchargeMinor int64 `json:"surcharge_minor,omitempty"`
	CashoutMinor   int64 `json:"cashout_minor,omitempty"`
	RefundMinor    int64 `json:"refund_minor,omitempty"`
	MOTOMinor      int64 `json:"moto_minor,omitempty"`
}

// TotalMinor returns the sum of all positive components (purchase +
// tip + surcharge + cashout + moto) minus refund. It is a convenience
// helper - always read individual components for receipts and
// reconciliation.
func (a Amounts) TotalMinor() int64 {
	return a.PurchaseMinor + a.TipMinor + a.SurchargeMinor + a.CashoutMinor + a.MOTOMinor - a.RefundMinor
}

// ReceiptOptions are the per-request flags that control whether the
// terminal or the POS owns receipt printing and signature
// verification. They map 1:1 to the MX51 SCI request fields with the
// same names.
//
// Recommended defaults (per MX51 docs):
//
//	PrintMerchantReceipt          : false
//	PromptCustomerReceipt         : false
//	VerifySignatureOnTerminal     : false
//	POSAutoPrintSignatureReceipt  : true
type ReceiptOptions struct {
	PrintMerchantReceipt         bool `json:"print_merchant_receipt,omitempty"`
	PromptCustomerReceipt        bool `json:"prompt_customer_receipt,omitempty"`
	VerifySignatureOnTerminal    bool `json:"verify_signature_on_terminal,omitempty"`
	POSAutoPrintSignatureReceipt bool `json:"pos_auto_print_signature_receipt,omitempty"`
}
