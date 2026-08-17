package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/money"
	analytics "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/insights/analytics"
)

// RefundRequestedEvent is emitted on the refund-events topic when a refund is
// raised against an order's captured payments.
type RefundRequestedEvent struct {
	RefundID    string      `json:"refund_id"`
	OrderID     string      `json:"order_id,omitempty"`
	OrderNumber string      `json:"order_number"`
	Amount      money.Money `json:"amount"`
	RequestedBy string      `json:"requested_by,omitempty"`
	// MarketID and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketID    string                `json:"market_id,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	RequestedAt time.Time             `json:"requested_at"`
	Reason      string                `json:"reason,omitempty"`
	RequestID   string                `json:"request_id,omitempty"`
}

// RefundCompletedEvent is emitted on the refund-events topic when a refund
// settles. Consumers reverse benefits/points and update order payment state.
// The optional restoration fields carry everything the pricing owner needs to
// reverse checkout benefits and points without a synchronous read-back.
type RefundCompletedEvent struct {
	RefundID                string                     `json:"refund_id"`
	OrderID                 string                     `json:"order_id,omitempty"`
	OrderNumber             string                     `json:"order_number"`
	PaymentID               string                     `json:"payment_id,omitempty"`
	Amount                  money.Money                `json:"amount"`
	Items                   []analytics.RefundItemFact `json:"items,omitempty"`
	RetailCustomerNumber    string                     `json:"retail_customer_number,omitempty"`
	OrganisationAccessID    string                     `json:"organisation_access_id,omitempty"`
	BenefitReservationID    string                     `json:"benefit_reservation_id,omitempty"`
	GiftCardRefundAmount    *money.Money               `json:"gift_card_refund_amount,omitempty"`
	QualifyingSpendReversal *money.Money               `json:"qualifying_spend_reversal,omitempty"`
	PointsToRestore         int                        `json:"points_to_restore,omitempty"`
	FullOrderRefund         bool                       `json:"full_order_refund,omitempty"`
	// MarketID and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketID    string                `json:"market_id,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	CompletedAt time.Time             `json:"completed_at"`
	RequestID   string                `json:"request_id,omitempty"`
}

// RefundFailedEvent is emitted on the refund-events topic when a refund
// attempt terminally fails.
type RefundFailedEvent struct {
	RefundID    string      `json:"refund_id"`
	OrderID     string      `json:"order_id,omitempty"`
	OrderNumber string      `json:"order_number"`
	Amount      money.Money `json:"amount,omitempty"`
	Reason      string      `json:"reason,omitempty"`
	// MarketID and CountryCode are the denormalized geography the event
	// belongs to. They are absent on every event published before v28.0.0;
	// a consumer that persists a geographically scoped record treats an
	// absent value as "no evidence" and fails closed rather than defaulting.
	MarketID    string                `json:"market_id,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	FailedAt    time.Time             `json:"failed_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
