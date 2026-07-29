package sales

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/common"
	promotionenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/promotion"
)

// VolumeDiscountTier defines a carton-quantity threshold for bulk purchase
// discounts. The system selects the single highest qualifying tier.
type VolumeDiscountTier struct {
	ID              string                                `json:"id"`
	Name            string                                `json:"name"`
	MinCartons      int                                   `json:"min_cartons"`
	DiscountPercent float64                               `json:"discount_percent"`
	AppliesTo       promotionenum.VolumeDiscountAppliesTo `json:"applies_to"`
	IsActive        bool                                  `json:"is_active"`
	common.AuditFields
}
