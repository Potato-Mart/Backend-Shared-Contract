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
)
