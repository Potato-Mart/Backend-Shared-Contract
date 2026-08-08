package order_test

import (
	"encoding/json"

	sales "github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/orders/order"

	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/payments/payment/payment_enums"
)

func TestOrderSummaryJSONShape(t *testing.T) {
	now := time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)
	summary := sales.OrderSummary{
		OrderNumber:       "MAMA260703ABC123",
		Status:            order_enums.SalesOrderStatusConfirmed,
		PaymentStatus:     payment_enums.PaymentStatusPaid,
		FulfillmentStatus: order_enums.FulfillmentStatusUnfulfilled,
		Channel:           commerce_enums.OrderTypeOnline,
		PlacedAt:          now,
		UpdatedAt:         now,
		Total:             money.Money{AmountMinor: 4200, Currency: "AUD"},
		ItemCount:         1,
		Items: []sales.OrderLineSummary{
			{
				SKUCode:        "A0001",
				Name:           "Potato 1kg",
				Components:     []sales.PricedPackageComponent{{RequestedPackageCount: 2, RequestedBaseUnits: 2, PackagePrice: money.Money{AmountMinor: 2100, Currency: "AUD"}}},
				TotalBaseUnits: 2,
				Total:          money.Money{AmountMinor: 4200, Currency: "AUD"},
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
		len(decoded.Items[0].Components) != 1 || decoded.Items[0].Components[0].PackagePrice.AmountMinor != 2100 {
		t.Fatalf("order summary did not round-trip: %+v", decoded)
	}
}
