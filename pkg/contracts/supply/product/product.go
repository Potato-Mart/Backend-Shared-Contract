package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/customers/retail/retail_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product/product_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse/warehouse_enums"
)

// Product is the canonical master record for one product SKU. Consumers that
// do not require a complete product use SKUCode instead of embedding Product.
type Product struct {
	SKUCode        string                        `json:"sku_code"`
	Content        ProductContent                `json:"content"`
	Classification ProductClassification         `json:"classification"`
	Packaging      ProductPackaging              `json:"packaging"`
	Commerce       ProductCommerce               `json:"commerce"`
	Metrics        ProductMetrics                `json:"metrics"`
	Supply         *classification.ProductSupply `json:"supply,omitempty"`
	Administration *ProductAdministration        `json:"administration,omitempty"`
}

// ProductContent contains the customer-facing, locale-aware product facts.
type ProductContent struct {
	Name         string                              `json:"name"`
	Descriptions []localization.LocalizedDescription `json:"descriptions,omitempty"`
	Localization *ProductLocalization                `json:"localization,omitempty"`
	Origin       *ProductOrigin                      `json:"origin,omitempty"`
	Images       *Images                             `json:"images,omitempty"`
}

// ProductLocalization groups secondary per-language display fields.
type ProductLocalization struct {
	OtherNames []localization.LocalizedName `json:"other_names,omitempty"`
}

// ProductOrigin is the customer-facing origin display block for a product.
type ProductOrigin struct {
	CountryCode geography.CountryCode        `json:"country_code,omitempty"`
	Label       []localization.LocalizedText `json:"label,omitempty"`
	Statement   []localization.LocalizedText `json:"statement,omitempty"`
}

// ProductClassification contains references to the product's catalogue
// classification. Category tags are deliberately lightweight references.
type ProductClassification struct {
	CategorySKUCode string                          `json:"category_sku_code"`
	BrandRef        *classification.BrandRef        `json:"brand_ref,omitempty"`
	CollectionRef   *classification.CollectionRef   `json:"collection_ref,omitempty"`
	CategoryTags    []classification.CategoryTagRef `json:"category_tags,omitempty"`
}

// ProductPackaging contains the canonical package, barcode, tax, and storage
// configuration for the product.
type ProductPackaging struct {
	PackageOptions     []ProductPackageOption      `json:"package_options"`
	BarcodeAssignments []ProductBarcodeAssignment  `json:"barcode_assignments,omitempty"`
	Taxed              bool                        `json:"taxed"`
	StorageType        warehouse_enums.StorageType `json:"storage_type,omitempty"`
}

// ProductCommerce contains lifecycle and sellability information. Its package
// entries are intentionally customer-safe and never disclose inventory source
// or raw quantity details.
type ProductCommerce struct {
	Status        product_enums.ProductStatus `json:"status,omitempty"`
	Selling       *Selling                    `json:"selling,omitempty"`
	FirstListedAt *time.Time                  `json:"first_listed_at,omitempty"`
	Packages      []ProductPackageCommerce    `json:"packages,omitempty"`
}

// ProductPackageCommerce is a safe commercial projection of one package
// option. Promotion applications are added in the later pricing phase.
type ProductPackageCommerce struct {
	PackageOptionID string                             `json:"package_option_id"`
	PackagePrice    money.Money                        `json:"package_price"`
	TaxAmount       money.Money                        `json:"tax_amount"`
	StockState      product_enums.StorefrontStockState `json:"stock_state,omitempty"`
	AsOf            time.Time                          `json:"as_of"`
}

// ProductMetrics contains backend-computed and curated selling measurements.
type ProductMetrics struct {
	SalesPerformance    *SalesPerformanceStats `json:"sales_performance,omitempty"`
	DisplaySellingCount *int64                 `json:"display_selling_count,omitempty"`
}

// ProductAdministration retains master-data history and audit information.
// Storefront and POS responses omit this optional component.
type ProductAdministration struct {
	History []security.HistoryEntry `json:"history,omitempty"`

	audit.AuditFields
}

// Selling groups the channel/buyer sellability rules for a product: which
// order channels it may be sold through, which buyer types may purchase it,
// and how its price/listing is exposed. Empty Channels and BuyerTypes mean no
// restriction; enforcement remains backend behavior.
type Selling struct {
	Channels   []commerce_enums.OrderType    `json:"channels,omitempty"`
	BuyerTypes []retail_enums.BuyerType      `json:"buyer_types,omitempty"`
	Visibility product_enums.PriceVisibility `json:"visibility,omitempty"`
}

// Images groups the customer-facing product imagery.
type Images struct {
	Cover   *security.ObjectMedia  `json:"cover,omitempty"`
	Gallery []security.ObjectMedia `json:"gallery,omitempty"`
	Details []security.ObjectMedia `json:"details,omitempty"`
}
