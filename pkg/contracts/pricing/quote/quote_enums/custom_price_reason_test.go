package quote_enums

import "testing"

func TestCustomPriceReasonValues(t *testing.T) {
	for reason, wire := range map[CustomPriceReason]string{
		CustomPriceReasonQuickSale:  "quick_sale",
		CustomPriceReasonSoonExpiry: "soon_expiry",
		CustomPriceReasonDamaged:    "damaged",
	} {
		if !reason.IsValid() || reason.String() != wire {
			t.Fatalf("reason %q: valid=%v string=%q, want %q", reason, reason.IsValid(), reason.String(), wire)
		}
	}
	for _, reason := range []CustomPriceReason{"", "unsellable", "QUICK_SALE", "custom"} {
		if reason.IsValid() {
			t.Errorf("unknown reason %q accepted", reason)
		}
	}
}
