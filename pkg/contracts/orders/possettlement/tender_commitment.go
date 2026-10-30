package possettlement

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

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
