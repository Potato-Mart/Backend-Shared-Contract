package sales_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/contracts/sales"
	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/contracts/warehouse"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/membership"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/payment"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/sales"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/warehouse"
)

func TestOrderJSONRoundTripWithHistory(t *testing.T) {
	occurredAt := time.Date(2026, 6, 17, 10, 20, 0, 0, time.UTC)
	order := sales.Order{
		ID:                "ord_1",
		OrderNumber:       "1001",
		Channel:           salesenum.OrderTypeOnline,
		Status:            salesenum.SalesOrderStatusPaid,
		PaymentStatus:     paymentenum.PaymentStatusPaid,
		PaymentMethod:     paymentenum.PaymentMethodCard,
		FulfillmentStatus: salesenum.FulfillmentStatusUnfulfilled,
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

	if decoded.Status != salesenum.SalesOrderStatusPaid || decoded.PaymentStatus != paymentenum.PaymentStatusPaid {
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

func TestV17LineItemsRemovePropertiesAndRequireOrderItemID(t *testing.T) {
	cartPayload, err := json.Marshal(sales.CartItem{Quantity: 1})
	if err != nil {
		t.Fatalf("marshal cart item: %v", err)
	}
	if strings.Contains(string(cartPayload), `"properties"`) {
		t.Fatalf("removed cart item properties remain: %s", cartPayload)
	}

	orderPayload, err := json.Marshal(sales.OrderItem{Quantity: 1})
	if err != nil {
		t.Fatalf("marshal order item: %v", err)
	}
	if !strings.Contains(string(orderPayload), `"id":""`) {
		t.Fatalf("order item id must remain a required JSON key: %s", orderPayload)
	}
	if strings.Contains(string(orderPayload), `"properties"`) {
		t.Fatalf("removed order item properties remain: %s", orderPayload)
	}
}

func TestOrderJSONRoundTripsPackingProgress(t *testing.T) {
	startedAt := time.Date(2026, 7, 9, 9, 30, 0, 0, time.UTC)
	updatedAt := startedAt.Add(15 * time.Minute)
	order := sales.Order{
		ID:                "ord_pack",
		OrderNumber:       "MAMA260709ABC123",
		FulfillmentStatus: salesenum.FulfillmentStatusPacking,
		Packing: &sales.OrderPackingProgress{
			Status:    salesenum.FulfillmentStatusPacking,
			Operator:  "packer@example.test",
			StartedAt: &startedAt,
			UpdatedAt: &updatedAt,
			Lines: []warehouse.PackingLine{
				{
					ProductSKUCode: "SKU-001",
					SKU:            "POT-001",
					ProductName:    "Washed potatoes",
					OrderedQty:     4,
					ScannedQty:     3,
					DamagedQty:     1,
				},
			},
			BoxPlan: &warehouse.PackingBoxPlan{AmbientBoxes: 1, UpdatedAt: updatedAt},
			Damages: []warehouse.PackingDamage{
				{
					ID:             "damage_1",
					ProductSKUCode: "SKU-001",
					DamagedQty:     1,
					Handling:       warehouseenum.PackingDamageShortShipRefund,
					CreatedAt:      updatedAt,
				},
			},
			Discrepancies: []warehouse.PackingDiscrepancy{
				{
					ID:             "disc_1",
					OrderNumber:    "MAMA260709ABC123",
					ProductSKUCode: "SKU-001",
					Kind:           warehouseenum.PackingDiscrepancyKindShortage,
					OrderedQty:     4,
					ScannedQty:     3,
					DiffQty:        1,
					RecordedAt:     updatedAt,
				},
			},
		},
	}

	payload, err := json.Marshal(order)
	if err != nil {
		t.Fatalf("marshal order with packing progress: %v", err)
	}
	for _, key := range []string{`"packing"`, `"scanned_qty"`, `"box_plan"`, `"discrepancies"`, `"started_at"`} {
		if !strings.Contains(string(payload), key) {
			t.Fatalf("order packing JSON missing %s: %s", key, payload)
		}
	}

	var decoded sales.Order
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal order with packing progress: %v", err)
	}
	if decoded.Packing == nil || len(decoded.Packing.Lines) != 1 || decoded.Packing.Lines[0].ScannedQty != 3 {
		t.Fatalf("packing progress did not round-trip: %+v", decoded.Packing)
	}
	if decoded.Packing.BoxPlan == nil || decoded.Packing.BoxPlan.AmbientBoxes != 1 {
		t.Fatalf("packing box plan did not round-trip: %+v", decoded.Packing.BoxPlan)
	}
}

func TestOrderJSONSnapshotsMembershipRedemptions(t *testing.T) {
	discount := common.Money{AmountMinor: 500, Currency: "AUD"}
	order := sales.Order{
		ID:          "ord_points",
		OrderNumber: "1002",
		PointRedemption: &sales.PointRedemptionSnapshot{
			MembershipAccountID: "mem_1",
			OwnerType:           membershipenum.MembershipOwnerTypeRetailCustomer,
			OwnerID:             "retail_1",
			ReservationID:       "res_1",
			LedgerEntryID:       "ledger_1",
			Points:              500,
			DiscountAmount:      discount,
		},
		RewardRedemptions: []sales.RewardRedemptionSnapshot{
			{
				RewardRedemptionID:  "reward_redemption_1",
				RewardCode:          "reward_1",
				MembershipAccountID: "mem_1",
				RewardType:          membershipenum.MembershipRewardTypeOrderDiscount,
				PointsSpent:         500,
				DiscountAmount:      &discount,
			},
		},
		VoucherRedemption: &sales.VoucherRedemptionSnapshot{
			VoucherCode:   "VOUCHER-1",
			AppliedAmount: discount,
			ReservationID: "benefit_res_1",
		},
		GiftCardRedemptions: []sales.GiftCardRedemptionSnapshot{
			{
				GiftCardCode:        "GC-1",
				AppliedAmount:       common.Money{AmountMinor: 2500, Currency: "AUD"},
				ReservationID:       "benefit_res_1",
				WalletTransactionID: "wallet_tx_1",
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
	for _, key := range []string{"point_redemption", "reward_redemptions", "voucher_redemption", "gift_card_redemptions"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("order JSON missing %q: %s", key, payload)
		}
	}
	if _, ok := got["coupon_code"]; ok {
		t.Fatalf("membership redemption should not require coupon_code: %s", payload)
	}
	if !strings.Contains(string(payload), `"applied_amount"`) || !strings.Contains(string(payload), `"wallet_transaction_id":"wallet_tx_1"`) {
		t.Fatalf("wallet redemption snapshots are incomplete: %s", payload)
	}
}
