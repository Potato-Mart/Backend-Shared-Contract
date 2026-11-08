package warehouse

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"

// ShopRefillLevel is the min/max stock policy for one product at one shop
// depot, in base units. When the shop's free stock falls below MinBaseUnits,
// Supply refills it from the shop's supporting depots up to MaxBaseUnits in
// whole cases. Supply owns validation and the refill movements.
type ShopRefillLevel struct {
	ID            string `json:"id"`
	ShopDepotCode string `json:"shop_depot_code"`
	SKUCode       string `json:"sku_code"`
	MinBaseUnits  int64  `json:"min_base_units"`
	MaxBaseUnits  int64  `json:"max_base_units"`
	IsActive      bool   `json:"is_active"`

	audit.AuditFields
}
