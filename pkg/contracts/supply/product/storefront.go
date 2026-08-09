package product

import (
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product/product_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse/warehouse_enums"
)

// StorefrontPromotionBadge is a customer-safe image-overlay promotion model.
type StorefrontPromotionBadge struct {
	PromotionID       string                       `json:"promotion_id"`
	SeriesKey         string                       `json:"series_key"`
	Labels            []localization.LocalizedName `json:"labels,omitempty"`
	DiscountPercent   *int                         `json:"discount_percent,omitempty"`
	StartsAt          *time.Time                   `json:"starts_at,omitempty"`
	ExpiresAt         *time.Time                   `json:"expires_at,omitempty"`
	ScheduleTimezone  string                       `json:"schedule_timezone"`
	GeographicContext geography.GeographicContext  `json:"geographic_context"`
}

// StorefrontOrigin is the customer-facing country-of-origin display block.
type StorefrontOrigin struct {
	CountryCode geography.CountryCode        `json:"country_code"`
	Label       []localization.LocalizedText `json:"label,omitempty"`
	Statement   []localization.LocalizedText `json:"statement,omitempty"`
}

// StorefrontProduct is the shared customer-safe package, offer, and
// availability projection consumed by retail and wholesale storefronts.
type StorefrontProduct struct {
	SKUCode             string                              `json:"sku_code"`
	CategorySKUCode     string                              `json:"category_sku_code"`
	Name                string                              `json:"name"`
	Description         []localization.LocalizedDescription `json:"description,omitempty"`
	BrandRef            *BrandRef                           `json:"brand_ref,omitempty"`
	StorageType         warehouse_enums.StorageType         `json:"storage_type,omitempty"`
	Status              product_enums.ProductStatus         `json:"status,omitempty"`
	Collection          *CollectionRef                      `json:"collection,omitempty"`
	CategoryTags        []CategoryTag                       `json:"category_tags,omitempty"`
	Supply              *ProductSupply                      `json:"supply,omitempty"`
	PackageOptions      []ProductPackageOptionSnapshot      `json:"package_options"`
	BarcodeAssignments  []ProductBarcodeAssignmentSnapshot  `json:"barcode_assignments,omitempty"`
	Offers              []SellableOfferSnapshot             `json:"offers"`
	Availability        *ProductStockSummary                `json:"availability,omitempty"`
	Commercial          *StorefrontCommercial               `json:"commercial,omitempty"`
	Audience            product_enums.PriceAudience         `json:"audience"`
	WholesaleMode       product_enums.WholesalePriceMode    `json:"wholesale_mode,omitempty"`
	StorefrontDisplay   StorefrontDisplay                   `json:"storefront_display"`
	CoverURL            string                              `json:"cover_url,omitempty"`
	ImageURLs           []string                            `json:"image_urls,omitempty"`
	DetailImages        []DetailImage                       `json:"detail_images,omitempty"`
	DisplaySellingCount *int64                              `json:"display_selling_count,omitempty"`
	PromotionBadge      *StorefrontPromotionBadge           `json:"promotion_badge,omitempty"`
	SalesPerformance    *SalesPerformanceStats              `json:"sales_performance,omitempty"`
	CountryOfOrigin     *StorefrontOrigin                   `json:"country_of_origin,omitempty"`
}

// StorefrontCommercial is the public AU commercial projection. It is safe for
// guests because it contains one canonical EACH package, one retail price, an
// aggregate stock state, market identity, and freshness only; it never carries
// depot, lot, offer, geographic-resolution, or raw quantity data.
type StorefrontCommercial struct {
	Price      *money.Money                       `json:"price,omitempty"`
	Package    ProductPackageOptionSnapshot       `json:"package_option"`
	StockState product_enums.StorefrontStockState `json:"stock_state"`
	Market     geography.CountryCode              `json:"market"`
	AsOf       time.Time                          `json:"as_of"`
}
