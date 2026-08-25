package classification

import (
	"time"
)

// FavouriteListProduct records product membership without cart quantities.
type FavouriteListProduct struct {
	SKUCode string    `json:"sku_code"`
	AddedAt time.Time `json:"added_at"`
}
