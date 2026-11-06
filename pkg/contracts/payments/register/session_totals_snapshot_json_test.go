package register_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/register"
)

func TestSessionTotalsSnapshotCarriesManualDiscountSummary(t *testing.T) {
	total := money.Money{AmountMinor: 1250, Currency: "AUD"}
	snapshot := register.SessionTotalsSnapshot{
		SessionID:           "sess_1",
		RefundTotal:         money.Money{AmountMinor: 0, Currency: "AUD"},
		CashMovementTotal:   money.Money{AmountMinor: 0, Currency: "AUD"},
		ManualDiscountTotal: &total,
		ManualDiscountCount: 3,
		GeneratedAt:         time.Date(2026, 11, 6, 8, 0, 0, 0, time.UTC),
	}
	payload, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	want := `"cash_movement_total":{"amount_minor":0,"currency":"AUD"},"manual_discount_total":{"amount_minor":1250,"currency":"AUD"},"manual_discount_count":3,"generated_at":"2026-11-06T08:00:00Z"`
	if !strings.Contains(string(payload), want) {
		t.Fatalf("session totals JSON = %s, want %s", payload, want)
	}
	var decoded register.SessionTotalsSnapshot
	if err := json.Unmarshal(payload, &decoded); err != nil || !reflect.DeepEqual(decoded, snapshot) {
		t.Fatalf("session totals round trip: %+v %v", decoded, err)
	}
}

func TestSessionTotalsSnapshotWithoutManualDiscountsKeepsLegacyBytes(t *testing.T) {
	legacy := `{"session_id":"sess_1","refund_total":{"amount_minor":0,"currency":"AUD"},"cash_movement_total":{"amount_minor":0,"currency":"AUD"},"generated_at":"2026-11-06T08:00:00Z"}`
	var decoded register.SessionTotalsSnapshot
	if err := json.Unmarshal([]byte(legacy), &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.ManualDiscountTotal != nil || decoded.ManualDiscountCount != 0 {
		t.Fatalf("legacy session totals gained a manual discount summary: %+v", decoded)
	}
	payload, err := json.Marshal(decoded)
	if err != nil {
		t.Fatal(err)
	}
	if string(payload) != legacy {
		t.Fatalf("legacy session totals JSON = %s, want %s", payload, legacy)
	}
}
