package order_test

import (
	"encoding/json"

	sales "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/orders/order"

	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/payments/payment/payment_enums"
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
				Image:          &security.ObjectMedia{ID: "media_1", URL: "https://cdn.example.test/products/A0001.png"},
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
		decoded.Items[0].Image == nil || decoded.Items[0].Image.ID != "media_1" || decoded.Items[0].Image.URL != "https://cdn.example.test/products/A0001.png" ||
		len(decoded.Items[0].Components) != 1 || decoded.Items[0].Components[0].PackagePrice.AmountMinor != 2100 {
		t.Fatalf("order summary did not round-trip: %+v", decoded)
	}
	line := got["items"].([]any)[0].(map[string]any)
	if _, exists := line["image_url"]; exists {
		t.Fatalf("order line retained image_url: %s", payload)
	}
}
