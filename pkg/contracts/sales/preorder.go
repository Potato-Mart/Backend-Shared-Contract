package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/notification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/contracts/shared"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/sales"
)

// Preorder records customer preorder interest for one product SKU. It is shared
// so Management, Operations, Commerce, and storefront consumers can agree on
// status, product identity, and customer-visible availability dates.
type Preorder struct {
	ID             string                   `json:"id"`
	PreorderNumber string                   `json:"preorder_number,omitempty"`
	Channel        salesenum.OrderType      `json:"channel,omitempty"`
	Status         salesenum.PreorderStatus `json:"status"`
	Customer       common.PartyRef          `json:"customer"`
	Buyer          *BuyerContext            `json:"buyer,omitempty"`

	ProductSKUCode string           `json:"product_sku_code"`
	Product        product.Snapshot `json:"product"`
	Quantity       int              `json:"quantity"`

	ExpectedAvailableAt  *time.Time            `json:"expected_available_at,omitempty"`
	RequestedAt          time.Time             `json:"requested_at"`
	AcceptedAt           *time.Time            `json:"accepted_at,omitempty"`
	RejectedAt           *time.Time            `json:"rejected_at,omitempty"`
	CancelledAt          *time.Time            `json:"cancelled_at,omitempty"`
	ConvertedAt          *time.Time            `json:"converted_at,omitempty"`
	FulfilledAt          *time.Time            `json:"fulfilled_at,omitempty"`
	ExpiredAt            *time.Time            `json:"expired_at,omitempty"`
	ConvertedOrderNumber string                `json:"converted_order_number,omitempty"`
	Availability         *PreorderAvailability `json:"availability,omitempty"`
	SourceDevice         SourceDevice          `json:"source_device,omitempty"`

	CustomerNote string `json:"customer_note,omitempty"`
	InternalNote string `json:"internal_note,omitempty"`
	ReasonCode   string `json:"reason_code,omitempty"`

	History []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields
}

// PreorderAvailability is the durable readiness and customer-notification
// projection for a paid preorder line. Commerce owns and persists it; Admin
// and storefront consumers can safely render it after a page refresh.
type PreorderAvailability struct {
	StockReserved       bool                                      `json:"stock_reserved"`
	ReservationIDs      []string                                  `json:"reservation_ids,omitempty"`
	ReservedAt          *time.Time                                `json:"reserved_at,omitempty"`
	NotificationEventID string                                    `json:"notification_event_id,omitempty"`
	Notification        *notification.NotificationDeliveryReceipt `json:"notification,omitempty"`
}

// PreorderSummary is a customer-facing projection for preorder history and
// profile checks. It intentionally omits audit actors and internal notes.
type PreorderSummary struct {
	PreorderNumber       string                   `json:"preorder_number"`
	Status               salesenum.PreorderStatus `json:"status"`
	Channel              salesenum.OrderType      `json:"channel,omitempty"`
	ProductSKUCode       string                   `json:"product_sku_code"`
	Product              product.Snapshot         `json:"product"`
	Quantity             int                      `json:"quantity"`
	ExpectedAvailableAt  *time.Time               `json:"expected_available_at,omitempty"`
	RequestedAt          time.Time                `json:"requested_at"`
	UpdatedAt            time.Time                `json:"updated_at"`
	ConvertedOrderNumber string                   `json:"converted_order_number,omitempty"`
	Availability         *PreorderAvailability    `json:"availability,omitempty"`
}
