package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/localization"
)

// CategorySalesRank records one product's position within one category tag.
// Rankings use the embedded window and are computed across the complete result
// set, never from one paginated storefront response.
type CategorySalesRank struct {
	CategoryTagCode string                       `json:"category_tag_code"`
	CategoryTagName []localization.LocalizedName `json:"category_tag_name,omitempty"`
	Rank            int                          `json:"rank"`
	Population      int                          `json:"population"`
	WindowDays      int                          `json:"window_days"`
	NetUnits        int64                        `json:"net_units"`
}
