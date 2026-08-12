package pos_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging/packaging_enums"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/security"
	sales "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/orders/order"
	pos "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/orders/pos"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/promotion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/product"
)

func TestReceiptSnapshotJSONUsesCustomerSafeFrozenLines(t *testing.T) {
	now := time.Date(2026, 8, 7, 3, 4, 5, 0, time.UTC)
	receipt := pos.ReceiptSnapshot{
		OrderNumber: "SO-1",
		Revision:    2,
		IssuedAt:    now,
		Attribution: sales.POSAttribution{StoreID: "store_1", RegisterID: "register_1"},
		Lines: []pos.ReceiptLine{{
			SKUID:        "A00001",
			ProductName:  "Potatoes 1kg",
			ProductImage: &security.ObjectMedia{ID: "media_1", URL: "https://cdn.example.test/products/A00001.png"},
			ProductPackageOption: product.ProductPackageOption{
				ID: "pkg_each", Code: "EACH", SKUID: "A00001",
				HandlingUnit: packaging_enums.PackageHandlingUnitEach, UnitsPerPackage: 1,
				IsCanonical: true, IsActive: true, EffectiveFrom: now,
			},
			CapturedAt:            now,
			PackageCount:          2,
			TotalBaseUnits:        2,
			PackagePrice:          money.Money{AmountMinor: 500, Currency: "AUD"},
			Subtotal:              money.Money{AmountMinor: 1000, Currency: "AUD"},
			TaxAmount:             money.Money{AmountMinor: 91, Currency: "AUD"},
			DiscountAmount:        money.Money{AmountMinor: 100, Currency: "AUD"},
			Total:                 money.Money{AmountMinor: 991, Currency: "AUD"},
			PromotionApplications: []promotion.PromotionApplication{},
		}},
		Subtotal:              money.Money{AmountMinor: 1000, Currency: "AUD"},
		Tax:                   money.Money{AmountMinor: 91, Currency: "AUD"},
		Total:                 money.Money{AmountMinor: 991, Currency: "AUD"},
		PromotionApplications: []promotion.PromotionApplication{},
	}

	payload, err := json.Marshal(receipt)
	if err != nil {
		t.Fatalf("marshal receipt: %v", err)
	}
	for _, want := range []string{`"sku_id":"A00001"`, `"product_name":"Potatoes 1kg"`, `"product_image"`, `"product_package_option"`, `"captured_at":"2026-08-07T03:04:05Z"`, `"package_count":2`, `"package_price"`, `"subtotal"`, `"tax_amount"`, `"discount_amount"`, `"promotion_applications":[]`} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("receipt JSON missing %s: %s", want, payload)
		}
	}
	for _, removed := range []string{`"accepted_package_pricing"`, `"available_package_count"`, `"available_base_units"`, `"source_bucket_id"`, `"source_stock_unit_id"`, `"stock_location"`, `"offers"`, `"offer"`} {
		if strings.Contains(string(payload), removed) {
			t.Fatalf("receipt JSON exposes removed pricing/inventory field %s: %s", removed, payload)
		}
	}
}
