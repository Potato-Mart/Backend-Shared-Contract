package order_test

import (
	"encoding/json"

	sales "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/orders/order"

	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging/packaging_enums"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/payments/payment/payment_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product"
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
				ProductSKUCode: "A0001",
				ProductName:    "Potato 1kg",
				ProductImage:   &security.ObjectMedia{ID: "media_1", URL: "https://cdn.example.test/products/A0001.png"},
				ProductPackageOption: product.ProductPackageOption{
					ID: "pkg_each", Code: "EACH", ProductSKUCode: "A0001",
					HandlingUnit: packaging_enums.PackageHandlingUnitEach, UnitsPerPackage: 1,
					IsCanonical: true, IsActive: true, EffectiveFrom: now,
				},
				CapturedAt:            now,
				PackageCount:          2,
				TotalBaseUnits:        2,
				PackagePrice:          money.Money{AmountMinor: 2100, Currency: "AUD"},
				Subtotal:              money.Money{AmountMinor: 4200, Currency: "AUD"},
				TaxAmount:             money.Money{AmountMinor: 382, Currency: "AUD"},
				DiscountAmount:        money.Money{AmountMinor: 100, Currency: "AUD"},
				PromotionApplications: []promotion.PromotionApplication{},
				Total:                 money.Money{AmountMinor: 4482, Currency: "AUD"},
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
	if len(decoded.Items) != 1 || decoded.Items[0].ProductSKUCode != "A0001" || decoded.Items[0].ProductName != "Potato 1kg" ||
		decoded.Items[0].ProductImage == nil || decoded.Items[0].ProductImage.ID != "media_1" || decoded.Items[0].ProductImage.URL != "https://cdn.example.test/products/A0001.png" ||
		decoded.Items[0].ProductPackageOption.ID != "pkg_each" || !decoded.Items[0].CapturedAt.Equal(now) ||
		decoded.Items[0].PackageCount != 2 || decoded.Items[0].PackagePrice.AmountMinor != 2100 ||
		decoded.Items[0].Subtotal.AmountMinor != 4200 || decoded.Items[0].TaxAmount.AmountMinor != 382 ||
		decoded.Items[0].DiscountAmount.AmountMinor != 100 || len(decoded.Items[0].PromotionApplications) != 0 {
		t.Fatalf("order summary did not round-trip: %+v", decoded)
	}
	line := got["items"].([]any)[0].(map[string]any)
	for _, key := range []string{"product_sku_code", "product_package_option", "captured_at", "package_count", "package_price", "subtotal", "tax_amount", "discount_amount", "promotion_applications"} {
		if _, exists := line[key]; !exists {
			t.Fatalf("order line missing %s: %s", key, payload)
		}
	}
	for _, removed := range []string{"sku_code", "name", "image", "image_url", "components", "accepted_package_pricing", "offer"} {
		if _, exists := line[removed]; exists {
			t.Fatalf("order line retained %s: %s", removed, payload)
		}
	}
}
