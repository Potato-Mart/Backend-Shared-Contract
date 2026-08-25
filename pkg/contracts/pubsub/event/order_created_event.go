package event

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	"time"
)

// OrderCreatedEvent is emitted on the order-events topic when a sales order
// is first recorded. AggregateID is the order number.
type OrderCreatedEvent struct {
	OrderID                   string                   `json:"order_id"`
	OrderNumber               string                   `json:"order_number"`
	OrderType                 commerce_enums.OrderType `json:"order_type,omitempty"`
	BuyerUserID               string                   `json:"buyer_user_id,omitempty"`
	RetailCustomerNumber      string                   `json:"retail_customer_number,omitempty"`
	WholesaleOrganisationCode string                   `json:"wholesale_organisation_code,omitempty"`
	// MarketCode, DepotCode, and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	DepotCode   string                `json:"depot_code,omitempty"`
	Total       money.Money           `json:"total"`
	CreatedAt   time.Time             `json:"created_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
