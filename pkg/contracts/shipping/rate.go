package shipping

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v2/pkg/enums"
)

type Rate struct {
	ID            string                 `json:"id"`
	ZoneID        string                 `json:"zone_id"`
	Name          enums.ShippingRateName `json:"name"`
	Price         float64                `json:"price"`
	FreeAbove     float64                `json:"free_above,omitempty"`
	EstimatedDays string                 `json:"estimated_days,omitempty"`
	StorageType   enums.StorageType      `json:"storage_type,omitempty"`
	IsActive      bool                   `json:"is_active"`
	SortOrder     int                    `json:"sort_order"`
	CreatedAt     time.Time              `json:"created_at"`
}
