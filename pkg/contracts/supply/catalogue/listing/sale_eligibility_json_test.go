package listing

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestSaleEligibilitySnapshotIsEvidenceOnly(t *testing.T) {
	capturedAt := time.Date(2026, 8, 12, 3, 4, 5, 0, time.UTC)
	value := SaleEligibilitySnapshot{
		MarketCode: "market_au", SKUCode: "sku_a00001", ListingRevision: 4,
		TaxCategoryCode: "tax_au_gst", DepotCode: "AU-VIC-MEL-DC-01",
		StockLocation: warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01"},
		BucketID:      "bucket_1", LotID: "lot_1",
		Condition:   warehouse_enums.InventoryConditionProductDamaged,
		Disposition: warehouse_enums.InventoryDispositionReducedSellable,
		DateMark: &warehouse.InventoryDateMark{
			Kind: warehouse_enums.InventoryDateMarkBestBefore, DateMarkAt: capturedAt, Timezone: "Australia/Melbourne",
		},
		ExpiryLeadDays: 30,
		DamageApproval: &DamageSaleApproval{
			QualityAssessmentID: "qa_1", Tier: warehouse_enums.DamageSaleTier50,
			ReasonCode: "CRUSHED_CARTON", ApprovedBy: "operator_1", ApprovedAt: capturedAt,
		},
		AvailableBaseUnits: 24, InventoryRevision: 11,
		ValidityToken: "eligibility_token_1", ValidUntil: capturedAt, CapturedAt: capturedAt,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal sale eligibility: %v", err)
	}
	for _, want := range []string{
		`"listing_revision":4`, `"condition":"PRODUCT_DAMAGED"`,
		`"disposition":"REDUCED_SELLABLE"`, `"date_mark"`, `"expiry_lead_days":30`,
		`"damage_approval":{"quality_assessment_id":"qa_1","tier":"tier_50"`,
		`"available_base_units":24`, `"validity_token":"eligibility_token_1"`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("SaleEligibilitySnapshot JSON = %s, want %s", payload, want)
		}
	}
	for _, forbidden := range []string{`"listing_id"`, `"amount_minor"`, `"price"`, `"discount"`, `"currency"`} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("Supply must return evidence and never a price, leaked %s: %s", forbidden, payload)
		}
	}

	var decoded SaleEligibilitySnapshot
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal sale eligibility: %v", err)
	}
	if decoded.DamageApproval == nil || decoded.DamageApproval.Tier != warehouse_enums.DamageSaleTier50 {
		t.Fatalf("damage approval did not round-trip: %+v", decoded.DamageApproval)
	}
}

func TestSaleEligibilitySnapshotPOSDeficitIsOptional(t *testing.T) {
	plain, err := json.Marshal(SaleEligibilitySnapshot{SKUCode: "sku_a00001"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(plain), `"pos_deficit"`) {
		t.Fatalf("absent POS deficit must be omitted: %s", plain)
	}

	value := SaleEligibilitySnapshot{
		SKUCode: "sku_a00001", DepotCode: "AU-VIC-MEL-DC-01", AvailableBaseUnits: 4,
		POSDeficit: &POSDeficitEligibility{
			DeficitLocationCode: "POS-DEFICIT",
			PackageOption:       packaging.PackageOptionRef{SKUCode: "sku_a00001", Code: "CASE6", Version: 1},
			HandlingUnit:        packaging_enums.PackageHandlingUnitCase,
			UnitsPerPackage:     6, ShortPackageCount: 1, ShortBaseUnits: 2,
		},
	}
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"pos_deficit":{"deficit_location_code":"POS-DEFICIT"`, `"short_package_count":1`, `"short_base_units":2`} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("SaleEligibilitySnapshot JSON = %s, want %s", payload, want)
		}
	}
	var decoded SaleEligibilitySnapshot
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.POSDeficit == nil || *decoded.POSDeficit != *value.POSDeficit {
		t.Fatalf("POS deficit did not round-trip: %+v", decoded.POSDeficit)
	}
}
