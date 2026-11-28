package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/payment/payment_enums"
)

// RefundAlternativeAnsweredEvent is emitted on the refund-events topic, with
// event type refund.alternative_answered v1, when staff record the customer's
// answer to an alternative offer on a pending refund request. It is published
// before any refund is executed and never announces money movement.
//
// It is the counterpart of RefundAlternativeEvent, which announces the offer.
// Revision is that offer's proposal revision, so an answer is matched to the
// offer it answers: a consumer that has already seen a higher revision for the
// same RefundRequestID must treat this answer as obsolete. Decision is the
// recorded outcome; a rejected or withdrawn offer leaves the refund request
// open, while refund.request_closed announces a request that ends without a
// refund.
//
// OrderNumber is opaque, including custom external order numbers; consumers
// must not classify it by prefix. AnsweredBy is the staff actor who recorded
// the answer, never the customer. MarketCode and CountryCode are the
// denormalized geography the event belongs to. Empty values provide no
// geographic evidence; a consumer that persists a geographically scoped record
// must fail closed rather than defaulting them.
type RefundAlternativeAnsweredEvent struct {
	RefundRequestID string                                  `json:"refund_request_id"`
	OrderNumber     string                                  `json:"order_number"`
	Revision        int                                     `json:"revision"`
	Decision        payment_enums.RefundAlternativeDecision `json:"decision"`
	AnsweredBy      string                                  `json:"answered_by,omitempty"`
	AnsweredAt      time.Time                               `json:"answered_at"`
	MarketCode      string                                  `json:"market_code,omitempty"`
	CountryCode     geography.CountryCode                   `json:"country_code,omitempty"`
}
