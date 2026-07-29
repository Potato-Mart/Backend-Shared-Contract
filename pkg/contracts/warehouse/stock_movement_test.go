package warehouse_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/warehouse"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/warehouse"
)

func TestStockMovementRoundTripWithPurchaseLinks(t *testing.T) {
	occurredAt := time.Date(2026, 6, 17, 9, 30, 0, 0, time.UTC)
	movement := warehouse.StockMovement{
		ID:                  "mov_1",
		ProductSKUCode:      "SKU-001",
		SKU:                 "SKU-001",
		ProductName:         "Potato Chips",
		DepotCode:           "DEPOT-1",
		LocationCode:        "A-01",
		Type:                warehouseenum.StockMovementTypePurchaseReceipt,
		QtyDelta:            30,
		BalanceAfter:        130,
		OccurredAt:          occurredAt,
		CreatedBy:           "ops@example.com",
		ReasonCode:          "supplier_delivery",
		Note:                "Received from supplier PO.",
		PurchaseOrderNumber: "PO-1001",
		PurchaseReceiptID:   "pr_1",
		SalesOrderNumber:    "SO-1001",
		ReferenceType:       "purchase_receipt",
		ReferenceID:         "pr_1",
		Metadata: common.Metadata{
			"batch": "B-20260617",
		},
	}

	payload, err := json.Marshal(movement)
	if err != nil {
		t.Fatalf("marshal stock movement: %v", err)
	}

	var decoded warehouse.StockMovement
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal stock movement: %v", err)
	}

	if decoded.PurchaseOrderNumber != "PO-1001" || decoded.PurchaseReceiptID != "pr_1" {
		t.Fatalf("purchase links did not round-trip: %+v", decoded)
	}
	if decoded.SalesOrderNumber != "SO-1001" {
		t.Fatalf("sales order number did not round-trip: %+v", decoded)
	}
	if decoded.ProductSKUCode != "SKU-001" || decoded.DepotCode != "DEPOT-1" {
		t.Fatalf("canonical codes did not round-trip: %+v", decoded)
	}
	if decoded.Type != warehouseenum.StockMovementTypePurchaseReceipt {
		t.Fatalf("type = %q, want %q", decoded.Type, warehouseenum.StockMovementTypePurchaseReceipt)
	}
	if decoded.QtyDelta != 30 || decoded.BalanceAfter != 130 {
		t.Fatalf("stock quantities did not round-trip: %+v", decoded)
	}
	if !decoded.OccurredAt.Equal(occurredAt) {
		t.Fatalf("occurred_at = %s, want %s", decoded.OccurredAt, occurredAt)
	}
	if decoded.Metadata["batch"] != "B-20260617" {
		t.Fatalf("metadata batch = %v, want B-20260617", decoded.Metadata["batch"])
	}
}
