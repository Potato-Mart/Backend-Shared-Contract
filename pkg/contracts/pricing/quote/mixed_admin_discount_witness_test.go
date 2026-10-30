package quote

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
)

func TestMixedAdminWitnessIsAdditiveAndRoundTrips(t *testing.T) {
	legacy := AmendmentAllocationAdjustment{Version: 1}
	raw, err := json.Marshal(legacy)
	if err != nil || strings.Contains(string(raw), `"admin_discount_witness"`) {
		t.Fatalf("legacy encoding changed: %s %v", raw, err)
	}
	amount := func(v int64) money.Money { return money.Money{Currency: "AUD", AmountMinor: v} }
	legacy.AdminDiscountWitness = &MixedAdminLineWitness{Version: 1, Operation: "new", LineID: "line", GlobalWitnessHash: "hash", PreAdminGross: amount(10), PreAdminTax: amount(10), CarriedGross: amount(0), CarriedTax: amount(0), AppliedGross: amount(10), AppliedTax: amount(10), Buckets: []MixedAdminBucketWitness{{TaxCategoryCode: "gst", FirstSourceUnit: 2, PreAdminBase: amount(0), PreAdminTax: amount(10), AppliedGross: amount(10), AppliedTax: amount(10), TaxLowerBound: amount(10), TaxUpperBound: amount(10)}}}
	raw, err = json.Marshal(legacy)
	if err != nil {
		t.Fatal(err)
	}
	var got AmendmentAllocationAdjustment
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got.AdminDiscountWitness == nil || got.AdminDiscountWitness.GlobalWitnessHash != "hash" || len(got.AdminDiscountWitness.Buckets) != 1 || got.AdminDiscountWitness.Buckets[0].FirstSourceUnit != 2 || got.AdminDiscountWitness.Buckets[0].AppliedTax.AmountMinor != 10 {
		t.Fatalf("witness lost: %+v", got)
	}
	global := MixedAdminDiscountWitness{Version: 1, Operation: "new", SourceVersionID: "v", SourceQuoteKey: "q", SourceQuoteRevision: "r", EditID: "edit", PricedLinesHash: "lines-hash", AdministrativeDiscountHash: "discount-hash", PreviousDesired: amount(0), CurrentDesired: amount(10), PreAdminTotal: amount(10), Lines: []MixedAdminDiscountLine{{LineID: "line", Kind: "retained", SourceVersionID: "v", FirstPackageOrdinal: 1, LastPackageOrdinal: 1, PreAdminGross: amount(10), PreAdminTax: amount(10), CarriedGross: amount(0), CarriedTax: amount(0), AppliedGross: amount(10), AppliedTax: amount(10)}}}
	raw, err = json.Marshal(global)
	if err != nil {
		t.Fatal(err)
	}
	var decoded MixedAdminDiscountWitness
	if err := json.Unmarshal(raw, &decoded); err != nil || decoded.PricedLinesHash != "lines-hash" || decoded.AdministrativeDiscountHash != "discount-hash" || len(decoded.Lines) != 1 || decoded.Lines[0].PreAdminTax.AmountMinor != 10 || decoded.Lines[0].SourceVersionID != "v" {
		t.Fatalf("global witness lost: %+v %v", decoded, err)
	}
}
