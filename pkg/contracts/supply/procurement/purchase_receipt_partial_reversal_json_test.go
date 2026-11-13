package procurement_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	purchase "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/procurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/procurement/purchase_enums"
)

// A partially reversed receipt carries the per-line reversed quantity and
// each reversal, and leaves the receipt-level reversal fields empty.
func TestPurchaseReceiptPartialReversalRoundTrip(t *testing.T) {
	at := time.Date(2026, 11, 14, 3, 0, 0, 0, time.UTC)
	receipt := purchase.PurchaseReceipt{
		ID: "r1", Status: purchase_enums.PurchaseReceiptStatusPartiallyReversed,
		Items: []purchase.PurchaseReceiptItem{{
			ID: "line-1", SKUCode: "SKU-1", LotID: "lot-1",
			ReceivedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 12},
			ReversedComposition: &packaging.PackageCompositionSnapshot{TotalBaseUnits: 4},
			Reversals: []purchase.PurchaseReceiptItemReversal{{
				ReversalID: "reversal-1", Composition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 4},
				Reason: "damaged on arrival", ReversedAt: at, ReversedBy: "staff-1",
			}},
			CorrectedByReceiptItemID: "line-2",
		}, {
			ID: "line-2", SKUCode: "SKU-2", LotID: "lot-2",
			ReceivedComposition:   packaging.PackageCompositionSnapshot{TotalBaseUnits: 4},
			CorrectsReceiptItemID: "line-1",
		}},
	}
	raw, err := json.Marshal(receipt)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	for _, fragment := range []string{
		`"status":"PARTIALLY_REVERSED"`,
		`"reversed_composition":{"total_base_units":4,"components":null}`,
		`"reversals":[{"reversal_id":"reversal-1","composition":{"total_base_units":4,"components":null},"reason":"damaged on arrival","reversed_at":"2026-11-14T03:00:00Z","reversed_by":"staff-1"}]`,
		`"corrected_by_receipt_item_id":"line-2"`,
		`"corrects_receipt_item_id":"line-1"`,
	} {
		if !strings.Contains(string(raw), fragment) {
			t.Fatalf("partial reversal wire shape changed: missing %s in %s", fragment, raw)
		}
	}
	shape := purchaseReceiptShape(t, receipt)
	for _, key := range []string{"reversed_at", "reversed_by", "reversal_reason"} {
		if _, ok := shape[key]; ok {
			t.Fatalf("a partial reversal must not set the receipt-level %q: %s", key, raw)
		}
	}
	var back purchase.PurchaseReceipt
	if err := json.Unmarshal(raw, &back); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !reflect.DeepEqual(back, receipt) || !back.Status.IsValid() {
		t.Fatalf("round trip = %+v", back)
	}
}

// A receipt line never reversed or corrected keeps its existing wire shape.
func TestPurchaseReceiptItemReversalFieldsAreOptional(t *testing.T) {
	raw, err := json.Marshal(purchase.PurchaseReceiptItem{ID: "line-1", SKUCode: "SKU-1", LotID: "lot-1"})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	for _, key := range []string{"reversed_composition", "reversals", "corrects_receipt_item_id", "corrected_by_receipt_item_id"} {
		if strings.Contains(string(raw), `"`+key+`"`) {
			t.Fatalf("absent %s must be omitted: %s", key, raw)
		}
	}
}
