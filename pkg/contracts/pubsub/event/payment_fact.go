package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
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
	// belongs to. Empty values provide no geographic evidence; a consumer
	// that persists a geographically scoped record must fail closed rather
	// than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	OccurredAt  time.Time             `json:"occurred_at"`
}
