package purchase

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/purchase/purchase_enums"
)

type Order struct {
	ID           string                             `json:"id"`
	OrderNumber  string                             `json:"order_number"`
	Status       purchase_enums.PurchaseOrderStatus `json:"status"`
	Currency     money.CurrencyCode                 `json:"currency"`
	Items        []OrderItem                        `json:"items"`
	Subtotal     money.Money                        `json:"subtotal"`
	TaxAmount    money.Money                        `json:"tax_amount"`
	ShippingCost *money.Money                       `json:"shipping_cost,omitempty"`
	Total        money.Money                        `json:"total"`
	Reference    string                             `json:"reference,omitempty"`
	SupplierCode string                             `json:"supplier_code"`
	SupplierName string                             `json:"supplier_name,omitempty"`
	ExpectedAt   *time.Time                         `json:"expected_at,omitempty"`
	SubmittedAt  *time.Time                         `json:"submitted_at,omitempty"`
	ConfirmedAt  *time.Time                         `json:"confirmed_at,omitempty"`
	CancelledAt  *time.Time                         `json:"cancelled_at,omitempty"`
	CompletedAt  *time.Time                         `json:"completed_at,omitempty"`
	Note         string                             `json:"note,omitempty"`
	InternalNote string                             `json:"internal_note,omitempty"`
	History      []security.HistoryEntry            `json:"history,omitempty"`

	audit.AuditFields
}

type OrderItem struct {
	ID    string `json:"id,omitempty"`
	SKUID string `json:"sku_id"`
	// SKUCode is the frozen SKU code captured when the purchase line was raised.
	SKUCode              string                               `json:"sku_code"`
	ProductName          string                               `json:"product_name"`
	ProductImage         *security.ObjectMedia                `json:"product_image,omitempty"`
	ProductPackageOption product.ProductPackageOption         `json:"product_package_option"`
	CapturedAt           time.Time                            `json:"captured_at"`
	PackageOptionID      string                               `json:"package_option_id"`
	UnitCost             money.Money                          `json:"unit_cost"`
	OrderedComposition   packaging.PackageCompositionSnapshot `json:"ordered_composition"`
	ReceivedComposition  packaging.PackageCompositionSnapshot `json:"received_composition"`
	RejectedComposition  packaging.PackageCompositionSnapshot `json:"rejected_composition"`
	LineTotal            money.Money                          `json:"line_total"`
	Note                 string                               `json:"note,omitempty"`
}
