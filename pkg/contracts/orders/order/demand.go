package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging"
)

type DemandBucket struct {
	ProductSKUCode  string                               `json:"product_sku_code"`
	CategorySKUCode string                               `json:"category_sku_code,omitempty"`
	ProductName     string                               `json:"product_name,omitempty"`
	DepotCode       string                               `json:"depot_code"`
	Channel         commerce_enums.OrderType             `json:"channel"`
	Date            time.Time                            `json:"date"`
	Composition     packaging.PackageCompositionSnapshot `json:"composition"`
}

type OpenDemandLine struct {
	OrderNumber     string                               `json:"order_number"`
	ProductSKUCode  string                               `json:"product_sku_code"`
	CategorySKUCode string                               `json:"category_sku_code,omitempty"`
	ProductName     string                               `json:"product_name,omitempty"`
	DepotCode       string                               `json:"depot_code"`
	Channel         commerce_enums.OrderType             `json:"channel"`
	Composition     packaging.PackageCompositionSnapshot `json:"composition"`
	CreatedAt       time.Time                            `json:"created_at"`
}
