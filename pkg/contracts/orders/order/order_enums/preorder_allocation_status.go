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

// FulfillmentReadiness gates warehouse intake independently from payment and
// order lifecycle status.
type FulfillmentReadiness string

const (
	FulfillmentReadinessReady                   FulfillmentReadiness = "ready"
	FulfillmentReadinessWaitingForPreorderStock FulfillmentReadiness = "waiting_for_preorder_stock"
)

func (s FulfillmentReadiness) IsValid() bool {
	switch s {
	case FulfillmentReadinessReady, FulfillmentReadinessWaitingForPreorderStock:
		return true
	default:
		return false
	}
}

func (s FulfillmentReadiness) String() string { return string(s) }
