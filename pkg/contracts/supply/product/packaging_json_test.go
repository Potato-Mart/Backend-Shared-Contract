package product

import (
	"encoding/json"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/geography"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/supply/product/product_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestProductPackageOptionsAndBarcodeAssignmentsJSON(t *testing.T) {
	effectiveFrom := time.Date(2026, 8, 4, 2, 3, 4, 0, time.UTC)
	product := Product{
		SKUCode: "A00001", CategorySKUCode: "CAT-1", Name: "Potatoes",
		StorageType: warehouse_enums.StorageDry,
		PackageOptions: []ProductPackageOption{
			{ID: "pkg_each", Code: "EACH", ProductSKUCode: "A00001", HandlingUnit: packaging_enums.PackageHandlingUnitEach, UnitsPerPackage: 1, IsCanonical: true, IsActive: true, EffectiveFrom: effectiveFrom},
			{ID: "pkg_case_6", Code: "CASE-6", ProductSKUCode: "A00001", HandlingUnit: packaging_enums.PackageHandlingUnitCase, UnitsPerPackage: 6, IsActive: true, EffectiveFrom: effectiveFrom},
			{ID: "pkg_case_12", Code: "CASE-12", ProductSKUCode: "A00001", HandlingUnit: packaging_enums.PackageHandlingUnitCase, UnitsPerPackage: 12, IsActive: true, EffectiveFrom: effectiveFrom},
		},
		BarcodeAssignments: []ProductBarcodeAssignment{
			{ID: "barcode_each", ProductSKUCode: "A00001", PackageOptionID: "pkg_each", Value: "930000000001", Format: product_enums.BarcodeFormatEAN13, IsPrimary: true, EffectiveFrom: effectiveFrom},
			{ID: "barcode_case", ProductSKUCode: "A00001", PackageOptionID: "pkg_case_12", Value: "19300000000018", Format: product_enums.BarcodeFormatEAN13, IsPrimary: true, EffectiveFrom: effectiveFrom},
		},
	}

	body, err := json.Marshal(product)
	if err != nil {
		t.Fatalf("marshal product packages: %v", err)
	}
	for _, want := range []string{`"category_sku_code":"CAT-1"`, `"storage_type":"DRY"`, `"handling_unit":"EACH"`, `"handling_unit":"CASE"`, `"units_per_package":6`, `"units_per_package":12`, `"barcode_assignments"`, `"effective_from":"2026-08-04T02:03:04Z"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("product package JSON missing %s: %s", want, body)
		}
	}
	for _, removed := range []string{`"sku":`, `"barcode":`, `"placing_area":`, `"current_stock":`, `"expired_at":`, `"restocked_at":`, `"display_status":`, `"physical":`, `"pricing":`, `"storage":`} {
		if strings.Contains(string(body), removed) {
			t.Fatalf("product package JSON contains removed field %s: %s", removed, body)
		}
	}
}

func TestHistoricalEndedPackageOptionSnapshotCoexistsWithCurrentOptionJSON(t *testing.T) {
	historicalFrom := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	historicalTo := time.Date(2026, 7, 31, 23, 59, 59, 0, time.UTC)
	capturedAt := time.Date(2026, 8, 4, 2, 3, 4, 0, time.UTC)
	snapshots := []ProductPackageOptionSnapshot{
		{
			ID:              "pkg_case_6_v1",
			Code:            "CASE-6-V1",
			ProductSKUCode:  "A00001",
			HandlingUnit:    packaging_enums.PackageHandlingUnitCase,
			UnitsPerPackage: 6,
			Dimensions:      &measurement.Dimensions{WidthMM: 200, LengthMM: 300, HeightMM: 150},
			Weight:          &measurement.Weight{Grams: 4200},
			EffectiveFrom:   historicalFrom,
			EffectiveTo:     &historicalTo,
			CapturedAt:      capturedAt,
		},
		{
			ID:              "pkg_case_8_v2",
			Code:            "CASE-8-V2",
			ProductSKUCode:  "A00001",
			HandlingUnit:    packaging_enums.PackageHandlingUnitCase,
			UnitsPerPackage: 8,
			Dimensions:      &measurement.Dimensions{WidthMM: 240, LengthMM: 360, HeightMM: 180},
			Weight:          &measurement.Weight{Grams: 5600},
			EffectiveFrom:   historicalTo,
			CapturedAt:      capturedAt,
		},
	}

	body, err := json.Marshal(snapshots)
	if err != nil {
		t.Fatalf("marshal historical package option snapshots: %v", err)
	}
	var shape []map[string]any
	if err := json.Unmarshal(body, &shape); err != nil {
		t.Fatalf("unmarshal historical package option snapshots: %v", err)
	}
	if len(shape) != 2 {
		t.Fatalf("package option snapshot count = %d, want 2: %s", len(shape), body)
	}

	historical, current := shape[0], shape[1]
	if historical["id"] != "pkg_case_6_v1" || historical["units_per_package"] != float64(6) || historical["effective_to"] != "2026-07-31T23:59:59Z" {
		t.Fatalf("historical ended option identity, quantity, or effective_to changed: %+v", historical)
	}
	dimensions, ok := historical["dimensions"].(map[string]any)
	if !ok {
		t.Fatalf("historical dimensions were not preserved: %+v", historical)
	}
	if dimensions["width_mm"] != float64(200) || dimensions["length_mm"] != float64(300) || dimensions["height_mm"] != float64(150) {
		t.Fatalf("historical dimensions changed: %+v", dimensions)
	}
	if historical["weight"].(map[string]any)["grams"] != float64(4200) {
		t.Fatalf("historical weight changed: %+v", historical["weight"])
	}
	if current["id"] != "pkg_case_8_v2" || current["units_per_package"] != float64(8) {
		t.Fatalf("current package option identity or quantity changed: %+v", current)
	}
	currentDimensions := current["dimensions"].(map[string]any)
	if currentDimensions["width_mm"] != float64(240) || currentDimensions["length_mm"] != float64(360) || currentDimensions["height_mm"] != float64(180) {
		t.Fatalf("current dimensions changed: %+v", currentDimensions)
	}
	if current["weight"].(map[string]any)["grams"] != float64(5600) {
		t.Fatalf("current weight changed: %+v", current["weight"])
	}
	if _, ok := current["effective_to"]; ok {
		t.Fatalf("current package option unexpectedly ended: %+v", current)
	}
}

func TestBarcodeValueMayBeAssignedAcrossProducts(t *testing.T) {
	now := time.Date(2026, 8, 4, 0, 0, 0, 0, time.UTC)
	products := []Product{
		{SKUCode: "A00001", Name: "First", BarcodeAssignments: []ProductBarcodeAssignment{{ID: "ba_1", ProductSKUCode: "A00001", PackageOptionID: "pkg_1", Value: "930000000001", Format: product_enums.BarcodeFormatEAN13, EffectiveFrom: now}}},
		{SKUCode: "A00002", Name: "Second", BarcodeAssignments: []ProductBarcodeAssignment{{ID: "ba_2", ProductSKUCode: "A00002", PackageOptionID: "pkg_2", Value: "930000000001", Format: product_enums.BarcodeFormatEAN13, EffectiveFrom: now}}},
	}

	body, err := json.Marshal(products)
	if err != nil {
		t.Fatalf("marshal shared barcode assignments: %v", err)
	}
	if strings.Count(string(body), `"value":"930000000001"`) != 2 {
		t.Fatalf("shared barcode assignments were not represented independently: %s", body)
	}
}

func TestSellableOfferSnapshotJSONFreezesPackageInventoryTimeAndGeography(t *testing.T) {
	now := time.Date(2026, 8, 4, 5, 6, 7, 0, time.UTC)
	offer := SellableOfferSnapshot{
		ID: "offer_1", ProductSKUCode: "A00001", DepotCode: "AU-VIC-MEL-DC-01", SourceBucketID: "bucket_1",
		PackageOption:         ProductPackageOptionSnapshot{ID: "pkg_case_6", Code: "CASE-6", ProductSKUCode: "A00001", HandlingUnit: packaging_enums.PackageHandlingUnitCase, UnitsPerPackage: 6, EffectiveFrom: now, CapturedAt: now},
		AvailablePackageCount: 3, AvailableBaseUnits: 18,
		Condition: warehouse_enums.InventoryConditionGood, Disposition: warehouse_enums.InventoryDispositionStandardSellable,
		DateMark: &SellableOfferDateMarkSnapshot{Kind: warehouse_enums.InventoryDateMarkBestBefore, DateMarkAt: now.Add(30 * 24 * time.Hour), Timezone: "Australia/Melbourne"},
		Revision: 7, InventoryRevision: 13,
		PackagePrice: money.Money{AmountMinor: 1200, Currency: "AUD"}, TaxAmount: money.Money{AmountMinor: 109, Currency: "AUD"},
		Discounts: []SellableOfferDiscountSnapshot{{ID: "promo_1", Amount: money.Money{AmountMinor: 200, Currency: "AUD"}}},
		ValidFrom: now, Timezone: "Etc/UTC", CapturedAt: now,
		GeographicContext: geography.GeographicContext{Source: geography_enums.GeographicContextSourceRetailCustomerProfile, CountryCode: "AU", SubdivisionCode: "AU-VIC", DepotCode: "AU-VIC-MEL-DC-01", ScopeRevision: 2, RuleRevision: 7, EvaluationTimezone: "Australia/Melbourne"},
	}

	body, err := json.Marshal(offer)
	if err != nil {
		t.Fatalf("marshal sellable offer snapshot: %v", err)
	}
	for _, want := range []string{`"available_base_units":18`, `"inventory_revision":13`, `"date_mark_at":"2026-09-03T05:06:07Z"`, `"timezone":"Australia/Melbourne"`, `"geographic_context"`, `"captured_at":"2026-08-04T05:06:07Z"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("sellable offer snapshot JSON missing %s: %s", want, body)
		}
	}
}
