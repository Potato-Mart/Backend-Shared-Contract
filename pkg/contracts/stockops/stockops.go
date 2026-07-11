// Package stockops defines the stable internal stock operation paths exposed
// by Backend-Operations.
package stockops

// Full request paths for the internal stock operations, as mounted by
// Backend-Operations (provider) under its service-authenticated
// /v1/internal group. Consumers (Commerce) POST to these paths with a
// service token carrying the scope noted on each constant. See ADR 0001.
const (
	// PathReserve reserves stock under a reference (provider: Operations;
	// requires scope serviceauth.ScopeStockReserve, "stock:reserve").
	PathReserve = "/v1/internal/stock/reservations"
	// PathCommit commits every active reservation under a reference
	// (provider: Operations; requires scope serviceauth.ScopeStockCommit,
	// "stock:commit").
	PathCommit = "/v1/internal/stock/commit"
	// PathRelease releases every active reservation under a reference
	// (provider: Operations; requires scope serviceauth.ScopeStockRelease,
	// "stock:release").
	PathRelease = "/v1/internal/stock/release"
	// PathPackingSettlement commits sold/damaged quantities and releases
	// short-shipped quantities for an order in one idempotent operation.
	PathPackingSettlement = "/v1/internal/stock/packing-settlement"
)

// ReservationLine is one SKU quantity in an atomic reservation command.
type ReservationLine struct {
	ProductSKUCode string `json:"product_sku_code"`
	Qty            int    `json:"qty"`
}

// ReservationCommand atomically reserves all lines under a reference. A
// unique idempotency key may allocate only these lines; it must not replace or
// release reservations created by other commands for the same order.
type ReservationCommand struct {
	DepotCode      string            `json:"depot_code,omitempty"`
	RefType        string            `json:"ref_type"`
	OrderNumber    string            `json:"order_number,omitempty"`
	RefNumber      string            `json:"ref_number,omitempty"`
	Lines          []ReservationLine `json:"lines"`
	IdempotencyKey string            `json:"idempotency_key"`
}

// ReservationAllocation identifies one persisted allocation made by a
// reservation command.
type ReservationAllocation struct {
	ReservationID  string `json:"reservation_id"`
	ProductSKUCode string `json:"product_sku_code"`
	Qty            int    `json:"qty"`
}

// ReservationResult is stable across an idempotent replay.
type ReservationResult struct {
	ReservationIDs []string                `json:"reservation_ids"`
	Allocations    []ReservationAllocation `json:"allocations,omitempty"`
	Replayed       bool                    `json:"replayed,omitempty"`
}

// PackingSettlementLine describes packing-time disposition of the order's
// reserved quantity for one SKU.
type PackingSettlementLine struct {
	ProductSKUCode string `json:"product_sku_code"`
	SaleQty        int    `json:"sale_qty,omitempty"`
	DamageQty      int    `json:"damage_qty,omitempty"`
	ReleaseQty     int    `json:"release_qty,omitempty"`
	DamageReportID string `json:"damage_report_id,omitempty"`
	DepotCode      string `json:"depot_code,omitempty"`
	LocationCode   string `json:"location_code,omitempty"`
	LotID          string `json:"lot_id,omitempty"`
}

// PackingSettlementCommand settles every supplied line atomically for one
// order and is replay-safe by idempotency key plus request fingerprint.
type PackingSettlementCommand struct {
	OrderNumber    string                  `json:"order_number"`
	Lines          []PackingSettlementLine `json:"lines"`
	IdempotencyKey string                  `json:"idempotency_key"`
}

// PackingSettlementMovement is a customer-service-safe summary of one stock
// movement created by settlement.
type PackingSettlementMovement struct {
	MovementID     string `json:"movement_id"`
	ProductSKUCode string `json:"product_sku_code"`
	Action         string `json:"action"`
	Qty            int    `json:"qty"`
}

// PackingSettlementResult is stable across an idempotent replay.
type PackingSettlementResult struct {
	OrderNumber string                      `json:"order_number"`
	Movements   []PackingSettlementMovement `json:"movements,omitempty"`
	Replayed    bool                        `json:"replayed,omitempty"`
}
