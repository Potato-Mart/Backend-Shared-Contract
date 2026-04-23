package purchase

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/pkg/enums"
)

type Order struct {
	ID              string                    `json:"id"`
	OrderNumber     string                    `json:"order_number"`
	SupplierID      string                    `json:"supplier_id,omitempty"`
	SupplierName    string                    `json:"supplier_name,omitempty"`
	SupplierContact string                    `json:"supplier_contact,omitempty"`
	SupplierPhone   string                    `json:"supplier_phone,omitempty"`
	SupplierEmail   string                    `json:"supplier_email,omitempty"`
	Status          enums.PurchaseOrderStatus `json:"status"`
	Currency        string                    `json:"currency"`
	Items           []OrderItem               `json:"items"`
	Subtotal        float64                   `json:"subtotal"`
	TaxAmount       float64                   `json:"tax_amount"`
	ShippingCost    float64                   `json:"shipping_cost,omitempty"`
	Total           float64                   `json:"total"`
	Reference       string                    `json:"reference,omitempty"`
	SupplierRef     string                    `json:"supplier_ref,omitempty"`
	ExpectedAt      *time.Time                `json:"expected_at,omitempty"`
	SubmittedAt     *time.Time                `json:"submitted_at,omitempty"`
	ConfirmedAt     *time.Time                `json:"confirmed_at,omitempty"`
	CancelledAt     *time.Time                `json:"cancelled_at,omitempty"`
	CompletedAt     *time.Time                `json:"completed_at,omitempty"`
	Note            string                    `json:"note,omitempty"`
	InternalNote    string                    `json:"internal_note,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID           string          `json:"id,omitempty"`
	Product      ProductSnapshot `json:"product"`
	UnitCost     float64         `json:"unit_cost"`
	OrderedQty   int             `json:"ordered_qty"`
	ReceivedQty  int             `json:"received_qty"`
	RejectedQty  int             `json:"rejected_qty,omitempty"`
	LineTotal    float64         `json:"line_total"`
	LocationCode string          `json:"location_code,omitempty"`
	Note         string          `json:"note,omitempty"`
}

type ProductSnapshot struct {
	ID       string `json:"id,omitempty"`
	SKU      string `json:"category,omitempty"`
	Name     string `json:"name"`
	EnName   string `json:"en_name,omitempty"`
	Brand    string `json:"brand,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	Storage  string `json:"storage,omitempty"`
	Barcode  string `json:"barcode,omitempty"`
}
