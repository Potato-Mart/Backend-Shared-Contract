package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/enums"
)

// Snapshot is the denormalised product summary embedded in carts, order
// lines, purchase orders, and subscription plans. It captures what the
// product looked like at the time of the transaction so historical rows
// survive later product edits.
//
// Price is intentionally not part of the snapshot: each consumer carries
// its own price field (unit price, unit cost, plan price) with its own
// semantics.
type Snapshot struct {
	ID         string                 `json:"id,omitempty"`
	SKU        string                 `json:"sku,omitempty"`
	Name       string                 `json:"name"`
	OtherNames []common.LocalizedName `json:"other_names,omitempty"`
	Brand      string                 `json:"brand,omitempty"`
	ImageURL   string                 `json:"image_url,omitempty"`
	Storage    enums.StorageType      `json:"storage,omitempty"`
	Barcode    string                 `json:"barcode,omitempty"`
}
