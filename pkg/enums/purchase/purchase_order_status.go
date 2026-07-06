package purchaseenum

// PurchaseOrderStatus is the lifecycle of an order. The enum is intentionally
// rich so that downstream analytics, fulfilment and accounting systems
// can share a single vocabulary without further migrations.
//
// Allowed transitions (enforced by the order service):
//		DRAFT    			-> SUBMITTED, CANCELLED
//	    SUBMITTED  			-> CONFIRMED, CANCELLED
//		CONFIRMED   		-> PARTIALLY_RECEIVED, RECEIVED, REFUNDED, CANCELLED
//		PARTIALLY_RECEIVED	-> RECEIVED, REFUNDED
//		RECEIVED  			-> REFUNDED
//		CANCELLED  			-> (terminal)
//		REFUNDED   			-> (terminal)

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

// IsValid reports whether s is a known PurchaseOrderStatus.
func (s PurchaseOrderStatus) IsValid() bool {
	switch s {
	case PurchaseOrderStatusDraft, PurchaseOrderStatusSubmitted, PurchaseOrderStatusConfirmed,
		PurchaseOrderStatusPartiallyReceived, PurchaseOrderStatusReceived,
		PurchaseOrderStatusCancelled, PurchaseOrderStatusRefunded:
		return true
	}
	return false
}

// IsTerminal reports whether an order in this status can still change.
func (s PurchaseOrderStatus) IsTerminal() bool {
	return s == PurchaseOrderStatusCancelled || s == PurchaseOrderStatusRefunded
}

// CanTransitionTo reports whether a natural (non-admin-override)
// transition from s to next is permitted.
func (s PurchaseOrderStatus) CanTransitionTo(next PurchaseOrderStatus) bool {
	for _, allowed := range allowedPurchasedTransitions[s] {
		if allowed == next {
			return true
		}
	}
	return false
}

var allowedPurchasedTransitions = map[PurchaseOrderStatus][]PurchaseOrderStatus{
	PurchaseOrderStatusDraft:             {PurchaseOrderStatusSubmitted, PurchaseOrderStatusCancelled},
	PurchaseOrderStatusSubmitted:         {PurchaseOrderStatusConfirmed, PurchaseOrderStatusCancelled},
	PurchaseOrderStatusConfirmed:         {PurchaseOrderStatusPartiallyReceived, PurchaseOrderStatusReceived, PurchaseOrderStatusRefunded, PurchaseOrderStatusCancelled},
	PurchaseOrderStatusReceived:          {PurchaseOrderStatusRefunded},
	PurchaseOrderStatusPartiallyReceived: {PurchaseOrderStatusReceived, PurchaseOrderStatusRefunded},
	PurchaseOrderStatusCancelled:         {},
	PurchaseOrderStatusRefunded:          {},
}

func (s PurchaseOrderStatus) String() string { return string(s) }
