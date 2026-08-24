package order_enums

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
