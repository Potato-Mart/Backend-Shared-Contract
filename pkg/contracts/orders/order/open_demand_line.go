package order

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/packaging"
	"time"
)

type OpenDemandLine struct {
	OrderNumber   string                               `json:"order_number"`
	SKUCode       string                               `json:"sku_code"`
	SKUSeriesCode string                               `json:"sku_series_code,omitempty"`
	ProductName   string                               `json:"product_name,omitempty"`
	DepotCode     string                               `json:"depot_code"`
	Channel       commerce_enums.OrderType             `json:"channel"`
	Composition   packaging.PackageCompositionSnapshot `json:"composition"`
	CreatedAt     time.Time                            `json:"created_at"`
}
