package orders

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v38/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v38/pkg/contracts/common/money"
)

// OrderEditedEvent is the completed fact of a committed product amendment on
// the order-events topic (event_version v2). EditID is the stable business
// dedupe key for one customer notification. RevisedItems is the whole requested
// order composition so Supply can invalidate old physical work and recheck the
// new generation; publication does not itself authorize dispatch or payment.
type OrderEditedEvent struct {
	OrderID                       string `json:"order_id"`
	OrderNumber                   string `json:"order_number"`
	EditID                        string `json:"edit_id"`
	PreviousFulfillmentGeneration int64  `json:"previous_fulfillment_generation"`
	FulfillmentGeneration         int64  `json:"fulfillment_generation"`
	RetailCustomerNumber          string `json:"retail_customer_number,omitempty"`
	OrganisationAccessID          string `json:"organisation_access_id,omitempty"`
	// Empty geography provides no scope evidence; consumers persisting a
	// geographically scoped record must fail closed rather than infer it.
	MarketCode    string                    `json:"market_code,omitempty"`
	CountryCode   geography.CountryCode     `json:"country_code,omitempty"`
	DepotCode     string                    `json:"depot_code,omitempty"`
	PreviousItems []OrderEditedItemSnapshot `json:"previous_items"`
	RevisedItems  []OrderEditedItemSnapshot `json:"revised_items"`
	PreviousTotal money.Money               `json:"previous_total"`
	RevisedTotal  money.Money               `json:"revised_total"`
	CommittedAt   time.Time                 `json:"committed_at"`
}
