package supply_enums

import "testing"

func TestInventorySaleCostStatus(t *testing.T) {
	for _, v := range []InventorySaleCostStatus{InventorySaleCostStatusValued, InventorySaleCostStatusUnresolved} {
		if !v.IsValid() || v.String() != string(v) {
			t.Fatalf("invalid status %q", v)
		}
	}
	for _, v := range []InventorySaleCostStatus{"", "unknown", "provisional"} {
		if v.IsValid() {
			t.Fatalf("unexpected valid status %q", v)
		}
	}
}
