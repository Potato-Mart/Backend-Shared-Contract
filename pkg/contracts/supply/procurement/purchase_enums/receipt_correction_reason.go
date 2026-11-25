package purchase_enums

// ReceiptCorrectionReason is why a confirmed receipt line's quantity was
// corrected after receiving: the goods were damaged, the delivery was short,
// the line was miscounted, or another reason given in the correction note.
type ReceiptCorrectionReason string

const (
	ReceiptCorrectionReasonDamaged  ReceiptCorrectionReason = "DAMAGED"
	ReceiptCorrectionReasonShort    ReceiptCorrectionReason = "SHORT"
	ReceiptCorrectionReasonMiscount ReceiptCorrectionReason = "MISCOUNT"
	ReceiptCorrectionReasonOther    ReceiptCorrectionReason = "OTHER"
)

// IsValid reports whether r is a known ReceiptCorrectionReason.
func (r ReceiptCorrectionReason) IsValid() bool {
	switch r {
	case ReceiptCorrectionReasonDamaged, ReceiptCorrectionReasonShort,
		ReceiptCorrectionReasonMiscount, ReceiptCorrectionReasonOther:
		return true
	default:
		return false
	}
}

func (r ReceiptCorrectionReason) String() string { return string(r) }
