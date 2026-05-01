package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/common"
)

type Cart struct {
	ID         string       `json:"id"`
	SessionID  string       `json:"session_id"`
	CustomerID string       `json:"customer_id,omitempty"`
	Items      []CartItem   `json:"items"`
	CouponCode string       `json:"coupon_code,omitempty"`
	Subtotal   common.Money `json:"subtotal"`
	ExpiresAt  time.Time    `json:"expires_at"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

type CartItem struct {
	Product  CartProduct  `json:"product"`
	Price    common.Money `json:"price"`
	Quantity int          `json:"quantity"`
}

type CartProduct struct {
	ID       string `json:"id,omitempty"`
	SKU      string `json:"sku,omitempty"`
	Name     string `json:"name"`
	Brand    string `json:"brand,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	Storage  string `json:"storage,omitempty"`
}
