package purchase_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v8/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v8/pkg/contracts/purchase"
	"github.com/Potato-Mart/Backend-Shared-Contract/v8/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v8/pkg/enums"
)

func TestOrderJSONRoundTripWithHistory(t *testing.T) {
	expectedAt := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	occurredAt := time.Date(2026, 6, 17, 9, 0, 0, 0, time.UTC)
	order := purchase.Order{
		ID:          "po_1",
		OrderNumber: "PO-1001",
		Supplier: purchase.SupplierSnapshot{
			PartyRef: common.PartyRef{
				ID:   "sup_1",
				Name: "Supplier One",
			},
		},
		Status:     enums.PurchaseOrderStatusSubmitted,
		Currency:   "AUD",
		Subtotal:   common.Money{AmountMinor: 10000, Currency: "AUD"},
		TaxAmount:  common.Money{AmountMinor: 1000, Currency: "AUD"},
		Total:      common.Money{AmountMinor: 11000, Currency: "AUD"},
		ExpectedAt: &expectedAt,
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

	if decoded.Status != enums.PurchaseOrderStatusSubmitted {
		t.Fatalf("status = %q, want %q", decoded.Status, enums.PurchaseOrderStatusSubmitted)
	}
	if decoded.Supplier.ID != "sup_1" || decoded.Supplier.Name != "Supplier One" {
		t.Fatalf("supplier did not round-trip: %+v", decoded.Supplier)
	}
	if decoded.ExpectedAt == nil || !decoded.ExpectedAt.Equal(expectedAt) {
		t.Fatalf("expected_at = %v, want %s", decoded.ExpectedAt, expectedAt)
	}
	if len(decoded.History) != 1 || decoded.History[0].Changes[0].ToValue != "SUBMITTED" {
		t.Fatalf("history did not round-trip: %+v", decoded.History)
	}
}

func TestOrderJSONLegacyPayloadWithoutHistory(t *testing.T) {
	var order purchase.Order
	if err := json.Unmarshal([]byte(`{"id":"po_1","order_number":"PO-1001","status":"SUBMITTED"}`), &order); err != nil {
		t.Fatalf("unmarshal legacy purchase order: %v", err)
	}
	if order.History != nil {
		t.Fatalf("legacy payload history = %+v, want nil", order.History)
	}
}
