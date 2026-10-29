package supply_test

import (
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/common/localization"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/pubsub/supply"
)

func TestStockRestockAlertTriggeredEventJSONShape(t *testing.T) {
	subscribed := time.Date(2026, 10, 1, 7, 8, 9, 0, time.UTC)
	returned := time.Date(2026, 10, 20, 7, 8, 9, 0, time.UTC)
	shape := marshalObject(t, event.StockRestockAlertTriggeredEvent{
		AlertID:              "restock_alert_1",
		SKUCode:              "A00001",
		MarketCode:           "AU",
		UserID:               "user_1",
		RetailCustomerNumber: "RC0000001",
		ProductName: []localization.LocalizedName{
			{Language: "en", Name: "Potato"},
			{Language: "zh-TW", Name: "馬鈴薯"},
			{Language: "zh-CN", Name: "马铃薯"},
		},
		SubscribedAt: subscribed,
		ReturnedAt:   returned,
	})
	for _, key := range []string{"alert_id", "sku_code", "market_code", "user_id", "retail_customer_number", "product_name", "subscribed_at", "returned_at"} {
		if _, ok := shape[key]; !ok {
			t.Fatalf("restock alert JSON missing %q: %+v", key, shape)
		}
	}
	if names, ok := shape["product_name"].([]any); !ok || len(names) != 3 {
		t.Fatalf("restock alert product names did not marshal: %+v", shape)
	}
	for _, retired := range []string{"occurred_at"} {
		if _, ok := shape[retired]; ok {
			t.Fatalf("restock alert payload must not carry %q: %+v", retired, shape)
		}
	}
}
