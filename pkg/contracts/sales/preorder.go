package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/enums"
)

// PreorderStatus is the lifecycle of a durable customer preorder record.
// Endpoint command payloads remain owned by the service exposing the API.
type PreorderStatus string

const (
	PreorderStatusRequested PreorderStatus = "requested"
	PreorderStatusAccepted  PreorderStatus = "accepted"
	PreorderStatusRejected  PreorderStatus = "rejected"
	PreorderStatusCancelled PreorderStatus = "cancelled"
	PreorderStatusConverted PreorderStatus = "converted"
	PreorderStatusFulfilled PreorderStatus = "fulfilled"
	PreorderStatusExpired   PreorderStatus = "expired"
)

// Preorder records customer preorder interest for one product SKU. It is shared
// so Management, Operations, Commerce, and storefront consumers can agree on
// status, product identity, and customer-visible availability dates.
type Preorder struct {
	ID             string          `json:"id"`
	PreorderNumber string          `json:"preorder_number,omitempty"`
	Channel        enums.OrderType `json:"channel,omitempty"`
	Status         PreorderStatus  `json:"status"`
	Customer       common.PartyRef `json:"customer"`
	Buyer          *BuyerContext   `json:"buyer,omitempty"`

	ProductSKUCode string           `json:"product_sku_code"`
	Product        product.Snapshot `json:"product"`
	Quantity       int              `json:"quantity"`

	ExpectedAvailableAt  *time.Time   `json:"expected_available_at,omitempty"`
	RequestedAt          time.Time    `json:"requested_at"`
	AcceptedAt           *time.Time   `json:"accepted_at,omitempty"`
	RejectedAt           *time.Time   `json:"rejected_at,omitempty"`
	CancelledAt          *time.Time   `json:"cancelled_at,omitempty"`
	ConvertedAt          *time.Time   `json:"converted_at,omitempty"`
	FulfilledAt          *time.Time   `json:"fulfilled_at,omitempty"`
	ExpiredAt            *time.Time   `json:"expired_at,omitempty"`
	ConvertedOrderNumber string       `json:"converted_order_number,omitempty"`
	SourceDevice         SourceDevice `json:"source_device,omitempty"`

	CustomerNote string `json:"customer_note,omitempty"`
	InternalNote string `json:"internal_note,omitempty"`
	ReasonCode   string `json:"reason_code,omitempty"`

	History []shared.HistoryEntry `json:"history,omitempty"`

	common.AuditFields
}

// PreorderSummary is a customer-facing projection for preorder history and
// profile checks. It intentionally omits audit actors and internal notes.
type PreorderSummary struct {
	PreorderNumber       string           `json:"preorder_number"`
	Status               PreorderStatus   `json:"status"`
	Channel              enums.OrderType  `json:"channel,omitempty"`
	ProductSKUCode       string           `json:"product_sku_code"`
	Product              product.Snapshot `json:"product"`
	Quantity             int              `json:"quantity"`
	ExpectedAvailableAt  *time.Time       `json:"expected_available_at,omitempty"`
	RequestedAt          time.Time        `json:"requested_at"`
	UpdatedAt            time.Time        `json:"updated_at"`
	ConvertedOrderNumber string           `json:"converted_order_number,omitempty"`
}
