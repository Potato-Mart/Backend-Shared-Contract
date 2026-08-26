// Package cost holds the Supply-owned private operational cost models. Base
// acquisition cost and carrying cost are internal values: they never appear in
// a storefront, order, invoice, or receipt response.
package cost

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
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

	audit.AuditFields
}
