package order

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"time"
)

type DemandBucket struct {
	SKUCode       string                               `json:"sku_code"`
	SKUSeriesCode string                               `json:"sku_series_code,omitempty"`
	ProductName   string                               `json:"product_name,omitempty"`
	DepotCode     string                               `json:"depot_code"`
	Channel       commerce_enums.OrderType             `json:"channel"`
	Date          time.Time                            `json:"date"`
	Composition   packaging.PackageCompositionSnapshot `json:"composition"`
}
