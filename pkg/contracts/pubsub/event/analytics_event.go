package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	analytics "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/insights/analytics"
)

// OrderFact is the immutable analytical projection of an order event.
type OrderFact struct {
	EventID              string                    `json:"event_id"`
	OrderNumber          string                    `json:"order_number"`
	RetailCustomerNumber string                    `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                    `json:"organisation_access_id,omitempty"`
	Status               string                    `json:"status"`
	Channel              string                    `json:"channel,omitempty"`
	ItemCount            int                       `json:"item_count"`
	Total                money.Money               `json:"total"`
	Items                []analytics.OrderItemFact `json:"items,omitempty"`
	OccurredAt           time.Time                 `json:"occurred_at"`
}

// PaymentFact is the immutable analytical projection of a payment event.
type PaymentFact struct {
	EventID     string      `json:"event_id"`
	PaymentID   string      `json:"payment_id"`
	OrderNumber string      `json:"order_number"`
	Method      string      `json:"method,omitempty"`
	Status      string      `json:"status"`
	Amount      money.Money `json:"amount"`
	OccurredAt  time.Time   `json:"occurred_at"`
}

// RefundFact is the immutable analytical projection of a refund event.
type RefundFact struct {
	EventID     string                     `json:"event_id"`
	RefundID    string                     `json:"refund_id"`
	OrderNumber string                     `json:"order_number"`
	Status      string                     `json:"status"`
	Amount      money.Money                `json:"amount"`
	Items       []analytics.RefundItemFact `json:"items,omitempty"`
	OccurredAt  time.Time                  `json:"occurred_at"`
}
