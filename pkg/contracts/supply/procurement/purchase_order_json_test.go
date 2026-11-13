package procurement_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/product"
	purchase "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/procurement"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/procurement/purchase_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestPurchaseOrderJSONRoundTripWithHistory(t *testing.T) {
	expectedAt := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	occurredAt := time.Date(2026, 6, 17, 9, 0, 0, 0, time.UTC)
	order := purchase.PurchaseOrder{
		ID:           "po_1",
		OrderNumber:  "PO-1001",
		SupplierCode: "sup_1",
		SupplierName: "Supplier One",
		Status:       purchase_enums.PurchaseOrderStatusSubmitted,
		Currency:     "AUD",
		Subtotal:     money.Money{AmountMinor: 10000, Currency: "AUD"},
		TaxAmount:    money.Money{AmountMinor: 1000, Currency: "AUD"},
		Total:        money.Money{AmountMinor: 11000, Currency: "AUD"},
		ExpectedAt:   &expectedAt,
		Items: []purchase.PurchaseOrderItem{
			{
				ID:           "po_line_1",
				SKUCode:      "A00001",
				ProductName:  "Potato Crisps",
				ProductImage: &security.ObjectMedia{Code: "media_1", URL: "https://cdn.example.test/products/A00001.png"},
				ProductPackageOption: product.ProductPackageOption{
					Code: "CASE-12", SKUCode: "A00001",
					HandlingUnit: packaging_enums.PackageHandlingUnitCase, UnitsPerPackage: 12,
					IsCanonical: true, IsActive: true, EffectiveFrom: occurredAt,
				},
				CapturedAt:    occurredAt,
				PackageOption: packaging.PackageOptionRef{SKUCode: "A00001", Code: "CASE6", Version: 1},
				UnitCost:      &money.Money{AmountMinor: 200, Currency: "AUD"},
				OrderedComposition: packaging.PackageCompositionSnapshot{
					TotalBaseUnits: 24,
					Components: []packaging.PackageComponentSnapshot{
						{PackageOption: packaging.PackageOptionRef{SKUCode: "A00001", Code: "CASE6", Version: 1}, HandlingUnit: packaging_enums.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24},
					},
				},
				ReceivedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []packaging.PackageComponentSnapshot{}},
				RejectedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []packaging.PackageComponentSnapshot{}},
				LineTotal:           money.Money{AmountMinor: 4800, Currency: "AUD"},
			},
		},
		History: []security.HistoryEntry{
			{
				OccurredAt: occurredAt,
				Type:       "status_change",
				Changes: []security.HistoryChange{
					{Field: "status", FromValue: "DRAFT", ToValue: "SUBMITTED"},
				},
			},
		},
	}

	payload, err := json.Marshal(order)
	if err != nil {
		t.Fatalf("marshal purchase order: %v", err)
	}

	var decoded purchase.PurchaseOrder
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal purchase order: %v", err)
	}

	if decoded.Status != purchase_enums.PurchaseOrderStatusSubmitted {
		t.Fatalf("status = %q, want %q", decoded.Status, purchase_enums.PurchaseOrderStatusSubmitted)
	}
	if decoded.SupplierCode != "sup_1" || decoded.SupplierName != "Supplier One" {
		t.Fatalf("supplier did not round-trip: code=%q name=%q", decoded.SupplierCode, decoded.SupplierName)
	}
	if decoded.ExpectedAt == nil || !decoded.ExpectedAt.Equal(expectedAt) {
		t.Fatalf("expected_at = %v, want %s", decoded.ExpectedAt, expectedAt)
	}
	if len(decoded.History) != 1 || decoded.History[0].Changes[0].ToValue != "SUBMITTED" {
		t.Fatalf("history did not round-trip: %+v", decoded.History)
	}
	if len(decoded.Items) != 1 || decoded.Items[0].OrderedComposition.TotalBaseUnits != 24 || decoded.Items[0].ReceivedComposition.TotalBaseUnits != 0 {
		t.Fatalf("package-aware order item did not round-trip: %+v", decoded.Items)
	}
	if decoded.Items[0].SKUCode != "A00001" || decoded.Items[0].ProductName != "Potato Crisps" || decoded.Items[0].ProductImage == nil || decoded.Items[0].ProductImage.Code != "media_1" || decoded.Items[0].ProductPackageOption.Code != "CASE-12" || !decoded.Items[0].CapturedAt.Equal(occurredAt) {
		t.Fatalf("purchase order item lost frozen product facts: %+v", decoded.Items[0])
	}
	for _, key := range []string{`"sku_code"`, `"product_name"`, `"product_image"`, `"product_package_option"`, `"captured_at"`} {
		if !strings.Contains(string(payload), key) {
			t.Fatalf("purchase order item missing %s: %s", key, payload)
		}
	}
	if strings.Contains(string(payload), `"product":`) {
		t.Fatalf("purchase order item retained product snapshot: %s", payload)
	}
	for _, removed := range []string{`"ordered_qty"`, `"received_qty"`, `"rejected_qty"`, `"location_code"`, `"expire_at"`} {
		if strings.Contains(string(payload), removed) {
			t.Fatalf("purchase order contains removed JSON key %s: %s", removed, payload)
		}
	}
}

