package purchase_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/purchase"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/warehouse"
	purchaseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/purchase"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
)

func TestOrderJSONRoundTripWithHistory(t *testing.T) {
	expectedAt := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	occurredAt := time.Date(2026, 6, 17, 9, 0, 0, 0, time.UTC)
	order := purchase.Order{
		ID:           "po_1",
		OrderNumber:  "PO-1001",
		SupplierCode: "sup_1",
		SupplierName: "Supplier One",
		Status:       purchaseenum.PurchaseOrderStatusSubmitted,
		Currency:     "AUD",
		Subtotal:     common.Money{AmountMinor: 10000, Currency: "AUD"},
		TaxAmount:    common.Money{AmountMinor: 1000, Currency: "AUD"},
		Total:        common.Money{AmountMinor: 11000, Currency: "AUD"},
		ExpectedAt:   &expectedAt,
		Items: []purchase.OrderItem{
			{
				ID:              "po_line_1",
				Product:         product.Snapshot{SKUCode: "A00001", Name: "Potato Crisps"},
				PackageOptionID: "pkg_case_12",
				UnitCost:        common.Money{AmountMinor: 2400, Currency: "AUD"},
				OrderedComposition: common.PackageCompositionSnapshot{
					TotalBaseUnits: 24,
					Components: []common.PackageComponentSnapshot{
						{PackageOptionID: "pkg_case_12", HandlingUnit: common.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24},
					},
				},
				ReceivedComposition: common.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []common.PackageComponentSnapshot{}},
				RejectedComposition: common.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []common.PackageComponentSnapshot{}},
				LineTotal:           common.Money{AmountMinor: 4800, Currency: "AUD"},
			},
		},
		History: []shared.HistoryEntry{
			{
				OccurredAt: occurredAt,
				Type:       "status_change",
				Changes: []shared.HistoryChange{
					{Field: "status", FromValue: "DRAFT", ToValue: "SUBMITTED"},
				},
			},
		},
	}

	payload, err := json.Marshal(order)
	if err != nil {
		t.Fatalf("marshal purchase order: %v", err)
	}

	var decoded purchase.Order
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal purchase order: %v", err)
	}

	if decoded.Status != purchaseenum.PurchaseOrderStatusSubmitted {
		t.Fatalf("status = %q, want %q", decoded.Status, purchaseenum.PurchaseOrderStatusSubmitted)
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
	for _, removed := range []string{`"ordered_qty"`, `"received_qty"`, `"rejected_qty"`, `"location_code"`, `"expire_at"`} {
		if strings.Contains(string(payload), removed) {
			t.Fatalf("purchase order contains removed JSON key %s: %s", removed, payload)
		}
	}
}

func TestReceiptJSONUsesLotBucketAndPackageComposition(t *testing.T) {
	receivedAt := time.Date(2026, 6, 30, 2, 0, 0, 0, time.UTC)
	dateMarkAt := time.Date(2027, 1, 31, 13, 0, 0, 0, time.UTC)
	receipt := purchase.Receipt{
		ID:          "receipt_1",
		OrderNumber: "PO-1001",
		DepotCode:   "AU-VIC-MEL-DC-01",
		Status:      purchaseenum.PurchaseOrderStatusReceived,
		ReceivedAt:  &receivedAt,
		Items: []purchase.ReceiptItem{
			{
				ID:                  "receipt_line_1",
				ProductSKUCode:      "A00001",
				PackageOptionID:     "pkg_case_12",
				LotID:               "lot_1",
				DestinationBucketID: "bucket_1",
				DestinationLocation: warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01-03"},
				DateMark: &warehouse.InventoryDateMark{
					Kind:       warehouseenum.InventoryDateMarkBestBefore,
					DateMarkAt: dateMarkAt,
					Timezone:   "Australia/Melbourne",
				},
				OrderedComposition:  common.PackageCompositionSnapshot{TotalBaseUnits: 24, Components: []common.PackageComponentSnapshot{{PackageOptionID: "pkg_case_12", HandlingUnit: common.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24}}},
				ReceivedComposition: common.PackageCompositionSnapshot{TotalBaseUnits: 24, Components: []common.PackageComponentSnapshot{{PackageOptionID: "pkg_case_12", HandlingUnit: common.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24}}},
				RejectedComposition: common.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []common.PackageComponentSnapshot{}},
			},
		},
	}

	payload, err := json.Marshal(receipt)
	if err != nil {
		t.Fatalf("marshal purchase receipt: %v", err)
	}
	text := string(payload)
	for _, expected := range []string{`"lot_id":"lot_1"`, `"destination_bucket_id":"bucket_1"`, `"location_code":"A-01-03"`, `"date_mark_at":"2027-01-31T13:00:00Z"`, `"received_composition"`, `"total_base_units":0`} {
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
