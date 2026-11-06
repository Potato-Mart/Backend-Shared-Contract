package quote_enums

import "testing"

func TestManualDiscountModeValues(t *testing.T) {
	for mode, wire := range map[ManualDiscountMode]string{
		ManualDiscountModePercent:  "percent",
		ManualDiscountModeAmount:   "amount",
		ManualDiscountModeSetTotal: "set_total",
	} {
		if !mode.IsValid() || mode.String() != wire {
			t.Fatalf("mode %q: valid=%v string=%q, want %q", mode, mode.IsValid(), mode.String(), wire)
		}
	}
	for _, mode := range []ManualDiscountMode{"", "PERCENT", "fixed", "set_price", "settotal"} {
		if mode.IsValid() {
			t.Errorf("unknown mode %q accepted", mode)
		}
	}
}
