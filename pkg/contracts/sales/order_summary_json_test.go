package sales_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/sales"
	paymentenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/payment"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/sales"
)

func TestOrderSummaryJSONShape(t *testing.T) {
	now := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	summary := sales.OrderSummary{
		OrderNumber:       "MAMA260703ABC123",
		Status:            salesenum.SalesOrderStatusConfirmed,
		PaymentStatus:     paymentenum.PaymentStatusPaid,
		FulfillmentStatus: salesenum.FulfillmentStatusUnfulfilled,
		Channel:           salesenum.OrderTypeOnline,
		PlacedAt:          now,
		UpdatedAt:         now,
		Total:             common.Money{AmountMinor: 4200, Currency: "AUD"},
		ItemCount:         1,
		Items: []sales.OrderLineSummary{
			{
				SKUCode:   "A0001",
				Name:      "Potato 1kg",
				Quantity:  2,
				UnitPrice: common.Money{AmountMinor: 2100, Currency: "AUD"},
				Total:     common.Money{AmountMinor: 4200, Currency: "AUD"},
			},
		},
	}

	payload, err := json.Marshal(summary)
	if err != nil {
		t.Fatalf("marshal order summary: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	for _, key := range []string{
		"order_number", "status", "payment_status", "fulfillment_status",
		"placed_at", "total", "item_count", "items",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("missing json key %q in %s", key, payload)
		}
	}

	var decoded sales.OrderSummary
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal order summary: %v", err)
	}
	if len(decoded.Items) != 1 || decoded.Items[0].SKUCode != "A0001" ||
		decoded.Items[0].UnitPrice.AmountMinor != 2100 {
		t.Fatalf("order summary did not round-trip: %+v", decoded)
	}
}
