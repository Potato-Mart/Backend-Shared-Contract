package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/contracts/product"
)

type Cart struct {
	ID         string       `json:"id"`
	SessionID  string       `json:"session_id"`
	CustomerID string       `json:"customer_id,omitempty"`
	Items      []CartItem   `json:"items"`
	CouponCode string       `json:"coupon_code,omitempty"`
	Subtotal   common.Money `json:"subtotal"`
	ExpiresAt  time.Time    `json:"expires_at"`

	common.AuditFields `bson:",inline"`
}

type CartItem struct {
	Product  product.Snapshot `json:"product"`
	Price    common.Money     `json:"price"`
	Quantity int              `json:"quantity"`
}
