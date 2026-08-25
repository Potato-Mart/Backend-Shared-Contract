package product

import "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/classification"

// ProductClassification contains code-only relationships to catalogue
// masters. SellingProduct resolves them into customer-safe display data.
type ProductClassification struct {
	SKUSeriesCode string                          `json:"sku_series_code"`
	Brands        []classification.BrandRef       `json:"brands,omitempty"`
	CollectionRef *classification.CollectionRef   `json:"collection_ref,omitempty"`
	CategoryTags  []classification.CategoryTagRef `json:"category_tags,omitempty"`
}
