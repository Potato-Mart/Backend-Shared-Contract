package orders

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads"
	"time"
)

// OrderCancelledEvent is emitted on the order-events topic when an order is
// cancelled before fulfilment completes.
type OrderCancelledEvent struct {
	OrderID              string `json:"order_id"`
	OrderNumber          string `json:"order_number"`
	RetailCustomerNumber string `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string `json:"organisation_access_id,omitempty"`
	CancelledBy          string `json:"cancelled_by,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. Empty values provide no geographic evidence; a consumer
	// that persists a geographically scoped record must fail closed rather
	// than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	CancelledAt time.Time             `json:"cancelled_at"`
	Reason      string                `json:"reason,omitempty"`
	RequestID   string                `json:"request_id,omitempty"`
	// AdConversionWithdrawal is the reversal evidence an advertising consumer
	// needs to withdraw or restate a purchase it already reported for this
	// order. A cancellation retracts the whole order, so Full is true and only
	// the withdrawn revenue is carried. It is absent when the order was never
	// attributed to an ad or was never paid, so there is no conversion to
	// withdraw; absence is never a zero amount.
	AdConversionWithdrawal *ads.AdConversionWithdrawal `json:"ad_conversion_withdrawal,omitempty"`
}
