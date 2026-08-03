package purchase_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/purchase"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/shared"
	purchaseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/purchase"
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
}
