// Package pos holds the in-store point-of-sale operational models: registers,
// shifts, cash movements, shift totals, and immutable receipt snapshots.
// POS sale/cart/customer/payment flows reuse the existing sales, customers,
// payments, wallet, and promotion models — nothing here duplicates them.
package pos

import (
	"time"

	sales "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/orders/order"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/orders/pos/pos_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/product"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/payments/payment/payment_enums"
)

// Register is one physical or virtual point-of-sale register.
type Register struct {
	ID      string `json:"id"`
	StoreID string `json:"store_id"`
	Name    string `json:"name"`
	Status  string `json:"status,omitempty"`

	audit.AuditFields
}

// RegisterShift is one operator shift on a register, from open to close-out.
type RegisterShift struct {
	ID             string                `json:"id"`
	RegisterID     string                `json:"register_id"`
	OperatorUserID string                `json:"operator_user_id"`
	OpenedAt       time.Time             `json:"opened_at"`
	ClosedAt       *time.Time            `json:"closed_at,omitempty"`
	OpeningFloat   money.Money           `json:"opening_float"`
	ClosingCount   *money.Money          `json:"closing_count,omitempty"`
	ExpectedCash   *money.Money          `json:"expected_cash,omitempty"`
	CashVariance   *money.Money          `json:"cash_variance,omitempty"`
	Status         pos_enums.ShiftStatus `json:"status"`

	audit.AuditFields
}

// CashMovement is one cash-drawer movement recorded during a shift.
type CashMovement struct {
	ID         string                     `json:"id"`
	ShiftID    string                     `json:"shift_id"`
	RegisterID string                     `json:"register_id"`
	Kind       pos_enums.CashMovementKind `json:"kind"`
	Amount     money.Money                `json:"amount"`
	Reason     string                     `json:"reason,omitempty"`
	RecordedBy string                     `json:"recorded_by,omitempty"`
	OccurredAt time.Time                  `json:"occurred_at"`
}

// MethodTotal is one payment-method line of a shift totals snapshot.
type MethodTotal struct {
	Method   payment_enums.PaymentMethod `json:"method"`
	Provider string                      `json:"provider,omitempty"`
	Amount   money.Money                 `json:"amount"`
	Count    int                         `json:"count"`
}

// ShiftTotalsSnapshot is the X/Z read model for one shift: per-method sale
// totals, refunds, and cash movements at generation time.
type ShiftTotalsSnapshot struct {
	ShiftID           string        `json:"shift_id"`
	MethodTotals      []MethodTotal `json:"method_totals,omitempty"`
	RefundTotal       money.Money   `json:"refund_total"`
	CashMovementTotal money.Money   `json:"cash_movement_total"`
	GeneratedAt       time.Time     `json:"generated_at"`
}

// ReceiptLine is the customer-safe, immutable product and monetary evidence
// recorded for one receipt line. It never exposes the inventory evidence in
// accepted package pricing.
type ReceiptLine struct {
	SKUID string `json:"sku_id"`
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
// completion. Lines and payment rows are frozen copies — later edits to the
// order or its payments never mutate an issued receipt.
type ReceiptSnapshot struct {
	OrderNumber           string                            `json:"order_number"`
	Revision              int64                             `json:"revision"`
	IssuedAt              time.Time                         `json:"issued_at"`
	Attribution           sales.POSAttribution              `json:"attribution"`
	Lines                 []ReceiptLine                     `json:"lines"`
	Subtotal              money.Money                       `json:"subtotal"`
	Tax                   money.Money                       `json:"tax"`
	Total                 money.Money                       `json:"total"`
	PaymentRows           []sales.CustomerPaymentAllocation `json:"payment_rows,omitempty"`
	PromotionApplications []promotion.PromotionApplication  `json:"promotion_applications"`
}
