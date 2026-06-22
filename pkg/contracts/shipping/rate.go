package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/enums"
)

type Rate struct {
	ID            string                 `json:"id"`
	ZoneID        string                 `json:"zone_id"`
	Name          enums.ShippingRateName `json:"name"`
	Price         common.Money           `json:"price"`
	FreeAbove     *common.Money          `json:"free_above,omitempty"`
	EstimatedDays string                 `json:"estimated_days,omitempty"`
	StorageType   enums.StorageType      `json:"storage_type,omitempty"`
	PackageLimits *PackageLimits         `json:"package_limits,omitempty"`
	IsActive      bool                   `json:"is_active"`
	SortOrder     int                    `json:"sort_order"`
	CreatedAt     time.Time              `json:"created_at"`
}
