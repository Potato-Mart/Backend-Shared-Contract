package orders

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	analytics "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/insights/sales"
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
	// MarketCode, DepotCode, and CountryCode are the denormalized geography
	// the event belongs to. Empty values provide no geographic evidence; a
	// consumer that persists a geographically scoped record must fail closed
	// rather than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	DepotCode   string                `json:"depot_code,omitempty"`
	OccurredAt  time.Time             `json:"occurred_at"`
}
