// Package stockops defines the wire DTOs for the internal stock operations
// that Backend-Operations exposes to other services (currently Commerce, at
// checkout) over the service-to-service auth flow. See ADR 0001.
//
// These are intentionally minimal and transport-only: the rich domain types
// (lots, items, movements) stay private to Operations.
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
)

// ReserveLine is one product line to reserve.
type ReserveLine struct {
	ProductID string `json:"product_id" binding:"required"`
	Qty       int    `json:"qty" binding:"required,gt=0"`
}

// ReserveRequest reserves stock for every line under a single reference
// (e.g. RefType="order", RefID=<order id>). DepotID may be empty to allocate
// across all depots (earliest-expiry-first). The reservation is all-or-nothing
// from the caller's perspective: a partial failure releases the whole ref.
type ReserveRequest struct {
	DepotID string        `json:"depot_id,omitempty"`
	RefType string        `json:"ref_type" binding:"required"`
	RefID   string        `json:"ref_id" binding:"required"`
	Lines   []ReserveLine `json:"lines" binding:"required"`
}

// ReserveResponse returns the created reservation IDs.
type ReserveResponse struct {
	ReservationIDs []string `json:"reservation_ids"`
}

// RefRequest commits or releases every active reservation under a reference.
type RefRequest struct {
	RefType string `json:"ref_type" binding:"required"`
	RefID   string `json:"ref_id" binding:"required"`
}
