package order

import (
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
)

type DemandBucket struct {
	ProductSKUCode  string                            `json:"product_sku_code"`
	CategorySKUCode string                            `json:"category_sku_code,omitempty"`
	ProductName     string                            `json:"product_name,omitempty"`
	DepotCode       string                            `json:"depot_code"`
	Channel         common.OrderType                  `json:"channel"`
	Date            time.Time                         `json:"date"`
	Composition     common.PackageCompositionSnapshot `json:"composition"`
}

type OpenDemandLine struct {
	OrderNumber     string                            `json:"order_number"`
	ProductSKUCode  string                            `json:"product_sku_code"`
	CategorySKUCode string                            `json:"category_sku_code,omitempty"`
	ProductName     string                            `json:"product_name,omitempty"`
	DepotCode       string                            `json:"depot_code"`
	Channel         common.OrderType                  `json:"channel"`
	Composition     common.PackageCompositionSnapshot `json:"composition"`
	CreatedAt       time.Time                         `json:"created_at"`
}
