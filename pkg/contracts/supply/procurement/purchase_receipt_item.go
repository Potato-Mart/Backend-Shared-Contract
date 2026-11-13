package procurement

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse"
)

// PurchaseReceiptItem is one received procurement line.
//
// ReversedComposition is the part of ReceivedComposition that reversals have
// taken back so far, and Reversals lists each of them, oldest first. Both are
// absent until the line is first reversed. A line is fully reversed when
// ReversedComposition.TotalBaseUnits equals ReceivedComposition.TotalBaseUnits.
//
// CorrectsReceiptItemID is set on a line recorded to correct another receipt
// line that was received against the wrong purchase-order line or product: it
// names that source line, whose units were reversed in the same correction.
// The source line carries CorrectedByReceiptItemID, naming the correcting
// line. Both are absent on ordinary lines.
type PurchaseReceiptItem struct {
	ID string `json:"id,omitempty"`
	// PurchaseOrderItemID identifies the exact line in the receipt's purchase
	// order, including when that order contains multiple lines for one SKU.
	PurchaseOrderItemID string `json:"purchase_order_item_id,omitempty"`
	// SKUCode is the frozen SKU code captured when the receipt line was recorded.
	SKUCode                  string                                `json:"sku_code"`
	ProductName              string                                `json:"product_name,omitempty"`
	PackageOption            packaging.PackageOptionRef            `json:"package_option"`
	LotID                    string                                `json:"lot_id"`
	SupplierLotCode          string                                `json:"supplier_lot_code,omitempty"`
	ManufacturerLotCode      string                                `json:"manufacturer_lot_code,omitempty"`
	DestinationBucketID      string                                `json:"destination_bucket_id"`
	DestinationLocation      warehouse.StockLocationRef            `json:"destination_location"`
	DateMark                 *warehouse.InventoryDateMark          `json:"date_mark,omitempty"`
	OrderedComposition       packaging.PackageCompositionSnapshot  `json:"ordered_composition"`
	ReceivedComposition      packaging.PackageCompositionSnapshot  `json:"received_composition"`
	RejectedComposition      packaging.PackageCompositionSnapshot  `json:"rejected_composition"`
	Note                     string                                `json:"note,omitempty"`
	ReversedComposition      *packaging.PackageCompositionSnapshot `json:"reversed_composition,omitempty"`
	Reversals                []PurchaseReceiptItemReversal         `json:"reversals,omitempty"`
	CorrectsReceiptItemID    string                                `json:"corrects_receipt_item_id,omitempty"`
	CorrectedByReceiptItemID string                                `json:"corrected_by_receipt_item_id,omitempty"`
}
