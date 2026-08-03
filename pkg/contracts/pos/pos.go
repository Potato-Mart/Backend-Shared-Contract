// Package pos holds the in-store point-of-sale operational models: registers,
// shifts, cash movements, shift totals, and immutable receipt snapshots.
// POS sale/cart/customer/payment flows reuse the existing sales, customers,
// payments, wallet, and promotion models — nothing here duplicates them.
package pos

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/sales"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/payment"
	posenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/pos"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/product"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

// CatalogProduct is the cashier-safe package, offer, and availability
// projection returned by the POS catalogue.
type CatalogProduct struct {
	SKUCode            string                                     `json:"sku_code"`
	CategorySKUCode    string                                     `json:"category_sku_code"`
	Name               string                                     `json:"name"`
	Taxed              bool                                       `json:"taxed"`
	StorageType        warehouseenum.StorageType                  `json:"storage_type,omitempty"`
	Status             productenum.ProductStatus                  `json:"status"`
	PackageOptions     []product.ProductPackageOptionSnapshot     `json:"package_options"`
	BarcodeAssignments []product.ProductBarcodeAssignmentSnapshot `json:"barcode_assignments,omitempty"`
	Offers             []product.SellableOfferSnapshot            `json:"offers"`
	Availability       *product.ProductStockSummary               `json:"availability,omitempty"`
	ImageURL           string                                     `json:"image_url,omitempty"`
	CategoryTags       []product.CategoryTag                      `json:"category_tags,omitempty"`
	UpdatedAt          time.Time                                  `json:"updated_at"`
}

// Register is one physical or virtual point-of-sale register.
type Register struct {
	ID      string `json:"id"`
	StoreID string `json:"store_id"`
	Name    string `json:"name"`
	Status  string `json:"status,omitempty"`

	common.AuditFields
}

// RegisterShift is one operator shift on a register, from open to close-out.
type RegisterShift struct {
	ID             string              `json:"id"`
	RegisterID     string              `json:"register_id"`
	OperatorUserID string              `json:"operator_user_id"`
	OpenedAt       time.Time           `json:"opened_at"`
	ClosedAt       *time.Time          `json:"closed_at,omitempty"`
	OpeningFloat   common.Money        `json:"opening_float"`
	ClosingCount   *common.Money       `json:"closing_count,omitempty"`
	ExpectedCash   *common.Money       `json:"expected_cash,omitempty"`
	CashVariance   *common.Money       `json:"cash_variance,omitempty"`
	Status         posenum.ShiftStatus `json:"status"`

	common.AuditFields
}

// CashMovement is one cash-drawer movement recorded during a shift.
type CashMovement struct {
	ID         string                   `json:"id"`
	ShiftID    string                   `json:"shift_id"`
	RegisterID string                   `json:"register_id"`
	Kind       posenum.CashMovementKind `json:"kind"`
	Amount     common.Money             `json:"amount"`
	Reason     string                   `json:"reason,omitempty"`
	RecordedBy string                   `json:"recorded_by,omitempty"`
	OccurredAt time.Time                `json:"occurred_at"`
}

// MethodTotal is one payment-method line of a shift totals snapshot.
type MethodTotal struct {
	Method   paymentenum.PaymentMethod `json:"method"`
	Provider string                    `json:"provider,omitempty"`
	Amount   common.Money              `json:"amount"`
	Count    int                       `json:"count"`
}

// ShiftTotalsSnapshot is the X/Z read model for one shift: per-method sale
// totals, refunds, and cash movements at generation time.
type ShiftTotalsSnapshot struct {
	ShiftID           string        `json:"shift_id"`
	MethodTotals      []MethodTotal `json:"method_totals,omitempty"`
	RefundTotal       common.Money  `json:"refund_total"`
	CashMovementTotal common.Money  `json:"cash_movement_total"`
	GeneratedAt       time.Time     `json:"generated_at"`
}

// ReceiptSnapshot is the immutable customer receipt captured at sale
// completion. Lines and payment rows are frozen copies — later edits to the
// order or its payments never mutate an issued receipt.
type ReceiptSnapshot struct {
	OrderNumber string                            `json:"order_number"`
	Revision    int64                             `json:"revision"`
	IssuedAt    time.Time                         `json:"issued_at"`
	Attribution sales.POSAttribution              `json:"attribution"`
	Lines       []sales.OrderItem                 `json:"lines"`
	Subtotal    common.Money                      `json:"subtotal"`
	Tax         common.Money                      `json:"tax"`
	Total       common.Money                      `json:"total"`
	PaymentRows []sales.CustomerPaymentAllocation `json:"payment_rows,omitempty"`
	Offers      []promotion.ReceiptOffer          `json:"offers,omitempty"`
}
