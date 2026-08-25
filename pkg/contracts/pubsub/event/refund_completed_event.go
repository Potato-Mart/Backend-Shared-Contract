package event

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/money"
	analytics "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/insights/analytics"
)

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
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. Empty values provide no geographic evidence; a consumer
	// that persists a geographically scoped record must fail closed rather
	// than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	CompletedAt time.Time             `json:"completed_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
