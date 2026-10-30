package possettlement

import "time"

// SettlementReceipt binds Supply's complete atomic outcome to the immutable
// obligation. ReceiptSHA256 identifies the full owner-held reservation, direct
// movement, deficit and valuation evidence, not merely these locator fields.
// It is accepted only after authenticated owner lookup and reconciliation.
type SettlementReceipt struct {
	Version               int32     `json:"version"`
	SaleID                string    `json:"sale_id"`
	OrderNumber           string    `json:"order_number"`
	FulfillmentGeneration int64     `json:"fulfillment_generation"`
	ObligationSHA256      string    `json:"obligation_sha256"`
	TenderSHA256          string    `json:"tender_sha256"`
	SettlementID          string    `json:"settlement_id"`
	ReceiptSHA256         string    `json:"receipt_sha256"`
	ReservationIDs        []string  `json:"reservation_ids"`
	MovementIDs           []string  `json:"movement_ids"`
	DeficitMovementIDs    []string  `json:"deficit_movement_ids"`
	RecordedAt            time.Time `json:"recorded_at"`
}
