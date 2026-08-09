package product

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product/product_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestProductPackagingJSONUsesCanonicalOptionsAndBarcodes(t *testing.T) {
	effectiveFrom := time.Date(2026, 8, 4, 2, 3, 4, 0, time.UTC)
	value := Product{
		SKUCode: "A00001",
		Content: ProductContent{Name: "Potatoes"},
		Packaging: ProductPackaging{
			StorageType: warehouse_enums.StorageAmbient,
			PackageOptions: []ProductPackageOption{
				{ID: "pkg_each", Code: "EACH", ProductSKUCode: "A00001", HandlingUnit: packaging_enums.PackageHandlingUnitEach, UnitsPerPackage: 1, IsCanonical: true, IsActive: true, EffectiveFrom: effectiveFrom},
				{ID: "pkg_case_6", Code: "CASE-6", ProductSKUCode: "A00001", HandlingUnit: packaging_enums.PackageHandlingUnitCase, UnitsPerPackage: 6, IsActive: true, EffectiveFrom: effectiveFrom},
			},
			BarcodeAssignments: []ProductBarcodeAssignment{
				{ID: "barcode_each", ProductSKUCode: "A00001", PackageOptionID: "pkg_each", Value: "930000000001", Format: product_enums.BarcodeFormatEAN13, IsPrimary: true, EffectiveFrom: effectiveFrom},
			},
		},
	}

	body, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal product packages: %v", err)
	}
	for _, want := range []string{`"packaging":{"package_options"`, `"storage_type":"AMBIENT"`, `"handling_unit":"EACH"`, `"handling_unit":"CASE"`, `"units_per_package":6`, `"barcode_assignments"`, `"effective_from":"2026-08-04T02:03:04Z"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("product package JSON missing %s: %s", want, body)
		}
	}
	for _, removed := range []string{`"package_option_snapshot"`, `"barcode_assignment_snapshot"`, `"current_stock"`, `"physical"`, `"pricing"`} {
		if strings.Contains(string(body), removed) {
			t.Fatalf("product package JSON contains removed field %s: %s", removed, body)
		}
	}
}

func TestCanonicalPackageOptionRetainsEffectiveHistory(t *testing.T) {
	historicalFrom := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	historicalTo := time.Date(2026, 7, 31, 23, 59, 59, 0, time.UTC)
	options := []ProductPackageOption{
		{ID: "pkg_case_6_v1", Code: "CASE-6-V1", ProductSKUCode: "A00001", HandlingUnit: packaging_enums.PackageHandlingUnitCase, UnitsPerPackage: 6, Dimensions: &measurement.Dimensions{WidthMM: 200, LengthMM: 300, HeightMM: 150}, Weight: &measurement.Weight{Grams: 4200}, IsActive: false, EffectiveFrom: historicalFrom, EffectiveTo: &historicalTo},
		{ID: "pkg_case_8_v2", Code: "CASE-8-V2", ProductSKUCode: "A00001", HandlingUnit: packaging_enums.PackageHandlingUnitCase, UnitsPerPackage: 8, Dimensions: &measurement.Dimensions{WidthMM: 240, LengthMM: 360, HeightMM: 180}, Weight: &measurement.Weight{Grams: 5600}, IsActive: true, EffectiveFrom: historicalTo},
	}

	body, err := json.Marshal(options)
	if err != nil {
		t.Fatalf("marshal canonical package options: %v", err)
	}
	var shape []map[string]any
	if err := json.Unmarshal(body, &shape); err != nil {
		t.Fatalf("unmarshal canonical package options: %v", err)
	}
	if len(shape) != 2 || shape[0]["effective_to"] != "2026-07-31T23:59:59Z" || shape[1]["units_per_package"] != float64(8) {
		t.Fatalf("canonical package option history changed: %s", body)
	}
	if _, ok := shape[0]["captured_at"]; ok {
		t.Fatalf("canonical package option must not carry duplicate snapshot capture time: %s", body)
	}
}

func TestBarcodeValueMayBeAssignedAcrossProducts(t *testing.T) {
	now := time.Date(2026, 8, 4, 0, 0, 0, 0, time.UTC)
	products := []Product{
		{SKUCode: "A00001", Packaging: ProductPackaging{BarcodeAssignments: []ProductBarcodeAssignment{{ID: "ba_1", ProductSKUCode: "A00001", PackageOptionID: "pkg_1", Value: "930000000001", Format: product_enums.BarcodeFormatEAN13, EffectiveFrom: now}}}},
		{SKUCode: "A00002", Packaging: ProductPackaging{BarcodeAssignments: []ProductBarcodeAssignment{{ID: "ba_2", ProductSKUCode: "A00002", PackageOptionID: "pkg_2", Value: "930000000001", Format: product_enums.BarcodeFormatEAN13, EffectiveFrom: now}}}},
	}

	body, err := json.Marshal(products)
	if err != nil {
		t.Fatalf("marshal shared barcode assignments: %v", err)
	}
	if strings.Count(string(body), `"value":"930000000001"`) != 2 {
		t.Fatalf("shared barcode assignments were not represented independently: %s", body)
	}
}
