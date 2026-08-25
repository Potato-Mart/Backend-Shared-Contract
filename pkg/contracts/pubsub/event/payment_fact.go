package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
)

// PaymentFact is the immutable analytical projection of a payment event.
type PaymentFact struct {
	EventID     string      `json:"event_id"`
	PaymentID   string      `json:"payment_id"`
	OrderNumber string      `json:"order_number"`
	Method      string      `json:"method,omitempty"`
	Status      string      `json:"status"`
	Amount      money.Money `json:"amount"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	OccurredAt  time.Time             `json:"occurred_at"`
}
