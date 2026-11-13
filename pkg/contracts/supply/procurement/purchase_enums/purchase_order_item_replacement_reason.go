package purchase_enums

// PurchaseOrderItemReplacementReason is why a purchase-order line's
// outstanding quantity was replaced by another line: the wrong product was
// ordered, the wrong product was received, the supplier substituted a
// product, or another reason given in the replacement note.
type PurchaseOrderItemReplacementReason string

const (
	PurchaseOrderItemReplacementReasonWrongProductOrdered  PurchaseOrderItemReplacementReason = "WRONG_PRODUCT_ORDERED"
	PurchaseOrderItemReplacementReasonWrongProductReceived PurchaseOrderItemReplacementReason = "WRONG_PRODUCT_RECEIVED"
	PurchaseOrderItemReplacementReasonSupplierSubstitution PurchaseOrderItemReplacementReason = "SUPPLIER_SUBSTITUTION"
	PurchaseOrderItemReplacementReasonOther                PurchaseOrderItemReplacementReason = "OTHER"
)

// IsValid reports whether r is a known PurchaseOrderItemReplacementReason.
func (r PurchaseOrderItemReplacementReason) IsValid() bool {
	switch r {
	case PurchaseOrderItemReplacementReasonWrongProductOrdered, PurchaseOrderItemReplacementReasonWrongProductReceived,
		PurchaseOrderItemReplacementReasonSupplierSubstitution, PurchaseOrderItemReplacementReasonOther:
		return true
	default:
		return false
	}
}

func (r PurchaseOrderItemReplacementReason) String() string { return string(r) }
