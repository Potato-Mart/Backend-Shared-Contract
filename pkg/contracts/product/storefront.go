package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/common"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/product"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/warehouse"
)

// StorefrontPricing contains only prices approved for the current storefront
// audience. It must never include cost or unrelated channel prices.
type StorefrontPricing struct {
	Audience      productenum.PriceAudience      `json:"audience"`
	Current       *common.Money                  `json:"current,omitempty"`
	Original      *common.Money                  `json:"original,omitempty"`
	WholesaleMode productenum.WholesalePriceMode `json:"wholesale_mode,omitempty"`
}

// StorefrontPromotionBadge is a customer-safe image-overlay promotion model.
type StorefrontPromotionBadge struct {
	PromotionID     string                 `json:"promotion_id"`
	Labels          []common.LocalizedName `json:"labels,omitempty"`
	DiscountPercent *int                   `json:"discount_percent,omitempty"`
	OriginalPrice   *common.Money          `json:"original_price,omitempty"`
	CurrentPrice    *common.Money          `json:"current_price,omitempty"`
	StartsAt        *time.Time             `json:"starts_at,omitempty"`
	ExpiresAt       *time.Time             `json:"expires_at,omitempty"`
}

// StorefrontProduct is the shared customer-safe catalogue projection consumed
// by retail and approved wholesale storefronts. Exact expiry is omitted unless
// StorefrontDisplay explicitly allows it.
// StorefrontOrigin is the customer-facing country-of-origin display block.
type StorefrontOrigin struct {
	CountryCode string                 `json:"country_code"`
	Label       []common.LocalizedText `json:"label,omitempty"`
	Statement   []common.LocalizedText `json:"statement,omitempty"`
}

type StorefrontProduct struct {
	SKUCode             string                        `json:"sku_code"`
	SKU                 string                        `json:"sku"`
	Name                string                        `json:"name"`
	Description         []common.LocalizedDescription `json:"description,omitempty"`
	BrandRef            *BrandRef                     `json:"brand_ref,omitempty"`
	Storage             warehouseenum.StorageType     `json:"storage,omitempty"`
	Status              productenum.ProductStatus     `json:"status,omitempty"`
	Collection          *CollectionRef                `json:"collection,omitempty"`
	CategoryTags        []CategoryTag                 `json:"category_tags,omitempty"`
	Supply              *ProductSupply                `json:"supply,omitempty"`
	CurrentStock        int                           `json:"current_stock"`
	Pricing             StorefrontPricing             `json:"pricing"`
	StorefrontDisplay   StorefrontDisplay             `json:"storefront_display"`
	ExpiryDate          *time.Time                    `json:"expiry_date,omitempty"`
	CoverURL            string                        `json:"cover_url,omitempty"`
	ImageURLs           []string                      `json:"image_urls,omitempty"`
	DetailImages        []DetailImage                 `json:"detail_images,omitempty"`
	DisplayStatus       string                        `json:"display_status,omitempty"`
	DisplaySellingCount *int64                        `json:"display_selling_count,omitempty"`
	PromotionBadge      *StorefrontPromotionBadge     `json:"promotion_badge,omitempty"`
	SalesPerformance    *SalesPerformanceStats        `json:"sales_performance,omitempty"`
	CountryOfOrigin     *StorefrontOrigin             `json:"country_of_origin,omitempty"`
	PhysicalWeight      *common.Weight                `json:"physical_weight,omitempty"`
}
