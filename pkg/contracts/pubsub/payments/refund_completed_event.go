package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	analytics "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/sales"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/membership"
)

// RefundCompletedEvent is emitted on the refund-events topic when a refund
// settles. Consumers reverse benefits/points and update order payment state.
// The optional restoration fields carry everything the pricing owner needs to
// reverse checkout benefits and points without a synchronous read-back. Tier
// reversal additionally requires Pricing-owned native proof verification.
type RefundCompletedEvent struct {
	RefundID             string                     `json:"refund_id"`
	OrderID              string                     `json:"order_id,omitempty"`
	OrderNumber          string                     `json:"order_number"`
	PaymentID            string                     `json:"payment_id,omitempty"`
	Amount               money.Money                `json:"amount"`
	Items                []analytics.RefundItemFact `json:"items,omitempty"`
	RetailCustomerNumber string                     `json:"retail_customer_number,omitempty"`
	OrganisationAccessID string                     `json:"organisation_access_id,omitempty"`
	// BenefitReservationID and GiftCardRefundAmount are the single-reservation
	// gift card restoration that events before v53.54.0 carry. When
	// GiftCardRefundParts is present it is authoritative, and these two fields
	// mirror its first part. Consumers that read the
	// parts must ignore these fields when parts are present.
	BenefitReservationID string       `json:"benefit_reservation_id,omitempty"`
	GiftCardRefundAmount *money.Money `json:"gift_card_refund_amount,omitempty"`
	// GiftCardRefundParts splits the money put back on gift cards used to pay
	// by benefit reservation, one part per reservation the refund restored.
	// Each part has a non-empty reservation id and a positive amount, all
	// parts share one currency, and no reservation appears twice. Omitted
	// when the refund restored no gift card money or Payments recorded only
	// the single fields above.
	GiftCardRefundParts []GiftCardRefundPart `json:"gift_card_refund_parts,omitempty"`
	// QualifyingSpendReversal retains the released gross earned-point clawback
	// evidence. Never reinterpret it as the NET tier-refund amount.
	QualifyingSpendReversal *money.Money `json:"qualifying_spend_reversal,omitempty"`
	// TierRefundProof is optional Pricing-verified NET tier evidence. Absence is
	// missing proof, never verified zero or nonretail classification. Service-owned
	// protected applicability decisions may omit it; retail without authority stays
	// pending. Existing points/benefit receipt legs remain independently durable.
	TierRefundProof *membership.TierRefundProof `json:"tier_refund_proof,omitempty"`
	PointsToRestore int                         `json:"points_to_restore,omitempty"`
	FullOrderRefund bool                        `json:"full_order_refund,omitempty"`
	// PaymentAllocations splits the refund across the order's payments, one
	// row per payment the refund returned money through, for example each
	// part of a POS split tender. Money put back on gift cards used to pay
	// stays in the gift card restoration fields and is not listed, so the rows
	// sum to Amount less the gift card restoration: the sum of
	// GiftCardRefundParts when present, otherwise GiftCardRefundAmount (zero
	// when absent). PaymentID keeps its existing meaning. Omitted when
	// Payments did not record the split.
	PaymentAllocations []RefundPaymentAllocation `json:"payment_allocations,omitempty"`
	// GiftCardReversals lists the sold gift cards this refund reversed, one
	// per refunded gift-card line. It is omitted when the refund covered no
	// sold gift card. It is not GiftCardRefundAmount, which is money put back
	// on gift cards used to pay.
	GiftCardReversals []RefundGiftCardReversal `json:"gift_card_reversals,omitempty"`
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. Empty values provide no geographic evidence; a consumer
	// that persists a geographically scoped record must fail closed rather
	// than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	CompletedAt time.Time             `json:"completed_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
