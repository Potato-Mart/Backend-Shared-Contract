package event

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/order/order_enums"
	"time"
)

// OrderStatusChangedEvent is emitted on the order-events topic for every
// order lifecycle transition (processing, picking, packed, shipped, …).
// Buyer identity, tracking, and invoice references are optional enrichment
// so notification consumers need no synchronous read-back.
type OrderStatusChangedEvent struct {
	OrderID              string                       `json:"order_id"`
	OrderNumber          string                       `json:"order_number"`
	PreviousStatus       order_enums.SalesOrderStatus `json:"previous_status,omitempty"`
	Status               order_enums.SalesOrderStatus `json:"status"`
	RetailCustomerNumber string                       `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                       `json:"organisation_access_id,omitempty"`
	TrackingNumber       string                       `json:"tracking_number,omitempty"`
	InvoiceNumber        string                       `json:"invoice_number,omitempty"`
	ChangedBy            string                       `json:"changed_by,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. Empty values provide no geographic evidence; a consumer
	// that persists a geographically scoped record must fail closed rather
	// than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	ChangedAt   time.Time             `json:"changed_at"`
	Reason      string                `json:"reason,omitempty"`
	RequestID   string                `json:"request_id,omitempty"`
}
