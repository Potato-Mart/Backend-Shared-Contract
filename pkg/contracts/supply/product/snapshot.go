package product

import (
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/warehouse"
)

// Snapshot is the immutable product summary carried by transaction records.
// Live availability and accepted prices are represented by their own
// revisioned snapshots.
type Snapshot struct {
	SKUCode            string                             `json:"sku_code,omitempty"`
	CategorySKUCode    string                             `json:"category_sku_code,omitempty"`
	Name               string                             `json:"name,omitempty"`
	OtherNames         []common.LocalizedName             `json:"other_names,omitempty"`
	Description        []common.LocalizedDescription      `json:"description,omitempty"`
	BrandRef           *BrandRef                          `json:"brand_ref,omitempty"`
	Collection         *CollectionRef                     `json:"collection,omitempty"`
	CategoryTags       []CategoryTag                      `json:"category_tags,omitempty"`
	Supply             *ProductSupply                     `json:"supply,omitempty"`
	ImageURL           string                             `json:"image_url,omitempty"`
	StorageType        warehouseenum.StorageType          `json:"storage_type,omitempty"`
	Status             ProductStatus                      `json:"status,omitempty"`
	PackageOptions     []ProductPackageOptionSnapshot     `json:"package_options,omitempty"`
	BarcodeAssignments []ProductBarcodeAssignmentSnapshot `json:"barcode_assignments,omitempty"`
	Taxed              bool                               `json:"taxed"`
}
