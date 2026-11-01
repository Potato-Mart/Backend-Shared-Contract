package supply

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/product/product_enums"
)

// CatalogProductStatusChangedEvent is emitted on the catalog-events topic when
// Supply changes a product's admin-controlled lifecycle status. AggregateID is
// the SKU code.
//
// It carries the revision that produced the status so consumers can apply it
// idempotently: a consumer applies the event only when Revision is newer than
// the product revision it already holds, so a redelivered or out-of-order
// message cannot restore a stale status. The status is lifecycle state only:
// it does not describe live stock availability or a market listing status.
type CatalogProductStatusChangedEvent struct {
	SKUCode string `json:"sku_code"`

	// PreviousStatus is empty on a product's first status write.
	PreviousStatus product_enums.ProductStatus `json:"previous_status,omitempty"`
	Status         product_enums.ProductStatus `json:"status"`

	PreviousRevision int64 `json:"previous_revision"`
	Revision         int64 `json:"revision"`

	ChangedAt time.Time `json:"changed_at"`
	ChangedBy string    `json:"changed_by,omitempty"`
}
