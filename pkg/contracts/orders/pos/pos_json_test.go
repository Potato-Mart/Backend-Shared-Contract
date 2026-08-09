package pos_test

import (
	"encoding/json"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/orders/pos"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/operations"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product"

	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging/packaging_enums"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product/product_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestCatalogProductJSONUsesPackageOffersAndAvailability(t *testing.T) {
	now := time.Date(2026, 8, 4, 0, 0, 0, 0, time.UTC)
	record := pos.CatalogProduct{
		SKUCode: "A00001", CategorySKUCode: "CAT-1", Name: "Potatoes",
		StorageType:        warehouse_enums.StorageAmbient,
		PackageOptions:     []product.ProductPackageOptionSnapshot{{ID: "pkg_each", Code: "EACH", ProductSKUCode: "A00001", HandlingUnit: packaging_enums.PackageHandlingUnitEach, UnitsPerPackage: 1, EffectiveFrom: now, CapturedAt: now}},
		BarcodeAssignments: []product.ProductBarcodeAssignmentSnapshot{{ID: "barcode_1", ProductSKUCode: "A00001", PackageOptionID: "pkg_each", Value: "930000000001", Format: product_enums.BarcodeFormatEAN13, EffectiveFrom: now, CapturedAt: now}},
		Offers:             []product.SellableOfferSnapshot{},
		Availability:       &operations.ProductStockSummary{ProductSKUCode: "A00001", AllDepots: operations.ProductStockQuantitySnapshot{AvailableBaseUnits: 5}, Revision: 2, Timezone: "Australia/Melbourne", AsOf: now},
		Image:              &security.ObjectMedia{ID: "media_1", URL: "https://cdn.example.test/products/A00001.png"},
		UpdatedAt:          now,
	}

	body, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("marshal POS catalog product: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal POS catalog product: %v", err)
	}
	for _, key := range []string{"category_sku_code", "storage_type", "package_options", "barcode_assignments", "offers", "availability", "image"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("POS catalog JSON missing %s: %s", key, body)
		}
	}
	for _, removed := range []string{"sku", "barcode", "price", "current_stock", "expiry_date", "display_status", "storage", "image_url"} {
		if _, ok := got[removed]; ok {
			t.Fatalf("POS catalog JSON contains removed %s: %s", removed, body)
		}
	}
}
