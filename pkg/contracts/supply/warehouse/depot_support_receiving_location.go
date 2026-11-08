package warehouse

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification/classification_enums"

// DepotSupportReceivingLocation names the shop location that receives stock
// of one storage type moved from a supporting depot, so ambient, chilled and
// frozen stock each land in a location of their own storage type. Supply
// validates that LocationCode is a standard location at the shop whose
// storage type equals StorageType.
type DepotSupportReceivingLocation struct {
	StorageType  classification_enums.StorageType `json:"storage_type"`
	LocationCode string                           `json:"location_code"`
}
