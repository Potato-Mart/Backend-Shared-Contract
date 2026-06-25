package wholesale

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

const (
	// PathApprovedStorefrontProducts is the authenticated wholesale storefront
	// catalogue endpoint for approved buyers.
	PathApprovedStorefrontProducts = "/v1/storefront/wholesale/approved/products"
	// PathInternalApprovedStorefrontProducts is the service-auth catalogue
	// resolver Commerce uses to reprice wholesale cart lines server-side.
	PathInternalApprovedStorefrontProducts = "/v1/internal/catalog/wholesale/products"
)

// ApprovedStorefrontProduct is the authenticated wholesale catalogue projection
// for approved B2B buyers. It intentionally exposes one effective wholesale
// price, not the product.Pricing object, and omits operational stock quantity.
type ApprovedStorefrontProduct struct {
	ID                 string                 `json:"id"`
	SKUCode            string                 `json:"sku_code,omitempty"`
	SKU                string                 `json:"sku,omitempty"`
	Name               string                 `json:"name"`
	Description        string                 `json:"description,omitempty"`
	Barcode            string                 `json:"barcode,omitempty"`
	OtherNames         []common.LocalizedName `json:"other_names,omitempty"`
	BrandKey           string                 `json:"brand_key,omitempty"`
	BrandNames         []common.LocalizedName `json:"brand_names,omitempty"`
	Vendor             string                 `json:"vendor,omitempty"`
	Catalogue          string                 `json:"catalogue,omitempty"`
	Storage            enums.StorageType      `json:"storage,omitempty"`
	SalesPerformance   enums.SalesPerformance `json:"sales_performance,omitempty"`
	DisplayStatus      string                 `json:"display_status,omitempty"`
	CoverURL           string                 `json:"cover_url,omitempty"`
	ImageURLs          []string               `json:"image_urls,omitempty"`
	CategoryKey        string                 `json:"category_key,omitempty"`
	CategoryPath       []string               `json:"category_path,omitempty"`
	CategoryTags       []product.CategoryTag  `json:"category_tags,omitempty"`
	LifecycleTags      []string               `json:"lifecycle_tags,omitempty"`
	NewExpiresAt       *time.Time             `json:"new_expires_at,omitempty"`
	RestockedExpiresAt *time.Time             `json:"restocked_expires_at,omitempty"`

	Price             common.Money `json:"price"`
	StockAvailable    bool         `json:"stock_available"`
	AvailabilityState string       `json:"availability_state"`
}

// ApprovedStorefrontLine is the minimal cart-safe projection needed to persist
// a Commerce cart line after the backend has resolved the product server-side.
type ApprovedStorefrontLine struct {
	ProductID string                    `json:"product_id"`
	Product   ApprovedStorefrontProduct `json:"product"`
	Price     common.Money              `json:"price"`
}
