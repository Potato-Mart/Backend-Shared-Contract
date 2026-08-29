package receipt

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	sales "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/merchant"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/payment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/payment/payment_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/promotion"
	"time"
)

// ReceiptSnapshot is the immutable customer receipt captured at sale
// completion. Lines, payment rows, and the issuer identity are frozen copies —
// later edits to the order, its payments, or the merchant profile never mutate
// an issued receipt.
//
// DocumentKind records what the document qualifies as. A digital receipt is
// always issued for a completed sale; "tax_invoice" is rendered only when the
// document qualifies, and Buyer is required when the market's tax-invoice
// threshold applies.
type ReceiptSnapshot struct {
	OrderNumber string    `json:"order_number"`
	MarketCode  string    `json:"market_code"`
	Revision    int64     `json:"revision"`
	IssuedAt    time.Time `json:"issued_at"`

	DocumentKind          payment_enums.DocumentKind          `json:"document_kind"`
	Issuer                merchant.MerchantLegalSnapshot      `json:"issuer"`
	Buyer                 *payment.BuyerLegalSnapshot         `json:"buyer,omitempty"`
	Attribution           sales.POSAttribution                `json:"attribution"`
	Lines                 []ReceiptLine                       `json:"lines"`
	Subtotal              money.Money                         `json:"subtotal"`
	Tax                   money.Money                         `json:"tax"`
	Total                 money.Money                         `json:"total"`
	PaymentRows           []payment.CustomerPaymentAllocation `json:"payment_rows,omitempty"`
	PromotionApplications []promotion.PromotionApplication    `json:"promotion_applications"`
	// CashRounding is present only when the sale was tendered entirely in
	// cash. Lines, Subtotal, Tax, and Total always stay at exact minor
	// units.
	CashRounding *CashRoundingSnapshot `json:"cash_rounding,omitempty"`
}
