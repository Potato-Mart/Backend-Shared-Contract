package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/sales"
)

// OrderPackingProjection is the durable packing snapshot shared between
// Supply and Orders. Transport acknowledgement types stay provider-local.
type OrderPackingProjection struct {
	OrderNumber string               `json:"order_number"`
	Revision    int64                `json:"revision"`
	Packing     OrderPackingProgress `json:"packing"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type DemandBucket struct {
	ProductSKUCode  string                            `json:"product_sku_code"`
	CategorySKUCode string                            `json:"category_sku_code,omitempty"`
	ProductName     string                            `json:"product_name,omitempty"`
	DepotCode       string                            `json:"depot_code"`
	Channel         salesenum.OrderType               `json:"channel"`
	Date            time.Time                         `json:"date"`
	Composition     common.PackageCompositionSnapshot `json:"composition"`
}

type OpenDemandLine struct {
	OrderNumber     string                            `json:"order_number"`
	ProductSKUCode  string                            `json:"product_sku_code"`
	CategorySKUCode string                            `json:"category_sku_code,omitempty"`
	ProductName     string                            `json:"product_name,omitempty"`
	DepotCode       string                            `json:"depot_code"`
	Channel         salesenum.OrderType               `json:"channel"`
	Composition     common.PackageCompositionSnapshot `json:"composition"`
	CreatedAt       time.Time                         `json:"created_at"`
}
