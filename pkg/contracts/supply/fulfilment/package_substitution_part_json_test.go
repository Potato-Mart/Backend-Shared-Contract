package fulfilment_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/fulfilment"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// An assembled case is one case made up from loose units of several lots and
// buckets: the snapshot names no single lot or bucket, and its parts add up to
// the units per case.
func TestPackageSubstitutionAssembledCaseRoundTrip(t *testing.T) {
	capturedAt := time.Date(2026, 11, 17, 9, 0, 0, 0, time.UTC)
	want := fulfilment.PackageSubstitutionSnapshot{
		ID:                           "sub-1",
		RequestedCasePackageOption:   packaging.PackageOptionRef{SKUCode: "A00001", Code: "CASE12", Version: 1},
		RequestedCaseCount:           1,
		RequestedUnitsPerCase:        12,
		ReplacementEachPackageOption: packaging.PackageOptionRef{SKUCode: "A00001", Code: "EACH", Version: 1},
		ReplacementBaseUnits:         12,
		Parts: []fulfilment.PackageSubstitutionPart{
			{AllocationID: "allocation-1", BucketID: "bucket-1", LotID: "lot-1", BaseUnits: 7, StockUnitIDs: []string{"unit-1"}},
			{AllocationID: "allocation-2", BucketID: "bucket-2", LotID: "lot-2", BaseUnits: 5},
		},
		ReasonCode: warehouse_enums.PackageSubstitutionReasonCaseAssembledFromLoose.String(),
		Operator:   "packer-1",
		CapturedAt: capturedAt,
	}
	raw, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	for _, fragment := range []string{
		`"lot_id":"","source_bucket_id":"",`,
		`"parts":[{"allocation_id":"allocation-1","bucket_id":"bucket-1","lot_id":"lot-1","base_units":7,"stock_unit_ids":["unit-1"]},{"allocation_id":"allocation-2","bucket_id":"bucket-2","lot_id":"lot-2","base_units":5}],`,
		`"reason_code":"CASE_ASSEMBLED_FROM_LOOSE"`,
	} {
		if !strings.Contains(string(raw), fragment) {
			t.Fatalf("assembled case wire shape changed: missing %s in %s", fragment, raw)
		}
	}
	var got fulfilment.PackageSubstitutionSnapshot
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("assembled case round trip = %+v, want %+v", got, want)
	}
	var total int64
	for _, part := range got.Parts {
		total += part.BaseUnits
	}
	if total != got.ReplacementBaseUnits {
		t.Fatalf("assembled case parts add up to %d, want %d", total, got.ReplacementBaseUnits)
	}
}

// Snapshots recorded before v53.41.0 have no parts: they decode with none, and
// an empty parts list stays off the wire.
func TestPackageSubstitutionPartsAreOptional(t *testing.T) {
	raw := []byte(`{"id":"sub-1","requested_case_package_option":{"sku_code":"A00001","code":"CASE12","version":1},"requested_case_count":1,"requested_units_per_case":12,"fulfilled_sealed_case_count":0,"replacement_each_package_option":{"sku_code":"A00001","code":"EACH","version":1},"replacement_base_units":12,"lot_id":"lot-1","source_bucket_id":"bucket-1","reason_code":"CASE_SHORTAGE_EXACT_EACH","operator":"system:supply","captured_at":"2026-08-05T12:34:56Z"}`)
	var got fulfilment.PackageSubstitutionSnapshot
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got.Parts != nil || got.LotID != "lot-1" || got.SourceBucketID != "bucket-1" {
		t.Fatalf("legacy exact-each snapshot decoded = %+v", got)
	}
	if got.ReasonCode != warehouse_enums.PackageSubstitutionReasonCaseShortageExactEach.String() {
		t.Fatalf("legacy reason code = %q", got.ReasonCode)
	}
	encoded, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), `"parts"`) {
		t.Fatalf("empty parts must stay off the wire: %s", encoded)
	}
}
