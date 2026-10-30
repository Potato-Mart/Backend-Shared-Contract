package possettlement

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
)

// TenderCommitment is the complete frozen tender basis independently reconciled
// by Payments and Pricing. The obligation digest binds the source partition;
// owner holds are separate evidence and are not part of this preimage.
type TenderCommitment struct {
	Schema                string `json:"schema"`
	OrderNumber           string `json:"order_number"`
	CheckoutID            string `json:"checkout_id"`
	SaleID                string `json:"sale_id"`
	RegisterID            string `json:"register_id"`
	SessionID             string `json:"session_id"`
	DepotCode             string `json:"depot_code"`
	OperatorUserID        string `json:"operator_user_id"`
	FulfillmentGeneration int64  `json:"fulfillment_generation"`
	ObligationSHA256      string `json:"obligation_sha256"`
	QuoteKey              string `json:"quote_key"`
	QuoteRevision         string `json:"quote_revision"`
	QuoteFingerprint      string `json:"quote_fingerprint"`
	DiscountedTotal       money.Money        `json:"discounted_total"`
	Payable               money.Money        `json:"payable"`
	CashRounding          money.Money        `json:"cash_rounding"`
	ZeroDue               bool               `json:"zero_due"`
	Instruments           []TenderInstrument `json:"instruments"`
}

// TenderInstrument identifies an original owner debit and its money valuation.
// A gift card binds both the Payments record and Pricing wallet transaction;
// points bind the original Pricing reservation and ledger entry.
type TenderInstrument struct {
	Kind                string      `json:"kind"`
	PaymentID           string      `json:"payment_id,omitempty"`
	ReservationID       string      `json:"reservation_id,omitempty"`
	GiftCardCode        string      `json:"gift_card_code,omitempty"`
	WalletTransactionID string      `json:"wallet_transaction_id,omitempty"`
	PointsLedgerEntryID string      `json:"points_ledger_entry_id,omitempty"`
	Amount              money.Money `json:"amount"`
}

// RefundBarrier identifies a complete order-scoped prefix of Payments-owned
// refund admissions. Pricing acknowledges a prefix only after every admitted
// effect has been applied or authoritatively resolved. A count, timestamp or
// empty query does not establish this evidence. Sequence zero is the initial
// prefix, not proof that a legacy order has registered all historical writers.
type RefundBarrier struct {
	OrderNumber      string `json:"order_number"`
	Sequence         int64  `json:"sequence"`
	PrefixSHA256     string `json:"prefix_sha256"`
	CoverageRevision string `json:"coverage_revision"`
}

// TenderHoldEvidence is a positive owner observation of a durable hold. Lease
// expiry is not a release. Every hold is bound to one obligation and tender
// digest, and Pricing's acknowledged barrier must equal Payments' frozen one.
type TenderHoldEvidence struct {
	Version               int32         `json:"version"`
	Authority             string        `json:"authority"`
	OrderNumber           string        `json:"order_number"`
	SaleID                string        `json:"sale_id"`
	FulfillmentGeneration int64         `json:"fulfillment_generation"`
	ObligationSHA256      string        `json:"obligation_sha256"`
	TenderSHA256          string        `json:"tender_sha256"`
	HoldID                string        `json:"hold_id"`
	OwnerRevision         int64         `json:"owner_revision"`
	RefundBarrier         RefundBarrier `json:"refund_barrier"`
	RecordedAt            time.Time     `json:"recorded_at"`
}
