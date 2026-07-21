package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/contracts/shared"
	customerenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/customer"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/product"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/sales"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/warehouse"
)

// Product is the master record for one sellable unit (v7.0.0).
//
// The struct is split into a small set of FLAT top-level fields -- the
// hot identity, filter, and sort keys that stay directly indexable for
// fast fetch/sort/search -- and a handful of nested GROUP structs that
// keep the remaining attributes tidy (pricing, localization, media,
// physical). Keep new descriptive attributes
// inside a group; only promote a field to the top level when it must be
// indexed or filtered hot.
type Product struct {
	ID          string                        `json:"id"`
	SKUCode     string                        `json:"sku_code"`
	SKU         string                        `json:"sku"`
	Name        string                        `json:"name"`
	Description []common.LocalizedDescription `json:"description,omitempty"`
	Brand       []common.LocalizedName        `json:"brand,omitempty"`
	BrandRef    *BrandRef                     `json:"brand_ref,omitempty"`
	Barcode     string                        `json:"barcode,omitempty"`
	Taxed       bool                          `json:"taxed"`

	// Storage is the physical storage zone (DRY/CHILLED/FROZEN)
	Storage warehouseenum.StorageType `json:"storage,omitempty"`
	// Status is the admin-controlled lifecycle state (draft/active/
	// archived/discontinued). Derived runtime states (new/restocked/
	// out_of_stock) are computed by DisplayStatus, never stored here.
	Status          productenum.ProductStatus `json:"status,omitempty"`
	Collection      *CollectionRef            `json:"collection,omitempty"`
	CategoryTags    []CategoryTag             `json:"category_tags,omitempty"`
	SupplierCode    string                    `json:"supplier_code,omitempty"`
	Supply          *ProductSupply            `json:"supply,omitempty"`
	PlacingAreaCode string                    `json:"placing_area_code,omitempty"`

	// CurrentStock is a denormalised cache of total sellable stock; the
	// authoritative quantities live in the warehouse subsystem. It backs
	// the out_of_stock display state.
	CurrentStock int `json:"current_stock"`
	// SalesPerformance contains backend-computed historical sales statistics.
	// It is never manually authored on product create/update operations.
	SalesPerformance *SalesPerformanceStats `json:"sales_performance,omitempty"`
	// DisplaySellingCount is an optional manually curated storefront count.
	// A non-nil pointer with value zero is distinct from an absent count.
	DisplaySellingCount *int64 `json:"display_selling_count,omitempty"`

	// FirstListedAt is set exactly once, when the product first becomes
	// active. It anchors the NEW (新品) 14-day window and is never reset
	// on delist/relist. RestockedAt is the most recent 0->(>=1) sellable-
	// stock transition; it refreshes the RESTOCKED (補貨) window.
	FirstListedAt *time.Time `json:"first_listed_at,omitempty"`
	RestockedAt   *time.Time `json:"restocked_at,omitempty"`
	ExpiredAt     time.Time  `json:"expired_at,omitempty"`

	// ── Nested groups ─────────────────────────────────────────────────
	Pricing Pricing `json:"pricing"`
	// Selling is an optional pointer (unlike the always-present groups
	// below) so a product with no channel/buyer sellability rules omits it.
	Selling      *Selling     `json:"selling,omitempty"`
	Localization Localization `json:"localization,omitempty"`
	Media        Media        `json:"media,omitempty"`
	Physical     Physical     `json:"physical,omitempty"`
	// StorefrontMerchandising carries admin-managed retail display policy.
	// Backend read models convert it into customer-safe display fields.
	StorefrontMerchandising *StorefrontMerchandising `json:"storefront_merchandising,omitempty"`

	// History is for product master-data changes only. Stock changes are
	// represented as warehouse.StockMovement records.
	History []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields
}

// Pricing groups the prices a product can carry. Every field is an
// optional *common.Money so each price can be absent and can even carry
// its own currency (international trading). Online is the canonical
// storefront price and is the one indexed for price sort
// (pricing.online.amount_minor).
type Pricing struct {
	Online    *common.Money `json:"online,omitempty"`    // online selling price
	Offline   *common.Money `json:"offline,omitempty"`   // offline / POS selling price
	Original  *common.Money `json:"original,omitempty"`  // original / RRP ("was") price
	Tag       *common.Money `json:"tag,omitempty"`       // temporary e-tag price
	Wholesale *common.Money `json:"wholesale,omitempty"` // wholesale price
	Cost      *common.Money `json:"cost,omitempty"`      // purchase cost
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

// Localization groups secondary per-language display fields. Product
// Description and Brand are the canonical localized display fields; these
// slices remain for compatibility with older clients that read localization
// as a nested group.
type Localization struct {
	OtherNames   []common.LocalizedName        `json:"other_names,omitempty"`
	BrandNames   []common.LocalizedName        `json:"brand_names,omitempty"`
	Descriptions []common.LocalizedDescription `json:"descriptions,omitempty"`
}

// Media groups the product imagery.
type Media struct {
	CoverURL     string        `json:"cover_url,omitempty"`
	ImageURLs    []string      `json:"image_urls,omitempty"`
	DetailImages []DetailImage `json:"detail_images,omitempty"`
}

// Physical groups the packaged-good physical attributes.
type Physical struct {
	Dimensions *common.Dimensions `json:"dimensions,omitempty"`
	Weight     *common.Weight     `json:"weight,omitempty"`
}
