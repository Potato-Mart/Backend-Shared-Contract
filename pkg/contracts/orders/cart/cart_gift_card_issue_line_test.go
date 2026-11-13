package cart

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
)

func TestCartGiftCardIssueLineRoundTrip(t *testing.T) {
	want := Cart{ID: "cart_1", GiftCardIssues: []CartGiftCardIssueLine{{
		LineID:                    "gift-line-1",
		Amount:                    money.Money{AmountMinor: 10000, Currency: "AUD"},
		CustomerNumber:            "RCABCDEF",
		DenominationPolicyVersion: 3,
		BonusAmountMinor:          1000,
		AddedByUserID:             "cashier_1",
		AddedAt:                   time.Date(2026, 11, 14, 2, 0, 0, 0, time.UTC),
	}}}
	encoded, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	wantJSON := `"gift_card_issues":[{"line_id":"gift-line-1","amount":{"amount_minor":10000,"currency":"AUD"},"customer_number":"RCABCDEF","denomination_policy_version":3,"bonus_amount_minor":1000,"added_by_user_id":"cashier_1","added_at":"2026-11-14T02:00:00Z"}]`
	if !strings.Contains(string(encoded), wantJSON) {
		t.Fatalf("cart JSON = %s, want %s", encoded, wantJSON)
	}
	var got Cart
	if err := json.Unmarshal(encoded, &got); err != nil || !reflect.DeepEqual(got.GiftCardIssues, want.GiftCardIssues) {
		t.Fatalf("gift card issue round trip: %+v %v", got.GiftCardIssues, err)
	}
}

func TestCartGiftCardIssuesAreOptional(t *testing.T) {
	encoded, err := json.Marshal(Cart{ID: "cart_1"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), `"gift_card_issues"`) {
		t.Fatalf("a cart selling no gift card must omit gift_card_issues: %s", encoded)
	}
	line, err := json.Marshal(CartGiftCardIssueLine{LineID: "gift-line-1"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(line), `"bonus_amount_minor"`) {
		t.Fatalf("a line without bonus must omit bonus_amount_minor: %s", line)
	}
}
