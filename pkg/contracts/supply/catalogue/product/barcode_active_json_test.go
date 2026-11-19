package product

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/product/product_enums"
)

func TestBarcodeActiveFlagIsOptionalOnTheWire(t *testing.T) {
	ref := packaging.PackageOptionRef{SKUCode: "A00001", Code: "EACH", Version: 1}
	unset := ProductBarcodeAssignment{Code: "BAR-A00001", SKUCode: "A00001", PackageOption: ref, Value: "0471234500417", Format: product_enums.BarcodeFormatCode128}
	body, err := json.Marshal(unset)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(body), `"is_active"`) {
		t.Fatalf("unset is_active must be omitted, got %s", body)
	}

	off := false
	set := unset
	set.IsActive = &off
	body, err = json.Marshal(set)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `"is_active":false`) {
		t.Fatalf("explicit is_active=false must be kept, got %s", body)
	}
	var back ProductBarcodeAssignment
	if err := json.Unmarshal(body, &back); err != nil {
		t.Fatal(err)
	}
	if back.IsActive == nil || *back.IsActive {
		t.Fatalf("is_active did not round-trip: %+v", back.IsActive)
	}

	on := true
	selling := SellingProductBarcode{PackageOption: ref, Value: "0471234500417", Format: product_enums.BarcodeFormatCode128, IsActive: &on}
	body, err = json.Marshal(selling)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `"is_active":true`) {
		t.Fatalf("selling barcode must carry is_active, got %s", body)
	}
}
