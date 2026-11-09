package procurement_test

import (
	"encoding/json"
	"testing"
	"time"

	purchase "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/procurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/procurement/purchase_enums"
)

func TestPurchaseReceiptReversalFieldsAreOptional(t *testing.T) {
	shape := purchaseReceiptShape(t, purchase.PurchaseReceipt{ID: "r1", Status: purchase_enums.PurchaseReceiptStatusDraft})
	for _, key := range []string{"reversed_at", "reversed_by", "reversal_reason"} {
		if _, ok := shape[key]; ok {
			t.Fatalf("draft receipt JSON carries %q: %+v", key, shape)
		}
	}
}

func TestPurchaseReceiptReversalFieldsRoundTrip(t *testing.T) {
	at := time.Date(2026, 11, 10, 3, 0, 0, 0, time.UTC)
	receipt := purchase.PurchaseReceipt{
		ID: "r1", Status: purchase_enums.PurchaseReceiptStatusReversed,
		ReversedAt: &at, ReversedBy: "staff_1", ReversalReason: "duplicate receipt",
	}
	shape := purchaseReceiptShape(t, receipt)
	if shape["status"] != "REVERSED" || shape["reversed_by"] != "staff_1" || shape["reversal_reason"] != "duplicate receipt" {
		t.Fatalf("reversed receipt JSON = %+v", shape)
	}
	raw, _ := json.Marshal(receipt)
	var back purchase.PurchaseReceipt
	if err := json.Unmarshal(raw, &back); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if back.ReversedAt == nil || !back.ReversedAt.Equal(at) || !back.Status.IsValid() {
		t.Fatalf("round trip = %+v", back)
	}
}

func purchaseReceiptShape(t *testing.T, value purchase.PurchaseReceipt) map[string]any {
	t.Helper()
	raw, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var shape map[string]any
	if err := json.Unmarshal(raw, &shape); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	return shape
}
