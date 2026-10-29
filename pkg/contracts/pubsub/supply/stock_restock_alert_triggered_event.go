package supply

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/localization"
)

// StockRestockAlertTriggeredEvent represents one customer's fired restock
// alert for one SKU in one market. AlertID is its aggregate and idempotency
// key; the subscription clears when the alert fires.
type StockRestockAlertTriggeredEvent struct {
	AlertID              string                       `json:"alert_id"`
	SKUCode              string                       `json:"sku_code"`
	MarketCode           string                       `json:"market_code"`
	UserID               string                       `json:"user_id"`
	RetailCustomerNumber string                       `json:"retail_customer_number"`
	ProductName          []localization.LocalizedName `json:"product_name"`
	SubscribedAt         time.Time                    `json:"subscribed_at"`
	ReturnedAt           time.Time                    `json:"returned_at"`
}
