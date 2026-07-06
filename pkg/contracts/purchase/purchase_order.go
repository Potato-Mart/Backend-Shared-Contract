package purchase

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/shared"
	purchaseenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/purchase"
)

type Order struct {
	ID           string                           `json:"id"`
	OrderNumber  string                           `json:"order_number"`
	Status       purchaseenum.PurchaseOrderStatus `json:"status"`
	Currency     string                           `json:"currency"`
	Items        []OrderItem                      `json:"items"`
	Subtotal     common.Money                     `json:"subtotal"`
	TaxAmount    common.Money                     `json:"tax_amount"`
	ShippingCost *common.Money                    `json:"shipping_cost,omitempty"`
	Total        common.Money                     `json:"total"`
	Reference    string                           `json:"reference,omitempty"`
	SupplierCode string                           `json:"supplier_code"`
	SupplierName string                           `json:"supplier_name,omitempty"`
	ExpectedAt   *time.Time                       `json:"expected_at,omitempty"`
	SubmittedAt  *time.Time                       `json:"submitted_at,omitempty"`
	ConfirmedAt  *time.Time                       `json:"confirmed_at,omitempty"`
	CancelledAt  *time.Time                       `json:"cancelled_at,omitempty"`
	CompletedAt  *time.Time                       `json:"completed_at,omitempty"`
	Note         string                           `json:"note,omitempty"`
	InternalNote string                           `json:"internal_note,omitempty"`
	History      []shared.HistoryEntry            `json:"history,omitempty"`

	common.AuditFields
}

type OrderItem struct {
	ID           string           `json:"id,omitempty"`
	Product      product.Snapshot `json:"product"`
	UnitCost     common.Money     `json:"unit_cost"`
	OrderedQty   int              `json:"ordered_qty"`
	ReceivedQty  int              `json:"received_qty"`
	RejectedQty  int              `json:"rejected_qty,omitempty"`
	LineTotal    common.Money     `json:"line_total"`
	LocationCode string           `json:"location_code,omitempty"`
	Note         string           `json:"note,omitempty"`
	ExpireAt     time.Time        `json:"expire_at"`
}
