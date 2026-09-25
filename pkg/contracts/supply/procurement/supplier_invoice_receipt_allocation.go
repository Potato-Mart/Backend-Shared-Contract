package procurement

import "github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/money"

// SupplierInvoiceReceiptAllocation is the exact portion of one invoice line
// attributed to one received item. It records evidence, not a stock movement.
// Supply validates aggregate quantities and amounts against the invoice line
// and receipt item before confirmation.
type SupplierInvoiceReceiptAllocation struct {
	ReceiptID     string `json:"receipt_id"`
	ReceiptItemID string `json:"receipt_item_id"`
	BaseUnits     int64  `json:"base_units"`
	// NetGoodsAmount is this allocation's frozen share after goods discounts,
	// excluding all tax, freight, and duty. Missing and explicit zero differ.
	NetGoodsAmount *money.Money `json:"net_goods_amount,omitempty"`
}
