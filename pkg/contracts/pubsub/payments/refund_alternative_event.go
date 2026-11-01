package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
)

// RefundAlternativeEvent is emitted on the refund-events topic, with
// event type refund.alternative_offered v1, when staff record an alternative
// offer on a pending refund request. It is published before any refund is
// executed and never announces money movement.
//
// One event is published per proposal revision. A consumer must treat an offer
// as obsolete when it has seen a higher ProposalRevision for the same
// RefundRequestID, and must suppress queued delivery for it. Staff record the
// customer's response; there is no customer acceptance endpoint.
type RefundAlternativeEvent struct {
	RefundRequestID string `json:"refund_request_id"`
	// ProposalRevision starts at 1 and increases with every new offer on the
	// same refund request.
	ProposalRevision int `json:"proposal_revision"`
	// OrderNumber is opaque, including custom external order numbers;
	// consumers must not classify it by prefix.
	OrderNumber          string `json:"order_number"`
	CustomerID           string `json:"customer_id,omitempty"`
	RetailCustomerNumber string `json:"retail_customer_number,omitempty"`
	// Locale is the BCP 47 language tag the customer copy should use.
	Locale string                  `json:"locale,omitempty"`
	Lines  []RefundAlternativeLine `json:"lines"`
	// OriginalPayable and ProposedPayable share one currency.
	OriginalPayable money.Money `json:"original_payable"`
	ProposedPayable money.Money `json:"proposed_payable"`
	// ContactNote is staff-written guidance on how to reply to staff.
	ContactNote string                `json:"contact_note,omitempty"`
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	OfferedBy   string                `json:"offered_by,omitempty"`
	OfferedAt   time.Time             `json:"offered_at"`
}
