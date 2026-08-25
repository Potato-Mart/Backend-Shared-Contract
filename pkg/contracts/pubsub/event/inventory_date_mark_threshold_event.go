package event

import (
	"time"

	warehouse "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/warehouse"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/warehouse/warehouse_enums"
)

type InventoryDateMarkThresholdEvent struct {
	LotID       string                                     `json:"lot_id"`
	SKUCode     string                                     `json:"sku_code"`
	DepotCode   string                                     `json:"depot_code"`
	DateMark    warehouse.InventoryDateMark                `json:"date_mark"`
	Threshold   warehouse_enums.InventoryDateMarkThreshold `json:"threshold"`
	ThresholdAt time.Time                                  `json:"threshold_at"`
	LotRevision int64                                      `json:"lot_revision"`
	OccurredAt  time.Time                                  `json:"occurred_at"`
}
