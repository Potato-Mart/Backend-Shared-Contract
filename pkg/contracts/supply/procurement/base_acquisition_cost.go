// Package procurement holds the Supply-owned private operational cost models. Base
// acquisition cost and carrying cost are internal values: they never appear in
// storefront, customer order, customer invoice, or customer receipt responses.
package procurement

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/common/money"
)

// BaseAcquisitionCost is the current tax-exclusive acquisition cost for one
// SKU in one currency. It is overwritten in place under an optimistic revision
// rather than versioned, and a change regenerates Pricing draft suggestions
// without disturbing approved sale prices.
type BaseAcquisitionCost struct {
	ID       string             `json:"id"`
	SKUCode  string             `json:"sku_code"`
	Currency money.CurrencyCode `json:"currency"`
	// Amount is tax exclusive. Recoverable input tax is never included.
	Amount        money.Money `json:"amount"`
	SourceType    string      `json:"source_type,omitempty"`
	SourceID      string      `json:"source_id,omitempty"`
	Revision      int64       `json:"revision"`
	EffectiveFrom time.Time   `json:"effective_from"`
	// Sources preserves the exact frozen quantity and net-goods amount for
	// each contributing source line; Amount is the per-base-unit result.
	Sources []NetGoodsCostSourceSnapshot `json:"sources,omitempty"`
	// ManualLock records the active manual cost decision. Its absence means
	// no manual lock is represented, not that the amount or source is zero.
	ManualLock *audit.LifecycleAction `json:"manual_lock,omitempty"`

	audit.AuditFields
}
