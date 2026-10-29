package fulfilment_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/packaging"
	pe "github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/packaging/packaging_enums"
	fulfillment "github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/supply/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/supply/inventory"
)

func TestPackingUnidentifiedTrackedQuantityPreservesReservationOnly(t *testing.T) {
	option := packaging.PackageOptionRef{SKUCode: "SKU-1", Code: "EACH", Version: 1}
	value := fulfillment.OutboundContainerContent{OrderItemID: "item-1", SKUCode: "SKU-1", AllocationID: "allocation-1", PackageOption: option, PackedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 1, Components: []packaging.PackageComponentSnapshot{{PackageOption: option, HandlingUnit: pe.PackageHandlingUnitEach, PackageCount: 1, UnitsPerPackage: 1, BaseUnits: 1}}}, UnidentifiedTrackedPhysical: &inventory.UnidentifiedTrackedPhysicalEvidence{}}
	if err := value.ValidatePhysicalEvidence(); err != nil {
		t.Fatal(err)
	}
	raw, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err = json.Unmarshal(raw, &fields); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"bucket_id", "lot_id", "stock_unit_ids"} {
		if _, ok := fields[name]; ok {
			t.Fatalf("invented %s: %s", name, raw)
		}
	}
	if string(fields["unidentified_tracked_physical"]) != "{}" {
		t.Fatalf("missing explicit variant: %s", raw)
	}
	for name, mutate := range map[string]func(*fulfillment.OutboundContainerContent){"bucket": func(v *fulfillment.OutboundContainerContent) { v.BucketID = "invented" }, "lot": func(v *fulfillment.OutboundContainerContent) { v.LotID = "invented" }, "substitution": func(v *fulfillment.OutboundContainerContent) {
		v.Substitutions = []fulfillment.PackageSubstitutionSnapshot{{}}
	}} {
		t.Run(name, func(t *testing.T) {
			bad := value
			mutate(&bad)
			if bad.ValidatePhysicalEvidence() == nil {
				t.Fatal("unverified identity accepted")
			}
		})
	}
	forged := strings.Replace(string(raw), `"unidentified_tracked_physical":{}`, `"unidentified_tracked_physical":{"lot_id":"made-up"}`, 1)
	var decoded fulfillment.OutboundContainerContent
	if err = json.Unmarshal([]byte(forged), &decoded); err == nil {
		t.Fatal("forged source accepted")
	}
}
