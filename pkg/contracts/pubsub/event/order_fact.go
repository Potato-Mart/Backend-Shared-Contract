package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	analytics "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/insights/analytics"
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
	// MarketCode, DepotCode, and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	DepotCode   string                `json:"depot_code,omitempty"`
	OccurredAt  time.Time             `json:"occurred_at"`
}
