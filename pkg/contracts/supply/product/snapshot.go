package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product/product_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse/warehouse_enums"
)

// Snapshot is the immutable product summary carried by transaction records.
// Live availability and accepted prices are represented by their own
// revisioned snapshots.
type Snapshot struct {
	SKUCode            string                              `json:"sku_code,omitempty"`
	CategorySKUCode    string                              `json:"category_sku_code,omitempty"`
	Name               string                              `json:"name,omitempty"`
	OtherNames         []localization.LocalizedName        `json:"other_names,omitempty"`
	Description        []localization.LocalizedDescription `json:"description,omitempty"`
	BrandRef           *classification.BrandRef            `json:"brand_ref,omitempty"`
	Collection         *classification.CollectionRef       `json:"collection,omitempty"`
	CategoryTags       []classification.CategoryTag        `json:"category_tags,omitempty"`
	Supply             *classification.ProductSupply       `json:"supply,omitempty"`
	ImageURL           string                              `json:"image_url,omitempty"`
	StorageType        warehouse_enums.StorageType         `json:"storage_type,omitempty"`
	Status             product_enums.ProductStatus         `json:"status,omitempty"`
	PackageOptions     []ProductPackageOptionSnapshot      `json:"package_options,omitempty"`
	BarcodeAssignments []ProductBarcodeAssignmentSnapshot  `json:"barcode_assignments,omitempty"`
	Taxed              bool                                `json:"taxed"`
}
