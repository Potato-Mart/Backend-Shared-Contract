package order_enums

// PreorderAllocationStatus is the stock-allocation state of an order-owned
// preorder line. A preorder is never an independent sales aggregate.
type PreorderAllocationStatus string

const (
	PreorderAllocationStatusWaitingForStock    PreorderAllocationStatus = "waiting_for_stock"
	PreorderAllocationStatusPartiallyAllocated PreorderAllocationStatus = "partially_allocated"
	PreorderAllocationStatusStockAllocated     PreorderAllocationStatus = "stock_allocated"
)

func (s PreorderAllocationStatus) IsValid() bool {
	switch s {
	case PreorderAllocationStatusWaitingForStock,
		PreorderAllocationStatusPartiallyAllocated,
		PreorderAllocationStatusStockAllocated:
		return true
	default:
		return false
	}
}

func (s PreorderAllocationStatus) String() string { return string(s) }
