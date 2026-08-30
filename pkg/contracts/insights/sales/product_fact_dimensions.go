package sales

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/classification/classification_enums"
)

// ProductFactDimensions is the frozen product, catalogue, and geographic
// dimensional evidence shared by order and refund item facts.
type ProductFactDimensions struct {
	SKUCode             string                               `json:"sku_code"`
	MarketCode          string                               `json:"market_code"`
	CountryCode         geography.CountryCode                `json:"country_code,omitempty"`
	BrandCode           string                               `json:"brand_code,omitempty"`
	StorageType         classification_enums.StorageType     `json:"storage_type,omitempty"`
	CollectionCode      string                               `json:"collection_code,omitempty"`
	CategoryTagCodes    []string                             `json:"category_tag_codes,omitempty"`
	Production          string                               `json:"production,omitempty"`
	PackageComposition  packaging.PackageCompositionSnapshot `json:"package_composition"`
}
