package listing

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/supply/warehouse/warehouse_enums"
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
