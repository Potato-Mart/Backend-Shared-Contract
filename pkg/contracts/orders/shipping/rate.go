package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/orders/shipping/shipping_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/supply/classification/classification_enums"
)

type Rate struct {
	ID            string                           `json:"id"`
	ZoneID        string                           `json:"zone_id"`
	Name          shipping_enums.ShippingRateName  `json:"name"`
	Price         money.Money                      `json:"price"`
	FreeAbove     *money.Money                     `json:"free_above,omitempty"`
	EstimatedDays string                           `json:"estimated_days,omitempty"`
	StorageType   classification_enums.StorageType `json:"storage_type,omitempty"`
	PackageLimits *PackageLimits                   `json:"package_limits,omitempty"`
	IsActive      bool                             `json:"is_active"`
	CreatedAt     time.Time                        `json:"created_at"`
}
