package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/customers/retail/retail_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product/product_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse/warehouse_enums"
)

// Product is the JSON master record for one product SKU.
type Product struct {
	SKUCode            string                              `json:"sku_code"`
	CategorySKUCode    string                              `json:"category_sku_code"`
	Name               string                              `json:"name"`
	Description        []localization.LocalizedDescription `json:"description,omitempty"`
	BrandRef           *BrandRef                           `json:"brand_ref,omitempty"`
	PackageOptions     []ProductPackageOption              `json:"package_options"`
	BarcodeAssignments []ProductBarcodeAssignment          `json:"barcode_assignments,omitempty"`
	Taxed              bool                                `json:"taxed"`

	StorageType warehouse_enums.StorageType `json:"storage_type,omitempty"`
	// Status is the admin-controlled product lifecycle state.
	Status       product_enums.ProductStatus `json:"status,omitempty"`
	Collection   *CollectionRef              `json:"collection,omitempty"`
	CategoryTags []CategoryTag               `json:"category_tags,omitempty"`
	Supply       *ProductSupply              `json:"supply,omitempty"`

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
	Images       *Images      `json:"images,omitempty"`
	// CountryOfOrigin is the customer-facing origin display block projected
	// onto the storefront product.
	CountryOfOrigin *StorefrontOrigin `json:"country_of_origin,omitempty"`
	// StorefrontMerchandising carries admin-managed retail display policy.
	// Backend read models convert it into customer-safe display fields.
	StorefrontMerchandising *StorefrontMerchandising `json:"storefront_merchandising,omitempty"`

	// History is for product master-data changes only. Stock changes are
	// represented as warehouse.StockMovement records.
	History []security.HistoryEntry `json:"history,omitempty"`

	audit.AuditFields
}

// Selling groups the channel/buyer sellability rules for a product: which
// order channels it may be sold through, which buyer types may purchase it,
// and how its price/listing is exposed. It reuses commerce_enums.OrderType for
// channel and customerenum.BuyerType for buyer type. Empty Channels/BuyerTypes
// mean "no restriction"; the contract defines the rules, not the
// enforcement — that lives in the backend.
type Selling struct {
	Channels   []commerce_enums.OrderType    `json:"channels,omitempty"`
	BuyerTypes []retail_enums.BuyerType      `json:"buyer_types,omitempty"`
	Visibility product_enums.PriceVisibility `json:"visibility,omitempty"`
}

// Localization groups secondary per-language display fields.
type Localization struct {
	OtherNames []localization.LocalizedName `json:"other_names,omitempty"`
}

// Images groups the customer-facing product imagery.
type Images struct {
	Cover   *security.ObjectMedia  `json:"cover,omitempty"`
	Gallery []security.ObjectMedia `json:"gallery,omitempty"`
	Details []security.ObjectMedia `json:"details,omitempty"`
}
