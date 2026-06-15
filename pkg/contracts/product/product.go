package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/enums"
)

type Product struct {
	ID              string                 `json:"id"`
	Code            string                 `json:"code"`
	SKU             string                 `json:"sku"`
	Name            string                 `json:"name"`
	Price           common.Money           `json:"price"`
	POSPrice        *common.Money          `json:"pos_price,omitempty"`
	Barcode         string                 `json:"barcode,omitempty"`
	OtherNames      []common.LocalizedName `json:"other_names,omitempty"`
	Brand           string                 `json:"brand,omitempty"`
	Vendor          string                 `json:"vendor,omitempty"`
	Catalogue       string                 `json:"catalogue,omitempty"`
	Storage         enums.StorageType      `json:"storage,omitempty"`
	Status          string                 `json:"status,omitempty"`
	FreshnessStatus string                 `json:"freshness_status,omitempty"`
	Dimensions      *common.Dimensions     `json:"dimensions,omitempty"`
	Weight          *common.Weight         `json:"weight,omitempty"`
	CurrentStock    int                    `json:"current_stock"`
	AvgWeeklySales  float64                `json:"avg_weekly_sales,omitempty"`
	CoverURL        string                 `json:"cover_url,omitempty"`
	ImageURLs       []string               `json:"image_urls,omitempty"`

	// ── Merchandising category & tags (additive, v5.2.0) ──────────────
	// CategoryKey references the canonical node of the ops product-
	// category tree (ops_product_categories). CategoryPath is the
	// denormalised root→leaf key chain (leaf last, includes CategoryKey)
	// maintained by Backend-Operations on every category assignment; it
	// exists so promotion targeting and storefront breadcrumbs never
	// need to re-walk the tree.
	CategoryKey  string   `json:"category_key,omitempty"`
	CategoryPath []string `json:"category_path,omitempty"`
	// CategoryTags are first-class merchandising labels (e.g. Hotpot,
	// Fresh Food). They complement — never replace — the canonical
	// category above, and stay separate from the SKU master.
	CategoryTags []CategoryTag `json:"category_tags,omitempty"`

	// ── Lifecycle (additive, v5.2.0; see lifecycle.go) ────────────────
	// FirstListedAt is set exactly once, when the product first becomes
	// publicly listed (status=active). It drives the NEW (新品) tag and
	// is never reset on delist/relist.
	FirstListedAt *time.Time `json:"first_listed_at,omitempty"`
	// RestockedAt is the most recent moment total sellable stock went
	// from 0 to >=1. It drives the RESTOCKED (補貨) tag and refreshes on
	// every such transition.
	RestockedAt *time.Time `json:"restocked_at,omitempty"`

	PlacingAreaCode string    `json:"placing_area_code,omitempty"`
	ExpiredAt       time.Time `json:"expired_at,omitempty"`

	common.AuditFields `bson:",inline"`
}
