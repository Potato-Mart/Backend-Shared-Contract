package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

type Depot struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Address   string    `json:"address,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	IsActive  bool      `json:"is_active"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostcodeRule struct {
	ID       string `json:"id"`
	DepotID  string `json:"depot_id"`
	Postcode string `json:"postcode"`
	Priority int    `json:"priority"`
}

type DepotProduct struct {
	DepotID      string    `json:"depot_id"`
	ProductID    string    `json:"product_id"`
	StockQty     int       `json:"stock_qty"`
	IsAvailable  bool      `json:"is_available"`
	LocationCode string    `json:"location_code,omitempty"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type StockLocation struct {
	ID       string            `json:"id"`
	DepotID  string            `json:"depot_id"`
	Code     string            `json:"code"`
	Name     string            `json:"name,omitempty"`
	Zone     enums.StorageType `json:"zone,omitempty"`
	IsActive bool              `json:"is_active"`
}
