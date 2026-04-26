package category

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v2/pkg/enums"
)

// SKU corresponds to one of the top-level SKU codes that
// identify a product family (e.g. A0 = 特色台灣商品/吊飾, F2 = 冷凍-肉品).
// Products reference a category by its SKU Code.
type SKU struct {
	ID        string            `json:"id"`
	Code      string            `json:"code"`
	Storage   enums.StorageType `json:"storage"`
	Label     string            `json:"label"`
	SortOrder int               `json:"sort_order"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
