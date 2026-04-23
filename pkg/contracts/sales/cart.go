package sales

import "time"

type Cart struct {
	ID         string     `json:"id"`
	SessionID  string     `json:"session_id"`
	CustomerID string     `json:"customer_id,omitempty"`
	Items      []CartItem `json:"items"`
	CouponCode string     `json:"coupon_code,omitempty"`
	Subtotal   float64    `json:"subtotal"`
	ExpiresAt  time.Time  `json:"expires_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type CartItem struct {
	Product  CartProduct `json:"product"`
	Price    float64     `json:"price"`
	Quantity int         `json:"quantity"`
}

type CartProduct struct {
	ID       string `json:"id,omitempty"`
	SKU      string `json:"category,omitempty"`
	Name     string `json:"name"`
	Brand    string `json:"brand,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	Storage  string `json:"storage,omitempty"`
}
