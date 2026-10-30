package possettlement

import (
	"time"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/receipt"
)

// Obligation is the immutable Orders-owned description of a paid POS sale.
// Version 2 binds the complete tender and exact reserved/direct source partition.
// It is a domain record, not permission to debit inventory: consumers separately
// authenticate its owner, generation, quote and durable tender holds. Digest and
// hold identities are deliberately absent from their own source representation.
type Obligation struct {
	Version               int32       `json:"version"`
	ObligationID          string      `json:"obligation_id"`
	SaleID                string      `json:"sale_id"`
	OrderNumber           string      `json:"order_number"`
	FulfillmentGeneration int64       `json:"fulfillment_generation"`
	CheckoutID            string      `json:"checkout_id"`
	QuoteID               string      `json:"quote_id"`
	QuoteKey              string      `json:"quote_key"`
	QuoteRevision         string      `json:"quote_revision"`
	QuoteFingerprint      string      `json:"quote_fingerprint"`
	MarketCode            string      `json:"market_code"`
	DepotCode             string      `json:"depot_code"`
	RegisterID            string      `json:"register_id"`
	SessionID             string      `json:"session_id"`
	OperatorID            string      `json:"operator_id"`
	PaidAt                time.Time   `json:"paid_at"`
	OrderTotal            money.Money `json:"order_total"`
	CouponDiscount        money.Money `json:"coupon_discount"`
	DiscountedPayable     money.Money `json:"discounted_payable"`
	// FinalPayable is after coupon and cash adjustment, before tenders.
	FinalPayable         money.Money                   `json:"final_payable"`
	ExactExternalDue     money.Money                   `json:"exact_external_due"`
	FinalExternalDue     money.Money                   `json:"final_external_due"`
	BenefitReservationID string                        `json:"benefit_reservation_id,omitempty"`
	CashRounding         *receipt.CashRoundingSnapshot `json:"cash_rounding,omitempty"`
	ZeroDue              bool                          `json:"zero_due"`
	Instruments          []Instrument                  `json:"instruments"`
	Lines                []Line                        `json:"lines"`
}
