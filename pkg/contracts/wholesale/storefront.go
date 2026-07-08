package wholesale

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/contracts/product"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/enums/product"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/enums/warehouse"
)

const (
	// PathApprovedStorefrontProducts is the authenticated wholesale storefront
	// product endpoint for approved buyers.
	PathApprovedStorefrontProducts = "/v1/storefront/wholesale/approved/products"
	// PathInternalApprovedStorefrontProducts is the service-auth product
	// resolver Commerce uses to reprice wholesale cart lines server-side.
	PathInternalApprovedStorefrontProducts = "/v1/internal/wholesale/products"
)

// ApprovedStorefrontProduct is the authenticated wholesale product projection
// for approved B2B buyers. It intentionally exposes one effective wholesale
// price, not the product.Pricing object, and omits operational stock quantity.
type ApprovedStorefrontProduct struct {
	ID                 string                        `json:"id"`
	SKUCode            string                        `json:"sku_code,omitempty"`
	SKU                string                        `json:"sku,omitempty"`
	Name               string                        `json:"name"`
	Description        []common.LocalizedDescription `json:"description,omitempty"`
	Barcode            string                        `json:"barcode,omitempty"`
	OtherNames         []common.LocalizedName        `json:"other_names,omitempty"`
	Brand              []common.LocalizedName        `json:"brand,omitempty"`
	Supplier           string                        `json:"supplier,omitempty"`
	BrandNames         []common.LocalizedName        `json:"brand_names,omitempty"`
	Storage            warehouseenum.StorageType     `json:"storage,omitempty"`
	SalesPerformance   productenum.SalesPerformance  `json:"sales_performance,omitempty"`
	DisplayStatus      string                        `json:"display_status,omitempty"`
	CoverURL           string                        `json:"cover_url,omitempty"`
	ImageURLs          []string                      `json:"image_urls,omitempty"`
	Collection         *product.CollectionRef        `json:"collection,omitempty"`
	CategoryTags       []product.CategoryTag         `json:"category_tags,omitempty"`
	LifecycleTags      []string                      `json:"lifecycle_tags,omitempty"`
	NewExpiresAt       *time.Time                    `json:"new_expires_at,omitempty"`
	RestockedExpiresAt *time.Time                    `json:"restocked_expires_at,omitempty"`
	StorefrontDisplay  *product.StorefrontDisplay    `json:"storefront_display,omitempty"`

	Price             common.Money `json:"price"`
	StockAvailable    bool         `json:"stock_available"`
	AvailabilityState string       `json:"availability_state"`
}

// ApprovedStorefrontLine is the minimal cart-safe projection needed to persist
// a Commerce cart line after the backend has resolved the product server-side.
type ApprovedStorefrontLine struct {
	ProductSKUCode string                    `json:"product_sku_code"`
	Product        ApprovedStorefrontProduct `json:"product"`
	Price          common.Money              `json:"price"`
}
