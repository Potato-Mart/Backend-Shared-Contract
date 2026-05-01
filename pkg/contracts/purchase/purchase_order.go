package purchase

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v2/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v2/pkg/enums"
)

type Order struct {
	ID           string                    `json:"id"`
	OrderNumber  string                    `json:"order_number"`
	Supplier     Supplier                  `json:"supplier"`
	Status       enums.PurchaseOrderStatus `json:"status"`
	Currency     string                    `json:"currency"`
	Items        []OrderItem               `json:"items"`
	Subtotal     common.Money              `json:"subtotal"`
	TaxAmount    common.Money              `json:"tax_amount"`
	ShippingCost *common.Money             `json:"shipping_cost,omitempty"`
	Total        common.Money              `json:"total"`
	Reference    string                    `json:"reference,omitempty"`
	SupplierRef  string                    `json:"supplier_ref,omitempty"`
	ExpectedAt   *time.Time                `json:"expected_at,omitempty"`
	SubmittedAt  *time.Time                `json:"submitted_at,omitempty"`
	ConfirmedAt  *time.Time                `json:"confirmed_at,omitempty"`
	CancelledAt  *time.Time                `json:"cancelled_at,omitempty"`
	CompletedAt  *time.Time                `json:"completed_at,omitempty"`
	Note         string                    `json:"note,omitempty"`
	InternalNote string                    `json:"internal_note,omitempty"`
	CreatedAt    time.Time                 `json:"created_at"`
	UpdatedAt    time.Time                 `json:"updated_at"`
}

type OrderItem struct {
	ID           string          `json:"id,omitempty"`
	Product      ProductSnapshot `json:"product"`
	UnitCost     common.Money    `json:"unit_cost"`
	OrderedQty   int             `json:"ordered_qty"`
	ReceivedQty  int             `json:"received_qty"`
	RejectedQty  int             `json:"rejected_qty,omitempty"`
	LineTotal    common.Money    `json:"line_total"`
	LocationCode string          `json:"location_code,omitempty"`
	Note         string          `json:"note,omitempty"`
	ExpireAt     time.Time       `json:"expire_at"`
}

type ProductSnapshot struct {
	ID         string                 `json:"id,omitempty"`
	SKU        string                 `json:"sku,omitempty"`
	Name       string                 `json:"name"`
	OtherNames []common.LocalizedName `json:"other_names,omitempty"`
	Brand      string                 `json:"brand,omitempty"`
	ImageURL   string                 `json:"image_url,omitempty"`
	Storage    string                 `json:"storage,omitempty"`
	Barcode    string                 `json:"barcode,omitempty"`
}
