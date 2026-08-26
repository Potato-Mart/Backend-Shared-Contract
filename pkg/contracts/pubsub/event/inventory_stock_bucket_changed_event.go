package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/operations"
	warehouse "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/warehouse"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/warehouse/warehouse_enums"
)

type InventoryStockBucketChangedEvent struct {
	BucketID                 string                               `json:"bucket_id"`
	Location                 warehouse.StockLocationRef           `json:"location"`
	SKUCode                  string                               `json:"sku_code"`
	LotID                    string                               `json:"lot_id,omitempty"`
	PackageOptionCode        string                               `json:"package_option_code"`
	HandlingUnit             packaging_enums.PackageHandlingUnit  `json:"handling_unit"`
	Condition                warehouse_enums.InventoryCondition   `json:"condition"`
	Disposition              warehouse_enums.InventoryDisposition `json:"disposition"`
	OnHandBeforeBaseUnits    int64                                `json:"on_hand_before_base_units"`
	OnHandAfterBaseUnits     int64                                `json:"on_hand_after_base_units"`
	ReservedBeforeBaseUnits  int64                                `json:"reserved_before_base_units"`
	ReservedAfterBaseUnits   int64                                `json:"reserved_after_base_units"`
	AvailableBeforeBaseUnits int64                                `json:"available_before_base_units"`
	AvailableAfterBaseUnits  int64                                `json:"available_after_base_units"`
	Cause                    operations.InventoryCauseRef         `json:"cause"`
	Revision                 int64                                `json:"revision"`
	OccurredAt               time.Time                            `json:"occurred_at"`
}
