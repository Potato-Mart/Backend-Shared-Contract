package payments

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/money"
	analytics "github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/insights/sales"
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/pricing/membership"
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
	BenefitReservationID string                     `json:"benefit_reservation_id,omitempty"`
	GiftCardRefundAmount *money.Money               `json:"gift_card_refund_amount,omitempty"`
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
	// MarketCode and CountryCode are the denormalized geography the event
	// belongs to. Empty values provide no geographic evidence; a consumer
	// that persists a geographically scoped record must fail closed rather
	// than defaulting them.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	CompletedAt time.Time             `json:"completed_at"`
	RequestID   string                `json:"request_id,omitempty"`
}
