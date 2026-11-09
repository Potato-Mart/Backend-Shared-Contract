package purchase_enums

// PurchaseReceiptStatus is the lifecycle of a supplier procurement receipt.
// A receipt may remain DRAFT without a warehouse link; CONFIRMED is permitted
// only after the owning service has established that link. REVERSED is set
// only by the reversal command on a CONFIRMED receipt.
type PurchaseReceiptStatus string

const (
	PurchaseReceiptStatusDraft     PurchaseReceiptStatus = "DRAFT"
	PurchaseReceiptStatusConfirmed PurchaseReceiptStatus = "CONFIRMED"
	PurchaseReceiptStatusReversed  PurchaseReceiptStatus = "REVERSED"
)

// IsValid reports whether s is a known PurchaseReceiptStatus.
func (s PurchaseReceiptStatus) IsValid() bool {
	return s == PurchaseReceiptStatusDraft || s == PurchaseReceiptStatusConfirmed || s == PurchaseReceiptStatusReversed
}

func (s PurchaseReceiptStatus) String() string { return string(s) }
