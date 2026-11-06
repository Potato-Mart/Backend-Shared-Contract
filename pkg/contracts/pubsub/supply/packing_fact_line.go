package supply

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification/classification_enums"

// PackingFactLine is one packed order line of a stage, with the storage type
// snapshotted when the order was marked packed.
type PackingFactLine struct {
	OrderItemID    string                           `json:"order_item_id"`
	SKUCode        string                           `json:"sku_code"`
	StorageType    classification_enums.StorageType `json:"storage_type"`
	TotalBaseUnits int64                            `json:"total_base_units"`
}
