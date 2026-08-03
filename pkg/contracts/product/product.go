package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/shared"
	customerenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/customer"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/product"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/sales"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

// Product is the JSON master record for one product SKU.
type Product struct {
	SKUCode            string                        `json:"sku_code"`
	CategorySKUCode    string                        `json:"category_sku_code"`
	Name               string                        `json:"name"`
	Description        []common.LocalizedDescription `json:"description,omitempty"`
	BrandRef           *BrandRef                     `json:"brand_ref,omitempty"`
	PackageOptions     []ProductPackageOption        `json:"package_options"`
	BarcodeAssignments []ProductBarcodeAssignment    `json:"barcode_assignments,omitempty"`
	Taxed              bool                          `json:"taxed"`

	StorageType warehouseenum.StorageType `json:"storage_type,omitempty"`
	// Status is the admin-controlled product lifecycle state.
	Status       productenum.ProductStatus `json:"status,omitempty"`
	Collection   *CollectionRef            `json:"collection,omitempty"`
	CategoryTags []CategoryTag             `json:"category_tags,omitempty"`
	Supply       *ProductSupply            `json:"supply,omitempty"`

	// SalesPerformance contains backend-computed historical sales statistics.
	// It is never manually authored on product create/update operations.
	SalesPerformance *SalesPerformanceStats `json:"sales_performance,omitempty"`
	// DisplaySellingCount is an optional manually curated storefront count.
	// A non-nil pointer with value zero is distinct from an absent count.
	DisplaySellingCount *int64 `json:"display_selling_count,omitempty"`

	// FirstListedAt is the absolute instant at which the product was first
	// listed. Producers serialize operational instants in UTC.
	FirstListedAt *time.Time `json:"first_listed_at,omitempty"`

	// Selling is an optional pointer so a product with no channel/buyer
	// sellability rules omits it.
	Selling      *Selling     `json:"selling,omitempty"`
	Localization Localization `json:"localization,omitempty"`
	Media        Media        `json:"media,omitempty"`
	// CountryOfOrigin is the customer-facing origin display block projected
	// onto the storefront product.
	CountryOfOrigin *StorefrontOrigin `json:"country_of_origin,omitempty"`
	// StorefrontMerchandising carries admin-managed retail display policy.
	// Backend read models convert it into customer-safe display fields.
	StorefrontMerchandising *StorefrontMerchandising `json:"storefront_merchandising,omitempty"`

	// History is for product master-data changes only. Stock changes are
	// represented as warehouse.StockMovement records.
	History []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields
}

// Selling groups the channel/buyer sellability rules for a product: which
// order channels it may be sold through, which buyer types may purchase it,
// and how its price/listing is exposed. It reuses salesenum.OrderType for
// channel and customerenum.BuyerType for buyer type. Empty Channels/BuyerTypes
// mean "no restriction"; the contract defines the rules, not the
// enforcement — that lives in the backend.
type Selling struct {
	Channels   []salesenum.OrderType       `json:"channels,omitempty"`
	BuyerTypes []customerenum.BuyerType    `json:"buyer_types,omitempty"`
	Visibility productenum.PriceVisibility `json:"visibility,omitempty"`
}

// Localization groups secondary per-language display fields.
type Localization struct {
	OtherNames []common.LocalizedName `json:"other_names,omitempty"`
}

// Media groups the product imagery.
type Media struct {
	CoverMediaID  string        `json:"cover_media_id,omitempty"`
	CoverURL      string        `json:"cover_url,omitempty"`
	ImageMediaIDs []string      `json:"image_media_ids,omitempty"`
	ImageURLs     []string      `json:"image_urls,omitempty"`
	DetailImages  []DetailImage `json:"detail_images,omitempty"`
}
