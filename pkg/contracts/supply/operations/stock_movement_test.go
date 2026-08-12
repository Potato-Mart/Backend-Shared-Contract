package operations_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/operations"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestPackageConversionMovementRoundTrip(t *testing.T) {
	occurredAt := time.Date(2026, 6, 17, 9, 30, 0, 0, time.UTC)
	sourceBalance := int64(0)
	destinationBalance := int64(25)
	caseComposition := packaging.PackageCompositionSnapshot{
		TotalBaseUnits: 12,
		Components: []packaging.PackageComponentSnapshot{{
			PackageOptionID: "pkg_case_12",
			HandlingUnit:    packaging_enums.PackageHandlingUnitCase,
			PackageCount:    1,
			UnitsPerPackage: 12,
			BaseUnits:       12,
		}},
	}
	eachComposition := packaging.PackageCompositionSnapshot{
		TotalBaseUnits: 12,
		Components: []packaging.PackageComponentSnapshot{{
			PackageOptionID: "pkg_each",
			HandlingUnit:    packaging_enums.PackageHandlingUnitEach,
			PackageCount:    12,
			UnitsPerPackage: 1,
			BaseUnits:       12,
		}},
	}
	movement := operations.StockMovement{
		ID:                               "mov_1",
		SKUID:                            "A00001",
		Type:                             warehouse_enums.StockMovementTypePackageConversion,
		SourceBucketID:                   "bucket_case",
		DestinationBucketID:              "bucket_each",
		LotID:                            "lot_1",
		SourcePackageOptionID:            "pkg_case_12",
		DestinationPackageOptionID:       "pkg_each",
		BaseUnits:                        12,
		SourcePackageComposition:         &caseComposition,
		DestinationPackageComposition:    &eachComposition,
		SourceBalanceAfterBaseUnits:      &sourceBalance,
		DestinationBalanceAfterBaseUnits: &destinationBalance,
		Cause:                            &operations.InventoryCauseRef{Type: "PACKING", ID: "packing_1"},
		PerformedBy:                      "operator_1",
		OccurredAt:                       occurredAt,
	}

	payload, err := json.Marshal(movement)
	if err != nil {
		t.Fatalf("marshal stock movement: %v", err)
	}
	if !strings.Contains(string(payload), `"occurred_at":"2026-06-17T09:30:00Z"`) {
		t.Fatalf("movement timestamp is not UTC RFC3339 JSON: %s", payload)
	}

	var shape map[string]any
	if err := json.Unmarshal(payload, &shape); err != nil {
		t.Fatalf("unmarshal movement shape: %v", err)
	}
	for _, key := range []string{"qty_delta", "balance_after", "sku", "location_code"} {
		if _, exists := shape[key]; exists {
			t.Fatalf("movement retained removed key %q: %s", key, payload)
		}
	}

	var decoded operations.StockMovement
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal stock movement: %v", err)
	}
	if decoded.Type != warehouse_enums.StockMovementTypePackageConversion || decoded.BaseUnits != 12 {
		t.Fatalf("movement identity did not round-trip: %+v", decoded)
	}
	if decoded.SourcePackageComposition == nil || decoded.DestinationPackageComposition == nil {
		t.Fatalf("conversion compositions did not round-trip: %+v", decoded)
	}
	if decoded.SourcePackageComposition.TotalBaseUnits != decoded.DestinationPackageComposition.TotalBaseUnits {
		t.Fatalf("conversion base-unit totals differ: %+v", decoded)
	}
}
