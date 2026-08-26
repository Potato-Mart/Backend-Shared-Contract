package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/warehouse/warehouse_enums"
)

// StockLocationCollectionEligibility identifies a collection's primary or
// ordered overflow placement at a location.
type StockLocationCollectionEligibility struct {
	CollectionCode string                                      `json:"collection_code"`
	Role           warehouse_enums.StockLocationCollectionRole `json:"role"`
	Priority       int                                         `json:"priority"`
}
