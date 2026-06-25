package sales_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/sales"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

func TestOrderJSONRoundTripWithHistory(t *testing.T) {
	occurredAt := time.Date(2026, 6, 17, 10, 20, 0, 0, time.UTC)
	order := sales.Order{
		ID:                "ord_1",
		OrderNumber:       "1001",
		Channel:           enums.OrderTypeOnline,
		Status:            enums.SalesOrderStatusPaid,
		PaymentStatus:     enums.PaymentStatusPaid,
		PaymentMethod:     enums.PaymentMethodCard,
		FulfillmentStatus: enums.FulfillmentStatusUnfulfilled,
		Customer: common.PartyRef{
			ID:    "cust_1",
			Name:  "Customer One",
			Email: "customer@example.com",
		},
		Subtotal:       common.Money{AmountMinor: 10000, Currency: "AUD"},
		DiscountAmount: common.Money{AmountMinor: 1000, Currency: "AUD"},
		ShippingAmount: common.Money{AmountMinor: 500, Currency: "AUD"},
		TaxAmount:      common.Money{AmountMinor: 864, Currency: "AUD"},
		Total:          common.Money{AmountMinor: 9500, Currency: "AUD"},
		History: []shared.HistoryEntry{
			{
				OccurredAt: occurredAt,
				Type:       "status_change",
				Changes: []shared.HistoryChange{
					{Field: "status", FromValue: "confirmed", ToValue: "paid"},
				},
			},
		},
	}

	payload, err := json.Marshal(order)
	if err != nil {
		t.Fatalf("marshal order: %v", err)
	}

	var decoded sales.Order
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal order: %v", err)
	}

	if decoded.Status != enums.SalesOrderStatusPaid || decoded.PaymentStatus != enums.PaymentStatusPaid {
		t.Fatalf("status fields did not round-trip: %+v", decoded)
	}
	if decoded.Total.AmountMinor != 9500 || decoded.Total.Currency != "AUD" {
		t.Fatalf("total did not round-trip: %+v", decoded.Total)
	}
	if len(decoded.History) != 1 || decoded.History[0].Changes[0].ToValue != "paid" {
		t.Fatalf("history did not round-trip: %+v", decoded.History)
	}
}

func TestOrderJSONOmitsEmptyHistory(t *testing.T) {
	payload, err := json.Marshal(sales.Order{ID: "ord_1"})
	if err != nil {
		t.Fatalf("marshal order: %v", err)
	}
	if strings.Contains(string(payload), `"history"`) {
		t.Fatalf("empty history should be omitted, got %s", payload)
	}
}

func TestOrderJSONSnapshotsMembershipRedemptions(t *testing.T) {
	discount := common.Money{AmountMinor: 500, Currency: "AUD"}
	order := sales.Order{
		ID:          "ord_points",
		OrderNumber: "1002",
		PointRedemption: &sales.PointRedemptionSnapshot{
			MembershipAccountID: "mem_1",
			OwnerType:           enums.MembershipOwnerTypeRetailCustomer,
			OwnerID:             "retail_1",
			ReservationID:       "res_1",
			LedgerEntryID:       "ledger_1",
			Points:              500,
			DiscountAmount:      discount,
		},
		RewardRedemptions: []sales.RewardRedemptionSnapshot{
			{
				RewardRedemptionID:  "reward_redemption_1",
				RewardID:            "reward_1",
				MembershipAccountID: "mem_1",
				RewardType:          enums.MembershipRewardTypeOrderDiscount,
				PointsSpent:         500,
				DiscountAmount:      &discount,
			},
		},
	}

	payload, err := json.Marshal(order)
	if err != nil {
		t.Fatalf("marshal order with membership redemptions: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal order JSON: %v", err)
	}
	for _, key := range []string{"point_redemption", "reward_redemptions"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("order JSON missing %q: %s", key, payload)
		}
	}
	if _, ok := got["coupon_code"]; ok {
		t.Fatalf("membership redemption should not require coupon_code: %s", payload)
	}
}

func TestOrderJSONLegacyPayloadWithoutHistory(t *testing.T) {
	var order sales.Order
	if err := json.Unmarshal([]byte(`{"id":"ord_1","order_number":"1001","status":"paid"}`), &order); err != nil {
		t.Fatalf("unmarshal legacy order: %v", err)
	}
	if order.History != nil {
		t.Fatalf("legacy payload history = %+v, want nil", order.History)
	}
}
