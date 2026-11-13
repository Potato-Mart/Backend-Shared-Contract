package purchase_enums

// PurchaseReceiptStatus is the lifecycle of a supplier procurement receipt.
// A receipt may remain DRAFT without a warehouse link; CONFIRMED is permitted
// only after the owning service has established that link.
//
// PARTIALLY_REVERSED is set only by the reversal command on a CONFIRMED
// receipt when some, but not all, of its received units have been reversed;
// the per-line detail is on PurchaseReceiptItem.Reversals. REVERSED is set
// only by the reversal command, on a CONFIRMED or PARTIALLY_REVERSED receipt,
// once every received unit has been reversed.
type PurchaseReceiptStatus string

const (
	PurchaseReceiptStatusDraft     PurchaseReceiptStatus = "DRAFT"
	PurchaseReceiptStatusConfirmed PurchaseReceiptStatus = "CONFIRMED"
	PurchaseReceiptStatusReversed  PurchaseReceiptStatus = "REVERSED"
	// PurchaseReceiptStatusPartiallyReversed marks a confirmed receipt with
	// some, but not all, received units reversed.
	PurchaseReceiptStatusPartiallyReversed PurchaseReceiptStatus = "PARTIALLY_REVERSED"
)

// IsValid reports whether s is a known PurchaseReceiptStatus.
func (s PurchaseReceiptStatus) IsValid() bool {
	switch s {
	case PurchaseReceiptStatusDraft, PurchaseReceiptStatusConfirmed, PurchaseReceiptStatusReversed,
		PurchaseReceiptStatusPartiallyReversed:
		return true
	default:
		return false
	}
}

func (s PurchaseReceiptStatus) String() string { return string(s) }
