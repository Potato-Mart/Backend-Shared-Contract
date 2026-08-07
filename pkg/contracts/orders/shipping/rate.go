package shipping

import (
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/warehouse"
)

type Rate struct {
	ID            string                    `json:"id"`
	ZoneID        string                    `json:"zone_id"`
	Name          ShippingRateName          `json:"name"`
	Price         common.Money              `json:"price"`
	FreeAbove     *common.Money             `json:"free_above,omitempty"`
	EstimatedDays string                    `json:"estimated_days,omitempty"`
	StorageType   warehouseenum.StorageType `json:"storage_type,omitempty"`
	PackageLimits *PackageLimits            `json:"package_limits,omitempty"`
	IsActive      bool                      `json:"is_active"`
	CreatedAt     time.Time                 `json:"created_at"`
}
