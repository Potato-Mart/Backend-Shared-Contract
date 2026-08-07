package pos_test

import (
	"encoding/json"
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/orders/pos"
	"github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/product"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/product"
	warehouseenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/warehouse"
	"testing"
	"time"
)

func TestCatalogProductJSONUsesPackageOffersAndAvailability(t *testing.T) {
	now := time.Date(2026, 8, 4, 0, 0, 0, 0, time.UTC)
	record := pos.CatalogProduct{
		SKUCode: "A00001", CategorySKUCode: "CAT-1", Name: "Potatoes",
		StorageType:        warehouseenum.StorageDry,
		PackageOptions:     []product.ProductPackageOptionSnapshot{{ID: "pkg_each", Code: "EACH", ProductSKUCode: "A00001", HandlingUnit: common.PackageHandlingUnitEach, UnitsPerPackage: 1, EffectiveFrom: now, CapturedAt: now}},
		BarcodeAssignments: []product.ProductBarcodeAssignmentSnapshot{{ID: "barcode_1", ProductSKUCode: "A00001", PackageOptionID: "pkg_each", Value: "930000000001", Format: productenum.BarcodeFormatEAN13, EffectiveFrom: now, CapturedAt: now}},
		Offers:             []product.SellableOfferSnapshot{},
		Availability:       &product.ProductStockSummary{ProductSKUCode: "A00001", AllDepots: product.ProductStockQuantitySnapshot{AvailableBaseUnits: 5}, Revision: 2, Timezone: "Australia/Melbourne", AsOf: now},
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
	for _, key := range []string{"category_sku_code", "storage_type", "package_options", "barcode_assignments", "offers", "availability"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("POS catalog JSON missing %s: %s", key, body)
		}
	}
	for _, removed := range []string{"sku", "barcode", "price", "current_stock", "expiry_date", "display_status", "storage"} {
		if _, ok := got[removed]; ok {
			t.Fatalf("POS catalog JSON contains removed %s: %s", removed, body)
		}
	}
}
