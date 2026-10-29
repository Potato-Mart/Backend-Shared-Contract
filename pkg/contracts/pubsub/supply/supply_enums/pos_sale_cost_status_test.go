package supply_enums

import "testing"

func TestPOSSaleCostStatus(t *testing.T) {
	for _, status := range []POSSaleCostStatus{POSSaleCostStatusValued, POSSaleCostStatusUnresolved} {
		if !status.IsValid() || status.String() != string(status) {
			t.Fatalf("invalid declared status: %q", status)
		}
	}
	for _, status := range []POSSaleCostStatus{"", "estimated", "unknown"} {
		if status.IsValid() {
			t.Fatalf("unexpected status: %q", status)
		}
	}
}
