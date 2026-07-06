package salesenum

// SalesOrderStatus is the lifecycle of an order. The enum is intentionally
// rich so that downstream analytics, fulfilment and accounting systems
// can share a single vocabulary without further migrations.
//
// Allowed transitions (enforced by the order service):
//
//		PENDING    	->	CONFIRMED, CANCELLED
//	    CONFIRMED  	->	PAID, CANCELLED
//		PAID       	->	PROCESSING, REFUNDED, CANCELLED
//		PROCESSING 	->	PICKING, REFUNDED (一旦我們開始拿貨品絕對不能更改，如果整單取消有歸位費用{比例扣費})
// 		PICKING		->	PACKED, REFUNDED (單據已列印)
//		PACKED		->	SHIPPED, REFUNDED
//		SHIPPED    	->	COMPLETED, DELIVERED, REFUNDED
//		DELIVERED  	->	COMPLETED, REFUNDED
//		COMPLETED  	->	REFUNDED
//		CANCELLED  	->	(terminal)
//		REFUNDED   	->	(terminal)

type SalesOrderStatus string

const (
	SalesOrderStatusPending    SalesOrderStatus = "pending"
	SalesOrderStatusConfirmed  SalesOrderStatus = "confirmed"
	SalesOrderStatusPaid       SalesOrderStatus = "paid"
	SalesOrderStatusProcessing SalesOrderStatus = "processing"
	SalesOrderStatusPicking    SalesOrderStatus = "picking"
	SalesOrderStatusPacked     SalesOrderStatus = "packed"
	SalesOrderStatusShipped    SalesOrderStatus = "shipped"
	SalesOrderStatusDelivered  SalesOrderStatus = "delivered"
	SalesOrderStatusCompleted  SalesOrderStatus = "completed"
	SalesOrderStatusCancelled  SalesOrderStatus = "cancelled"
	SalesOrderStatusRefunded   SalesOrderStatus = "refunded"
)

// IsValid reports whether s is a known SalesOrderStatus.
func (s SalesOrderStatus) IsValid() bool {
	switch s {
	case SalesOrderStatusPending, SalesOrderStatusConfirmed, SalesOrderStatusPaid, SalesOrderStatusProcessing,
		SalesOrderStatusPicking, SalesOrderStatusPacked, SalesOrderStatusShipped, SalesOrderStatusDelivered,
		SalesOrderStatusCompleted, SalesOrderStatusCancelled, SalesOrderStatusRefunded:
		return true
	}
	return false
}

// IsTerminal reports whether an order in this status can still change.
func (s SalesOrderStatus) IsTerminal() bool {
	return s == SalesOrderStatusCancelled || s == SalesOrderStatusRefunded
}

// CanTransitionTo reports whether a natural (non-admin-override)
// transition from s to next is permitted.
func (s SalesOrderStatus) CanTransitionTo(next SalesOrderStatus) bool {
	for _, allowed := range allowedSalesTransitions[s] {
		if allowed == next {
			return true
		}
	}
	return false
}

var allowedSalesTransitions = map[SalesOrderStatus][]SalesOrderStatus{
	SalesOrderStatusPending:    {SalesOrderStatusConfirmed, SalesOrderStatusCancelled},
	SalesOrderStatusConfirmed:  {SalesOrderStatusPaid, SalesOrderStatusCancelled},
	SalesOrderStatusPaid:       {SalesOrderStatusProcessing, SalesOrderStatusRefunded, SalesOrderStatusCancelled},
	SalesOrderStatusProcessing: {SalesOrderStatusPicking, SalesOrderStatusRefunded},
	SalesOrderStatusPicking:    {SalesOrderStatusPacked, SalesOrderStatusRefunded},
	SalesOrderStatusPacked:     {SalesOrderStatusShipped, SalesOrderStatusRefunded},
	SalesOrderStatusShipped:    {SalesOrderStatusDelivered, SalesOrderStatusCompleted, SalesOrderStatusRefunded},
	SalesOrderStatusDelivered:  {SalesOrderStatusCompleted, SalesOrderStatusRefunded},
	SalesOrderStatusCompleted:  {SalesOrderStatusRefunded},
	SalesOrderStatusCancelled:  {},
	SalesOrderStatusRefunded:   {},
}

func (s SalesOrderStatus) String() string { return string(s) }
