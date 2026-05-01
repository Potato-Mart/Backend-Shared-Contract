package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
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
	PlacingAreaCode string                 `json:"placing_area_code,omitempty"`
	ExpiredAt       time.Time              `json:"expired_at,omitempty"`
	CreatedAt       time.Time              `json:"created_at,omitempty"`
	UpdatedAt       time.Time              `json:"updated_at,omitempty"`
}
