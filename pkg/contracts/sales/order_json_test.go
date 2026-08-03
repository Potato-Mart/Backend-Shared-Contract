package sales_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/sales"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/warehouse"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/membership"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/payment"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/sales"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/warehouse"
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

func TestV22LineItemsUsePackageComponentsAndRequireOrderItemID(t *testing.T) {
	cartPayload, err := json.Marshal(sales.CartItem{TotalBaseUnits: 1})
	if err != nil {
		t.Fatalf("marshal cart item: %v", err)
	}
	if strings.Contains(string(cartPayload), `"properties"`) {
		t.Fatalf("removed cart item properties remain: %s", cartPayload)
	}

	orderPayload, err := json.Marshal(sales.OrderItem{TotalBaseUnits: 1})
	if err != nil {
		t.Fatalf("marshal order item: %v", err)
	}
	if !strings.Contains(string(orderPayload), `"id":""`) {
		t.Fatalf("order item id must remain a required JSON key: %s", orderPayload)
	}
	for _, removed := range []string{`"properties"`, `"quantity"`, `"unit_price"`, `"carton_qty"`, `"carton_size"`} {
		if strings.Contains(string(cartPayload), removed) || strings.Contains(string(orderPayload), removed) {
			t.Fatalf("removed scalar order field %s remains: cart=%s order=%s", removed, cartPayload, orderPayload)
		}
	}
	for _, required := range []string{`"components"`, `"total_base_units":1`, `"substitution_policy"`} {
		if !strings.Contains(string(orderPayload), required) {
			t.Fatalf("package-aware order item missing %s: %s", required, orderPayload)
		}
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
					ID:                   "pack_line_1",
					OrderItemID:          "item_1",
					ProductSKUCode:       "A00001",
					ProductName:          "Washed potatoes",
					RequestedComposition: common.PackageCompositionSnapshot{TotalBaseUnits: 4},
					AllocatedComposition: common.PackageCompositionSnapshot{TotalBaseUnits: 4},
					PickedComposition:    common.PackageCompositionSnapshot{TotalBaseUnits: 4},
					PackedComposition:    common.PackageCompositionSnapshot{TotalBaseUnits: 3},
				},
			},
			Containers: []warehouse.OutboundContainerPlan{{ID: "container_1", ContainerCode: "OUT-1", StorageType: warehouseenum.StorageDry, UpdatedAt: updatedAt}},
			Damages: []warehouse.PackingDamage{
				{
					ID:                  "damage_1",
					ProductSKUCode:      "A00001",
					SourceBucketID:      "bucket_1",
					QualityAssessmentID: "qa_1",
					AffectedComposition: common.PackageCompositionSnapshot{TotalBaseUnits: 1},
					Handling:            warehouseenum.PackingDamageShortShipRefund,
					CreatedAt:           updatedAt,
				},
			},
			Discrepancies: []warehouse.PackingDiscrepancy{
				{
					ID:                   "disc_1",
					OrderNumber:          "MAMA260709ABC123",
					ProductSKUCode:       "A00001",
					Kind:                 warehouseenum.PackingDiscrepancyKindShortage,
					RequestedComposition: common.PackageCompositionSnapshot{TotalBaseUnits: 4},
					ObservedComposition:  common.PackageCompositionSnapshot{TotalBaseUnits: 3},
					RecordedAt:           updatedAt,
				},
			},
		},
	}

	payload, err := json.Marshal(order)
	if err != nil {
		t.Fatalf("marshal order with packing progress: %v", err)
	}
	for _, key := range []string{`"packing"`, `"packed_composition"`, `"containers"`, `"discrepancies"`, `"started_at"`} {
		if !strings.Contains(string(payload), key) {
			t.Fatalf("order packing JSON missing %s: %s", key, payload)
		}
	}

	var decoded sales.Order
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal order with packing progress: %v", err)
	}
	if decoded.Packing == nil || len(decoded.Packing.Lines) != 1 || decoded.Packing.Lines[0].PackedComposition.TotalBaseUnits != 3 {
		t.Fatalf("packing progress did not round-trip: %+v", decoded.Packing)
	}
	if len(decoded.Packing.Containers) != 1 || decoded.Packing.Containers[0].ContainerCode != "OUT-1" {
		t.Fatalf("outbound containers did not round-trip: %+v", decoded.Packing.Containers)
	}
}

func TestOrderJSONSnapshotsMembershipRedemptions(t *testing.T) {
	discount := common.Money{AmountMinor: 500, Currency: "AUD"}
	pointOccurredAt := time.Date(2026, 7, 30, 1, 2, 3, 0, time.UTC)
	rewardOccurredAt := pointOccurredAt.Add(time.Second)
	voucherOccurredAt := pointOccurredAt.Add(2 * time.Second)
	giftCardOccurredAt := pointOccurredAt.Add(3 * time.Second)
	order := sales.Order{
		ID:          "ord_points",
		OrderNumber: "1002",
		PointRedemption: &sales.PointRedemptionSnapshot{
			CustomerNumber: "RC-20260727-ABCDEF",
			ReservationID:  "res_1",
			LedgerEntryID:  "ledger_1",
			Points:         500,
			DiscountAmount: discount,
			OccurredAt:     &pointOccurredAt,
		},
		RewardRedemptions: []sales.RewardRedemptionSnapshot{
			{
				RewardRedemptionID: "reward_redemption_1",
				RewardCode:         "reward_1",
				CustomerNumber:     "RC-20260727-ABCDEF",
				RewardType:         membershipenum.MembershipRewardTypeOrderDiscount,
				PointsSpent:        500,
				DiscountAmount:     &discount,
				OccurredAt:         &rewardOccurredAt,
			},
		},
		VoucherRedemption: &sales.VoucherRedemptionSnapshot{
			VoucherCode:   "VOUCHER-1",
			AppliedAmount: discount,
			ReservationID: "benefit_res_1",
			OccurredAt:    &voucherOccurredAt,
		},
		GiftCardRedemptions: []sales.GiftCardRedemptionSnapshot{
			{
				GiftCardCode:        "GC-1",
				AppliedAmount:       common.Money{AmountMinor: 2500, Currency: "AUD"},
				ReservationID:       "benefit_res_1",
				WalletTransactionID: "wallet_tx_1",
				OccurredAt:          &giftCardOccurredAt,
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
	for _, timestamp := range []string{
		`"occurred_at":"2026-07-30T01:02:03Z"`,
		`"occurred_at":"2026-07-30T01:02:04Z"`,
		`"occurred_at":"2026-07-30T01:02:05Z"`,
		`"occurred_at":"2026-07-30T01:02:06Z"`,
	} {
		if !strings.Contains(string(payload), timestamp) {
			t.Fatalf("redemption commit timestamp missing %s: %s", timestamp, payload)
		}
	}
}
