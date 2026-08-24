package payment_enums

// DocumentKind names the customer document a transaction produced. A tax
// invoice is rendered only when the document qualifies; a receipt is always
// issued for a completed sale.
type DocumentKind string

const (
	DocumentKindReceipt        DocumentKind = "receipt"
	DocumentKindTaxInvoice     DocumentKind = "tax_invoice"
	DocumentKindInvoice        DocumentKind = "invoice"
	DocumentKindAdjustmentNote DocumentKind = "adjustment_note"
)

// IsValid reports whether k is a known DocumentKind.
func (k DocumentKind) IsValid() bool {
	switch k {
	case DocumentKindReceipt, DocumentKindTaxInvoice, DocumentKindInvoice, DocumentKindAdjustmentNote:
		return true
	}
	return false
}

func (k DocumentKind) String() string { return string(k) }
