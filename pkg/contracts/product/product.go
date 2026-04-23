package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/pkg/enums"
)

type Product struct {
	ID             string            `json:"id"`
	SKU            string            `json:"category"`
	Name           string            `json:"name"`
	EnName         string            `json:"en_name,omitempty"`
	Brand          string            `json:"brand,omitempty"`
	Catalogue      string            `json:"catalogue,omitempty"`
	Storage        enums.StorageType `json:"storage,omitempty"`
	Price          float64           `json:"price"`
	POSPrice       float64           `json:"pos_price,omitempty"`
	Status         string            `json:"status,omitempty"`
	CurrentStock   int               `json:"current_stock,omitempty"`
	AvgWeeklySales float64           `json:"avg_weekly_sales,omitempty"`
	ImageURL       string            `json:"image_url,omitempty"`
	ImageURLs      []string          `json:"image_urls,omitempty"`
	CreatedAt      time.Time         `json:"created_at,omitempty"`
	UpdatedAt      time.Time         `json:"updated_at,omitempty"`
}
