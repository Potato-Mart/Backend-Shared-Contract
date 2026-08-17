package product

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/localization"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/product/product_enums"
)

// Product is the global conceptual product: what the item is, independently of
// any market, currency, channel, or depot. It carries content, origin,
// classification, provenance, lifecycle status, and administration only.
//
// A Product holds no price, no tax flag, no channel sellability, and no
// market-dependent metric. Sellable identity lives on SKU, market availability
// on listing.MarketListing, and authoritative prices in Pricing price books.
type Product struct {
	ID             string                        `json:"id"`
	Status         product_enums.ProductStatus   `json:"status"`
	Content        ProductContent                `json:"content"`
	Classification ProductClassification         `json:"classification"`
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
	ProductCategoryCode string                          `json:"product_category_code"`
	BrandRef            *classification.BrandRef        `json:"brand_ref,omitempty"`
	CollectionRef       *classification.CollectionRef   `json:"collection_ref,omitempty"`
	CategoryTags        []classification.CategoryTagRef `json:"category_tags,omitempty"`
}

// ProductAdministration retains master-data history and audit information.
// Storefront and POS responses omit this optional component.
type ProductAdministration struct {
	History []security.HistoryEntry `json:"history,omitempty"`

	audit.AuditFields
}

// Images groups the customer-facing product imagery.
type Images struct {
	Cover   *security.ObjectMedia  `json:"cover,omitempty"`
	Gallery []security.ObjectMedia `json:"gallery,omitempty"`
	Details []security.ObjectMedia `json:"details,omitempty"`
}
