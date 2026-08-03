package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/product"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

// StorefrontPromotionBadge is a customer-safe image-overlay promotion model.
type StorefrontPromotionBadge struct {
	PromotionID       string                   `json:"promotion_id"`
	SeriesKey         string                   `json:"series_key"`
	Labels            []common.LocalizedName   `json:"labels,omitempty"`
	DiscountPercent   *int                     `json:"discount_percent,omitempty"`
	StartsAt          *time.Time               `json:"starts_at,omitempty"`
	ExpiresAt         *time.Time               `json:"expires_at,omitempty"`
	ScheduleTimezone  string                   `json:"schedule_timezone"`
	GeographicContext common.GeographicContext `json:"geographic_context"`
}

// StorefrontOrigin is the customer-facing country-of-origin display block.
type StorefrontOrigin struct {
	CountryCode common.CountryCode     `json:"country_code"`
	Label       []common.LocalizedText `json:"label,omitempty"`
	Statement   []common.LocalizedText `json:"statement,omitempty"`
}

// StorefrontProduct is the shared customer-safe package, offer, and
// availability projection consumed by retail and wholesale storefronts.
type StorefrontProduct struct {
	SKUCode             string                             `json:"sku_code"`
	CategorySKUCode     string                             `json:"category_sku_code"`
	Name                string                             `json:"name"`
	Description         []common.LocalizedDescription      `json:"description,omitempty"`
	BrandRef            *BrandRef                          `json:"brand_ref,omitempty"`
	StorageType         warehouseenum.StorageType          `json:"storage_type,omitempty"`
	Status              productenum.ProductStatus          `json:"status,omitempty"`
	Collection          *CollectionRef                     `json:"collection,omitempty"`
	CategoryTags        []CategoryTag                      `json:"category_tags,omitempty"`
	Supply              *ProductSupply                     `json:"supply,omitempty"`
	PackageOptions      []ProductPackageOptionSnapshot     `json:"package_options"`
	BarcodeAssignments  []ProductBarcodeAssignmentSnapshot `json:"barcode_assignments,omitempty"`
	Offers              []SellableOfferSnapshot            `json:"offers"`
	Availability        *ProductStockSummary               `json:"availability,omitempty"`
	Audience            productenum.PriceAudience          `json:"audience"`
	WholesaleMode       productenum.WholesalePriceMode     `json:"wholesale_mode,omitempty"`
	StorefrontDisplay   StorefrontDisplay                  `json:"storefront_display"`
	CoverURL            string                             `json:"cover_url,omitempty"`
	ImageURLs           []string                           `json:"image_urls,omitempty"`
	DetailImages        []DetailImage                      `json:"detail_images,omitempty"`
	DisplaySellingCount *int64                             `json:"display_selling_count,omitempty"`
	PromotionBadge      *StorefrontPromotionBadge          `json:"promotion_badge,omitempty"`
	SalesPerformance    *SalesPerformanceStats             `json:"sales_performance,omitempty"`
	CountryOfOrigin     *StorefrontOrigin                  `json:"country_of_origin,omitempty"`
}
