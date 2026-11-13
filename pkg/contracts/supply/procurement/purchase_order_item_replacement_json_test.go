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

// A replaced line records what was replaced and by which line; the new line
// names the line it replaces.
func TestPurchaseOrderItemReplacementRoundTrip(t *testing.T) {
	at := time.Date(2026, 11, 14, 4, 0, 0, 0, time.UTC)
	items := []purchase.PurchaseOrderItem{{
		ID: "po_line_1", SKUCode: "SKU-1",
		OrderedComposition:  packaging.PackageCompositionSnapshot{TotalBaseUnits: 24},
		ReceivedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 6},
		ReplacedComposition: &packaging.PackageCompositionSnapshot{TotalBaseUnits: 18},
		Replacements: []purchase.PurchaseOrderItemReplacement{{
			ID: "replacement-1", ReplacementItemID: "po_line_2",
			ReplacedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 18},
			ReasonCode:          purchase_enums.PurchaseOrderItemReplacementReasonWrongProductOrdered,
			Note:                "ordered the 500 g pack", ReplacedBy: "buyer-1", ReplacedAt: at,
		}},
	}, {
		ID: "po_line_2", SKUCode: "SKU-2", ReplacesItemID: "po_line_1",
		OrderedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 18},
	}}
	raw, err := json.Marshal(items)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	for _, fragment := range []string{
		`"replaced_composition":{"total_base_units":18,"components":null},"replacements":[{"id":"replacement-1","replacement_item_id":"po_line_2","replaced_composition":{"total_base_units":18,"components":null},"reason_code":"WRONG_PRODUCT_ORDERED","note":"ordered the 500 g pack","replaced_by":"buyer-1","replaced_at":"2026-11-14T04:00:00Z"}]`,
		`"replaces_item_id":"po_line_1"`,
	} {
		if !strings.Contains(string(raw), fragment) {
			t.Fatalf("replacement wire shape changed: missing %s in %s", fragment, raw)
		}
	}
	var back []purchase.PurchaseOrderItem
	if err := json.Unmarshal(raw, &back); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !reflect.DeepEqual(back, items) || !back[0].Replacements[0].ReasonCode.IsValid() {
		t.Fatalf("round trip = %+v", back)
	}
	// outstanding = ordered - received - replaced
	outstanding := back[0].OrderedComposition.TotalBaseUnits - back[0].ReceivedComposition.TotalBaseUnits - back[0].ReplacedComposition.TotalBaseUnits
	if outstanding != 0 {
		t.Fatalf("outstanding = %d, want 0", outstanding)
	}
}

// A line never replaced keeps its existing wire shape.
func TestPurchaseOrderItemReplacementFieldsAreOptional(t *testing.T) {
	raw, err := json.Marshal(purchase.PurchaseOrderItem{ID: "po_line_1", SKUCode: "SKU-1"})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	for _, key := range []string{"replaced_composition", "replaces_item_id", "replacements"} {
		if strings.Contains(string(raw), `"`+key+`"`) {
			t.Fatalf("absent %s must be omitted: %s", key, raw)
		}
	}
}
