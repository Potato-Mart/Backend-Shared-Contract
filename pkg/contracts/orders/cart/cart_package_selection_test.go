package cart

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/common/packaging"
)

func TestUnpricedSelectionCarriesNoPriceOrStockClaim(t *testing.T) {
	want := CartItem{SKUCode: "SKU-1", UnpricedSelection: &CartPackageSelection{PackageOption: packaging.PackageOptionRef{SKUCode: "SKU-1", Code: "CASE6", Version: 2}, PackageCount: 2, BaseUnits: 12}, TotalBaseUnits: 12}
	raw, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `"unpriced_selection":{"package_option":{"sku_code":"SKU-1","code":"CASE6","version":2},"package_count":2,"base_units":12}`) || strings.Contains(string(raw), `"price_snapshot"`) {
		t.Fatalf("unexpected selection wire shape: %s", raw)
	}
	var got CartItem
	if err := json.Unmarshal(raw, &got); err != nil || !reflect.DeepEqual(want, got) {
		t.Fatalf("selection round trip: %+v %v", got, err)
	}
}
