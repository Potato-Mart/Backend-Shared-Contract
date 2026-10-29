package quote

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/supply/catalogue/listing"
)

func TestPriceSnapshotSaleEvidenceVariantsRoundTrip(t *testing.T) {
	at := time.Date(2026, 10, 29, 1, 2, 3, 0, time.UTC)
	option := packaging.PackageOptionRef{SKUCode: "SKU", Code: "CASE6", Version: 2}
	tests := []struct {
		name     string
		evidence SaleEvidence
		variant  string
	}{
		{"stocked", SaleEvidence{Stocked: &listing.SaleEligibilitySnapshot{SKUCode: "SKU", MarketCode: "AU", DepotCode: "D", BucketID: "B", AvailableBaseUnits: 12, InventoryRevision: 3, ValidityToken: "stock-token", CapturedAt: at, ValidUntil: at.Add(time.Minute)}}, "stocked"},
		{"wholesale backorder", SaleEvidence{WholesaleBackorder: &listing.BackorderEvidence{Allowed: true, SKUCode: "SKU", MarketCode: "AU", PackageOption: option, HandlingUnit: packaging_enums.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 6, BaseUnits: 12, ListingRevision: 4, EvidenceToken: "backorder-token", CapturedAt: at, ValidUntil: at.Add(time.Minute)}}, "wholesale_backorder"},
		{"POS merchandise", SaleEvidence{POSMerchandise: &listing.POSMerchandiseEvidence{SKUCode: "SKU", MarketCode: "AU", DepotCode: "D", PackageOption: option, HandlingUnit: packaging_enums.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 6, BaseUnits: 12, ListingRevision: 4, MarketRevision: 5, TaxCategoryRevision: 6, TaxCategoryCode: "GST", TaxRuleID: "tax-rule", TaxRuleRevision: 7, ProductFactsDigest: "facts-digest", Token: "pos-token", CapturedAt: at, ValidUntil: at.Add(time.Minute)}}, "pos_merchandise"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			original := PriceSnapshot{SaleEvidence: test.evidence}
			raw, err := json.Marshal(original)
			if err != nil {
				t.Fatal(err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(raw, &fields); err != nil {
				t.Fatal(err)
			}
			if _, exists := fields["eligibility"]; exists {
				t.Fatalf("retired field appeared: %s", raw)
			}
			var evidence map[string]json.RawMessage
			if err := json.Unmarshal(fields["sale_evidence"], &evidence); err != nil {
				t.Fatal(err)
			}
			if len(evidence) != 1 || evidence[test.variant] == nil {
				t.Fatalf("unexpected variant: %s", raw)
			}
			if test.variant != "stocked" {
				var merchandise map[string]json.RawMessage
				if err := json.Unmarshal(evidence[test.variant], &merchandise); err != nil {
					t.Fatal(err)
				}
				for _, forbidden := range []string{"bucket_id", "stock_location", "stock_unit_id", "lot_id", "available_base_units", "inventory_revision", "reservation_id"} {
					if _, exists := merchandise[forbidden]; exists {
						t.Fatalf("non-stock evidence contains %s: %s", forbidden, raw)
					}
				}
			}
			var decoded PriceSnapshot
			if err := json.Unmarshal(raw, &decoded); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(original, decoded) {
				t.Fatalf("round trip changed evidence: %+v", decoded.SaleEvidence)
			}
		})
	}
}