func TestReceiptJSONUsesLotBucketAndPackageComposition(t *testing.T) {
	receivedAt := time.Date(2026, 6, 30, 2, 0, 0, 0, time.UTC)
	dateMarkAt := time.Date(2027, 1, 31, 13, 0, 0, 0, time.UTC)
	receipt := purchase.PurchaseReceipt{
		ID:                  "receipt_1",
		PurchaseOrderNumber: "PO-1001",
		DepotCode:           "AU-VIC-MEL-DC-01",
		Status:              purchase_enums.PurchaseReceiptStatusConfirmed,
		ReceivedAt:          &receivedAt,
		Items: []purchase.PurchaseReceiptItem{
			{
				ID:                  "receipt_line_1",
				SKUCode:             "A00001",
				PackageOption:       packaging.PackageOptionRef{SKUCode: "A00001", Code: "CASE6", Version: 1},
				LotID:               "lot_1",
				SupplierLotCode:     "supplier_lot_1",
				ManufacturerLotCode: "manufacturer_lot_1",
				DestinationBucketID: "bucket_1",
				DestinationLocation: warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01-03"},
				DateMark: &warehouse.InventoryDateMark{
					Kind:       warehouse_enums.InventoryDateMarkBestBefore,
					DateMarkAt: dateMarkAt,
					Timezone:   "Australia/Melbourne",
				},
				OrderedComposition:  packaging.PackageCompositionSnapshot{TotalBaseUnits: 24, Components: []packaging.PackageComponentSnapshot{{PackageOption: packaging.PackageOptionRef{SKUCode: "A00001", Code: "CASE6", Version: 1}, HandlingUnit: packaging_enums.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24}}},
				ReceivedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 24, Components: []packaging.PackageComponentSnapshot{{PackageOption: packaging.PackageOptionRef{SKUCode: "A00001", Code: "CASE6", Version: 1}, HandlingUnit: packaging_enums.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24}}},
				RejectedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []packaging.PackageComponentSnapshot{}},
			},
		},
	}

	payload, err := json.Marshal(receipt)
	if err != nil {
		t.Fatalf("marshal purchase receipt: %v", err)
	}
	text := string(payload)
	for _, expected := range []string{`"lot_id":"lot_1"`, `"supplier_lot_code":"supplier_lot_1"`, `"manufacturer_lot_code":"manufacturer_lot_1"`, `"destination_bucket_id":"bucket_1"`, `"location_code":"A-01-03"`, `"date_mark_at":"2027-01-31T13:00:00Z"`, `"received_composition"`, `"total_base_units":0`} {
		if !strings.Contains(text, expected) {
			t.Fatalf("purchase receipt missing %s: %s", expected, payload)
		}
	}
	for _, removed := range []string{`"sku"`, `"ordered_qty"`, `"received_qty"`, `"rejected_qty"`, `"expire_at"`} {
		if strings.Contains(text, removed) {
			t.Fatalf("purchase receipt contains removed JSON key %s: %s", removed, payload)
		}
	}
}

func TestPurchaseReceiptItemLotCodesAreOptional(t *testing.T) {
	payload, err := json.Marshal(purchase.PurchaseReceiptItem{})
	if err != nil {
		t.Fatalf("marshal receipt item: %v", err)
	}
	for _, field := range []string{`"supplier_lot_code"`, `"manufacturer_lot_code"`} {
		if strings.Contains(string(payload), field) {
			t.Fatalf("empty receipt item retained optional %s: %s", field, payload)
		}
	}
}

func TestPurchaseOrderItemExpectedDateMarkRoundTrip(t *testing.T) {
	bare, err := json.Marshal(purchase.PurchaseOrderItem{SKUCode: "A00001"})
	if err != nil {
		t.Fatalf("marshal bare item: %v", err)
	}
	if strings.Contains(string(bare), `"expected_date_mark"`) {
		t.Fatalf("empty expected_date_mark must be omitted: %s", bare)
	}
	var old purchase.PurchaseOrderItem
	if err := json.Unmarshal(bare, &old); err != nil {
		t.Fatalf("unmarshal bare item: %v", err)
	}
	if old.ExpectedDateMark != nil {
		t.Fatalf("an item stored without expected_date_mark must decode nil: %+v", old.ExpectedDateMark)
	}

	// The first instant of 1 December 2026 in Melbourne (UTC+11).
	mark := &warehouse.InventoryDateMark{
		Kind:       warehouse_enums.InventoryDateMarkBestBefore,
		DateMarkAt: time.Date(2026, 11, 30, 13, 0, 0, 0, time.UTC),
		Timezone:   "Australia/Melbourne",
	}
	raw, err := json.Marshal(purchase.PurchaseOrderItem{SKUCode: "A00001", ExpectedDateMark: mark})
	if err != nil {
		t.Fatalf("marshal item: %v", err)
	}
	want := `"expected_date_mark":{"kind":"BEST_BEFORE","date_mark_at":"2026-11-30T13:00:00Z","timezone":"Australia/Melbourne"}`
	if !strings.Contains(string(raw), want) {
		t.Fatalf("item JSON missing %s: %s", want, raw)
	}
	var back purchase.PurchaseOrderItem
	if err := json.Unmarshal(raw, &back); err != nil {
		t.Fatalf("unmarshal item: %v", err)
	}
	if !reflect.DeepEqual(back.ExpectedDateMark, mark) {
		t.Fatalf("expected date mark did not round-trip: %s", raw)
	}

	// The receipt line carries the same type under its own key.
	receiptLine, err := json.Marshal(purchase.PurchaseReceiptItem{DateMark: mark})
	if err != nil {
		t.Fatalf("marshal receipt item: %v", err)
	}
	if !strings.Contains(string(receiptLine), strings.Replace(want, "expected_date_mark", "date_mark", 1)) {
		t.Fatalf("receipt line date mark differs in shape from the expected one: %s", receiptLine)
	}
}
