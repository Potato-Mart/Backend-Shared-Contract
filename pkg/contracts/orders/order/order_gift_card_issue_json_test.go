package order_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order/order_enums"
)

func TestOrderGiftCardIssueRoundTrip(t *testing.T) {
	issuedAt := time.Date(2026, 11, 14, 2, 30, 0, 0, time.UTC)
	want := order.Order{OrderNumber: "POS-1", GiftCardIssues: []order.OrderGiftCardIssue{{
		LineID: "gift-line-1", ItemID: "item-1",
		Amount:           money.Money{AmountMinor: 10000, Currency: "AUD"},
		BonusAmountMinor: 1000, CustomerNumber: "RCABCDEF", DenominationPolicyVersion: 3,
		IssuanceID: "issuance-1", GiftCardCode: "GC-0001",
		Status:   order_enums.GiftCardIssueStatusIssued,
		IssuedAt: &issuedAt,
	}, {
		LineID: "gift-line-2", ItemID: "item-2",
		Amount:         money.Money{AmountMinor: 5000, Currency: "AUD"},
		CustomerNumber: "RCABCDEF", DenominationPolicyVersion: 3,
		Status: order_enums.GiftCardIssueStatusFailed, FailureReasonCode: "GIFT_CARD_ISSUE_FAILED",
	}}}
	encoded, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	for _, fragment := range []string{
		`{"line_id":"gift-line-1","item_id":"item-1","amount":{"amount_minor":10000,"currency":"AUD"},"bonus_amount_minor":1000,"customer_number":"RCABCDEF","denomination_policy_version":3,"issuance_id":"issuance-1","gift_card_code":"GC-0001","status":"issued","issued_at":"2026-11-14T02:30:00Z"}`,
		`{"line_id":"gift-line-2","item_id":"item-2","amount":{"amount_minor":5000,"currency":"AUD"},"customer_number":"RCABCDEF","denomination_policy_version":3,"status":"failed","failure_reason_code":"GIFT_CARD_ISSUE_FAILED"}`,
	} {
		if !strings.Contains(string(encoded), fragment) {
			t.Fatalf("gift card issue wire shape changed: missing %s in %s", fragment, encoded)
		}
	}
	var got order.Order
	if err := json.Unmarshal(encoded, &got); err != nil || !reflect.DeepEqual(got.GiftCardIssues, want.GiftCardIssues) {
		t.Fatalf("gift card issue round trip: %+v %v", got.GiftCardIssues, err)
	}
	for _, issue := range got.GiftCardIssues {
		if !issue.Status.IsValid() {
			t.Fatalf("status must validate: %+v", issue)
		}
	}
}

// The issued card's PIN is never part of the order record.
func TestOrderGiftCardIssueCarriesNoPIN(t *testing.T) {
	model := reflect.TypeOf(order.OrderGiftCardIssue{})
	for index := 0; index < model.NumField(); index++ {
		field := model.Field(index)
		key := strings.Split(field.Tag.Get("json"), ",")[0]
		for _, word := range strings.Split(key, "_") {
			if word == "pin" {
				t.Fatalf("OrderGiftCardIssue must not carry a PIN field: %s (%s)", field.Name, key)
			}
		}
		if strings.Contains(field.Name, "PIN") || strings.Contains(field.Name, "Pin") {
			t.Fatalf("OrderGiftCardIssue must not carry a PIN field: %s", field.Name)
		}
	}
}

func TestOrderGiftCardIssuesAreOptional(t *testing.T) {
	encoded, err := json.Marshal(order.Order{OrderNumber: "SO-1"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), `"gift_card_issues"`) {
		t.Fatalf("an order selling no gift card must omit gift_card_issues: %s", encoded)
	}
}
