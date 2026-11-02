package register_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/register"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/payments/register/register_enums"
)

func TestRegisterSessionSequenceRoundTrip(t *testing.T) {
	session := register.RegisterSession{
		ID:             "sess_2",
		RegisterID:     "reg_1",
		DepotCode:      "SYD",
		BusinessDate:   "2026-11-03",
		Sequence:       2,
		Status:         register_enums.SessionStatusOpen,
		OpenedAt:       time.Date(2026, 11, 3, 4, 0, 0, 0, time.UTC),
		OpenedByUserID: "user_1",
		OpeningFloat:   money.Money{AmountMinor: 20000, Currency: "AUD"},
	}
	payload, err := json.Marshal(session)
	if err != nil {
		t.Fatalf("marshal register session: %v", err)
	}
	if !strings.Contains(string(payload), `"business_date":"2026-11-03","sequence":2,"status":"open"`) {
		t.Fatalf("register session JSON is missing its sequence: %s", payload)
	}
	var decoded register.RegisterSession
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal register session: %v", err)
	}
	if decoded.Sequence != 2 || decoded.RegisterID != "reg_1" || decoded.BusinessDate != "2026-11-03" {
		t.Fatalf("register session did not round-trip: %+v", decoded)
	}
}

func TestRegisterSessionWithoutSequenceDecodesAsLegacyRecord(t *testing.T) {
	var decoded register.RegisterSession
	if err := json.Unmarshal([]byte(`{"id":"sess_1","register_id":"reg_1","depot_code":"SYD","business_date":"2026-11-02","status":"closed","opened_at":"2026-11-02T00:00:00Z","opened_by_user_id":"user_1","opening_float":{"amount_minor":0,"currency":"AUD"}}`), &decoded); err != nil {
		t.Fatalf("unmarshal legacy register session: %v", err)
	}
	if decoded.Sequence != 0 {
		t.Fatalf("legacy register session sequence = %d, want 0 (absent)", decoded.Sequence)
	}
	payload, err := json.Marshal(decoded)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(payload), "sequence") {
		t.Fatalf("absent sequence must stay omitted: %s", payload)
	}
}
