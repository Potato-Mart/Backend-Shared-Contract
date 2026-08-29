package purchase_enums

// PurchaseOrderStatus is the persisted lifecycle state of a purchase order.
type PurchaseOrderStatus string

const (
	PurchaseOrderStatusDraft             PurchaseOrderStatus = "DRAFT"
	PurchaseOrderStatusSubmitted         PurchaseOrderStatus = "SUBMITTED"
	PurchaseOrderStatusConfirmed         PurchaseOrderStatus = "CONFIRMED"
	PurchaseOrderStatusPartiallyReceived PurchaseOrderStatus = "PARTIALLY_RECEIVED"
	PurchaseOrderStatusReceived          PurchaseOrderStatus = "RECEIVED"
	PurchaseOrderStatusCancelled         PurchaseOrderStatus = "CANCELLED"
	PurchaseOrderStatusRefunded          PurchaseOrderStatus = "REFUNDED"
)

func (s PurchaseOrderStatus) IsValid() bool {
	switch s {
	case PurchaseOrderStatusDraft, PurchaseOrderStatusSubmitted, PurchaseOrderStatusConfirmed,
		PurchaseOrderStatusPartiallyReceived, PurchaseOrderStatusReceived,
		PurchaseOrderStatusCancelled, PurchaseOrderStatusRefunded:
		return true
	default:
		return false
	}
}

func (s PurchaseOrderStatus) String() string { return string(s) }
