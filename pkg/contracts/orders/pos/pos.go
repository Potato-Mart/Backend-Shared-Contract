// Package pos holds the in-store point-of-sale operational models: registers,
// daily register sessions, cash movements, session totals, and immutable
// receipt snapshots. POS sale/cart/customer/payment flows reuse the existing
// sales, customers, payments, wallet, and promotion models — nothing here
// duplicates them.
//
// A register trades from a depot: depots are the only site identity in the
// platform, so POS records carry depot_code and never a store code.
package pos

import (
	"time"

	sales "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/orders/order"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/money"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/orders/pos/pos_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/supply/product"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/payments/payment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/payments/payment/payment_enums"
)

// Register is one physical or virtual point-of-sale register standing in a
// depot. TerminalID links the register to its paired payment terminal when
// one is bound.
type Register struct {
	ID          string                `json:"id"`
	DepotCode   string                `json:"depot_code"`
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	TerminalID  string                `json:"terminal_id,omitempty"`
	Name        string                `json:"name"`
	Status      string                `json:"status,omitempty"`

	audit.AuditFields
}

// RegisterSession is the single trading session one register runs for one
// business day, from open to close-out.
//
// A session belongs to the register, not to an operator: every operator
// working that register on that business date shares the same session, and
// every workforce rank with POS access may open, manage, and close it. One
// register therefore has at most one session per business date, and
// (register_id, business_date) is the record's natural key.
type RegisterSession struct {
	ID           string                  `json:"id"`
	RegisterID   string                  `json:"register_id"`
	TerminalID   string                  `json:"terminal_id,omitempty"`
	DepotCode    string                  `json:"depot_code"`
	MarketCode   string                  `json:"market_code,omitempty"`
	CountryCode  geography.CountryCode   `json:"country_code,omitempty"`
	BusinessDate temporal.Date           `json:"business_date"`
	Status       pos_enums.SessionStatus `json:"status"`

	OpenedAt       time.Time  `json:"opened_at"`
	OpenedByUserID string     `json:"opened_by_user_id"`
	ClosedAt       *time.Time `json:"closed_at,omitempty"`
	ClosedByUserID string     `json:"closed_by_user_id,omitempty"`

	OpeningFloat money.Money  `json:"opening_float"`
	ClosingCount *money.Money `json:"closing_count,omitempty"`
	ExpectedCash *money.Money `json:"expected_cash,omitempty"`
	CashVariance *money.Money `json:"cash_variance,omitempty"`

	audit.AuditFields
}

// CashMovement is one cash-drawer movement recorded during a session.
type CashMovement struct {
	ID         string                     `json:"id"`
	SessionID  string                     `json:"session_id"`
	RegisterID string                     `json:"register_id"`
	DepotCode  string                     `json:"depot_code,omitempty"`
	Kind       pos_enums.CashMovementKind `json:"kind"`
	Amount     money.Money                `json:"amount"`
	Reason     string                     `json:"reason,omitempty"`
	RecordedBy string                     `json:"recorded_by,omitempty"`
	OccurredAt time.Time                  `json:"occurred_at"`
}

// MethodTotal is one payment-method line of a session totals snapshot.
type MethodTotal struct {
	Method   payment_enums.PaymentMethod `json:"method"`
	Provider string                      `json:"provider,omitempty"`
	Amount   money.Money                 `json:"amount"`
	Count    int                         `json:"count"`
}

// SessionTotalsSnapshot is the X/Z read model for one register session:
// per-method sale totals, refunds, and cash movements at generation time.
type SessionTotalsSnapshot struct {
	SessionID         string        `json:"session_id"`
	RegisterID        string        `json:"register_id,omitempty"`
	DepotCode         string        `json:"depot_code,omitempty"`
	BusinessDate      temporal.Date `json:"business_date,omitempty"`
	MethodTotals      []MethodTotal `json:"method_totals,omitempty"`
	RefundTotal       money.Money   `json:"refund_total"`
	CashMovementTotal money.Money   `json:"cash_movement_total"`
	GeneratedAt       time.Time     `json:"generated_at"`
}

// ReceiptLine is the customer-safe, immutable product and monetary evidence
// recorded for one receipt line. It never exposes the inventory evidence in
// accepted package pricing.
type ReceiptLine struct {
	// SKUCode is the frozen SKU code captured when the receipt was issued.
	SKUCode               string                           `json:"sku_code"`
	ProductName           string                           `json:"product_name"`
	ProductImage          *security.ObjectMedia            `json:"product_image,omitempty"`
	ProductPackageOption  product.ProductPackageOption     `json:"product_package_option"`
	CapturedAt            time.Time                        `json:"captured_at"`
	PackageCount          int64                            `json:"package_count"`
	TotalBaseUnits        int64                            `json:"total_base_units"`
	PackagePrice          money.Money                      `json:"package_price"`
	Subtotal              money.Money                      `json:"subtotal"`
	TaxAmount             money.Money                      `json:"tax_amount"`
	DiscountAmount        money.Money                      `json:"discount_amount"`
	Total                 money.Money                      `json:"total"`
	PromotionApplications []promotion.PromotionApplication `json:"promotion_applications"`
}

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

	DocumentKind          payment_enums.DocumentKind        `json:"document_kind"`
	Issuer                payment.MerchantLegalSnapshot     `json:"issuer"`
	Buyer                 *payment.BuyerLegalSnapshot       `json:"buyer,omitempty"`
	Attribution           sales.POSAttribution              `json:"attribution"`
	Lines                 []ReceiptLine                     `json:"lines"`
	Subtotal              money.Money                       `json:"subtotal"`
	Tax                   money.Money                       `json:"tax"`
	Total                 money.Money                       `json:"total"`
	PaymentRows           []sales.CustomerPaymentAllocation `json:"payment_rows,omitempty"`
	PromotionApplications []promotion.PromotionApplication  `json:"promotion_applications"`
	// CashRounding is present only when the sale was tendered entirely in
	// cash. Lines, Subtotal, Tax, and Total always stay at exact minor
	// units.
	CashRounding *CashRoundingSnapshot `json:"cash_rounding,omitempty"`
}
