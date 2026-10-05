package supply

import "github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/packaging"

type InventoryPackageConvertedEvent struct {
	MovementID                    string                               `json:"movement_id"`
	SKUCode                       string                               `json:"sku_code"`
	DepotCode                     string                               `json:"depot_code"`
	LotID                         string                               `json:"lot_id"`
	SourceBucketID                string                               `json:"source_bucket_id"`
	DestinationBucketID           string                               `json:"destination_bucket_id"`
	SourcePackageOption           packaging.PackageOptionRef           `json:"source_package_option"`
	DestinationPackageOption      packaging.PackageOptionRef           `json:"destination_package_option"`
	BaseUnits                     int64                                `json:"base_units"`
	SourcePackageComposition      packaging.PackageCompositionSnapshot `json:"source_package_composition"`
	DestinationPackageComposition packaging.PackageCompositionSnapshot `json:"destination_package_composition"`
	SourceBucketRevision          int64                                `json:"source_bucket_revision"`
	DestinationBucketRevision     int64                                `json:"destination_bucket_revision"`
}
