package possettlement

// AbortEvidence is Supply's durable tombstone excluding settlement of this
// exact obligation. It does not prove financial hold release, physical return
// of handed-over goods, or cancellation of a different sale or generation.
type AbortEvidence struct {
	Version               int32     `json:"version"`
	SaleID                string    `json:"sale_id"`
	OrderNumber           string    `json:"order_number"`
	FulfillmentGeneration int64     `json:"fulfillment_generation"`
	ObligationSHA256      string    `json:"obligation_sha256"`
	TenderSHA256          string    `json:"tender_sha256"`
	AbortIntentSHA256     string    `json:"abort_intent_sha256"`
	RequestedAt           time.Time `json:"requested_at"`
	RecordedAt            time.Time `json:"recorded_at"`
}
