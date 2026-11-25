package pkg_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification/classification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/inventory"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/procurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/procurement/purchase_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// V53.51 adds three optional typed fields. Each is omitted when empty, so
// records written before the release read back unchanged, and each carries
// its enum type rather than a free string.
func TestV5351OptionalTypedFields(t *testing.T) {
	cases := []struct {
		name  string
		model any
		field string
		json  string
		typ   reflect.Type
		set   any
	}{
		{"stock location role", warehouse.StockLocation{}, "Role", "role,omitempty", reflect.TypeOf(warehouse_enums.StockLocationRole("")),
			warehouse.StockLocation{Role: warehouse_enums.StockLocationRoleDisplay}},
		{"stock movement correction reason", inventory.StockMovement{}, "CorrectionReason", "correction_reason,omitempty", reflect.TypeOf(purchase_enums.ReceiptCorrectionReason("")),
			inventory.StockMovement{CorrectionReason: purchase_enums.ReceiptCorrectionReasonMiscount}},
		{"receipt item correction reason", procurement.PurchaseReceiptItem{}, "CorrectionReason", "correction_reason,omitempty", reflect.TypeOf(purchase_enums.ReceiptCorrectionReason("")),
			procurement.PurchaseReceiptItem{CorrectionReason: purchase_enums.ReceiptCorrectionReasonDamaged}},
		{"supplier freight mode", classification.Supplier{}, "FreightMode", "freight_mode,omitempty", reflect.TypeOf(classification_enums.FreightMode("")),
			classification.Supplier{FreightMode: classification_enums.FreightModeMixed}},
	}
	for _, tc := range cases {
		field, ok := reflect.TypeOf(tc.model).FieldByName(tc.field)
		if !ok || field.Type != tc.typ || field.Tag.Get("json") != tc.json {
			t.Errorf("%s: field %s = %+v, want %s `json:%q`", tc.name, tc.field, field, tc.typ, tc.json)
			continue
		}
		key := `"` + strings.TrimSuffix(tc.json, ",omitempty") + `"`
		empty, err := json.Marshal(tc.model)
		if err != nil {
			t.Fatalf("%s: marshal empty: %v", tc.name, err)
		}
		if strings.Contains(string(empty), key) {
			t.Errorf("%s: empty value must be omitted: %s", tc.name, empty)
		}
		set, err := json.Marshal(tc.set)
		if err != nil {
			t.Fatalf("%s: marshal set: %v", tc.name, err)
		}
		if !strings.Contains(string(set), key+":") {
			t.Errorf("%s: set value missing %s: %s", tc.name, key, set)
		}
	}
}
