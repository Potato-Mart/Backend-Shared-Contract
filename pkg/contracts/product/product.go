package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v2/pkg/enums"
)

type Product struct {
	ID              string                `json:"id"`
	Code            string                `json:"code"`
	SKU             string                `json:"sku"`
	Barcode         string                `json:"barcode"`
	Name            string                `json:"name"`
	OtherNames      []string              `json:"other_names,omitempty"`
	Brand           string                `json:"brand,omitempty"`
	Catalogue       string                `json:"catalogue,omitempty"`
	Storage         enums.StorageType     `json:"storage,omitempty"`
	Price           float64               `json:"price"`
	POSPrice        float64               `json:"pos_price,omitempty"`
	Status          string                `json:"status,omitempty"`
	FreshnessStatus enums.FreshnessStatus `json:"freshness_status,omitempty"`
	CurrentStock    int                   `json:"current_stock"`
	AvgWeeklySales  float64               `json:"avg_weekly_sales,omitempty"`
	CoverURL        string                `json:"cover_url,omitempty"`
	ImageURLs       []string              `json:"image_urls,omitempty"`
	PlacingAreaCode string                `json:"placing_area_code,omitempty"`
	ExpiredAt       time.Time             `json:"expired_at,omitempty"`
	CreatedAt       time.Time             `json:"created_at,omitempty"`
	UpdatedAt       time.Time             `json:"updated_at,omitempty"`
}
