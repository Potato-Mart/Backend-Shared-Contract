package salesenum

// SalesOrderStatus is the persisted lifecycle state of a sales order.
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

func (s SalesOrderStatus) IsValid() bool {
	switch s {
	case SalesOrderStatusPending, SalesOrderStatusConfirmed, SalesOrderStatusPaid,
		SalesOrderStatusProcessing, SalesOrderStatusPicking, SalesOrderStatusPacked,
		SalesOrderStatusShipped, SalesOrderStatusDelivered, SalesOrderStatusCompleted,
		SalesOrderStatusCancelled, SalesOrderStatusRefunded:
		return true
	default:
		return false
	}
}

func (s SalesOrderStatus) String() string { return string(s) }
