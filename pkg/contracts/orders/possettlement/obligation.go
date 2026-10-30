package possettlement

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/receipt"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/inventory"
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

// Instrument references the original authoritative debit, never a balance.
// Kind is external, cash, gift_card or points. Coupons are discounts, not tender.
// Optional original-debit locators are absent in historical draft representations.
type Instrument struct {
	Kind                 string      `json:"kind"`
	Authority            string      `json:"authority"`
	ID                   string      `json:"id"`
	Amount               money.Money `json:"amount"`
	BenefitReservationID string      `json:"benefit_reservation_id,omitempty"`
	WalletTransactionID  string      `json:"wallet_transaction_id,omitempty"`
	PointReservationID   string      `json:"point_reservation_id,omitempty"`
	GiftCardCode         string      `json:"gift_card_code,omitempty"`
	PointsLedgerEntryID  string      `json:"points_ledger_entry_id,omitempty"`
}

// Line preserves one original priced package grain. Reserved lines carry the
// complete authenticated native allocation; direct merchandise has none and
// cannot prescribe a bucket, lot or physical unit to consume.
type Line struct {
	OrderItemID   string                               `json:"order_item_id"`
	QuoteLineID   string                               `json:"quote_line_id"`
	SKUCode       string                               `json:"sku_code"`
	PackageOption packaging.PackageOptionRef           `json:"package_option"`
	Composition   packaging.PackageCompositionSnapshot `json:"composition"`
	SourceMode    string                               `json:"source_mode"`
	Reservations  []ReservationAllocation              `json:"reservations,omitempty"`
}

// ReservationAllocation retains Supply's native source union and revision.
// The owner must verify it is still the exact allocation for this order and
// generation; a historical reservation receipt alone is insufficient.
type ReservationAllocation struct {
	Allocation             *inventory.StockReservationAllocation `json:"allocation,omitempty"`
	GroupOrderCode         string                                `json:"group_order_code,omitempty"`
	GroupOrderRole         string                                `json:"group_order_role,omitempty"`
	OrderItemID            string                                `json:"order_item_id,omitempty"`
	ParentAllocationLineID string                                `json:"parent_allocation_line_id,omitempty"`
	ParentFulfilmentID     string                                `json:"parent_fulfilment_id,omitempty"`
	ParentOrderNumber      string                                `json:"parent_order_number,omitempty"`
	ReservationID          string                                `json:"reservation_id,omitempty"`
}
