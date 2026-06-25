package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// Snapshot is the denormalised product summary embedded in carts, order
// lines, purchase orders, and membership subscription plans. It captures what the
// product looked like at the time of the transaction so historical rows
// survive later product edits.
//
// Price is intentionally not part of the snapshot: each consumer carries
// its own price field (unit price, unit cost, plan price) with its own
// semantics.
//
// DisplayStatus is the read-time merge of status + recency + stock (see
// Product.DisplayStatus). It is populated only on LIVE snapshots built
// for display (e.g. a SKU's product lineup, replenishment suggestions);
// it is left empty on snapshots persisted into orders/carts/POs, because
// freezing a transient state like "new" or "out_of_stock" into a
// historical row would be wrong.
type Snapshot struct {
	ID            string                 `json:"id,omitempty"`
	SKUCode       string                 `json:"sku_code,omitempty"`
	SKU           string                 `json:"sku,omitempty"`
	Name          string                 `json:"name,omitempty"`
	OtherNames    []common.LocalizedName `json:"other_names,omitempty"`
	Brand         string                 `json:"brand,omitempty"`
	Vendor        string                 `json:"vendor,omitempty"`
	ImageURL      string                 `json:"image_url,omitempty"`
	Storage       enums.StorageType      `json:"storage,omitempty"`
	Status        enums.ProductStatus    `json:"status,omitempty"`
	DisplayStatus string                 `json:"display_status,omitempty"`
	Barcode       string                 `json:"barcode,omitempty"`
	Taxed         bool                   `json:"taxed"`
}
