package possettlement

import "time"

// TenderHoldEvidence is a positive owner observation of a durable hold. Lease
// expiry is not a release. Every hold is bound to one obligation and tender
// digest, and Pricing's acknowledged barrier must equal Payments' frozen one.
type TenderHoldEvidence struct {
	Version               int32         `json:"version"`
	Authority             string        `json:"authority"`
	OrderNumber           string        `json:"order_number"`
	SaleID                string        `json:"sale_id"`
	FulfillmentGeneration int64         `json:"fulfillment_generation"`
	ObligationSHA256      string        `json:"obligation_sha256"`
	TenderSHA256          string        `json:"tender_sha256"`
	HoldID                string        `json:"hold_id"`
	OwnerRevision         int64         `json:"owner_revision"`
	RefundBarrier         RefundBarrier `json:"refund_barrier"`
	RecordedAt            time.Time     `json:"recorded_at"`
}
