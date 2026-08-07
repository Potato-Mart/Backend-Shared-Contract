package event

import (
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	analytics "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/insights/analytics"
	"time"
)

// RefundRequestedEvent is emitted on the refund-events topic when a refund is
// raised against an order's captured payments.
type RefundRequestedEvent struct {
	RefundID    string       `json:"refund_id"`
	OrderID     string       `json:"order_id,omitempty"`
	OrderNumber string       `json:"order_number"`
	Amount      common.Money `json:"amount"`
	RequestedBy string       `json:"requested_by,omitempty"`
	RequestedAt time.Time    `json:"requested_at"`
	Reason      string       `json:"reason,omitempty"`
	RequestID   string       `json:"request_id,omitempty"`
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
	Amount                  common.Money               `json:"amount"`
	Items                   []analytics.RefundItemFact `json:"items,omitempty"`
	RetailCustomerNumber    string                     `json:"retail_customer_number,omitempty"`
	OrganisationAccessID    string                     `json:"organisation_access_id,omitempty"`
	BenefitReservationID    string                     `json:"benefit_reservation_id,omitempty"`
	GiftCardRefundAmount    *common.Money              `json:"gift_card_refund_amount,omitempty"`
	QualifyingSpendReversal *common.Money              `json:"qualifying_spend_reversal,omitempty"`
	PointsToRestore         int                        `json:"points_to_restore,omitempty"`
	FullOrderRefund         bool                       `json:"full_order_refund,omitempty"`
	CompletedAt             time.Time                  `json:"completed_at"`
	RequestID               string                     `json:"request_id,omitempty"`
}

// RefundFailedEvent is emitted on the refund-events topic when a refund
// attempt terminally fails.
type RefundFailedEvent struct {
	RefundID    string       `json:"refund_id"`
	OrderID     string       `json:"order_id,omitempty"`
	OrderNumber string       `json:"order_number"`
	Amount      common.Money `json:"amount,omitempty"`
	Reason      string       `json:"reason,omitempty"`
	FailedAt    time.Time    `json:"failed_at"`
	RequestID   string       `json:"request_id,omitempty"`
}
