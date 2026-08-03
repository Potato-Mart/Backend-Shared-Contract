package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

// Depot is a fulfilment site within a depot region. Code is its canonical
// cross-service reference and Timezone is an IANA time-zone identifier.
type Depot struct {
	ID         string         `json:"id"`
	Code       string         `json:"code"`
	Name       string         `json:"name"`
	RegionCode string         `json:"region_code"`
	Address    common.Address `json:"address"`
	Timezone   string         `json:"timezone"`
	Phone      string         `json:"phone,omitempty"`
	IsActive   bool           `json:"is_active"`
	LayoutID   string         `json:"layout_id,omitempty"`

	common.AuditFields
}

type DepotProduct struct {
	DepotCode      string    `json:"depot_code"`
	ProductSKUCode string    `json:"product_sku_code"`
	StockQty       int       `json:"stock_qty"`
	IsAvailable    bool      `json:"is_available"`
	LocationCode   string    `json:"location_code,omitempty"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// StockLocation is the inventory bin a product physically lives in.
type StockLocation struct {
	ID           string                    `json:"id"`
	DepotCode    string                    `json:"depot_code"`
	Code         string                    `json:"code"`
	Name         string                    `json:"name,omitempty"`
	Zone         warehouseenum.StorageType `json:"zone,omitempty"`
	IsActive     bool                      `json:"is_active"`
	LayoutNodeID string                    `json:"layout_node_id,omitempty"`
	Transform    *common.Transform         `json:"transform,omitempty"`
	Size         *common.Size3D            `json:"size,omitempty"`
	Shape        warehouseenum.ShapeType   `json:"shape,omitempty"`
	Color        string                    `json:"color,omitempty"`
}
