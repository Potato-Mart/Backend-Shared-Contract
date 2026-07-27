package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/common"
	shippingenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/shipping"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/warehouse"
)

type Rate struct {
	ID            string                        `json:"id"`
	ZoneID        string                        `json:"zone_id"`
	Name          shippingenum.ShippingRateName `json:"name"`
	Price         common.Money                  `json:"price"`
	FreeAbove     *common.Money                 `json:"free_above,omitempty"`
	EstimatedDays string                        `json:"estimated_days,omitempty"`
	StorageType   warehouseenum.StorageType     `json:"storage_type,omitempty"`
	PackageLimits *PackageLimits                `json:"package_limits,omitempty"`
	IsActive      bool                          `json:"is_active"`
	CreatedAt     time.Time                     `json:"created_at"`
}
